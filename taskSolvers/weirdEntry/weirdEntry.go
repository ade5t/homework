package weirdEntry

// Решение задачи "Чудные вхождения в массив"
func Solution(inputArray []int) int {
	numberWithoutPair := 0

	for _, value := range inputArray {
		numberWithoutPair = numberWithoutPair ^ value
	}

	return numberWithoutPair
}
