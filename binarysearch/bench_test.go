package main

import "testing"

func BenchmarkCompareSearches(b *testing.B) {
	orderedList := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		orderedList[i] = i + 1
	}

	b.Run("BinarySearch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			binarySearch(orderedList, 999_999)
		}
	})

	b.Run("LinearSearch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linearSearch(orderedList, 999_999)
		}
	})
}
