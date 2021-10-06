package matrix

func SimpleMult(amat, bmat MInt) MInt {
	rmat := formResMat(amat.n, bmat.m)

	for i := 0; i < rmat.n; i++ {
		for j := 0; j < rmat.m; j++ {
			for k := 0; k < amat.m; k++ {
				rmat.mat[i][j] += amat.mat[i][k] * bmat.mat[k][j]
			}
		}
	}

	return rmat
}
