package levenshtein

func (l Levenshtein) DamerauLevenshteinRec() int {
	return l.getDamerauLevenshteinRec(len(l.s1), len(l.s2))
}

func (l *Levenshtein) getDamerauLevenshteinRec(i, j int) int {
	if i == 0 || j == 0 {
		return abs(len(l.s1) - len(l.s2))
	}
	eq := 1
	if l.s1[i-1] == l.s2[j-1] {
		eq = 0
	}
	res := minFromThree(
		l.getDamerauLevenshteinRec(i, j-1)+1,
		l.getDamerauLevenshteinRec(i-1, j)+1,
		l.getDamerauLevenshteinRec(i-1, j-1)+eq,
	)
	if i > 1 && j > 1 &&
		l.s1[i-1] == l.s2[j-2] &&
		l.s1[i-2] == l.s2[j-1] {
		res = minFromTwo(res, l.getDamerauLevenshteinRec(i-2, j-2)+1)
	}
	return res
}
