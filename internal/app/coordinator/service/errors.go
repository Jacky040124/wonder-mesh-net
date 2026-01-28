package service

import "errors"

// OIDC service errors.
var (
	ErrInvalidState    = errors.New("invalid or expired state")
	ErrStateExpired    = errors.New("state expired")
	ErrTokenExchange   = errors.New("token exchange failed")
	ErrInvalidIDToken  = errors.New("invalid ID token")
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
)

// Worker service errors.
var (
	ErrInvalidToken = errors.New("invalid or expired token")
)
