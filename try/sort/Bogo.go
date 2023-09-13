package try

import (
	"math/rand"
	"time"
)

func In_order(numbers []int) bool {
	for i := 0; i < len(numbers) - 1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func Bogo_sort(numbers []int) []int {

	rand.Seed(time.Now().UnixNano())

	for !In_order(numbers) {
	rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
	}
	return numbers
}

