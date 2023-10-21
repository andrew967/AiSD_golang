package quickSortLomutoGorutines

import (
	"alg_go/pkg/services"
	"sync"
)

type QuickSortLomutoGorutines struct{}

func (qsl *QuickSortLomutoGorutines) Sort(arr []float32) ([]float32, error) {
	err := services.ValidateParams(arr)
	if err != nil {
		return arr, err
	}

	quickSortLomutoGorutines(arr, 0, len(arr)-1)

	return arr, nil
}

func quickSortLomutoGorutines(arr []float32, low, high int) {
	if low < high {
		pivot_index := partition(arr, low, high)
		wg := sync.WaitGroup{}

		wg.Add(2)

		go func() {
			defer wg.Done()
			quickSortLomutoGorutines(arr, low, pivot_index-1)
		}()

		go func() {
			defer wg.Done()
			quickSortLomutoGorutines(arr, pivot_index+1, high)
		}()

		wg.Wait()
	}
}

func partition(arr []float32, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
