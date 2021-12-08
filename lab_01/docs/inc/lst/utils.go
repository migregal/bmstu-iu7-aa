package levenshtein

func minFromThree(a, b, c int) int {
	return minFromTwo(a, minFromTwo(b, c))
}

func minFromTwo(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
