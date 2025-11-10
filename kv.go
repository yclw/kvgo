package kvgo

import (
	"fmt"
)

var kv *KV

type KV struct {
	manager *Manager
}

func NewKV(manager *Manager) *KV {
	return &KV{manager: manager}
}

func Init(kvManager KvManager) {
	manager := NewManager(kvManager)
	kv = NewKV(manager)
}

func GetManager() *Manager {
	return kv.manager
}

func Set(namespace string, key string, val Value, opts ...Option) error {
	if kv == nil {
		return fmt.Errorf("kv not initialized")
	}
	valType := val.Type()
	handler, ok := kv.manager.GetHandler(valType)
	if !ok {
		return fmt.Errorf("handler for value type %s not registered", valType)
	}
	setter, ok := handler.(Setter)
	if !ok {
		return fmt.Errorf("handler for value type %s is not a setter", valType)
	}
	err := Delete(namespace, key, opts...)
	if err != nil {
		return err
	}
	err = kv.manager.UpsertType(namespace, key, valType, opts...)
	if err != nil {
		return err
	}
	return setter.Set(namespace, key, val, opts...)
}

func Get(namespace string, key string, opts ...Option) (Value, error) {
	if kv == nil {
		return nil, fmt.Errorf("kv not initialized")
	}
	exists, err := kv.manager.Exists(namespace, key, opts...)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("key %s not found", key)
	}
	valType, err := kv.manager.Type(namespace, key, opts...)
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
	return getter.Get(namespace, key, opts...)
}

func Delete(namespace string, key string, opts ...Option) error {
	if kv == nil {
		return fmt.Errorf("kv not initialized")
	}
	exists, err := kv.manager.Exists(namespace, key, opts...)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	valType, err := kv.manager.Type(namespace, key, opts...)
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
	err = kv.manager.Delete(namespace, key, opts...)
	if err != nil {
		return err
	}
	return deleter.Delete(namespace, key, opts...)
}

func ListKeys(namespace string, opts ...Option) ([]string, error) {
	if kv == nil {
		return nil, fmt.Errorf("kv not initialized")
	}
	return kv.manager.ListKeys(namespace, opts...)
}
