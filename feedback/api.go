package feedback

import (
	"context"
	"encore.app/authentication"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("feedback", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// API defines the API for the user service
// encore: service
type API struct {
	Service UseCase
}

func initAPI() (*API, error) {
	return &API{Service: NewService(db)}, nil
}

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
//encore:api auth method=POST path=/v1/feedback
func (a *API) StoreFeedback(ctx context.Context, p *StoreFeedbackParams) (*StoreFeedbackResponse, error) {
	eb := errs.B().Meta("store_feedback", p.Title)
	var email string
	data := auth.Data()
	if data != nil {
		email = data.(*authentication.Data).Email
	}
	if email == "" {
		return nil, eb.Code(errs.Unauthenticated).Msg("unauthenticated").Err()
	}
	f := &Feedback{
		Email: email,
		Title: p.Title,
		Body:  p.Body,
	}
	id, err := a.Service.Store(ctx, f)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	return &StoreFeedbackResponse{ID: id}, nil
}
