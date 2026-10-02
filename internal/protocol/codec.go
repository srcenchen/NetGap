package protocol

import (
	"io"
)

// NetGap 编解码器
// 我们规定 NetGap 协议的前4字节为后续data长度
// [4/header][data]

type Codec interface {
	Decode(r io.Reader, v any) error
	Encode(w io.Writer, v any) error
}
