package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionManager struct {
	lock       sync.RWMutex
	sessionMap map[string]Session
}

func NewMemorySessionManager() SessionManager {
	m := &MemorySessionManager{
		sessionMap: make(map[string]Session, 1024),
	}
	return m
}

func (m *MemorySessionManager) Init(addr string, options ...string) error {
	return nil
}

func (m *MemorySessionManager) CreateSession() (session Session, id string, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	sessionId := uuid.NewV4().String()
	memorySession := NewMemorySession(sessionId)
	m.sessionMap[sessionId] = memorySession
	return memorySession, sessionId, nil
}

func (m *MemorySessionManager) Get(sessionId string) (session *Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	get, ok := m.sessionMap[sessionId]
	if !ok {
		return nil, errors.New("不存在此Session")
	}
	return &get, nil
}

func (m *MemorySessionManager) Save() error {
	return nil
}

func (m *MemorySessionManager) Delete(sessionId string) error {
	delete(m.sessionMap, sessionId)
	return nil
}
