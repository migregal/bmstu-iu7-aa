package levenshtein

import "testing"

func BenchmarkDamerauLevenshteinCacheLen1(b *testing.B) {
	l := NewLevenstein("a", "a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen2(b *testing.B) {
	l := NewLevenstein("ab", "ab")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen3(b *testing.B) {
	l := NewLevenstein("abo", "abo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen4(b *testing.B) {
	l := NewLevenstein("abou", "abov")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen5(b *testing.B) {
	l := NewLevenstein("about", "above")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen10(b *testing.B) {
	l := NewLevenstein("abbanition", "abaptiston")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen15(b *testing.B) {
	l := NewLevenstein("characteristics", "recommendations")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen20(b *testing.B) {
	l := NewLevenstein("abdominohysterectomy", "acetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen30(b *testing.B) {
	l := NewLevenstein("chlorobenzylidenemalononitrile", "abdominalexternalobliquemuscle")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen50(b *testing.B) {
	l := NewLevenstein("chlorobenzylidenemalononitrileabdominohysterectomy", "abdominalexternalobliquemuscleacetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen100(b *testing.B) {
	l := NewLevenstein("chlorobenzylidenemalononitrileabdominohysterectomy"+
		"chlorobenzylidenemalononitrileabdominohysterectomy",
		"abdominalexternalobliquemuscleacetylcholinesterase"+
			"abdominalexternalobliquemuscleacetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}

func BenchmarkDamerauLevenshteinCacheLen200(b *testing.B) {
	l := NewLevenstein(
		"chlorobenzylidenemalononitrileabdominohysterectomy"+
			"chlorobenzylidenemalononitrileabdominohysterectomy"+
			"chlorobenzylidenemalononitrileabdominohysterectomy"+
			"chlorobenzylidenemalononitrileabdominohysterectomy",
		"abdominalexternalobliquemuscleacetylcholinesterase"+
			"abdominalexternalobliquemuscleacetylcholinesterase"+
			"abdominalexternalobliquemuscleacetylcholinesterase"+
			"abdominalexternalobliquemuscleacetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DamerauLevenshteinCache()
	}
}
