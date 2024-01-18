package transform

import (
	"log"
	"testing"
	"time"

	"github.com/trajectoryjp/spatial_id_go/common/object"
)

func TestGetSpatialIdsWithinRadiusOfLine(t *testing.T) {

	// sumida river, 200m
	startPoint, error := object.NewPoint(139.788452, 35.670935, 100)
	if error != nil {
		t.Fatal(error)
	}
	endPoint, error := object.NewPoint(139.788074, 35.672711, 100)
	if error != nil {
		t.Fatal(error)
	}

	start := time.Now()
	spatialIds, error := GetSpatialIdsWithinRadiusOfLine(startPoint, endPoint, 10, 31, 31)
	elapsed := time.Since(start)
	if error != nil {
		t.Fatal(error)
	}

	log.Printf("\n%v Spatial IDs found in %v \n", len(spatialIds), elapsed)

}
