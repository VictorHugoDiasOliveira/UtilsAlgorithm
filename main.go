package main

import (
	"fmt"

	arraygenerator "github.com/VictorHugoDiasOliveira/UtilsAlgorithm/arrayGenerator"
)

func main() {
	array := arraygenerator.GenerateArrayWithRandomNumbers(20, 100)
	fmt.Println(array)
}
