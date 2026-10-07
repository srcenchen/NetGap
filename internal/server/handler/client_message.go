package handler

type client struct {
	ClientId        string `json:"client_id"`
	LastConnectTime int64  `json:"last_connect_time"`
	Status          bool   `json:"status"`
}

type listReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"` // default 20
}

type listResp struct {
	Code int      `json:"code"`
	Data []client `json:"list"`
}
