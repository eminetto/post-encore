package user_test

import (
	"context"
	"testing"

	"encore.app/user"
)

type ServiceMock struct{}

func (s *ServiceMock) ValidateUser(ctx context.Context, email string, password string) error {
	return nil
}

func (s *ServiceMock) ValidatePassword(ctx context.Context, u *user.User, password string) error {
	return nil
}

func TestIntegration(t *testing.T) {
	api := &user.API{
		Service: &ServiceMock{},
	}
	email := "eminetto@email.com"
	resp, err := api.Auth(context.Background(), &user.AuthParams{
		Email:    email,
		Password: "12345",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Token == "" {
		t.Fatalf("expected token to be non-empty")
	}
	r, err := api.ValidateToken(context.Background(), &user.ValidateTokenParams{
		Token: resp.Token,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Email != email {
		t.Fatalf("expected email to be %q, got %q", email, r.Email)
	}
}
