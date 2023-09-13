package try

func InsertionSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
		}
		numbers[j+1] = temp
	}
	return numbers
}

