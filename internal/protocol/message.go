package protocol

type HandshakeReq struct {
	ClientID string
	Token    string
}

type HandshakeResp struct {
	Code int    // 0 为失败 1 为成功
	Msg  string // 消息通知
}
