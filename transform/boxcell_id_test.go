package transform

import (
	"fmt"
	"reflect"
	"testing"
)

// this test will check to see if GetBoxcellIDfromSpatialID() returns the
// expected BoxcellIDvalue from a known Spatial-Boxcell ID set
// 拡張空間ID： 25/200/29803148/13212522
// expectVal := []int64{200, 29803148, 0}

func TestGetBoxcellIDfromSpatialID(t *testing.T) {

	spatialID := "25/200/29803148/25/0" // known spatialID

	// expected value
	expectVal := []int64{200, 29803148, 0}

	// result from GetBoxcell...() function
	resultVal := GetBoxcellIDfromSpatialID(spatialID)

	// check if result and expected values are the same
	if !reflect.DeepEqual(expectVal, resultVal) {
		// if error, print the values
		t.Errorf("BoxcellID expected value: %v, BoxcellID return value: %v", expectVal, resultVal)
	} else {
		// if success, print them anyway, lol
		fmt.Printf("Congrats, your BoxcellID expected value(%v) and result value (%v) are 'Deeply Equal'.\n", expectVal, resultVal)
	}

	t.Log("Test completed.")

}
