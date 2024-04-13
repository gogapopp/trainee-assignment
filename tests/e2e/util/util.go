package util

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
