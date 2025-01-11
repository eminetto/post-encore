package feedback

import (
	"context"

	"encore.app/user"
	"encore.dev/beta/errs"
	"encore.dev/middleware"
)

// IsAuthenticated checks if the user is authenticated.
//
//encore:middleware target=tag:authenticated
func IsAuthenticated(req middleware.Request, next middleware.Next) middleware.Response {
	eb := errs.B().Meta("is_authenticated", "true")
	token := req.Data().Headers["Authorization"]
	if len(token) == 0 || token[0] == "" {
		return middleware.Response{Err: eb.Code(errs.Unauthenticated).Msg("missing token").Err()}
	}
	resp, err := user.ValidateToken(req.Context(), &user.ValidateTokenParams{Token: token[0]})
	if err != nil {
		return middleware.Response{Err: eb.Code(errs.Unauthenticated).Msg("invalid token").Err()}
	}
	newCTX := context.WithValue(req.Context(), emailKey, resp.Email)
	return next(req.WithContext(newCTX))
}
