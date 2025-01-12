package vote

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("vote", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// EmailKey is the key used to store the email in the context
type EmailKey string

const emailKey = EmailKey("email")

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
//encore:api public method=POST path=/v1/vote tag:authenticated
func StoreVote(ctx context.Context, p *StoreVoteParams) (*StoreVoteResponse, error) {
	eb := errs.B().Meta("store_vote", p.TalkName)
	email := ctx.Value(emailKey).(string)
	s := NewService(db)
	v := &Vote{
		Email:    email,
		TalkName: p.TalkName,
		Score:    p.Score,
	}
	id, err := s.Store(ctx, v)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	return &StoreVoteResponse{ID: id}, nil
}
