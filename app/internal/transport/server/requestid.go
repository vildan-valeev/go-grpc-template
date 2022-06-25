package server

import (
	"context"
	"go-grpc-template/pkg/requestid"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const requestIDKey = 1

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent context.Context, key interface{}, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

// WithRequestID returns a copy of parent in which the RequestID value is set.
func WithRequestID(parent context.Context, id string) context.Context {
	return WithValue(parent, requestIDKey, id)
}

// RequestIDFrom returns the value of the RequestID key on the ctx.
func RequestIDFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

func ZerologCtxUnaryServerInterceptor(l zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(l.WithContext(ctx), req)
	}
}

func RequestIDCtxUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(WithRequestID(ctx, requestid.Get(ctx)), req)
	}
}

func InjectRequestIDCtxUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if id, ok := RequestIDFrom(ctx); ok {
			ctx = logging.InjectFields(ctx, logging.Fields{"grpc.request_id", id})
		}

		return handler(ctx, req)
	}
}
