package insertionSort

import (
	"alg_go/test/GeneralTest"
	"alg_go/test/utils"
	"testing"
)

var insertionSortTest = InsertionSort{}

func Test_RunAll_GeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &insertionSortTest)
}

func Test_optimisticScenario(t *testing.T) {
	//given
	var inputArr = []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//when
	sortedArr, err := insertionSortTest.Sort(inputArr)
	if err != nil {
		t.Fatalf("Expected error nil but get : %s", err)
	}

	if !utils.IsSorted(sortedArr) {
		t.Fatalf("Array isn't sorted.")
	}
	if !utils.ContainElements(sortedArr, inputArr) {
		t.Fatalf("Elements in arrays are not the same.")
	}

}
