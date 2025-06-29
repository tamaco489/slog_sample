package logger

import (
	"context"
	"log/slog"
	"net/http"
)

// DebugContext: output debug log
func (l *AppLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// InfoContext: output info log
func (l *AppLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// WarnContext: output warn log
func (l *AppLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// ErrorContext: output error log
func (l *AppLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// SetLogContext: Set log context with appropriate level based on status code
func (l *AppLogger) SetLogContext(
	ctx context.Context,
	statusCode int,
	httpRequestInfo HTTPRequestInfo,
	systemInfo SystemInfo,
	authorizedInfo AuthorizedInfo,
) {
	// Create logger with context
	lw := l.Logger.With(
		slog.Any(LogFieldKeyHTTPRequest.String(), httpRequestInfo),
		slog.Any(LogFieldKeySystem.String(), systemInfo),
		slog.Any(LogFieldKeyAuthorized.String(), authorizedInfo),
	)

	// Determine log level and log with appropriate method
	switch {
	// status: 5xx, level: error
	case statusCode >= http.StatusInternalServerError:
		lw.ErrorContext(ctx, "request failed")

	// status: 4xx, level: info (client errors are not warnings)
	case statusCode >= http.StatusBadRequest:
		lw.InfoContext(ctx, "request completed")

	// status: 2xx, level: info
	default:
		lw.InfoContext(ctx, "request completed")
	}
}
