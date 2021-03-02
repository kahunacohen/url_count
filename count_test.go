package main

import (
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetBodyLens()
	}
}
