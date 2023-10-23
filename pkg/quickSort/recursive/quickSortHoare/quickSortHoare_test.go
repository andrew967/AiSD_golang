package quickSortHoare

import (
	"alg_go/test/GeneralTest"
	"testing"
)

var QuickSortHoareTest = QuickSortHoare{}

func Test_RunAllGeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &QuickSortHoareTest)
}
