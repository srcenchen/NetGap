package sdk

import (
	"NetGap/internal/session"
	"sync"
)

// 服务端 session 会话管理

type SessionManager struct {
	mu         sync.RWMutex
	sessionMap map[string]session.Session // session 键值对 Key: ClientId Val: Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}
