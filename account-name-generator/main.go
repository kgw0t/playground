package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			generate()
		}
		fmt.Println()
	}
}

func generate() {
	const consonants = "bcdfghjklmnpqrstvwxyz"
	const vowels = "aiueo"

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 7; i++ {
		target := consonants
		if i%2 == 1 {
			target = vowels
		}

		point := rand.Intn(len(target))
		fmt.Print(target[point : point+1])
	}
	fmt.Print(" ")
}
