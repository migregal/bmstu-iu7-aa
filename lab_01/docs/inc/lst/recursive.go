package levenshtein

func (l Levenshtein) Recursive() int {
	return l.getDistance(len(l.s1), len(l.s2))
}

func (l *Levenshtein) getDistance(i, j int) int {
	if i == 0 || j == 0 {
		return abs(len(l.s1) - len(l.s2))
	}
	eq := 1
	if l.s1[i-1] == l.s2[j-1] {
		eq = 0
	}
	return minFromThree(
		l.getDistance(i, j-1)+1,
		l.getDistance(i-1, j)+1,
		l.getDistance(i-1, j-1)+eq)
}
