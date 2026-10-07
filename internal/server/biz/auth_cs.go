package biz

import (
	"NetGap/internal/server/data"
	"NetGap/internal/server/data/model"
)

// 客户端-服务器 身份校验

func VerityClient(clientId string, token string) bool {
	db := data.GetServerDb()
	var cnt int64
	db.Where("client_id = ? and token = ?", clientId, token).Find(&model.Client{}).Count(&cnt)
	if cnt == 0 {
		return false
	}
	return true
}
