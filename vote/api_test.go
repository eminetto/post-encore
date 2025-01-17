package vote_test

import (
	"context"
	"testing"

	"encore.app/vote"
	"github.com/google/uuid"
)

type ServiceMock struct{}

func (s *ServiceMock) Store(ctx context.Context, v *vote.Vote) (string, error) {
	return uuid.New().String(), nil
}

func TestIntegration(t *testing.T) {
	api := vote.API{
		Service: &ServiceMock{},
	}
	p := vote.StoreVoteParams{
		TalkName: "talk",
		Score:    5,
	}
	// ctx := auth.WithContext(context.Background(), "uuid", &authentication.Data{Email:"eminetto@email.com"})
	ctx := context.WithValue(context.Background(), "Email", "email@email.com")
	resp, err := api.StoreVote(ctx, &p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID == "" {
		t.Fatalf("expected ID to be non-empty")
	}
}
