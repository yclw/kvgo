package kvstring

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yclw/kvgo"
)

var _ kvgo.Handler = (*StringHandler)(nil)
var _ kvgo.Setter = (*StringHandler)(nil)
var _ kvgo.Getter = (*StringHandler)(nil)
var _ kvgo.Deleter = (*StringHandler)(nil)

type StringHandler struct {
	client *redis.Client
}

func NewStringHandler(client *redis.Client) kvgo.Handler {
	return &StringHandler{client: client}
}

func (h *StringHandler) Set(namespace string, key string, val kvgo.Value, opts ...kvgo.Option) error {
	stringValue, ok := val.(*StringValue)
	if !ok {
		return fmt.Errorf("value is not a string")
	}
	option := GetOptions(opts...)
	ctx := option.Context
	key_ := kvgo.FormatKey(namespace, key)
	return h.client.Set(ctx, key_, stringValue.Value, option.TTL).Err()
}

func (h *StringHandler) Get(namespace string, key string, opts ...kvgo.Option) (kvgo.Value, error) {
	ctx := GetContext(opts...)
	key_ := kvgo.FormatKey(namespace, key)
	value, err := h.client.Get(ctx, key_).Result()
	if err != nil {
		return nil, err
	}
	return &StringValue{Value: value}, nil
}

func (h *StringHandler) Delete(namespace string, key string, opts ...kvgo.Option) error {
	ctx := GetContext(opts...)
	key_ := kvgo.FormatKey(namespace, key)
	return h.client.Del(ctx, key_).Err()
}
