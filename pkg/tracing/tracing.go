package tracing

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(serviceName string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().
		Timestamp().
		Str("service", serviceName).
		Logger().
		Hook(TracingHook{})
	log.Logger = logger
}

func GenerateTraceID() string {
	return uuid.New().String()
}

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceID := GetTraceIDFromContext(ctx)
	if traceID != "" {
		e.Str("request_id", traceID)
	}
}

func GetTraceIDFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value("request_id").(string)
	if !ok {
		traceID = GenerateTraceID()
	}
	return traceID
}
