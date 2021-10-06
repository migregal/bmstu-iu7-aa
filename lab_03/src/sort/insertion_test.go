package sort

import (
	"math/rand"
	"testing"
)

// Insertion sort benchmarks.

func BenchmarkInsertionSorted10(b *testing.B) {
	InsertionSorted(b, 10)
}

func BenchmarkInsertionSorted100(b *testing.B) {
	InsertionSorted(b, 100)
}

func BenchmarkInsertionSorted200(b *testing.B) {
	InsertionSorted(b, 200)
}

func BenchmarkInsertionSorted300(b *testing.B) {
	InsertionSorted(b, 300)
}

func BenchmarkInsertionSorted400(b *testing.B) {
	InsertionSorted(b, 400)
}

func BenchmarkInsertionSorted500(b *testing.B) {
	InsertionSorted(b, 500)
}

func BenchmarkInsertionSorted1000(b *testing.B) {
	InsertionSorted(b, 1000)
}

func InsertionSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionReverseSorted10(b *testing.B) {
	InsertionReversedSorted(b, 10)
}

func BenchmarkInsertionReverseSorted100(b *testing.B) {
	InsertionReversedSorted(b, 100)
}

func BenchmarkInsertionReverseSorted200(b *testing.B) {
	InsertionReversedSorted(b, 200)
}

func BenchmarkInsertionReverseSorted300(b *testing.B) {
	InsertionReversedSorted(b, 300)
}

func BenchmarkInsertionReverseSorted400(b *testing.B) {
	InsertionReversedSorted(b, 400)
}

func BenchmarkInsertionReverseSorted500(b *testing.B) {
	InsertionReversedSorted(b, 500)
}

func BenchmarkInsertionReverseSorted1000(b *testing.B) {
	InsertionReversedSorted(b, 1000)
}

func InsertionReversedSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}

func BenchmarkInsertionRandom10(b *testing.B) {
	InsertionRandom(b, 10)
}

func BenchmarkInsertionRandom100(b *testing.B) {
	InsertionRandom(b, 100)
}

func BenchmarkInsertionRandom200(b *testing.B) {
	InsertionRandom(b, 200)
}

func BenchmarkInsertionRandom300(b *testing.B) {
	InsertionRandom(b, 300)
}

func BenchmarkInsertionRandom400(b *testing.B) {
	InsertionRandom(b, 400)
}

func BenchmarkInsertionRandom500(b *testing.B) {
	InsertionRandom(b, 500)
}
func BenchmarkInsertionRandom1000(b *testing.B) {
	InsertionRandom(b, 1000)
}

func InsertionRandom(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Insertion(narr)
	}
}
