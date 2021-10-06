package matrix

import (
	"testing"
)

func BenchmarkWinogradImp100(b *testing.B) {
	WinogradImpN(b, 100)
}

func BenchmarkWinogradImp200(b *testing.B) {
	WinogradImpN(b, 200)
}

func BenchmarkWinogradImp300(b *testing.B) {
	WinogradImpN(b, 300)
}

func BenchmarkWinogradImp400(b *testing.B) {
	WinogradImpN(b, 400)
}

func BenchmarkWinogradImp500(b *testing.B) {
	WinogradImpN(b, 500)
}

func BenchmarkWinogradImp600(b *testing.B) {
	WinogradImpN(b, 600)
}

func BenchmarkWinogradImp700(b *testing.B) {
	WinogradImpN(b, 700)
}

func BenchmarkWinogradImp800(b *testing.B) {
	WinogradImpN(b, 800)
}

func BenchmarkWinogradImp101(b *testing.B) {
	WinogradImpN(b, 101)
}

func BenchmarkWinogradImp201(b *testing.B) {
	WinogradImpN(b, 201)
}

func BenchmarkWinogradImp301(b *testing.B) {
	WinogradImpN(b, 301)
}

func BenchmarkWinogradImp401(b *testing.B) {
	WinogradImpN(b, 401)
}

func BenchmarkWinogradImp501(b *testing.B) {
	WinogradImpN(b, 501)
}

func BenchmarkWinogradImp601(b *testing.B) {
	WinogradImpN(b, 601)
}

func BenchmarkWinogradImp701(b *testing.B) {
	WinogradImpN(b, 701)
}

func BenchmarkWinogradImp801(b *testing.B) {
	WinogradImpN(b, 801)
}

func WinogradImpN(b *testing.B, N int) {
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}
