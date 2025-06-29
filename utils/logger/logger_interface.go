package logger

import (
	"context"
)

// Logger: Interface for application logging
type Logger interface {
	// Basic logging methods
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)

	// Request logging method
	SetLogContext(ctx context.Context, statusCode int, httpInfo HTTPRequestInfo, systemInfo SystemInfo, authInfo AuthorizedInfo)
}

var _ Logger = (*AppLogger)(nil)
