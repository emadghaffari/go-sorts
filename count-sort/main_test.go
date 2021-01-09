package main

import (
	"strconv"
	"testing"

	"github.com/emadghaffari/go-sorts/utils/random"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		b.Run(str, func(b *testing.B) {
			items := random.Slice(1000000, 1, 9999999)

			for i := 0; i < b.N; i++ {
				count(items)
			}
		})
	}
}
