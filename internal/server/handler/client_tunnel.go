package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"uuid"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (r *Router) handleNewTunnel(c *gin.Context) {
	var tunnelReq NewTunnelReq
	err := c.ShouldBindJSON(&tunnelReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrParmsNotAble)
		return
	}
	sess, err := r.sessionManager.GetSession(tunnelReq.ClientId)
	if err != nil {
		logrus.Error(tunnelReq, err.Error())
		c.JSON(http.StatusBadRequest, ErrClientNotFound)
		return
	}
	port := (rand.Intn(99) + 1) * 3000
	sess.RunNewRelay(uuid.NewV4().String(), strconv.Itoa(port), tunnelReq.ClientAddr)
}
