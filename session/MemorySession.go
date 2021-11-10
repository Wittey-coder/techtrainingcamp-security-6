package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	lock      sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
	return nil
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err := errors.New("此键无对应的值！")
		return nil, err
	}
	return value, nil
}

func (m *MemorySession) Delete(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, key)
	return nil
}

func (m *MemorySession) Save() error {
	return nil
}
