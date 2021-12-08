package levenshtein

func NewLevenstein(s1, s2 string) Levenshtein {
	return Levenshtein{[]rune(s1), []rune(s2)}
}

type Levenshtein struct {
	s1 []rune
	s2 []rune
}
