package main

import (
	"testing"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := random.Slice(100, 1, 500)

		insertion(items)
	}
}
