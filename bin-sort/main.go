package main

import (
	"fmt"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func main() {
	items := random.Slice(20, 1, 15)
	fmt.Println("Before: ", items)

	result := bucket(items, len(items))

	fmt.Println("After : ", result)
}
func insertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func bucket(array []int, bucketSize int) []int {
	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}
