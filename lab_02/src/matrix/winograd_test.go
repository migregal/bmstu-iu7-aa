package matrix

import (
	"testing"
)

func BenchmarkWinograd100(b *testing.B) {
	WinogradN(b, 100)
}

func BenchmarkWinograd200(b *testing.B) {
	WinogradN(b, 200)
}

func BenchmarkWinograd300(b *testing.B) {
	WinogradN(b, 300)
}

func BenchmarkWinograd400(b *testing.B) {
	WinogradN(b, 400)
}

func BenchmarkWinograd500(b *testing.B) {
	WinogradN(b, 500)
}

func BenchmarkWinograd600(b *testing.B) {
	WinogradN(b, 600)
}

func BenchmarkWinograd700(b *testing.B) {
	WinogradN(b, 700)
}

func BenchmarkWinograd800(b *testing.B) {
	WinogradN(b, 800)
}

func BenchmarkWinograd101(b *testing.B) {
	WinogradN(b, 101)
}

func BenchmarkWinograd201(b *testing.B) {
	WinogradN(b, 201)
}

func BenchmarkWinograd301(b *testing.B) {
	WinogradN(b, 301)
}

func BenchmarkWinograd401(b *testing.B) {
	WinogradN(b, 401)
}

func BenchmarkWinograd501(b *testing.B) {
	WinogradN(b, 501)
}

func BenchmarkWinograd601(b *testing.B) {
	WinogradN(b, 601)
}

func BenchmarkWinograd701(b *testing.B) {
	WinogradN(b, 701)
}

func BenchmarkWinograd801(b *testing.B) {
	WinogradN(b, 801)
}

func WinogradN(b *testing.B, N int) {
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}
