package sdk

var Version = "0.0.2"

type CryptoType int

const (
	CryptoNone CryptoType = iota
	CryptoAES128
	CryptoAES192
)
