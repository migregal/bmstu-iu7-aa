package levenshtein

import "testing"

func BenchmarkRecursiveDLLen1(b *testing.B) {
	l := NewLevenstein("a", "a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen2(b *testing.B) {
	l := NewLevenstein("ab", "ab")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen3(b *testing.B) {
	l := NewLevenstein("abo", "abo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen4(b *testing.B) {
	l := NewLevenstein("abou", "abov")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen5(b *testing.B) {
	l := NewLevenstein("about", "above")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen10(b *testing.B) {
	l := NewLevenstein("abbanition", "abaptiston")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}

func BenchmarkRecursiveDLLen15(b *testing.B) {
	l := NewLevenstein("characteristics", "recommendations")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinRec()
	}
}
