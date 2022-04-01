package searchMissingElement

func Solution(inputArray []int) int {
	sum := 0

	for _, value := range inputArray {
		sum += value
	}

	return (len(inputArray)+1)*(len(inputArray)+2)/2 - sum
}
