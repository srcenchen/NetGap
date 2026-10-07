package handler

import (
	"NetGap/internal/server/biz/http_biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpBiz *http_biz.HttpBiz

// clientRouter /client/ 组
func (r *Router) clientRouter(router *gin.RouterGroup) {
	httpBiz = http_biz.NewHttpBiz()
	router.GET("/list", r.handleList) // /client/list
	router.POST("/tunnel/new", r.handleNewTunnel)
}

func (r *Router) handleList(c *gin.Context) {
	cl := httpBiz.ClientBiz.GetClientList()
	resp := listResp{
		Code: http.StatusOK,
		Data: make([]client, 0),
	}
	for _, cl := range cl {
		status := false
		_, err := r.sessionManager.GetSession(cl.ClientID)
		if err == nil {
			status = true
		}
		resp.Data = append(resp.Data, client{
			ClientId:        cl.ClientID,
			LastConnectTime: cl.LastConnected.Unix(),
			Status:          status,
		})
	}
	c.JSON(http.StatusOK, resp)
}
