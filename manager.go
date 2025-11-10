package kvgo

import "sync"

type Manager struct {
	mu       sync.RWMutex
	handlers map[Type]Handler
	KvManager
}

func NewManager(kvManager KvManager) *Manager {
	return &Manager{handlers: make(map[Type]Handler), KvManager: kvManager}
}

func (m *Manager) SetKvManager(kvManager KvManager) {
	m.KvManager = kvManager
}

func (m *Manager) RegisterHandler(valType Type, handler Handler) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.handlers[valType] = handler
}

func (m *Manager) GetHandler(valType Type) (Handler, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	handler, ok := m.handlers[valType]
	return handler, ok
}

type KvManager interface {
	ListKeys(namespace string, opts ...Option) ([]string, error)
	Type(namespace string, key string, opts ...Option) (Type, error)
	Exists(namespace string, key string, opts ...Option) (bool, error)
	UpsertType(namespace string, key string, valType Type, opts ...Option) error
	Delete(namespace string, key string, opts ...Option) error
}
