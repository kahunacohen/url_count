package main

import (
	"testing"
)

func BenchmarkNotParallel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetBodyLens()
	}
}
func BenchmarkParallel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetBodyLens2()
	}
}
