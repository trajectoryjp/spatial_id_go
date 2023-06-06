package main

import (
	"fmt"
	"os"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID1 := "0/0/0/0"
	spatialID2 := "1/0/0/0"
	spatialID3 := "2/0/0/0"
	spatialID4 := "3/0/0/0"
	spatialID5 := "4/0/0/0"
	spatialID6 := "5/0/0/0"
	spatialID7 := "6/0/0/0"
	spatialID8 := "7/0/0/0"
	spatialID9 := "8/0/0/0"
	spatialID10 := "9/0/0/0"
	spatialID11 := "10/0/0/0"
	spatialID12 := "11/0/0/0"
	spatialID13 := "12/0/0/0"
	spatialID14 := "13/0/0/0"
	spatialID15 := "14/0/0/0"
	spatialID16 := "15/0/0/0"
	spatialID17 := "16/0/0/0"
	spatialID18 := "17/0/0/0"
	spatialID19 := "18/0/0/0"
	spatialID20 := "19/0/0/0"
	spatialID21 := "20/0/0/0"
	spatialID22 := "21/0/0/0"
	spatialID23 := "22/0/0/0"
	spatialID24 := "23/0/0/0"
	spatialID25 := "24/0/0/0"
	spatialID26 := "25/0/0/0"
	spatialID27 := "26/0/0/0"
	spatialID28 := "27/0/0/0"
	spatialID29 := "28/0/0/0"
	spatialID30 := "29/0/0/0"
	spatialID31 := "30/0/0/0"
	spatialID32 := "31/0/0/0"
	spatialID33 := "32/0/0/0"
	spatialID34 := "33/0/0/0"
	spatialID35 := "34/0/0/0"
	spatialID36 := "35/0/0/0"

	spatialIDs := []string{spatialID1, spatialID2, spatialID3, spatialID4, spatialID5, spatialID6, spatialID7, spatialID8, spatialID9,
		spatialID10, spatialID11, spatialID12, spatialID13, spatialID14, spatialID15, spatialID16, spatialID17, spatialID18, spatialID19,
		spatialID20, spatialID21, spatialID22, spatialID23, spatialID24, spatialID25, spatialID26, spatialID27, spatialID28, spatialID29,
		spatialID30, spatialID31, spatialID32, spatialID33, spatialID34, spatialID35, spatialID36}

	for index, spatialID := range spatialIDs {
		// ボクセル中心
		centers, err := shape.GetPointOnSpatialId(spatialID, enum.Center)
		if err != nil {
			fmt.Println(fmt.Errorf("ボクセル中心算出時にエラー発生: %w", err))
			os.Exit(1)
		}

		// 結果を出力
		for _, center := range centers {
			fmt.Printf("空間ID: %d/0/0/0 ボクセルの中心: %.16f\n", index, *center)
		}
	}

}
