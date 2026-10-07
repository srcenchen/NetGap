package handler

import (
	"NetGap/internal/server/session"

	"github.com/gin-gonic/gin"
)

type Router struct {
	sessionManager *session.Manager
}

func NewRouter(sm *session.Manager) *Router {
	return &Router{
		sessionManager: sm,
	}
}

func (r *Router) AuthNeededRouter(router *gin.RouterGroup) {
	clientG := router.Group("/client")
	r.clientRouter(clientG)
}
