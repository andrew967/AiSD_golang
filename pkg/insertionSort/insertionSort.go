package insertionSort

import (
	"alg_go/pkg/services"
	"errors"
)

type InsertionSort struct{}

func (is *InsertionSort) Sort(arr []float32) ([]float32, error) {
	err := services.ValidateParams(arr)
	if err != nil {
		return nil, errors.New("Array is empty")
	}

	insertionSort(arr)

	return arr, nil
}

func insertionSort(arr []float32) []float32 {
	for i := 1; i < len(arr); i++ {

		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}

	}

	return arr
}
