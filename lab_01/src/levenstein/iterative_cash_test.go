package levenshtein

import "testing"

func BenchmarkIterativeCachedLen1(b *testing.B) {
	l := NewLevenstein("a", "a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen2(b *testing.B) {
	l := NewLevenstein("ab", "ab")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}
func BenchmarkIterativeCachedLen3(b *testing.B) {
	l := NewLevenstein("abo", "abo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}
func BenchmarkIterativeCachedLen4(b *testing.B) {
	l := NewLevenstein("abou", "abov")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen5(b *testing.B) {
	l := NewLevenstein("about", "above")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen10(b *testing.B) {
	l := NewLevenstein("abbanition", "abaptiston")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen15(b *testing.B) {
	l := NewLevenstein("characteristics", "recommendations")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen20(b *testing.B) {
	l := NewLevenstein("abdominohysterectomy", "acetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}
func BenchmarkIterativeCachedLen30(b *testing.B) {
	l := NewLevenstein("chlorobenzylidenemalononitrile", "abdominalexternalobliquemuscle")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen50(b *testing.B) {
	l := NewLevenstein("chlorobenzylidenemalononitrileabdominohysterectomy",
		"abdominalexternalobliquemuscleacetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen100(b *testing.B) {
	l := NewLevenstein(
		"chlorobenzylidenemalononitrileabdominohysterectomy"+
			"chlorobenzylidenemalononitrileabdominohysterectomy",
		"abdominalexternalobliquemuscleacetylcholinesterase"+
			"abdominalexternalobliquemuscleacetylcholinesterase")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.IterativeCache()
	}
}

func BenchmarkIterativeCachedLen200(b *testing.B) {
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
		l.IterativeCache()
	}
}
