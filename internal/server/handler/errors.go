package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrParmsNotAble = gin.H{
		"code":    http.StatusBadRequest,
		"message": "Bad Request. 不被允许的参数列表",
	}
	ErrClientNotFound = gin.H{
		"code":    http.StatusBadRequest,
		"message": "目标客户端目前不在线。",
	}
)
