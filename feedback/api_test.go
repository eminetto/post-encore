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

	newCTX := context.WithValue(context.Background(), feedback.EmailKeyValue, "email@email.com")
	resp, err := api.StoreFeedback(newCTX, &p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID == "" {
		t.Fatalf("expected ID to be non-empty")
	}
}
