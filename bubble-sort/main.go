package main

import "fmt"

var (
	items []int
)

func main() {
	items = append(items, 8)
	items = append(items, 6)
	items = append(items, 1)
	items = append(items, 10)
	items = append(items, 4)
	items = append(items, 3)
	items = append(items, 2)
	items = append(items, 5)
	items = append(items, 9)
	items = append(items, 12)
	items = append(items, 2)
	items = append(items, 7)
	fmt.Println(bubble(items))
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
