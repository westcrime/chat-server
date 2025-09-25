package model

import (
	"time"
)

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
