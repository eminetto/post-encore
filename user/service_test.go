package user_test

import (
	"context"
	"testing"

	"encore.app/user"
	"encore.dev/et"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	et.EnableServiceInstanceIsolation()
	testDB, err := et.NewTestDatabase(ctx, "user")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	s := user.NewService(testDB)
	t.Run("valid user", func(t *testing.T) {
		err := s.ValidateUser(ctx, "eminetto@email.com", "12345")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("invalid user", func(t *testing.T) {
		err := s.ValidateUser(ctx, "e@email.com", "12345")
		if err == nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("invalid password", func(t *testing.T) {
		err := s.ValidateUser(ctx, "eminetto@email.com", "111")
		if err == nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
