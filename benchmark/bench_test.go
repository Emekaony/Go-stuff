package benchmark

import "testing"

// benchmark the sumInt
func BenchmarkSumInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumInt(1, 2, 3)
	}
}
