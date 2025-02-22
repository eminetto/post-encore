package feedback_test

import (
	"context"
	"encore.app/authentication"
	"encore.app/feedback"
	"encore.dev/et"
	"github.com/google/uuid"
	"testing"
)

type ServiceMock struct{}

func (s *ServiceMock) Store(ctx context.Context, f *feedback.Feedback) (string, error) {
	return uuid.New().String(), nil
}

func TestStoreFeedback(t *testing.T) {
	api := feedback.API{
		Service: &ServiceMock{},
	}
	et.OverrideAuthInfo("uuid", &authentication.Data{Email: "eminetto@email.com"})
	p := feedback.StoreFeedbackParams{
		Title: "title",
		Body:  "body",
	}

	resp, err := api.StoreFeedback(context.Background(), &p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID == "" {
		t.Fatalf("expected ID to be non-empty")
	}
}
