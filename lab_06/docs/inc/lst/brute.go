package graph

import "math"

func SearchBrute(w [][]int) []int {
	var (
		path = make([]int, 0)
		r    = make([]int, len(w))
	)

	for i := 0; i < len(w); i++ {
		var (
			rts  = make([][]int, 0)
			sum  = math.MaxInt64
			curr = 0
		)
		getRoutes(i, w, path, &rts)

		for j := 0; j < len(rts); j++ {
			curr = 0

			for k := 0; k < len(rts[j])-1; k++ {
				curr += w[rts[j][k]][rts[j][k+1]]
			}

			if curr < sum {
				sum = curr
			}
		}

		r[i] = sum
	}

	return r
}

func getRoutes(pos int, w [][]int, path []int, rts *[][]int) {
	path = append(path, pos)

	if len(path) < len(w) {
		for i := 0; i < len(w); i++ {
			if !isExist(path, i) {
				getRoutes(i, w, path, rts)
			}
		}
	} else {
		*rts = append(*rts, path)
	}
}

func isExist(a []int, v int) bool {
	for _, val := range a {
		if v == val {
			return true
		}
	}

	return false
}
