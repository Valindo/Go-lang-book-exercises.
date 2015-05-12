package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(500))

	s1 := rand.NewSource(21)
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(500))
}
