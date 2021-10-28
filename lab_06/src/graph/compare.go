package graph

import (
	"lab_06/utils"
	"time"
)

func Compare(fn string) {
	ant := make([]time.Duration, 0)
	brute := make([]time.Duration, 0)
	for i := 2; i <= 11; i++ {
		genData(fn, i)
		w := utils.GetWeights(fn)
		col := CreateColony(w)

		start := time.Now()
		col.SearchAnt(100)
		end := time.Now()
		ant = append(ant, end.Sub(start))

		start = time.Now()
		SearchBrute(w)
		end = time.Now()
		brute = append(brute, end.Sub(start))
	}

	utils.LogTime("ANT ALGORITHM", ant)
	utils.LogTime("BRUTE ALGORITHM", brute)
}
