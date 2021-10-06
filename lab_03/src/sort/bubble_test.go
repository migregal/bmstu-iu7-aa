package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSorted10(b *testing.B) {
	BubbleSorted(b, 10)
}

func BenchmarkBubbleSorted100(b *testing.B) {
	BubbleSorted(b, 100)
}

func BenchmarkBubbleSorted200(b *testing.B) {
	BubbleSorted(b, 200)
}

func BenchmarkBubbleSorted300(b *testing.B) {
	BubbleSorted(b, 300)
}

func BenchmarkBubbleSorted400(b *testing.B) {
	BubbleSorted(b, 400)
}

func BenchmarkBubbleSorted500(b *testing.B) {
	BubbleSorted(b, 500)
}

func BenchmarkBubbleSorted1000(b *testing.B) {
	BubbleSorted(b, 1000)
}

func BubbleSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}

func BenchmarkBubbleReverseSorted10(b *testing.B) {
	BubbleReverseSorted(b, 10)
}

func BenchmarkBubbleReverseSorted100(b *testing.B) {
	BubbleReverseSorted(b, 100)
}

func BenchmarkBubbleReverseSorted200(b *testing.B) {
	BubbleReverseSorted(b, 200)
}

func BenchmarkBubbleReverseSorted300(b *testing.B) {
	BubbleReverseSorted(b, 300)
}

func BenchmarkBubbleReverseSorted400(b *testing.B) {
	BubbleReverseSorted(b, 400)
}

func BenchmarkBubbleReverseSorted500(b *testing.B) {
	BubbleReverseSorted(b, 500)
}

func BenchmarkBubbleReverseSorted1000(b *testing.B) {
	BubbleReverseSorted(b, 1000)
}

func BubbleReverseSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = N - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}

func BenchmarkBubbleRandom10(b *testing.B) {
	BubbleRandomSorted(b, 10)
}

func BenchmarkBubbleRandom100(b *testing.B) {
	BubbleRandomSorted(b, 100)
}

func BenchmarkBubbleRandom200(b *testing.B) {
	BubbleRandomSorted(b, 200)
}

func BenchmarkBubbleRandom300(b *testing.B) {
	BubbleRandomSorted(b, 300)
}

func BenchmarkBubbleRandom400(b *testing.B) {
	BubbleRandomSorted(b, 400)
}

func BenchmarkBubbleRandom500(b *testing.B) {
	BubbleRandomSorted(b, 500)
}

func BenchmarkBubbleRandom1000(b *testing.B) {
	BubbleRandomSorted(b, 1000)
}

func BubbleRandomSorted(b *testing.B, N int) {
	arr := make([]int, N)
	narr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N-i) + i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(narr, arr)
		Bubble(narr)
	}
}
