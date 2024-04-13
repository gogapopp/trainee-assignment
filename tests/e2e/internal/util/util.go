package util

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func GenerateRandomNumber() int {
	return r.Intn(100) + 1
}
