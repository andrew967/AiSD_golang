package selectionSort

import (
	"alg_go/test/GeneralTest"
	"alg_go/test/utils"
	"testing"
)

var selectionSortTest = SelectionSort{}

func Test_RunAllGeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &selectionSortTest)
}

func Test_OptimisticScenario(t *testing.T) {
	//given
	var inputArr = []float32{1, 3, 5, 6, 8, 8, 9}

	//when
	sortedArr, err := selectionSortTest.Sort(inputArr)

	//then
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
