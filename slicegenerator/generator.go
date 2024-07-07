package slicegenerator

import (
	"math/rand"
	"time"
)

func GenerateSliceWithRandomNumbers(slice_length int, max_number int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []int

	for i := 0; i < slice_length; i++ {
		result = append(result, rand.Intn(max_number))
	}

	return result
}
