package main

import "fmt"

type matrix [][]int

func main() {
	src := []rune("aaaa")
	dest := []rune("bbbb")

	DMtx := make(matrix, len(src)+1)
	for i := range DMtx {
		DMtx[i] = make([]int, len(dest)+1)
	}

	rows := len(src) + 1  // (1)
	cols := len(dest) + 1 // (2)

	for i := 1; i < rows; i++ { // (3)
		DMtx[i][0] = i // (4)
	}
	for j := 1; j < cols; j++ { // (5)
		DMtx[0][j] = j // (6)
	}

	match := 0                  // (7)
	for i := 1; i < rows; i++ { // (8)
		for j := 1; j < cols; j++ { // (9)
			if src[i-1] == dest[j-1] { // (10)
				match = 0 // (11)
			} else {
				match = 1 // (12)
			}
			insert := DMtx[i][j-1] + 1        // (13)
			delete := DMtx[i-1][j] + 1        // (14)
			replace := DMtx[i-1][j-1] + match // (15)

			min := insert      // (16)
			if replace < min { // (17)
				min = replace // (18)
			}
			if delete < min { // (19)
				min = delete // (20)
			}
			DMtx[i][j] = min // (21)
		}
	}
	fmt.Println(DMtx[rows-1][cols-1])
}
