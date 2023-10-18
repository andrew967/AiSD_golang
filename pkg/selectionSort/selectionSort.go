package selectionSort

import "alg_go/pkg/services"

type SelectionSort struct{}

func (ss SelectionSort) Sort(arr []float32) ([]float32, error) {
	err := services.ValidateParams(arr)
	if err != nil {
		return arr, err
	}

	selectionSort(arr)

	return arr, nil
}

func selectionSort(arr []float32) []float32 {
	for i := 0; i < len(arr); i++ {
		minValID := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minValID] > arr[j] {
				minValID = j
			}
		}
		arr[minValID], arr[i] = arr[i], arr[minValID]
	}

	return arr
}
