package interfaces

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, name string, reader io.Reader, size int64) error
}
