package levenshtein

import "testing"

func BenchmarkRecursiveLen1(b *testing.B) {
	l := NewLevenstein("a", "a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen2(b *testing.B) {
	l := NewLevenstein("ab", "ab")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen3(b *testing.B) {
	l := NewLevenstein("abo", "abo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen4(b *testing.B) {
	l := NewLevenstein("abou", "abov")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen5(b *testing.B) {
	l := NewLevenstein("about", "above")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen10(b *testing.B) {
	l := NewLevenstein("abbanition", "abaptiston")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}

func BenchmarkRecursiveLen15(b *testing.B) {
	l := NewLevenstein("characteristics", "recommendations")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Recursive()
	}
}
