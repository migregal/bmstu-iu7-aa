package matrix

import (
	"testing"
)

func BenchmarkSimple100(b *testing.B) {
	SimpleN(b, 100)
}

func BenchmarkSimple200(b *testing.B) {
	SimpleN(b, 200)
}

func BenchmarkSimple300(b *testing.B) {
	SimpleN(b, 300)
}

func BenchmarkSimple400(b *testing.B) {
	SimpleN(b, 400)
}

func BenchmarkSimple500(b *testing.B) {
	SimpleN(b, 500)
}

func BenchmarkSimple600(b *testing.B) {
	SimpleN(b, 600)
}

func BenchmarkSimple700(b *testing.B) {
	SimpleN(b, 700)
}

func BenchmarkSimple800(b *testing.B) {
	SimpleN(b, 800)
}

func BenchmarkSimple101(b *testing.B) {
	SimpleN(b, 101)
}

func BenchmarkSimple201(b *testing.B) {
	SimpleN(b, 201)
}

func BenchmarkSimple301(b *testing.B) {
	SimpleN(b, 301)
}

func BenchmarkSimple401(b *testing.B) {
	SimpleN(b, 401)
}

func BenchmarkSimple501(b *testing.B) {
	SimpleN(b, 501)
}

func BenchmarkSimple601(b *testing.B) {
	SimpleN(b, 601)
}

func BenchmarkSimple701(b *testing.B) {
	SimpleN(b, 701)
}

func BenchmarkSimple801(b *testing.B) {
	SimpleN(b, 801)
}

func SimpleN(b *testing.B, N int) {
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}
