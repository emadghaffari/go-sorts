package main

import (
	"fmt"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func main() {
	items := random.Slice(20, 1, 15)
	fmt.Println("Before: ", items)

	result := count(items)

	fmt.Println("After : ", result)
}

func count(arr []int) []int {

	// 1. Create an array(slice) of the size of the maximum value + 1
	k := max(arr)
	count := make([]int, k+1)

	// 2. Count each element
	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}
	// fmt.Println(count)

	// 3. Add up the elements
	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}
	// fmt.Println(count)

	// 4. Put them back to result
	result := make([]int, len(arr)+1)
	for j := 0; j < len(arr); j++ {
		result[count[arr[j]]] = arr[j]
		count[arr[j]] = count[arr[j]] - 1
	}
	// fmt.Println(count)

	return result
}

func max(arr []int) int {
	max := arr[0]
	for _, elem := range arr {
		if max < elem {
			max = elem
		}
	}
	return max
}

func countIntArray(arr []int) map[int]int {
	model := make(map[int]int)
	for _, elem := range arr {
		// first element is already initialized with 0
		model[elem]++
	}
	return model
}
