package sort

import "testing"

func TestIntQuickSort(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1, 9, 8, 7, 6, 3, 4, 5, 6}
	valid := []int{1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9}

	IntQuickSort(slice, 0, 12)
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Failed IntQuickSort(). Got %v, expected %v", slice, valid)
		}
	}
}

func BenchmarkIntQuickSort(b *testing.B) {
	b.StopTimer()
	slice := []int{5, 4, 3, 2, 1, 9, 8, 7, 6, 3, 4, 5, 6}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s := make([]int, 13)
		copy(s, slice)
		IntQuickSort(s, 3, 12)
	}
}
