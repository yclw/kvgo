package redis

import (
	"context"

	"github.com/yclw/kvgo"
)

type ScanOptions struct {
	Cursor         uint64
	Count          int64
	IterateAll     bool
	nextCursorSink *uint64
}

func GetContext(opts ...kvgo.Option) context.Context {
	option := kvgo.GetCommonOptions(&kvgo.Options{Context: context.Background()}, opts...)
	return option.Context
}

func GetOptions(opts ...kvgo.Option) *kvgo.Options {
	return kvgo.GetCommonOptions(&kvgo.Options{Context: context.Background(), TTL: 0}, opts...)
}

func WithScanCursor(cursor uint64) kvgo.Option {
	return kvgo.WithImplFn(func(opts *ScanOptions) {
		opts.Cursor = cursor
	})
}

func WithScanCount(count int64) kvgo.Option {
	return kvgo.WithImplFn(func(opts *ScanOptions) {
		opts.Count = count
	})
}

func WithScanIterateAll(iterateAll bool) kvgo.Option {
	return kvgo.WithImplFn(func(opts *ScanOptions) {
		opts.IterateAll = iterateAll
	})
}

func WithScanCursorSink(sink *uint64) kvgo.Option {
	return kvgo.WithImplFn(func(opts *ScanOptions) {
		opts.nextCursorSink = sink
	})
}

func GetScanOptions(opts ...kvgo.Option) *ScanOptions {
	return kvgo.GetImplSpecificOptions(&ScanOptions{IterateAll: true}, opts...)
}
