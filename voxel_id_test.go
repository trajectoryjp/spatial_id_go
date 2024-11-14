package spatialID

import (
	"fmt"
	"reflect"
	"testing"
)

// this test will check to see if GetVoxelIDfromSpatialID() returns the
// expected VoxelIDvalue from a known Spatial-Voxel ID set
// 拡張空間ID： 25/200/29803148/13212522
// expectVal := []int64{200, 29803148, 0}

func TestGetVoxelIDfromSpatialID(t *testing.T) {

	spatialID := "25/200/29803148/25/0" // known spatialID

	// expected value
	expectVal := []int64{200, 29803148, 0}

	// result from GetVoxel...() function
	resultVal := GetVoxelIDfromSpatialID(spatialID)

	// check if result and expected values are the same
	if !reflect.DeepEqual(expectVal, resultVal) {
		// if error, print the values
		t.Errorf("VoxelID expected value: %v, VoxelID return value: %v", expectVal, resultVal)
	} else {
		// if success, print them anyway, lol
		fmt.Printf("Congrats, your VoxelID expected value(%v) and result value (%v) are 'Deeply Equal'.\n", expectVal, resultVal)
	}

	t.Log("Test completed.")

}
