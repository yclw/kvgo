package kvfile

import (
	"context"
	"time"

	"github.com/yclw/kvgo"
)

type FileOptions struct {
	Size        int64
	ContentType string
}

func WithSize(size int64) kvgo.Option {
	return kvgo.WithImplFn(func(opts *FileOptions) {
		opts.Size = size
	})
}

func WithContentType(contentType string) kvgo.Option {
	return kvgo.WithImplFn(func(opts *FileOptions) {
		opts.ContentType = contentType
	})
}

func WithContext(ctx context.Context) kvgo.Option {
	return kvgo.WithContext(ctx)
}

func WithTTL(ttl time.Duration) kvgo.Option {
	return kvgo.WithTTL(ttl)
}

func GetOptions(opts ...kvgo.Option) *kvgo.Options {
	return kvgo.GetCommonOptions(&kvgo.Options{Context: context.Background(), TTL: 0}, opts...)
}

func GetContext(opts ...kvgo.Option) context.Context {
	return kvgo.GetCommonOptions(&kvgo.Options{Context: context.Background()}, opts...).Context
}

func GetFileOptions(opts ...kvgo.Option) *FileOptions {
	return kvgo.GetImplSpecificOptions(&FileOptions{Size: -1}, opts...)
}
