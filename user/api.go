package user

import (
	"context"

	"encore.app/user/security"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("user", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// AuthParams are the parameters to the Auth method
type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse is the response to the Auth method
type AuthResponse struct {
	Token string `json:"token"`
}

// Auth authenticates a user and returns a token
//
//encore:api public method=POST path=/v1/auth
func Auth(ctx context.Context, p *AuthParams) (*AuthResponse, error) {
	s := NewService(db)
	// Construct a new error builder with errs.B()
	eb := errs.B().Meta("auth", p.Email)

	err := s.ValidateUser(ctx, p.Email, p.Password)
	if err != nil {
		return nil, eb.Code(errs.Unauthenticated).Msg("invalid credentials").Err()
	}
	var response AuthResponse
	response.Token, err = security.NewToken(p.Email)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	return &response, nil
}

// ValidateTokenParams are the parameters to the ValidateToken method
type ValidateTokenParams struct {
	Token string `json:"token"`
}

// ValidateTokenResponse is the response to the ValidateToken method
type ValidateTokenResponse struct {
	Email string `json:"email"`
}

// ValidateToken validates a token
//
//encore:api public method=POST path=/v1/validate-token
func ValidateToken(ctx context.Context, p *ValidateTokenParams) (*ValidateTokenResponse, error) {
	// Construct a new error builder with errs.B()
	eb := errs.B().Meta("validate_token", p.Token)
	t, err := security.ParseToken(p.Token)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return nil, eb.Code(errs.Internal).Msg("internal error").Err()
	}
	response := ValidateTokenResponse{
		Email: tData["email"].(string),
	}
	return &response, nil
}
