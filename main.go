package main

import (
	"fmt"

	slicegenerator "github.com/VictorHugoDiasOliveira/UtilsAlgorithm/slicegenerator"
)

func main() {
	array := slicegenerator.GenerateSliceWithRandomNumbers(20, 100)
	fmt.Println(array)
}
