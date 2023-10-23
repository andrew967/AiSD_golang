package quickSort

import "alg_go/pkg/services"

type QuickSortIterative struct{}

func (qsi *QuickSortIterative) Sort(arr []float32) ([]float32, error) {
	err := services.ValidateParams(arr)
	if err != nil {
		return arr, err
	}

	quicksortIteratve(arr)

	return arr, nil
}

func quicksortIteratve(arr []float32) {
	stack := make([]int, len(arr))
	top := -1

	//init stack
	top++
	stack[top] = 0
	top++
	stack[top] = len(arr) - 1

	for top >= 0 {
		high := stack[top]
		top--
		low := stack[top]
		top--

		pivotIndex := partition(arr, low, high)

		if pivotIndex-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = pivotIndex - 1
		}
		if pivotIndex+1 < high {
			top++
			stack[top] = pivotIndex + 1
			top++
			stack[top] = high
		}
	}

}

// Lomuto partition
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
