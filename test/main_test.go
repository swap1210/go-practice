package main

import (
	"testing"
)

func BenchmarkGetMinParcel_V(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMinParcel_V([]int{2, 3, 6, 10, 11}, 9)
	}
}
