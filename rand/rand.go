package rand

import (
	"math"
	"math/rand"
)

// Sample is a (not 1:1) analog to `np.random.choice` with pseudo-random number sampling.
// https://www.wikiwand.com/en/Pseudo-random_number_sampling
func Sample(p []float64, accuracy float64) int {
	intervals := []int{}

	for i, probability := range p {
		intervalSize := int(math.Round(accuracy * probability))
		for j := 0; j < intervalSize; j++ {
			intervals = append(intervals, i)
		}
	}

	return intervals[rand.Intn(len(intervals))]
}
