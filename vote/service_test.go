package vote_test

import (
	"context"
	"testing"

	"encore.app/vote"
	"encore.dev/et"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	et.EnableServiceInstanceIsolation()
	testDB, err := et.NewTestDatabase(ctx, "vote")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	s := vote.NewService(testDB)
	t.Run("store vote", func(t *testing.T) {
		id, err := s.Store(ctx, &vote.Vote{
			Email:    "email",
			TalkName: "talk",
			Score:    5,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id == "" {
			t.Fatalf("expected ID to be non-empty")
		}
	})
}
