package sort

import (
	"math/rand"
	"testing"
)

// Selection sort benchmarks.

func BenchmarkSelectionSorted10(b *testing.B) {
	SelectionSorted(b, 10)
}

func BenchmarkSelectionSorted100(b *testing.B) {
	SelectionSorted(b, 100)
}

func BenchmarkSelectionSorted200(b *testing.B) {
	SelectionSorted(b, 200)
}

func BenchmarkSelectionSorted300(b *testing.B) {
	SelectionSorted(b, 300)
}

func BenchmarkSelectionSorted400(b *testing.B) {
	SelectionSorted(b, 400)
}

func BenchmarkSelectionSorted500(b *testing.B) {
	SelectionSorted(b, 500)
}

func BenchmarkSelectionSorted1000(b *testing.B) {
	SelectionSorted(b, 1000)
}

func SelectionSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionReverseSorted10(b *testing.B) {
	SelectionReverseSorted(b, 10)
}

func BenchmarkSelectionReverseSorted100(b *testing.B) {
	SelectionReverseSorted(b, 100)
}

func BenchmarkSelectionReverseSorted200(b *testing.B) {
	SelectionReverseSorted(b, 200)
}

func BenchmarkSelectionReverseSorted300(b *testing.B) {
	SelectionReverseSorted(b, 300)
}

func BenchmarkSelectionReverseSorted400(b *testing.B) {
	SelectionReverseSorted(b, 400)
}

func BenchmarkSelectionReverseSorted500(b *testing.B) {
	SelectionReverseSorted(b, 500)
}

func BenchmarkSelectionReverseSorted1000(b *testing.B) {
	SelectionReverseSorted(b, 1000)
}

func SelectionReverseSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}

func BenchmarkSelectionRandom10(b *testing.B) {
	SelectionRandomSorted(b, 10)
}

func BenchmarkSelectionRandom100(b *testing.B) {
	SelectionRandomSorted(b, 100)
}

func BenchmarkSelectionRandom200(b *testing.B) {
	SelectionRandomSorted(b, 200)
}

func BenchmarkSelectionRandom300(b *testing.B) {
	SelectionRandomSorted(b, 300)
}

func BenchmarkSelectionRandom400(b *testing.B) {
	SelectionRandomSorted(b, 400)
}

func BenchmarkSelectionRandom500(b *testing.B) {
	SelectionRandomSorted(b, 500)
}

func BenchmarkSelectionRandom1000(b *testing.B) {
	SelectionRandomSorted(b, 1000)
}

func SelectionRandomSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Selection(narr)
	}
}
