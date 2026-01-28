package jwtauth

import (
	"context"
)

// contextKey is a type for context keys.
type contextKey string

// Context keys for JWT claims.
const (
	ContextKeyClaims contextKey = "jwt_claims"
)

// ClaimsFromContext retrieves JWT claims from the request context.
func ClaimsFromContext(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(ContextKeyClaims).(*Claims); ok {
		return claims
	}
	return nil
}
