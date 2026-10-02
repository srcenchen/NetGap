package protocol

import (
	"encoding/binary"
	"io"
)

const maxPacketSize = 128

// readFrame 读取帧
func readFrame(r io.Reader) ([]byte, error) {
	var hdr [4]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	dataLen := int(binary.BigEndian.Uint32(hdr[:]))
	if dataLen == 0 || dataLen > maxPacketSize {
		return nil, ErrPacketTooLarge
	}
	data := make([]byte, dataLen)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}
	return data, nil
}

// writeFrame 写入帧
func writeFrame(w io.Writer, v []byte) error {
	if len(v) > maxPacketSize {
		return ErrPacketTooLarge
	}
	var hdr [4]byte
	binary.BigEndian.PutUint32(hdr[:], uint32(len(v)))
	_, err := w.Write(hdr[:])
	if err != nil {
		return err
	}
	_, err = w.Write(v)
	return err
}
