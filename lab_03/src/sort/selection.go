package sort

func Selection(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
