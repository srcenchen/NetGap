package handler

type NewTunnelReq struct {
	ClientId   string `json:"client_id"`
	TunnelName string `json:"tunnel_name"`
	ServerPort string `json:"server_port"`
	ClientAddr string `json:"client_addr"`
}

type NewTunnelResp struct {
	TunnelId string `json:"tunnel_id"`
}
