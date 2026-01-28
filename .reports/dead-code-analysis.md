# Dead Code Analysis Report

Generated: 2026-01-27

## Summary

| Category | Count | Status |
|----------|-------|--------|
| Unused Functions | 6 | **REMOVED** |
| Unused Error Variables | 3 | **REMOVED** |
| Unused Dependencies | 0 | N/A |

**Total lines removed**: ~75 lines of dead code

## Changes Applied

All dead code identified below has been safely removed with test verification.

### Removed Functions

#### 1. `jwtauth.OptionalMiddleware` - REMOVED
- **File**: `pkg/jwtauth/middleware.go`
- **Description**: HTTP middleware that validates JWT tokens optionally
- **Reason**: Never called. Coordinator uses custom auth middleware in `server.go`

#### 2. `jwtauth.Middleware` - REMOVED
- **File**: `pkg/jwtauth/middleware.go`
- **Description**: HTTP middleware that requires valid JWT tokens
- **Reason**: Never called. Coordinator implements its own auth middleware

#### 3. `jwtauth.extractBearerToken` (private) - REMOVED
- **File**: `pkg/jwtauth/middleware.go`
- **Description**: Helper to extract Bearer token from Authorization header
- **Reason**: Only used by removed middleware functions

#### 4. `jointoken.EncodeForCLI` - REMOVED
- **File**: `pkg/jointoken/token.go`
- **Description**: URL-safe base64 encoding for CLI copy-paste
- **Reason**: Never called. CLI uses raw JWT tokens directly

#### 5. `jointoken.DecodeFromCLI` - REMOVED
- **File**: `pkg/jointoken/token.go`
- **Description**: Decodes CLI-encoded tokens
- **Reason**: Companion to unused `EncodeForCLI`

#### 6. `jwtauth.Claims.HasRole` - REMOVED
- **File**: `pkg/jwtauth/validator.go`
- **Description**: Check for Keycloak realm role
- **Reason**: Never called. Only `IsServiceAccount()` was used

#### 7. `jwtauth.Claims.HasClientRole` - REMOVED
- **File**: `pkg/jwtauth/validator.go`
- **Description**: Check for Keycloak client role
- **Reason**: Never called

### Removed Error Variables

#### 1. `service.ErrProviderNotFound` - REMOVED
- **File**: `internal/app/coordinator/service/errors.go`
- **Reason**: Defined but never returned or checked

#### 2. `service.ErrInvalidRedirectURI` - REMOVED
- **File**: `internal/app/coordinator/service/errors.go`
- **Reason**: Defined but never returned or checked

#### 3. `service.ErrMissingCode` - REMOVED
- **File**: `internal/app/coordinator/service/errors.go`
- **Reason**: Defined but never returned or checked

## Verified Still In Use

The following items were investigated and confirmed to be in use:

- `jwtauth.ClaimsFromContext` - Used in `server.go:221`
- `jwtauth.Claims.IsServiceAccount` - Used in `wondernet.go:152`
- `jwtauth.ContextKeyClaims` - Used in `server.go` for context keys
- `jointoken.GetJoinInfo` - Used in `worker/join.go:72`
- `jointoken.ParseUnsafe` - Used by `GetJoinInfo`
- `headscale.ACLManager` - Used in `server.go:111` and `wondernet.go`
- `wondersdk.Client` - Used in `examples/kubeadm-deployer/deployer`
- All remaining service errors are properly used

## Test Verification

All tests pass after changes:

```
ok   github.com/strrl/wonder-mesh-net/cmd/wonder/commands/worker
ok   github.com/strrl/wonder-mesh-net/internal/app/coordinator
ok   github.com/strrl/wonder-mesh-net/internal/app/coordinator/controller
ok   github.com/strrl/wonder-mesh-net/internal/app/coordinator/service
```

Build verification:
```
go build ./... # SUCCESS
go vet ./...   # PASS
```

## Pre-existing Issues (Not Related to Dead Code)

The `make check` has 12 pre-existing `errcheck` issues in:
- `examples/kubeadm-deployer/` - unchecked error returns
- `internal/app/coordinator/service/oidc.go` - unchecked `resp.Body.Close()`

These are unrelated to the dead code cleanup and existed before this analysis.
