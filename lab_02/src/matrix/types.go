package matrix

import "fmt"

type MInt struct {
	mat  [][]int
	n, m int
}

func (mat MInt) PrintMatrix() {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Printf("%3d ", mat.mat[i][j])
		}
		fmt.Printf("\n")
	}
}
