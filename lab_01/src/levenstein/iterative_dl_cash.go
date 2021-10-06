package levenshtein

func (l Levenshtein) DamerauLevenshteinCache() int {
	prev2 := make([]int, len(l.s2)+1)
	prev1 := make([]int, len(l.s2)+1)
	cur := make([]int, len(l.s2)+1)

	for i := 0; i <= len(l.s2); i++ {
		prev1[i] = i
	}

	for i := 1; i <= len(l.s1); i++ {
		cur[0] = i
		for j := 1; j <= len(l.s2); j++ {
			eq := 1
			if l.s1[i-1] == l.s2[j-1] {
				eq = 0
			}
			cur[j] = minFromThree(
				cur[j-1]+1,
				prev1[j]+1,
				prev1[j-1]+eq,
			)
			if i > 1 && j > 1 &&
				l.s1[i-1] == l.s2[j-2] &&
				l.s1[i-2] == l.s2[j-1] {
				cur[j] = minFromTwo(
					cur[j],
					prev2[j-2]+1,
				)
			}
		}
		copy(prev2, prev1)
		copy(prev1, cur)
	}

	return cur[len(l.s2)]
}
