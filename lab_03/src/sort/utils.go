package sort

import "fmt"

func ReadArray(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}
	return arr
}

func ReadNum() int {
	var num int
	fmt.Scanln(&num)
	return num
}
