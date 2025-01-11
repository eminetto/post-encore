package feedback

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("feedback", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// EmailKey is the key used to store the email in the context
type EmailKey string

const emailKey = EmailKey("email")

// StoreFeedbackParams represents the response of the StoreFeedback function
type StoreFeedbackParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// StoreFeedbackResponse represents the response of the StoreFeedback function
type StoreFeedbackResponse struct {
	ID string `json:"id"`
}

// StoreFeedback stores feedback
//
//encore:api public method=POST path=/v1/feedback tag:authenticated
func StoreFeedback(ctx context.Context, p *StoreFeedbackParams) (*StoreFeedbackResponse, error) {
	eb := errs.B().Meta("store_feedback", p.Title)
	email := ctx.Value(emailKey).(string)
	s := NewService(db)
	f := &Feedback{
		Email: email,
		Title: p.Title,
		Body:  p.Body,
	}
	id, err := s.Store(ctx, f)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	return &StoreFeedbackResponse{ID: id}, nil
}
