package try

func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := len(numbers) / 2
	left := numbers[:center]
	right := numbers[center:]

	left = MergeSort(left)
	right = MergeSort(right)

	i, j, k := 0, 0, 0
	temp := make([]int, len(numbers))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			temp[k] = left[i]
			i += 1
		} else {
			temp[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		temp[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		temp[k] = right[j]
		j += 1
		k += 1
	}

	copy(numbers, temp)
	return numbers
}
