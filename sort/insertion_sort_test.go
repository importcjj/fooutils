package sort

import "testing"

func TestIntInsertionSort(t *testing.T) {
	slice := []int{5, 5, 6, 6, 2, 3, 1, 4, 99, 23, 23, 4, 5}
	valid := []int{1, 2, 3, 4, 4, 5, 5, 5, 6, 6, 23, 23, 99}

	IntInsertionSort(slice, 0, 12)
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Falid IntInsertSort(). Got %v, expected %v", slice, valid)
		}
	}
}

func BenchmarkIntInsertionSort(b *testing.B) {
	slice := []int{5, 5, 6, 6, 2, 3, 1, 4, 99, 23, 23, 4, 5}

	for i := 0; i < b.N; i++ {
		b := make([]int, 13)
		copy(b, slice)
		IntInsertionSort(b, 0, 12)
	}
}
