package user

import (
	"context"
	"crypto/sha1"
	"fmt"

	"encore.dev/storage/sqldb"
)

// UseCase is user logic interface
type UseCase interface {
	ValidateUser(ctx context.Context, email, password string) error
	ValidatePassword(ctx context.Context, u *User, password string) error
}

// Service is the service for the user package
type Service struct {
	DB *sqldb.Database
}

// NewService creates a new user service
func NewService(db *sqldb.Database) *Service {
	return &Service{DB: db}
}

// ValidateUser validates a user
func (s *Service) ValidateUser(ctx context.Context, email, password string) error {
	var u User
	err := db.QueryRow(ctx, `
        select id, email, password, first_name, last_name from users where email = $1
    `, email).Scan(&u.ID, &u.Email, &u.Password, &u.FirstName, &u.LastName)
	if err != nil {
		return fmt.Errorf("invalid user %w", err)
	}
	err = s.ValidatePassword(ctx, &u, password)
	if err != nil {
		return fmt.Errorf("invalid user")
	}
	return nil
}

// ValidatePassword validates a password
func (s *Service) ValidatePassword(ctx context.Context, u *User, password string) error {
	h := sha1.New()
	h.Write([]byte(password))
	p := fmt.Sprintf("%x", h.Sum(nil))
	if p != u.Password {
		return fmt.Errorf("invalid password")
	}
	return nil
}
