package matrix

func WinogradMult(A, B MInt) MInt {
	n, m, q := A.n, B.m, B.n
	C := formResMat(n, m)
	tmpA, tmpB := precomputeRows(A), precomputeCols(B)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			C.mat[i][j] = -(tmpA[i] + tmpB[j])
			for k := 0; k < q/2; k++ {
				C.mat[i][j] = C.mat[i][j] +
					(A.mat[i][k*2]+B.mat[k*2+1][j])*
						(A.mat[i][k*2+1]+B.mat[k*2][j])
			}
		}
	}
	if q%2 != 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				C.mat[i][j] = C.mat[i][j] +
					A.mat[i][q-1]*B.mat[q-1][j]
			}
		}
	}
	return C
}

func precomputeRows(M MInt) []int {
	res := make([]int, M.n)
	for i := 0; i < M.n; i++ {
		for j := 0; j < M.m/2; j++ {
			res[i] = res[i] + M.mat[i][j*2]*M.mat[i][j*2+1]
		}
	}
	return res
}

func precomputeCols(M MInt) []int {
	res := make([]int, M.m)
	for i := 0; i < M.m; i++ {
		for j := 0; j < M.n/2; j++ {
			res[i] = res[i] + M.mat[j*2][i]*M.mat[j*2+1][i]
		}
	}
	return res
}
