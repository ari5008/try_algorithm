package try

func ProcessPairs(pairs [][]int) <-chan []int {
	result := make(chan []int)

	go func() {
			cache := make(map[int]int)

			for _, pair := range pairs {
					first, second := pair[0], pair[1]
					value, exists := cache[second]

					if !exists {
							cache[first] = second
					} else if value == first {
							result <- pair
					}
			}

			close(result)
	}()

	return result
}