package quickSort

import (
	"alg_go/test/GeneralTest"
	"testing"
)

var QuickSortIterativeTest = QuickSortIterative{}

func Test_RunAllGeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &QuickSortIterativeTest)
}
