package kvgo

type Handler interface{}

type Setter interface {
	Set(namespace string, key string, val Value, opts ...Option) error
}

type Getter interface {
	Get(namespace string, key string, opts ...Option) (Value, error)
}

type Deleter interface {
	Delete(namespace string, key string, opts ...Option) error
}
