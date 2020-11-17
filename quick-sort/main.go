package main

import (
	"fmt"
	"math/rand"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func main() {
	items := random.Slice(20, 1, 15)
	fmt.Println("Before: ", items)

	result := quick(items)

	fmt.Println("After : ", result)
}

func quick(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := (rand.Int() % len(a))

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quick(a[:left])
	quick(a[left+1:])

	return a
}
