package sdk

var Version = "0.0.1"

type CryptoType int

const (
	CryptoNone CryptoType = iota
	CryptoAES128
	CryptoAES192
)
