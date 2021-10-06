package levenshtein

func (l Levenshtein) IterativeCache() int {
	prev := make([]int, len(l.s2)+1)
	for i := 0; i <= len(l.s2); i++ {
		prev[i] = i
	}
	cur := make([]int, len(l.s2)+1)
	for i := 1; i <= len(l.s1); i++ {
		cur[0] = i
		for j := 1; j <= len(l.s2); j++ {
			eq := 1
			if l.s1[i-1] == l.s2[j-1] {
				eq = 0
			}
			cur[j] = minFromThree(
				cur[j-1]+1,
				prev[j]+1,
				prev[j-1]+eq,
			)
		}
		copy(prev, cur)
	}
	return cur[len(l.s2)]
}
