package http_biz

import (
	"NetGap/internal/server/data/model"

	"gorm.io/gorm"
)

type clientBiz struct {
	db *gorm.DB
}

func (c *clientBiz) GetClientList() (r []model.Client) {
	c.db.Find(&r)
	return
}
