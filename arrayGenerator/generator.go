package arraygenerator

import (
	"math/rand"
	"time"
)

func GenerateArrayWithRandomNumbers(array_length int, max_number int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []int

	for i := 0; i < array_length; i++ {
		result = append(result, rand.Intn(max_number))
	}

	return result
}
