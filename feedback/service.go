package feedback

import (
	"context"

	"encore.dev/storage/sqldb"
	"github.com/google/uuid"
)

// Service is the service for the feedback package
type Service struct {
	DB *sqldb.Database
}

// NewService creates a new feedback service
func NewService(db *sqldb.Database) *Service {
	return &Service{DB: db}
}

// Store stores feedback
func (s *Service) Store(ctx context.Context, f *Feedback) (string, error) {
	f.ID = uuid.New().String()
	_, err := s.DB.Exec(ctx, `
		INSERT INTO feedbacks (id, email, title, body, created_at)
		VALUES ($1, $2, $3, $4, now())
	`, f.ID, f.Email, f.Title, f.Body)

	return f.ID, err
}
