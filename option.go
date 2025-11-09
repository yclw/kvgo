package kvgo

type Option struct {
	implFn any
}

func WithImplFn[T any](optFn func(*T)) Option {
	return Option{
		implFn: optFn,
	}
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
