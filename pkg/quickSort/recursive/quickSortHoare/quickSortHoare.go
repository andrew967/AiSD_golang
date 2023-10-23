package quickSortHoare

import (
	"alg_go/pkg/services"
)

type QuickSortHoare struct{}

func (qsh *QuickSortHoare) Sort(arr []float32) ([]float32, error) {
	err := services.ValidateParams(arr)
	if err != nil {
		return arr, err
	}

	quickSortHoare(arr, 0, len(arr)-1)

	return arr, nil
}

func quickSortHoare(arr []float32, low, high int) {
	if low < high {
		pivot_index := partition(arr, low, high)

		quickSortHoare(arr, low, pivot_index)
		quickSortHoare(arr, pivot_index+1, high)
	}
}

func partition(arr []float32, low, high int) int {
	pivot := arr[(low+high)/2]
	i := low
	j := high
	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}
