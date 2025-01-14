package feedback_test

import (
	"context"
	"testing"

	"encore.app/feedback"
	"github.com/google/uuid"
)

type ServiceMock struct{}

func (s *ServiceMock) Store(ctx context.Context, f *feedback.Feedback) (string, error) {
	return uuid.New().String(), nil
}

func TestIntegration(t *testing.T) {
	api := feedback.API{
		Service: &ServiceMock{},
	}
	p := feedback.StoreFeedbackParams{
		Title: "title",
		Body:  "body",
	}

	ctx := context.WithValue(context.Background(), "Email", "email@email.com")
	resp, err := api.StoreFeedback(ctx, &p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID == "" {
		t.Fatalf("expected ID to be non-empty")
	}
}
