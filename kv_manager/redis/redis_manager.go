package redis

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/yclw/kvgo"
)

type RedisManager struct {
	prefix string
	client *redis.Client
}

func NewRedisManager(prefix string, client *redis.Client) kvgo.KvManager {
	return &RedisManager{prefix: prefix, client: client}
}

func (m *RedisManager) ListKeys(namespace string, opts ...kvgo.Option) ([]string, error) {
	ctx := GetContext(opts...)
	scanOptions := GetScanOptions(opts...)
	cursor := scanOptions.Cursor
	pattern := kvgo.FormatKey(m.prefix, namespace, "*")
	trimPrefix := kvgo.FormatKey(m.prefix, namespace, "")
	var (
		keys []string
	)
	for {
		res, nextCursor, err := m.scanKeys(ctx, cursor, pattern, scanOptions.Count)
		if err != nil {
			return nil, err
		}
		for i := range res {
			keys = append(keys, strings.TrimPrefix(res[i], trimPrefix))
		}
		cursor = nextCursor
		if !scanOptions.IterateAll || cursor == 0 {
			break
		}
	}
	if scanOptions.nextCursorSink != nil {
		*scanOptions.nextCursorSink = cursor
	}
	return keys, nil
}

func (m *RedisManager) Type(namespace string, key string, opts ...kvgo.Option) (kvgo.Type, error) {
	ctx := GetContext(opts...)
	key_ := kvgo.FormatKey(m.prefix, namespace, key)
	typ, err := m.client.Get(ctx, key_).Result()
	return kvgo.Type(typ), err
}

func (m *RedisManager) Exists(namespace string, key string, opts ...kvgo.Option) (bool, error) {
	ctx := GetContext(opts...)
	key_ := kvgo.FormatKey(m.prefix, namespace, key)
	exists, err := m.client.Exists(ctx, key_).Result()
	return exists > 0, err
}

func (m *RedisManager) UpsertType(namespace string, key string, valType kvgo.Type, opts ...kvgo.Option) error {
	option := GetOptions(opts...)
	ctx := option.Context
	key_ := kvgo.FormatKey(m.prefix, namespace, key)
	return m.client.Set(ctx, key_, string(valType), option.TTL).Err()
}

func (m *RedisManager) Delete(namespace string, key string, opts ...kvgo.Option) error {
	ctx := GetContext(opts...)
	key_ := kvgo.FormatKey(m.prefix, namespace, key)
	return m.client.Del(ctx, key_).Err()
}

func (m *RedisManager) scanKeys(ctx context.Context, cursor uint64, pattern string, count int64) ([]string, uint64, error) {
	return m.client.Scan(ctx, cursor, pattern, count).Result()
}
