package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	labsort "lab_03/sort"
)

func main() {
	fmt.Println(aurora.Magenta("Array sorting methods"))

	fmt.Print(aurora.Cyan("Input array size: "))
	n := labsort.ReadNum()
	fmt.Print(aurora.Cyan("Input array: "))
	arr := labsort.ReadArray(n)
	barr := make([]int, n)
	iarr := make([]int, n)
	sarr := make([]int, n)

	fmt.Print(aurora.Yellow("Bubble sort: "))
	copy(barr, arr)
	labsort.Bubble(barr)
	fmt.Println(barr)

	fmt.Print(aurora.Yellow("Insertion sort: "))
	copy(iarr, arr)
	labsort.Insertion(iarr)
	fmt.Println(iarr)

	fmt.Print(aurora.Yellow("Selection sort: "))
	copy(sarr, arr)
	labsort.Selection(sarr)
	fmt.Println(sarr)
}
