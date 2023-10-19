package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(randomBool())
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
