package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSelectionSort(b *testing.B) {
	size := 10_000

	testData := make([]int, size)

	for i := 0; i < size; i++ {
		testData[i] = rand.Intn(size) // Random numbers between 0 and 9999
	}

	for i := 0; i < b.N; i++ {
		// Run the selectionSort function
		selectionSort(testData)
	}
}
