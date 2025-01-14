package authentication

import (
	"context"

	"encore.app/user"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"github.com/google/uuid"
)

// Data is the auth data
type Data struct {
	Email string
}

// AuthHandler handle auth information
//
//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, *Data, error) {
	if token == "" {
		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid token",
		}
	}
	resp, err := user.ValidateToken(ctx, &user.ValidateTokenParams{Token: token})
	if err != nil {
		return "", nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid token",
		}
	}
	return auth.UID(uuid.New().String()), &Data{Email: resp.Email}, nil
}
