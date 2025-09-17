package model

import (
	"time"
)

type Role int32

type CreateChat struct {
	Name      string
	Usernames []string
}

type Message struct {
	From      string
	Text      string
	ChatId    int64
	Timestamp time.Time
}
