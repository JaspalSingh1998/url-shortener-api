package service

import (
	"math/rand"
	"time"
)

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)

	for i := range b {
		b[i] = base62[rand.Intn(len(base62))]
	}

	return string(b)
}
