package config

import (
	"context"
	"log/slog"
	"os"
)

const (
	TraceIdContextKey = "traceId"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(TraceIdContextKey).(string); ok {
		r.AddAttrs(slog.String(TraceIdContextKey, requestID))
	}

	return h.Handler.Handle(ctx, r)
}

func ConfigureLogger(env string) {
	defaultAttrs := []slog.Attr{
		slog.String("service", "gororoba"),
		slog.String("environment", env),
		slog.String("version", "1.0.0"),
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	baseHandler := slog.NewJSONHandler(os.Stderr, handlerOptions).WithGroup("metadata").WithAttrs(defaultAttrs)
	customHandler := &ContextHandler{baseHandler}
	logger := slog.New(customHandler)

	slog.SetDefault(logger)
}
