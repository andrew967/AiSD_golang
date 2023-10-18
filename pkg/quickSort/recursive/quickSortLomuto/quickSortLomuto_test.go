package quickSortLomuto

import (
	"alg_go/test/GeneralTest"
	"testing"
)

var quickSortLomutoTest = QuickSortLomuto{}

func Test_RunAllGeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &quickSortLomutoTest)
}
