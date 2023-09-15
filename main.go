package main

import (
	"backend/try/sort"
	"fmt"
)

func main() {
	numbers := []int{5, 4, 1, 8, 7, 3, 2, 9}
	fmt.Println(try.MergeSort(numbers))
}
