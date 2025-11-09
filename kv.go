package kvgo

import (
	"fmt"
)

var kv = NewKV()

type KV struct {
	manager *Manager
}

func NewKV() *KV {
	return &KV{manager: NewManager()}
}

func Set(namespace string, key string, val Value, opts ...Option) error {
	valType := val.Type()
	handler, ok := kv.manager.GetHandler(valType)
	if !ok {
		return fmt.Errorf("handler for value type %s not registered", valType)
	}
	setter, ok := handler.(Setter)
	if !ok {
		return fmt.Errorf("handler for value type %s is not a setter", valType)
	}
	err := kv.manager.kvManager.UpsertType(namespace, key, valType, opts...)
	if err != nil {
		return err
	}
	return setter.Set(namespace, key, val, opts...)
}

func Get(namespace string, key string, opts ...Option) (Value, error) {
	valType, err := kv.manager.kvManager.Type(namespace, key, opts...)
	if err != nil {
		return nil, err
	}
	handler, ok := kv.manager.GetHandler(valType)
	if !ok {
		return nil, fmt.Errorf("handler for value type %s not registered", valType)
	}
	getter, ok := handler.(Getter)
	if !ok {
		return nil, fmt.Errorf("handler for value type %s is not a getter", valType)
	}
	err = kv.manager.kvManager.UpsertType(namespace, key, valType, opts...)
	if err != nil {
		return nil, err
	}
	return getter.Get(namespace, key, opts...)
}

func Delete(namespace string, key string, opts ...Option) error {
	valType, err := kv.manager.kvManager.Type(namespace, key, opts...)
	if err != nil {
		return err
	}
	handler, ok := kv.manager.GetHandler(valType)
	if !ok {
		return fmt.Errorf("handler for value type %s not registered", valType)
	}
	deleter, ok := handler.(Deleter)
	if !ok {
		return fmt.Errorf("handler for value type %s is not a deleter", valType)
	}
	err = kv.manager.kvManager.Delete(namespace, key, opts...)
	if err != nil {
		return err
	}
	return deleter.Delete(namespace, key, opts...)
}
