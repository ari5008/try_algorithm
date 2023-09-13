package try

import "fmt"

func Hanoi(disk int, src string, dest string, center string) {
	if disk < 1 {
		return
	}
	Hanoi(disk-1, src, center, dest)
	fmt.Printf("move %v from %v to %v", disk, src, dest)
	fmt.Println()
	Hanoi(disk-1, center, dest, src)
}
