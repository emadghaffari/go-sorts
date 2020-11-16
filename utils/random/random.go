package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Slice random Slice of numbers
func Slice(num, min, max int) []int {
	if min > max {
		fmt.Println("Error min and max")
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var items []int
	for i := 0; i < num; i++ {
		random := (rand.Intn((max - min)) + min)
		items = append(items, random)
	}
	return items
}
