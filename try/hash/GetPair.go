package try

func GetPair(numbers []int, target int) (int, int) {
	cache := make(map[int]bool)
	for _, num := range numbers {
			val := target - num
			if cache[val] {
					return val, num
			}
			cache[num] = true
	}
	return 0, 0
}

func GetPairHalfSum(numbers []int) (int, int) {
sumNumbers := 0
	for _, num := range numbers {
			sumNumbers += num
	}
	halfSum, remainder := sumNumbers/2, sumNumbers%2
	if remainder != 0 {
			return 0, 0
	}

	cache := make(map[int]bool)
	for _, num := range numbers {
			cache[num] = true
			val := halfSum - num
			if cache[val] {
					return val, num
			}
	}
	return 0, 0
}