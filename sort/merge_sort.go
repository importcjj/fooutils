package sort

func IntMergeSort(slice []int, left, right int) {
	if left >= right {
		return
	}
	cache := make([]int, len(slice))
	intMergeSort(slice, cache, left, right)
}

func intMergeSort(slice, cache []int, left, right int) {
	length := right - left
	if length <= 0 {
		return
	}

	middle := length>>1 + left

	left1, right1 := left, middle
	left2, right2 := middle+1, right

	intMergeSort(slice, cache, left1, right1)
	intMergeSort(slice, cache, left2, right2)

	index := left
	for left1 <= right1 && left2 <= right2 {
		if slice[left1] < slice[left2] {
			cache[index] = slice[left1]
			left1++
		} else {
			cache[index] = slice[left2]
			left2++
		}
		index++

	}

	for left1 <= right1 {
		cache[index] = slice[left1]
		left1++
		index++
	}
	for left2 <= right2 {
		cache[index] = slice[left2]
		left2++
		index++
	}

	for i := left; i <= right; i++ {
		slice[i] = cache[i]
	}
}
