package main

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	if length < 1 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteRune(letters[rand.Intn(len(letters)-1)])
	}
	return builder.String()
}
