package helpers

import "math/rand"

func Random(n int) int {
	value := rand.Intn(n)
	return value
}