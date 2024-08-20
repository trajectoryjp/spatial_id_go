package main

import (
	"fmt"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/transform"
)

func main() {
	inputData := []struct {
		hZoom uint16
		x     int64
		y     int64
		vZoom uint16
		z     int64
	}{
		{
			22,
			85263,
			65423,
			23,
			0,
		},
		{
			22,
			85263,
			65423,
			23,
			1,
		},
	}
	var inputXYZ []*object.TileXYZ
	for _, in := range inputData {
		tile, err := object.NewTileXYZ(in.hZoom, in.x, in.y, in.vZoom, in.z)
		if err != nil {
			panic(err)
		}
		inputXYZ = append(inputXYZ, tile)
	}
	outputData, err := transform.ConvertTileXYZsToSpatialIDs(
		inputXYZ,
		25,
		-1,
		23,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inputData: %v\n", inputData)
	// {
	//	"23/0/170526/130846",
	//	"23/0/170527/130846",
	//	"23/0/170526/130847",
	//	"23/0/170527/130847",
	//	"23/1/170526/130846",
	//	"23/1/170526/130847",
	//	"23/1/170527/130846",
	//	"23/1/170527/130847",
	//	"23/2/170526/130846",
	//	"23/2/170526/130847",
	//	"23/2/170527/130846",
	//	"23/2/170527/130847",
	//	}
	fmt.Printf("outputData: %v\n", outputData)
}
