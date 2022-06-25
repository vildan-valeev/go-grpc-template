package requestid

import (
	"context"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
)

// headerRequestID is metadata key name for request ID.
const headerRequestID = "x-request-id"

func Get(ctx context.Context) string {
	val := metadata.ExtractIncoming(ctx).Get(headerRequestID)

	if val == "" {
		val = uuid.NewString()
	}

	return val
}
