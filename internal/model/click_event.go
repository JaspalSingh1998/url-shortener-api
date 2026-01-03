package model

import "time"

type ClickEvent struct {
	ID        int64
	LinkID    int64
	ShortCode string

	IPAddress string
	UserAgent string
	Referrer  string

	CreatedAt time.Time
}
