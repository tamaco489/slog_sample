package logging

import (
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// RequestMiddleware: Manage request start and end with single log output
func RequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record request start time
		startTime := time.Now()

		// Initialize request context with anonymous status
		r = initializeRequestContext(r)

		// Create ResponseWriterWrapper and set initial context
		wrappedWriter := logger.NewResponseWriterWrapper(w)
		wrappedWriter.UpdateContext(r.Context())

		// defer for request end logging (single log output)
		defer func() {
			// Get final context from ResponseWriterWrapper (may be updated by authorization middleware)
			finCtx := wrappedWriter.GetContext()
			requestID, _ := logger.GetRequestIDContext(finCtx)
			systemInfo, _ := logger.GetSystemInfoContext(finCtx)
			authorizedInfo, _ := logger.GetAuthorizedInfoContext(finCtx)

			// Calculate processing time
			latency := time.Since(startTime)

			// Create HTTP information
			httpInfo := logger.NewHTTPRequestInfo(
				r,
				wrappedWriter.GetStatusCode(),
				latency.String(),
				requestID,
			)

			// Single log output with final state (including authorization result)
			env := configuration.GetEnvironment()
			log := logger.New(env)
			log.SetLogContext(
				finCtx,
				wrappedWriter.GetStatusCode(),
				httpInfo,
				systemInfo,
				authorizedInfo,
			)
		}()

		// Execute next handler
		next.ServeHTTP(wrappedWriter, r)
	}
}

// initializeRequestContext: Initialize context with anonymous status
func initializeRequestContext(r *http.Request) *http.Request {
	// Generate request ID and set it in context
	ctx := logger.SetRequestIDContext(r.Context())

	// Initialize system information
	env := configuration.GetEnvironment()
	systemInfo := logger.NewSystemInfo(env)
	ctx = logger.SetSystemInfoContext(ctx, systemInfo)

	// Set initial authorized information (anonymous)
	authInfo := logger.NewInitialAuthorizedInfo()
	ctx = logger.SetAuthorizedInfoContext(ctx, authInfo)

	// Update request with updated context
	return r.WithContext(ctx)
}
