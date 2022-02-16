package main

import (
	"fmt"
	"homework/arrayShift"
	"homework/searchMissingElement"
	"homework/sequenceCheck"
	"homework/weirdEntry"
)

var testArray = []int{1, 7, 3, 2, 6, 5}

func main() {
	fmt.Println(arrayShift.Solution(testArray, 4))
	fmt.Println(sequenceCheck.Solution(testArray))
	fmt.Println(weirdEntry.Solution(testArray))
	fmt.Println(searchMissingElement.Solution(testArray))
}
