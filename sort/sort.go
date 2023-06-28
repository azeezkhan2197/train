package sort

import "github.com/azeezkhan2197/train/model"

func Partition(a []model.BOGIE, low, high int) int {
	pivot := a[low]
	i, j := low, high

	for i < j {
		for a[i].Distance >= pivot.Distance && i < high {
			i++
		}

		for a[j].Distance <= pivot.Distance && j > low {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[low] = a[j]
	a[j] = pivot
	return j
}

func Quick_sort(a []model.BOGIE, low, high int) {
	if low < high {
		pivot_pos := Partition(a, low, high)
		Quick_sort(a, low, pivot_pos)
		Quick_sort(a, pivot_pos+1, high)
	}
}
