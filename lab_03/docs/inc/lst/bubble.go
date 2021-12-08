package sort

func Bubble(arr []int) {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
}
