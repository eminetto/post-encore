package audit

import (
	"context"

	"encore.app/user"
	"encore.dev/pubsub"
	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
)

var db = sqldb.NewDatabase("audit", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

var _ = pubsub.NewSubscription(
	user.AuthEvents, "auth-audit",
	pubsub.SubscriptionConfig[*user.AuthEvent]{
		Handler: Auth,
	},
)

// Auth logs user authentication events
func Auth(ctx context.Context, event *user.AuthEvent) error {
	id := uuid.New().String()
	_, err := db.Exec(ctx, `
		INSERT INTO audit_auths (id, email, created_at)
		VALUES ($1, $2,now())
	`, id, event.UserEmail)

	return err
}
