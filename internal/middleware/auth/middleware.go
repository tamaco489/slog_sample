package auth

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error)
}

// WithAuth: Authorization middleware
func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// NOTE: Return authorization failure for testing
		// authInfo, err := authorizer.Authorize(r.Context(), r)
		// logger.GetLogger().DebugContext(r.Context(), "Authorization failed", "error", fmt.Errorf("Authorization failed"))
		// http.Error(w, fmt.Errorf("Authorization failed").Error(), http.StatusUnauthorized)
		// return

		// Execute authorization process
		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			// Set authentication failed status in context for logging
			failedAuthInfo := logger.AuthorizedInfo{
				TenantID: "unknown",
				MemberID: "unknown",
				Role:     "failed",
			}
			ctx := logger.SetAuthorizedInfoContext(r.Context(), failedAuthInfo)
			r = r.WithContext(ctx)

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Set authorized information to context
		ctx := logger.SetAuthorizedInfoContext(r.Context(), *authInfo)

		// Update request context
		r = r.WithContext(ctx)

		// Update ResponseWriterWrapper context if it's a wrapped writer
		if wrappedWriter, ok := w.(*logger.ResponseWriterWrapper); ok {
			wrappedWriter.UpdateContext(ctx)
		}

		// Pass the updated request to the next handler
		next.ServeHTTP(w, r)
	}
}
