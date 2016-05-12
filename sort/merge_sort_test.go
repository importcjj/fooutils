package sort

import "testing"

func TestIntMergeSort(t *testing.T) {
	slice := []int{5, 5, 6, 6, 2, 3, 1, 4, 99, 23, 23, 4, 5}
	valid := []int{1, 2, 3, 4, 4, 5, 5, 5, 6, 6, 23, 23, 99}

	IntMergeSort(slice, 0, 12)
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Faild IntMergeSort(), Got %v, expected %v", slice, valid)
		}
	}

}

func BenchmarkIntMergeSort(b *testing.B) {
	b.StopTimer()
	slice := []int{5, 4, 3, 2, 1, 9, 8, 7, 6, 3, 4, 5, 6}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s := make([]int, 13)
		copy(s, slice)
		IntMergeSort(slice, 0, 12)
	}
}
