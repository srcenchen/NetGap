package model

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClientID      string
	Token         string
	LastConnected time.Time
}
