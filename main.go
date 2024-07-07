package main

import (
	"fmt"

	slicegenerator "github.com/VictorHugoDiasOliveira/UtilsAlgorithm/sliceGenerator"
)

func main() {
	array := slicegenerator.GenerateSliceWithRandomNumbers(20, 100)
	fmt.Println(array)
}
