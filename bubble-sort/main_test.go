package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		random := rand.Intn(1000-1) + 1
		for i := 0; i < 1000; i++ {
			items = append(items, random)
		}
		bubble(items)
	}
}
