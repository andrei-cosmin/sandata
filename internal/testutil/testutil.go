package testutil

import "math/rand/v2"

func RandomUInts(size int, maxValue uint) []uint {
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
