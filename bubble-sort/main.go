package main

import (
	"fmt"

	"github.com/emadghaffari/go-sorts/utils/random"
)

var (
	items []int
)

func main() {
	items := random.Slice(20, 1, 15)
	fmt.Println("Before: ", items)

	result := bubble(items)

	fmt.Println("After: ", result)
}

func bubble(items []int) []int {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(items)-1; i++ {
			for j := 0; j < len(items)-i-1; j++ {
				if items[j] > items[j+1] {
					items[j], items[j+1] = items[j+1], items[j]
					keepWorking = true
				}
			}
		}
	}
	return items
}
