package transform

import (
	"fmt"
	"reflect"
	"testing"
)

// this test will check to see if GetVoxcellIDfromSpatialID() returns the
// expected VoxcellIDvalue from a known Spatial-Voxcell ID set
// 拡張空間ID： 25/200/29803148/13212522
// expectVal := []int64{200, 29803148, 0}

func TestGetVoxcellIDfromSpatialID(t *testing.T) {

	spatialID := "25/200/29803148/25/0" // known spatialID

	// expected value
	expectVal := []int64{200, 29803148, 0}

	// result from GetVoxcell...() function
	resultVal := GetVoxcellIDfromSpatialID(spatialID)

	// check if result and expected values are the same
	if !reflect.DeepEqual(expectVal, resultVal) {
		// if error, print the values
		t.Errorf("VoxcellID expected value: %v, VoxcellID return value: %v", expectVal, resultVal)
	} else {
		// if success, print them anyway, lol
		fmt.Printf("Congrats, your VoxcellID expected value(%v) and result value (%v) are 'Deeply Equal'.\n", expectVal, resultVal)
	}

	t.Log("Test completed.")

}
