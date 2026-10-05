package session

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

// 服务端 session 会话管理

type Manager struct {
	mu         sync.RWMutex
	sessionMap map[string]*Session // session 键值对 Key: ClientId Val: Session
}

var (
	ErrSessionNotFound = errors.New("session not found")
)

func NewSessionManager() *Manager {
	return &Manager{
		sessionMap: make(map[string]*Session),
	}
}

// SetSession 将新的 Client 连接放进 SM 中
func (m *Manager) SetSession(clientId string, sess *Session) {
	m.mu.Lock()
	m.sessionMap[clientId] = sess
	m.mu.Unlock()
	// session 生命周期管理
	m.alive(sess)
}

// GetSession 将 Client Session 拉出
func (m *Manager) GetSession(clientId string) (*Session, error) {
	m.mu.RLock()
	cs, ok := m.sessionMap[clientId]
	m.mu.RUnlock()
	if !ok {
		return nil, ErrSessionNotFound
	}
	return cs, nil
}

// remove 清理 Session 会话
// remove 并不应该被其他地方调用，只能够在session离线时自动清除
func (m *Manager) remove(clientId string) {
	m.mu.Lock()
	delete(m.sessionMap, clientId)
	m.mu.Unlock()
}

func (m *Manager) alive(s *Session) {
	go func() {
		<-s.CloseChan()
		logrus.Infof("[TunnelManager]: 客户端[%s] 断开连接", s.GetClientId())
		m.remove(s.GetClientId())
	}()
}
