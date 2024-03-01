package data

import "math/rand/v2"

const (
	defaultSize      = 1000
	numArrayElements = 10000
	maxRandomValue   = 100000
)

func getRandomUints(size int, maxValue uint) []uint {
	mapped := make(map[uint]bool)
	random := make([]uint, 0)

	for len(mapped) < size {
		inserted := rand.UintN(maxValue)
		if _, exists := mapped[inserted]; !exists {
			mapped[inserted] = true
			random = append(random, inserted)
		}
	}

	return random
}
