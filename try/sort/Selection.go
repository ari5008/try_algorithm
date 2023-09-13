package try

func SelectionSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		minIdx := i
		for j := i+1; j < lenNumbers; j++ {
			if numbers[minIdx] > numbers[j] {
				minIdx = j
			}
		}
		numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
	}

	return numbers
}