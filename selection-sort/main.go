package main

import (
	"fmt"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func main() {
	items := random.Slice(20, 1, 15)
	fmt.Println("Before: ", items)
	result := selection(items)
	fmt.Println("After : ", result)
}

func selection(items []int) []int {
	n := len(items)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i; j < n; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
	return items
}
