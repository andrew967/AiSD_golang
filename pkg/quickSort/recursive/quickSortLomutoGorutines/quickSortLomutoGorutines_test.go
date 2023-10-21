package quickSortLomutoGorutines

import (
	"alg_go/test/GeneralTest"
	"testing"
)

var quickSortLomutoGourutinesTest QuickSortLomutoGorutines

func Test_RunAllGeneralTests(t *testing.T) {
	GeneralTest.RunAllGeneralTests(t, &quickSortLomutoGourutinesTest)
}
