package http_biz

import "NetGap/internal/server/data"

type HttpBiz struct {
	ClientBiz *clientBiz
}

func NewHttpBiz() *HttpBiz {
	return &HttpBiz{
		ClientBiz: &clientBiz{db: data.GetServerDb()},
	}
}
