package vote

import (
	"context"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
)

// Service is the service for the vote package
type Service struct {
	DB *sqldb.Database
}

// NewService creates a new vote service
func NewService(db *sqldb.Database) *Service {
	return &Service{DB: db}
}

// Store stores a vote
func (s *Service) Store(ctx context.Context, v *Vote) (string, error) {
	v.ID = uuid.New().String()
	_, err := s.DB.Exec(ctx, `
		INSERT INTO votes (id, email, talk_name, score, created_at)
		VALUES ($1, $2, $3, $4, now())
	`, v.ID, v.Email, v.TalkName, v.Score)

	return v.ID, err
}
