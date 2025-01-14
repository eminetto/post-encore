package vote

import (
	"context"

	"encore.app/authentication"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("vote", sqldb.DatabaseConfig{
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

// EmailKey is the key used to store the email in the context
type EmailKey string

// EmailKeyValue is the key used to store the email in the context
const EmailKeyValue = EmailKey("email")

// StoreVoteParams represents the response of the StoreVote function
type StoreVoteParams struct {
	TalkName string `json:"talk_name"`
	Score    int    `json:"score,string"`
}

// StoreVoteResponse represents the response of the StoreVote function
type StoreVoteResponse struct {
	ID string `json:"id"`
}

// StoreVote stores vote
//
//encore:api auth method=POST path=/v1/vote tag:authenticated
func (a *API) StoreVote(ctx context.Context, p *StoreVoteParams) (*StoreVoteResponse, error) {
	eb := errs.B().Meta("store_vote", p.TalkName)
	var email string
	data := auth.Data()
	if data != nil {
		email = data.(*authentication.Data).Email
	}
	if email == "" {
		email = ctx.Value("Email").(string)
	}
	if email == "" {
		return nil, eb.Code(errs.Unauthenticated).Msg("unauthenticated").Err()
	}
	v := &Vote{
		Email:    email,
		TalkName: p.TalkName,
		Score:    p.Score,
	}
	id, err := a.Service.Store(ctx, v)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	return &StoreVoteResponse{ID: id}, nil
}
