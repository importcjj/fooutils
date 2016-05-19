package sort

func IntInsertionSort(slice []int, left, right int) {
	if right-left <= 0 {
		return
	}
	for i := left + 1; i <= right; i++ {
		temp := slice[i]

		j := i
		for ; j > left && temp < slice[j-1]; j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = temp
	}
}
