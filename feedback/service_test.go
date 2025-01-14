package feedback_test

import (
	"context"
	"testing"

	"encore.app/feedback"
	"encore.dev/et"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	et.EnableServiceInstanceIsolation()
	testDB, err := et.NewTestDatabase(ctx, "feedback")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	s := feedback.NewService(testDB)
	t.Run("store vote", func(t *testing.T) {
		id, err := s.Store(ctx, &feedback.Feedback{
			Email: "email",
			Title: "title",
			Body:  "body",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id == "" {
			t.Fatalf("expected ID to be non-empty")
		}
	})
}
