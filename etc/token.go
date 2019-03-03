//usr/bin/env go run $0 "$@" ; exit

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CHARACTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SIZE = 24
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateKey() string {
	bytes := make([]byte, SIZE)
	for index := 0; index < SIZE; index++ {
		bytes[index] = CHARACTERS[rand.Intn(len(CHARACTERS))]
	}
	return string(bytes)
}

func main() {
	fmt.Println(GenerateKey())
}
