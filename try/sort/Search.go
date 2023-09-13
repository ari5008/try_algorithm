package try

func LineaSearch(numbers []int, value int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == value {
			return i
		}
	}
	return -1
}

func BinarySearch(numbers []int, value int) int {
	left, right := 0, len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2 
		if numbers[mid] == value {
			return mid
		} else if numbers[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
