package sdk

import (
	"NetGap/internal/session"
	"errors"
	"sync"
)

// 服务端 session 会话管理

type SessionManager struct {
	mu         sync.RWMutex
	sessionMap map[string]*session.Session // session 键值对 Key: ClientId Val: Session
}

var (
	ErrSessionNotFound = errors.New("session not found")
)

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: make(map[string]*session.Session),
	}
}

// SetSession 将新的 Client 连接放进 SM 中
func (m *SessionManager) SetSession(clientId string, session *session.Session) {
	m.mu.Lock()
	m.sessionMap[clientId] = session
	m.mu.Unlock()
}

// GetSession 将 Client Session 拉出
func (m *SessionManager) GetSession(clientId string) (*session.Session, error) {
	m.mu.RLock()
	cs, ok := m.sessionMap[clientId]
	m.mu.RUnlock()
	if !ok {
		return nil, ErrSessionNotFound
	}
	return cs, nil
}
