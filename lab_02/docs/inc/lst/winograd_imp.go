package matrix

func WinogradMultImp(A, B MInt) MInt {
	n, m, q := A.n, B.m, B.n
	C := formResMat(n, m)
	tmpA, tmpB := precomputeRowsImp(A), precomputeColsImp(B)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp := -(tmpA[i] + tmpB[j])
			for k := 1; k < q; k += 2 {
				temp += (A.mat[i][k-1] + B.mat[k][j]) *
					(A.mat[i][k] + B.mat[k-1][j])
			}
			C.mat[i][j] = temp
		}
	}
	if q&1 != 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				C.mat[i][j] += A.mat[i][q-1] * B.mat[q-1][j]
			}
		}
	}
	return C
}

func precomputeRowsImp(M MInt) []int {
	res := make([]int, M.n)
	for i := 0; i < M.n; i++ {
		for j := 1; j < M.m; j += 2 {
			res[i] += M.mat[i][j-1] * M.mat[i][j]
		}
	}
	return res
}

func precomputeColsImp(M MInt) []int {
	res := make([]int, M.m)
	for i := 0; i < M.m; i++ {
		for j := 1; j < M.n; j += 2 {
			res[i] += M.mat[j-1][i] * M.mat[j][i]
		}
	}
	return res
}
