package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"lab_02/matrix"
)

func main() {
	fmt.Println(aurora.Magenta("Matrix multiplication\n"))

	fmt.Println(aurora.Cyan("Matrix A"))
	fmt.Print(aurora.Cyan("Input rows count: "))
	an := matrix.ReadNum()
	fmt.Print(aurora.Cyan("Input cols count: "))
	am := matrix.ReadNum()
	fmt.Println(aurora.Cyan("Input matrix data:"))
	A := matrix.ReadMatrix(an, am)

	fmt.Println(aurora.Cyan("\nMatrix B"))
	fmt.Print(aurora.Cyan("Input rows count: "))
	bn := matrix.ReadNum()
	if bn != am {
		fmt.Println(aurora.Red("Multiplication is imposible. " +
			"Matrix B rows count is not equal " +
			" to Matrix A cols count."))
		return
	}

	fmt.Print(aurora.Cyan("Input cols count: "))
	bm := matrix.ReadNum()
	fmt.Println(aurora.Cyan("Input matrix data:"))
	B := matrix.ReadMatrix(bn, bm)

	fmt.Println(aurora.Yellow("\nResult"))
	fmt.Println(aurora.Cyan("Simple multiplication:"))
	matrix.SimpleMult(A, B).PrintMatrix()
	fmt.Println(aurora.Cyan("Winograd multiplication:"))
	matrix.WinogradMult(A, B).PrintMatrix()
	fmt.Println(aurora.Cyan("Winograd multiplication (imporved):"))
	matrix.WinogradMultImp(A, B).PrintMatrix()
}
