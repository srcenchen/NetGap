package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientID string
	Token    string
}
