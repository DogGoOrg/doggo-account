package middleware

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func UnaryCallLogger(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	requestID := uuid.New().String()

	resp, err := handler(ctx, req)

	end := time.Now()
	latency := end.Sub(start)

	attributes := []slog.Attr{
		slog.String("path", info.FullMethod),
		slog.String("ID", requestID),
		slog.Duration("latency", latency),
		slog.Time("time", end),
	}

	log.Default().Println(attributes)

	return resp, err
}
