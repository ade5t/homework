package arrayShift

func Solution(inputArray []int, offset int) []int {
	if offset == 0 || offset == len(inputArray) {
		return inputArray
	}

	if offset > len(inputArray) {
		offset = (offset % len(inputArray))
	}

	return append(inputArray[len(inputArray)-offset:], inputArray[:len(inputArray)-offset]...)
}
