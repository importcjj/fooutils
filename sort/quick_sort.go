package sort

func IntQuickSort(slice []int, left, right int) {

	if right-left <= 0 {
		return
	}

	// 默认将第一位作为基准位
	pivotIndex := left

	// 把基准交换到最后一位
	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]

	storeIndex := left
	for i := left; i < right; i++ {
		if slice[i] < slice[right] {
			slice[storeIndex], slice[i] = slice[i], slice[storeIndex]
			storeIndex++
		}
	}
	slice[storeIndex], slice[right] = slice[right], slice[storeIndex]

	IntQuickSort(slice, left, storeIndex-1)
	IntQuickSort(slice, storeIndex+1, right)
}
