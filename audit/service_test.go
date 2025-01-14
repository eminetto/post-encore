package audit

import (
	"context"
	"testing"

	"encore.app/user"
	"encore.dev/et"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	et.EnableServiceInstanceIsolation()
	var err error
	db, err = et.NewTestDatabase(ctx, "audit")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = Auth(ctx, &user.AuthEvent{UserEmail: "email"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
