package try

func QuickSort(numbers []int) []int {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func quickSort(numbers []int, low int, high int) {
	if low < high {
		partitionIndex := partition(numbers, low, high)
		quickSort(numbers, low, partitionIndex-1)
		quickSort(numbers, partitionIndex+1, high)
	}
}

func partition(numbers []int, low int, high int) int {
	pivot := numbers[high]
	i := low - 1

	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}
