package main

import (
	"testing"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func BenchmarkMain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		items := random.Slice(100000000, 1, 500)

		count(items)
	}
}
