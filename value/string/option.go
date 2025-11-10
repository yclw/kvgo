package string

import (
	"context"
	"time"

	"github.com/yclw/kvgo"
)

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
