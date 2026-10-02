package protocol

import (
	"encoding/json"
	"io"
)

type JSONCodec struct{}

func (J JSONCodec) Decode(r io.Reader, v any) error {
	data, err := readFrame(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (J JSONCodec) Encode(w io.Writer, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return writeFrame(w, data)
}

func NewJSONCodec() Codec {
	return JSONCodec{}
}
