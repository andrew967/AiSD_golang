package GeneralTest

import (
	"alg_go/pkg/services"
	"alg_go/test/utils"
	"testing"
)

func RunAllGeneralTests(t *testing.T, sorter services.Sorting) {
	t.Run("Should_ReturnError_WithInputNull", func(t *testing.T) {
		Generaltest_Should_ReturnError_WithInputNull(t, sorter)
	})
	t.Run("Generaltest_Should_ReturnSortedArray_WithInputRandomData", func(t *testing.T) {
		Generaltest_Should_ReturnSortedArray_WithInputRandomData(t, sorter)
	})
}

func Generaltest_Should_ReturnError_WithInputNull(t *testing.T, sorter services.Sorting) {
	//given
	var inputArr []float32

	//when
	_, err := sorter.Sort(inputArr)

	//then
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func Generaltest_Should_ReturnSortedArray_WithInputRandomData(t *testing.T, sorter services.Sorting) {
	//given
	var inputArr = utils.GenerateRandomData(100)

	//when
	sortedArr, err := sorter.Sort(inputArr)

	//then
	if err != nil {
		t.Fatalf("Expected nil error but ger %s", err)
	}
	if !utils.IsSorted(sortedArr) {
		t.Fatalf("Array is not sorted")
	}
	if !utils.ContainElements(sortedArr, inputArr) {
		t.Fatalf("Elements in arrays are not the same.")
	}
}
