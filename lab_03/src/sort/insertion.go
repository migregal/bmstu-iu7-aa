package sort

func Insertion(arr []int) {
	for i := 0; i < len(arr); i++ {
		x, j := arr[i], i
		for j > 0 && arr[j-1] > x {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = x
	}
}
