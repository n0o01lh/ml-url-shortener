package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GetRandomString() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 7

	rand.Seed(uint64(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortKey)
}
