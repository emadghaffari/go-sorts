package random

import "math/rand"

// Slice random Slice of numbers
func Slice(num, min, max int) []int {
	items := make([]int, num)
	for i := 0; i < num; i++ {
		random := rand.Intn(max-min) + min
		items = append(items, random)
	}
	return items
}
