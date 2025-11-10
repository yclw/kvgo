package kvgo

import (
	"context"
	"time"
)

type Options struct {
	Context context.Context
	TTL     time.Duration
}

type Option struct {
	apply  func(*Options)
	implFn any
}

func WithContext(ctx context.Context) Option {
	return Option{
		apply: func(opts *Options) {
			opts.Context = ctx
		},
	}
}

func WithTTL(ttl time.Duration) Option {
	return Option{
		apply: func(opts *Options) {
			opts.TTL = ttl
		},
	}
}

func WithImplFn[T any](optFn func(*T)) Option {
	return Option{
		implFn: optFn,
	}
}

func GetCommonOptions(base *Options, opts ...Option) *Options {
	if base == nil {
		base = &Options{}
	}

	for i := range opts {
		opt := opts[i]
		if opt.apply != nil {
			opt.apply(base)
		}
	}

	return base
}

func GetImplSpecificOptions[T any](base *T, opts ...Option) *T {
	if base == nil {
		base = new(T)
	}
	for i := range opts {
		opt := opts[i]
		if opt.implFn != nil {
			optFn, ok := opt.implFn.(func(*T))
			if ok {
				optFn(base)
			}
		}
	}
	return base
}
