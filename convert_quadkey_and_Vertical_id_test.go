// 空間IDパッケージ
package spatialID

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestConvertTileXYZsToSpatialIDs_01(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{"23/-2/85263/65423"},

		23, 23, 85263, 65423, 0,

		23, 8,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_02(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{"25/1/170526/130846", "25/1/170526/130847", "25/1/170527/130846", "25/1/170527/130847"},

		24, 25, 85263, 65423, 3,

		25, 2,

		25,
	)
}

func TestConvertTileXYZsToSpatialIDs_03(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{"4/0/63/23", "4/1/63/23"},

		4, 3, 63, 23, 3,

		3, 2,

		3,
	)
}

func TestConvertTileXYZsToSpatialIDs_04(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{"23/-2/85263/65423", "23/-1/85263/65423"},

		23, 23, 85263, 65423, 0,

		25, 7,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_05(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,

		[]string{"26/6/85263/65423", "26/7/85263/65423"},

		26, 26, 85263, 65423, 3,

		25, -2,

		26,
	)
}

func TestConvertTileXYZsToSpatialIDs_06_01(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/0/170526/130846",
			"23/1/170526/130846",
			"23/2/170526/130846",
			"23/0/170526/130847",
			"23/1/170526/130847",
			"23/2/170526/130847",
		},

		22, 23, 85263, 65423, 0,

		25, -1,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_06_02(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/0/170527/130846",
			"23/1/170527/130846",
			"23/2/170527/130846",
			"23/0/170527/130847",
			"23/1/170527/130847",
			"23/2/170527/130847",
		},

		22, 23, 85263, 65423, 1,

		25, -1,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_07_01(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/0/85263/65423",
			"23/1/85263/65423",
		},

		23, 23, 85263, 65423, 0,

		25, -1,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_07_02(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/1/85263/65423",
			"23/2/85263/65423",
		},

		23, 23, 85263, 65423, 1,

		25, -1,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_07_03(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/0/85264/65424",
			"23/1/85264/65424",
		},

		23, 23, 85264, 65424, 0,

		25, -1,

		23,
	)
}

func TestConvertTileXYZsToSpatialIDs_07_04(t *testing.T) {
	testNewSpatialIDBoxFromTileXYZBox(
		t,
		[]string{
			"23/1/85264/65424",
			"23/2/85264/65424",
		},

		23, 23, 85264, 65424, 1,

		25, -1,

		23,
	)
}

func testNewSpatialIDBoxFromTileXYZBox(
	t *testing.T,
	expected []string,
	quadkeyZoomLevel int8,
	altitudekeyZoomLevel int8,
	x int64,
	y int64,
	z int64,
	tileXYZZBaseExponent int8,
	tileXYZZBaseOffset int64,
	spatialIDZoomLevel int8,
	) {
	tileXYZ, theError := NewTileXYZ(quadkeyZoomLevel, altitudekeyZoomLevel, x, y, z)

	tileXYZBox, theError := NewTileXYZBox(*tileXYZ, *tileXYZ)
	if theError != nil {
		t.Fatal(theError)
	}

	oldZBaseExponent := TileXYZZBaseExponent
	oldZBaseOffset := TileXYZZBaseOffset
	defer func() {
		TileXYZZBaseExponent = oldZBaseExponent
		TileXYZZBaseOffset = oldZBaseOffset
	}()
	TileXYZZBaseExponent = tileXYZZBaseExponent
	TileXYZZBaseOffset = tileXYZZBaseOffset

	spatialIDBox, theError := NewSpatialIDBoxFromTileXYZBox(*tileXYZBox)
	if theError != nil {
		t.Fatal(theError)
	}
	spatialIDBox.AddZ(spatialIDZoomLevel-spatialIDBox.GetMin().GetZ())

	i := 0
	for id := range spatialIDBox.AllXYF() {
		if id.String() != expected[i] {
			t.Fatal(id)
		}
		i += 1
	}
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_Max_UpFlat(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 10,
				x: 170526,
				y: 130846,
				z: 511, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 10,
				x: 170527,
				y: 130846,
				z: 511, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 10,
				x: 170526,
				y: 130847,
				z: 511, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 10,
				x: 170527,
				y: 130847,
				z: 511, // Max
			},
		},
		nil,

		"20/32768/85263/65423", // 32768 = 1 << (1 + TileXYZZBaseExponent)

		21, 9,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_Max_UpUp(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 11,
				x: 170526,
				y: 130846,
				z: 1023, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 11,
				x: 170527,
				y: 130846,
				z: 1023, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 11,
				x: 170526,
				y: 130847,
				z: 1023, // Max
			},
			{
				quadkeyZoomLevel: 21,
				altitudekeyZoomLevel: 11,
				x: 170527,
				y: 130847,
				z: 1023, // Max
			},
		},
		nil,

		"20/32768/85263/65423", // 32768 = 1 << (1 + TileXYZZBaseExponent)

		21, 10,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_Max_DownDown(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 19,
				altitudekeyZoomLevel: 9,
				x: 42631,
				y: 32711,
				z: 255, // Max
			},
		},
		nil,

		"20/32768/85263/65423", // 32768 = 1 << (1 + TileXYZZBaseExponent)

		19, 8,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_Max_DownUp(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 19,
				altitudekeyZoomLevel: 11,
				x: 42631,
				y: 32711,
				z: 1023, // Max
			},
		},
		nil,

		"20/32768/85263/65423", // 32768 = 1 << (1 + TileXYZZBaseExponent)

		19, 10,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_MinZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 0,
				altitudekeyZoomLevel: 0,
				x: 0,
				y: 0,
				z: 0,
			},
		},
		nil,

		"0/-1/0/0",

		0, 0,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_MaxZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: MaxQuadkeyZoomLevel,
				altitudekeyZoomLevel: 24, // MaxZ - SpatialIDZBaseExponent + TileXYZZBaseExponent
				x: 34359738367, // = (1 << MaxZ) - 1
				y: 34359738367, // = (1 << MaxZ) - 1
				z: 34359738879, // = (1 << MaxZ) - 1 - SpatialIDZOffset + TileXYZZBaseOffset
			},
		},
		nil,

		"35/34359738367/34359738367/34359738367", // 34359738367 = (1 << MaxZ) - 1

		MaxQuadkeyZoomLevel, 24,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_UnderMinQuadkeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		-1, 0,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_OverMaxQuadkeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		36, 0, // 36 = MaxQuadkeyZoomLevel + 1
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_UnderMinAltitudekeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		0, -1,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_OverMaxAltitudekeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		0, 36, // 36 = MaxAltitudekeyZoomLevel + 1
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs(t *testing.T) {
	quadkeyAndVerticalIDs := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID := object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(21, [][2]int64{{29728048124, 1023}, {29728048125, 1023}, {29728048126, 1023}, {29728048127, 1023}}, 10, 500, 0)
	quadkeyAndVerticalIDs = append(quadkeyAndVerticalIDs, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsUpup := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(21, [][2]int64{{29728048124, 2047}, {29728048125, 2047}, {29728048126, 2047}, {29728048127, 2047}}, 11, 500, 0)
	quadkeyAndVerticalIDsUpup = append(quadkeyAndVerticalIDsUpup, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsDwdw := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(19, [][2]int64{{1858003007, 511}}, 9, 500, 0)
	quadkeyAndVerticalIDsDwdw = append(quadkeyAndVerticalIDsDwdw, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsDwup := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(19, [][2]int64{{1858003007, 2047}}, 11, 500, 0)
	quadkeyAndVerticalIDsDwup = append(quadkeyAndVerticalIDsDwup, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsSpatialIDs := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(21, [][2]int64{{29728048124, 112}, {29728048124, 113}, {29728048125, 112}, {29728048125, 113}, {29728048126, 112}, {29728048126, 113}, {29728048127, 112}, {29728048127, 113}}, 21, 0, 0)
	quadkeyAndVerticalIDsSpatialIDs = append(quadkeyAndVerticalIDsSpatialIDs, newQuadkeyAndVerticalID)
	quadkeyAndVerticalIDsSpatialIDsUpdw := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(21, [][2]int64{{29728048124, 28}, {29728048125, 28}, {29728048126, 28}, {29728048127, 28}}, 19, 0, 0)
	quadkeyAndVerticalIDsSpatialIDsUpdw = append(quadkeyAndVerticalIDsSpatialIDsUpdw, newQuadkeyAndVerticalID)
	quadkeyAndVerticalIDsSpatialIDsDwup := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(19, [][2]int64{{1858003007, 112}, {1858003007, 113}}, 21, 0, 0)
	quadkeyAndVerticalIDsSpatialIDsDwup = append(quadkeyAndVerticalIDsSpatialIDsDwup, newQuadkeyAndVerticalID)
	quadkeyAndVerticalIDsSpatialIDsDwdw := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(19, [][2]int64{{1858003007, 28}}, 19, 0, 0)
	quadkeyAndVerticalIDsSpatialIDsDwdw = append(quadkeyAndVerticalIDsSpatialIDsDwdw, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsHBorders1 := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(1, [][2]int64{{0, 56}}, 26, 0, 0)
	quadkeyAndVerticalIDsHBorders1 = append(quadkeyAndVerticalIDsHBorders1, newQuadkeyAndVerticalID)
	quadkeyAndVerticalIDsHBorders31 := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(31, [][2]int64{{29031296, 0}}, 10, 500, 0)
	quadkeyAndVerticalIDsHBorders31 = append(quadkeyAndVerticalIDsHBorders31, newQuadkeyAndVerticalID)

	// quadkeyAndVerticalIDsValueE := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	// newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(31, [][2]int64{{29031296, 1}, {29031296, 0}}, 10, 500, 0)
	// quadkeyAndVerticalIDsValueE = append(quadkeyAndVerticalIDsValueE, newQuadkeyAndVerticalID)

	_, err := strconv.ParseInt("test", 10, 64)
	datas := []struct {
		spatialIds   []string
		ToHZoom      int64
		ToVZoom      int64
		maxHeight    float64
		minHeight    float64
		result       []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID
		resultLength int
		pattern      int64 // 0:正常 1:異常 2:個数(水平) 3:個数(垂直)
		e            error
	}{
		// 異常系(精度エラー)
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 0, ToVZoom: 10, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 20, ToVZoom: -1, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"35/56/85263/65423"}, ToHZoom: 32, ToVZoom: 10, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 20, ToVZoom: 36, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"36/56/85263/65423"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},

		// 異常系(高度エラー)
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 1, ToVZoom: 10, maxHeight: -500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		// 異常系(入力エラー)
		{spatialIds: []string{"20/56/test/65423"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, err.Error())},
	}
	for _, p := range datas {
		result, e := ConvertSpatialIDsToQuadkeysAndVerticalIDs(p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight)
		if p.pattern == 0 && !SortCheck(result, p.result) {
			t.Log(t.Name())
			t.Errorf("ConvertSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, p.result[0], result[0])
		}
		if p.pattern == 1 && e != p.e {
			t.Log(t.Name())
			t.Errorf("ConvertSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, e, p.e)
		}

		if p.pattern == 2 && p.resultLength != len(result[0].InnerIDList()) {
			t.Log(t.Name())
			t.Errorf("ConvertSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, len(result[0].InnerIDList()), p.resultLength)
		}
		if p.pattern == 3 && p.resultLength != len(result[0].InnerIDList()) {
			t.Log(t.Name())
			t.Errorf("ConvertSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, len(result[0].InnerIDList()), p.resultLength)
		}

	}
}

func testNewTileXYZBoxFromSpatialIDBox(
	t *testing.T,
	expected []*TileXYZ,
	expectedError error,
	spatialIDString string,
	quadkeyZoomLevel int8,
	altitudekeyZoomLevel int8,
	) {
	spatialID, theError := NewSpatialIDFromString(spatialIDString)
	if theError != nil {
		if theError.Error() != expectedError.Error() {
			t.Errorf("expectedError: %+v, result: %+v", expectedError, theError)
		}
		return
	}

	spatialIDBox, theError := NewSpatialIDBox(*spatialID, *spatialID)
	if theError != nil {
		if theError.Error() != expectedError.Error() {
			t.Errorf("expectedError: %+v, result: %+v", expectedError, theError)
		}
		return
	}

	tileXYZBox, theError := NewTileXYZBoxFromSpatialIDBox(*spatialIDBox)
	if theError != nil {
		if theError.Error() != expectedError.Error() {
			t.Errorf("expectedError: %+v, result: %+v", expectedError, theError)
		}
		return
	}

	tileXYZBox.AddZoomLevel(quadkeyZoomLevel - tileXYZBox.GetMin().GetQuadkeyZoomLevel(), altitudekeyZoomLevel - tileXYZBox.GetMin().GetAltitudekeyZoomLevel())

	i := 0
	for tileXYZ := range tileXYZBox.AllXYZ() {
		if !reflect.DeepEqual(&tileXYZ, expected[i]) {
			t.Fatal(tileXYZ)
		}
		i += 1
	}
}

func SortCheck(r []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID, rt []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID) bool {
	// 実行結果とテスト値が一致しない場合、false。それ以外の場合、true
	for index, rTmp := range r {
		rtTmp := rt[index]
		if rtTmp.QuadkeyZoom() != rTmp.QuadkeyZoom() {
			return false
		}
		if rtTmp.VerticalZoom() != rTmp.VerticalZoom() {
			return false
		}
		if rtTmp.MaxHeight() != rTmp.MaxHeight() {
			return false
		}
		if rtTmp.MinHeight() != rTmp.MinHeight() {
			return false
		}
		for _, li := range rTmp.InnerIDList() {
			b := false
			for _, li2 := range rtTmp.InnerIDList() {
				if li2 == li {
					b = true
				}
			}
			// テスト値の内部形式IDが関数の実行結果に存在しない場合、false
			if !b {
				return b
			}
		}
	}
	// 最後に到達した場合はTrue
	return true
}

func TestDeleteDuplicationList(t *testing.T) {
	datas := []struct {
		duplicationList []string
		result          []string
	}{
		{duplicationList: []string{"20/562356/78451/23/956", "20/562356/78451/23/957", "20/562356/78451/23/956", "20/562356/78451/23/957"}, result: []string{"20/562356/78451/23/956", "20/562356/78451/23/957"}},
	}
	for _, p := range datas {
		result := deleteDuplicationList(p.duplicationList)
		sort.Strings(result)
		sort.Strings(p.result)
		if !reflect.DeepEqual(result, p.result) {
			t.Log(t.Name())
			t.Errorf("deleteDuplicationList(%s) == %s, result: %s", p.duplicationList, p.result, result)
		}

	}
}

func TestConvertHorizontalIDToQuadkey(t *testing.T) {
	datas := []struct {
		horizontalID string
		result       int64
	}{
		{horizontalID: "20/45621/43566", result: 3448507833},         //"00003031203000312321"
		{horizontalID: "26/4562451/2343566", result: 26508024119725}, //"00012001233201113020012231"
		{horizontalID: "26/1/2", result: 9},                          //"00000000000000000000000021"
		{horizontalID: "26/2/1", result: 6},                          //"00000000000000000000000012"
		{horizontalID: "5/4562451/2343566", result: 429},             //"12231"
		///{horizontalID: "31/2147483647/2147483647", result: 5},
	}
	for _, p := range datas {
		result := convertHorizontalIDToQuadkey(p.horizontalID)
		if result != p.result {
			t.Log(t.Name())
			t.Errorf("convertHorizontalIDToQuadkey(%s) == %d, result: %d", p.horizontalID, p.result, result)
		}
	}
}
func TestConvertQuadkeyToHorizontalID(t *testing.T) {
	datas := []struct {
		quadkey int64
		zoom    int64
		resultX int64
		resultY int64
	}{
		{quadkey: 2914, zoom: 6, resultX: 24, resultY: 53},   //2914-""231202""
		{quadkey: 438, zoom: 6, resultX: 22, resultY: 13},    //438-"012312"
		{quadkey: 14628, zoom: 7, resultX: 82, resultY: 100}, //14628-"3210210"
	}
	for _, p := range datas {
		resultX, resultY := convertQuadkeyToHorizontalID(p.quadkey, p.zoom)
		if resultX != p.resultX && resultY != p.resultY {
			t.Log(t.Name())
			t.Errorf("convertQuadkeyToHorizontalID(%d,%d) == %d,%d, resultX: %d,resultY: %d,", p.quadkey, p.zoom, p.resultX, p.resultY, resultX, resultY)
		}

	}
}
func TestConvertVerticallIDToBit(t *testing.T) {
	datas := []struct {
		vZoom      int64
		vIndex     int64
		outputZoom int64
		maxHeight  float64
		minHeight  float64
		result     []int64
	}{
		// 正常系 1111−0000すべて
		{vZoom: 16, vIndex: 0, outputZoom: 4, maxHeight: 500.0, minHeight: 0.0, result: []int64{15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
		// 正常系 bit形式の枠外(正方向)
		{vZoom: 13, vIndex: 6, outputZoom: 5, maxHeight: 0.0, minHeight: -500.0, result: []int64{31}},
		// 正常系 bit形式の枠外(負方向)
		{vZoom: 13, vIndex: -6, outputZoom: 5, maxHeight: 500.0, minHeight: 0.0, result: []int64{0}},
		{vZoom: 20, vIndex: 0, outputZoom: 8, maxHeight: 500.0, minHeight: -500.0, result: []int64{136, 128, 129, 130, 131, 132, 133, 134, 135}},
		//{vZoom: 35, vIndex: 0, outputZoom: 35, maxHeight: 1.0, minHeight: 0.0, result: []int64{136, 128, 129, 130, 131, 132, 133, 134, 135}},
	}
	for _, p := range datas {
		result := convertVerticallIDToBit(p.vZoom, p.vIndex, p.outputZoom, p.maxHeight, p.minHeight)
		if !reflect.DeepEqual(result, p.result) {
			t.Log(t.Name())
			t.Errorf("convertVerticallIDToBit(%d,%d, %d,%f,%f) == %d, result: %d", p.vZoom, p.vIndex, p.outputZoom, p.maxHeight, p.minHeight, p.result, result)
		}

	}
}

func TestCalcBitIndex(t *testing.T) {
	datas := []struct {
		altitude   float64
		outputZoom int64
		maxHeight  float64
		minHeight  float64
		pattern    int64
		result     int64
	}{
		// 正常(地表)
		{altitude: 256.0, outputZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 0, result: 524},
		// 正常(桁数補正)
		{altitude: 0.0, outputZoom: 10, maxHeight: 256.0, minHeight: -256.0, pattern: 0, result: 512},
		// 正常(地中)
		{altitude: -200.0, outputZoom: 10, maxHeight: 0.0, minHeight: -500.0, pattern: 0, result: 614},
		// 正常 All1
		{altitude: 2560.0, outputZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 0, result: 1023},
		// 正常 All0
		{altitude: -256.0, outputZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 0, result: 0},
		// 正常 0
		{altitude: -256.0, outputZoom: 1, maxHeight: 500.0, minHeight: 0.0, pattern: 0, result: 0},
		// 正常 1
		{altitude: 512.0, outputZoom: 1, maxHeight: 500.0, minHeight: 0.0, pattern: 0, result: 1},
	}
	for _, p := range datas {
		result := calcBitIndex(p.altitude, p.outputZoom, p.maxHeight, p.minHeight)
		if result != p.result {
			t.Log(t.Name())
			t.Errorf("calcBitIndex(%f, %d,%f,%f) == %d, result: %d", p.altitude, p.outputZoom, p.maxHeight, p.minHeight, p.result, result)
		}
	}
}

func TestConvertBitToVerticalID(t *testing.T) {
	datas := []struct {
		vZoom      int64
		vIndex     int64
		outputZoom int64
		maxHeight  float64
		minHeight  float64
		pattern    int64
		result     []string
	}{
		// 正常系
		// 01010101010101010101010101→5592405
		{vZoom: 26, vIndex: 5592405, outputZoom: 26, maxHeight: 500, minHeight: 0, pattern: 0, result: []string{"26/83", "26/83"}},
		{vZoom: 2, vIndex: 4, outputZoom: 25, maxHeight: 500, minHeight: 0, pattern: 1, result: []string{"126", "500", "375"}},
		{vZoom: 26, vIndex: 5592405, outputZoom: 25, maxHeight: 0, minHeight: -500, pattern: 0, result: []string{"25/-459", "25/-459"}},
		// インデックスの補完
		{vZoom: 8, vIndex: 85, outputZoom: 26, maxHeight: 1000, minHeight: 0, pattern: 0, result: []string{"26/671", "26/664", "26/665", "26/666", "26/667", "26/668", "26/669", "26/670"}},
	}
	for _, p := range datas {
		result := convertBitToVerticalID(p.vZoom, p.vIndex, p.outputZoom, p.maxHeight, p.minHeight)
		if p.pattern == 0 {
			if !reflect.DeepEqual(result, p.result) {
				t.Log(t.Name())
				t.Errorf("convertBitToVerticalID(%d,%d, %d,%f,%f) == %s, result: %s", p.vZoom, p.vIndex, p.outputZoom, p.maxHeight, p.minHeight, p.result, result)
			}
		} else {
			// スライスの長さ,桁数
			length, _ := strconv.Atoi(p.result[0])
			if len(result) != length {
				t.Log(t.Name())
				t.Errorf("convertBitToVerticalID(%d,%d, %d,%f,%f) == %s, result: %s", p.vZoom, p.vIndex, p.outputZoom, p.maxHeight, p.minHeight, p.result[0], result)
			}
		}
	}

}

func TestQuadkeyCheckZoom(t *testing.T) {
	// テストデータのテーブルを定義
	datas := []struct {
		HZoom  int64
		VZoom  int64
		result bool
	}{
		{HZoom: 1, VZoom: 12, result: true},
		{HZoom: 2, VZoom: 12, result: true},
		{HZoom: 31, VZoom: 12, result: true},
		{HZoom: 12, VZoom: 0, result: true},
		{HZoom: 12, VZoom: 1, result: true},
		{HZoom: 12, VZoom: 35, result: true},
		{HZoom: 0, VZoom: 12, result: false},
		{HZoom: 32, VZoom: 12, result: false},
		{HZoom: 12, VZoom: -1, result: false},
		{HZoom: 12, VZoom: 36, result: false},
	}

	for _, p := range datas {
		result := quadkeyCheckZoom(p.HZoom, p.VZoom)
		if result != p.result {
			t.Log(t.Name())
			t.Errorf("quadkeyCheckZoom(%d, %d) == %s, result: %s", p.HZoom, p.VZoom, strconv.FormatBool(p.result), strconv.FormatBool(result))
		}
	}
}

func TestSpatialIDCheckZoom(t *testing.T) {
	// テストデータのテーブルを定義
	datas := []struct {
		HZoom  int64
		VZoom  int64
		result bool
	}{
		{HZoom: 0, VZoom: 12, result: true},
		{HZoom: 1, VZoom: 12, result: true},
		{HZoom: 35, VZoom: 12, result: true},
		{HZoom: 12, VZoom: 0, result: true},
		{HZoom: 12, VZoom: 1, result: true},
		{HZoom: 12, VZoom: 35, result: true},
		{HZoom: -1, VZoom: 12, result: false},
		{HZoom: 36, VZoom: 12, result: false},
		{HZoom: 12, VZoom: -1, result: false},
		{HZoom: 12, VZoom: 36, result: false},
	}

	for _, p := range datas {
		result := extendedSpatialIDCheckZoom(p.HZoom, p.VZoom)
		if result != p.result {
			t.Log(t.Name())
			t.Errorf("extendedSpatialIDCheckZoom(%d, %d) == %s, result: %s", p.HZoom, p.VZoom, strconv.FormatBool(p.result), strconv.FormatBool(result))
		}
	}

}

type argSetForConvertZToMinMaxAltitudekey struct {
	inputIndex    int64
	inputZoom     int64
	outputZoom    int64
	zBaseExponent int64
	zBaseOffset   int64
}

func assertConvertZToMinMaxAltitudekey(t *testing.T, expectedMin int64, expectedMax int64, testInput argSetForConvertZToMinMaxAltitudekey) {
	resultMin, resultMax, err := ConvertZToMinMaxAltitudekey(
		testInput.inputIndex,
		testInput.inputZoom,
		testInput.outputZoom,
		testInput.zBaseExponent,
		testInput.zBaseOffset,
	)
	if err != nil {
		t.Fatal(err)
	}

	if resultMin != expectedMin || resultMax != expectedMax {
		t.Errorf("expected=[%v, %v], result=[%v, %v], input=%+v", expectedMin, expectedMax, resultMin, resultMax, testInput)
	}
}

func TestConvertZToMinMaxAltitudekey_1(t *testing.T) {
	var expectedMin int64 = 400
	var expectedMax int64 = 403
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		27,
		25,
		0,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_2(t *testing.T) {
	var expectedMin int64 = 50
	expectedMax := expectedMin
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		24,
		25,
		0,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_3(t *testing.T) {
	var expectedMin int64 = 100
	expectedMax := expectedMin
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		25,
		25,
		0,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_4(t *testing.T) {
	var expectedMin int64 = 53
	expectedMax := expectedMin
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		25,
		25,
		-47,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_5(t *testing.T) {
	expectedError := errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")

	result, _, error := ConvertZToMinMaxAltitudekey(
		100,
		25,
		21,
		25,
		-272,
	)
	if error != expectedError {
		t.Fatal(result, error)
	}
}

func TestConvertZToMinMaxAltitudekey_6(t *testing.T) {
	expectedError := errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")

	result, _, error := ConvertZToMinMaxAltitudekey(
		100,
		25,
		25,
		25,
		-512,
	)
	if error != expectedError {
		t.Fatal(result, error)
	}
}

func TestConvertZToMinMaxAltitudekey_7(t *testing.T) {
	var expectedMin int64 = 1000
	expectedMax := expectedMin
	args := argSetForConvertZToMinMaxAltitudekey{
		28,
		25,
		14,
		25,
		2048000,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_8(t *testing.T) {
	var expectedMin int64 = 100
	expectedMax := expectedMin
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		24,
		24,
		0,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinMaxAltitudekey_9(t *testing.T) {
	var expectedMin int64 = 800
	var expectedMax int64 = 807
	args := argSetForConvertZToMinMaxAltitudekey{
		100,
		25,
		27,
		24,
		0,
	}
	assertConvertZToMinMaxAltitudekey(t, expectedMin, expectedMax, args)
}

func TestConvertZToMinAltitudekey_1(t *testing.T) {
	expected := int64(47)

	result, error := convertZToMinAltitudekey(
		0,
		25,
		25,
		25,
		47,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_2(t *testing.T) {
	expected := int64(0)

	result, error := convertZToMinAltitudekey(
		0,
		25,
		25,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_3(t *testing.T) {
	expected := int64(0)

	result, error := convertZToMinAltitudekey(
		0,
		25,
		27,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_4(t *testing.T) {
	expected := int64(4)

	result, error := convertZToMinAltitudekey(
		1,
		25,
		27,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_5(t *testing.T) {
	expected := int64(3276800)

	result, error := convertZToMinAltitudekey(
		100,
		10,
		25,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_6(t *testing.T) {
	expectedError := errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")

	result, error := convertZToMinAltitudekey(
		100,
		10,
		25,
		25,
		-3276801,
	)
	if error != expectedError {
		t.Fatal(result, error)
	}
}

func TestConvertZToMinAltitudekey_7(t *testing.T) {
	expected := int64(24)

	result, error := convertZToMinAltitudekey(
		47,
		25,
		24,
		25,
		2,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_8(t *testing.T) {
	expected := int64(2)

	result, error := convertZToMinAltitudekey(
		47,
		25,
		20,
		25,
		32,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_9(t *testing.T) {
	expected := int64(12)

	result, error := convertZToMinAltitudekey(
		47,
		25,
		12,
		14,
		4,
	)
	if error != nil {
		t.Fatal(error)
	}

	if result != expected {
		t.Fatal(result)
	}
}

func TestConvertZToMinAltitudekey_10(t *testing.T) {
	expectedError := errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")

	result, error := convertZToMinAltitudekey(
		-1,
		25,
		25,
		25,
		0,
	)
	if error != expectedError {
		t.Fatal(result, error)
	}
}

func TestConvertZToMinAltitudekey_11(t *testing.T) {
	expectedError := errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")

	result, error := convertZToMinAltitudekey(
		-100,
		25,
		26,
		24,
		51,
	)
	if error != expectedError {
		t.Fatal(result, error)
	}
}

type argsForConvertAltitudekeyToMinMaxZ struct {
	altitudekey          int64
	altitudekeyZoomLevel int64
	outputZoomLevel      int64
	zBaseExponent        int64
	zBaseOffset          int64
}

func TestConvertAltitudekeyToMinMaxZ_OffseMustBeConverted(t *testing.T) {
	var expectedMin int64 = -2
	expectedMax := expectedMin
	args := argsForConvertAltitudekeyToMinMaxZ{
		altitudekey:          0,
		altitudekeyZoomLevel: 23,
		outputZoomLevel:      23,
		zBaseExponent:        23,
		zBaseOffset:          8,
	}
	assertConvertAltitudekeyToMinMaxZ(t, args, expectedMin, expectedMax, false)
}

func TestConvertAltitudekeyToMinMaxZ_ZoomLevelMustBeConverted(t *testing.T) {
	var expectedMin int64 = 0
	expectedMax := expectedMin
	args := argsForConvertAltitudekeyToMinMaxZ{
		altitudekey:          0,
		altitudekeyZoomLevel: 23,
		outputZoomLevel:      22,
		zBaseExponent:        25,
		zBaseOffset:          -2,
	}
	assertConvertAltitudekeyToMinMaxZ(t, args, expectedMin, expectedMax, false)
}

func TestConvertAltitudekeyToMinMaxZ_MinMaxDiffersWhenZBaseOffsetIsNotPowerOf2(t *testing.T) {
	var expectedMin int64 = -2
	var expectedMax int64 = -1
	args := argsForConvertAltitudekeyToMinMaxZ{
		altitudekey:          0,
		altitudekeyZoomLevel: 23,
		outputZoomLevel:      23,
		zBaseExponent:        25,
		zBaseOffset:          7,
	}
	assertConvertAltitudekeyToMinMaxZ(t, args, expectedMin, expectedMax, false)
}

func TestConvertAltitudekeyToMinMaxZ_MinMaxDiffersWhenZBaseExponentLessThanInputZoomLevelOrOutputZoomLevel(t *testing.T) {
	var expectedMin int64 = 6
	var expectedMax int64 = 7
	args := argsForConvertAltitudekeyToMinMaxZ{
		altitudekey:          3,
		altitudekeyZoomLevel: 26,
		outputZoomLevel:      26,
		zBaseExponent:        25,
		zBaseOffset:          -2,
	}
	assertConvertAltitudekeyToMinMaxZ(t, args, expectedMin, expectedMax, false)
}

func assertConvertAltitudekeyToMinMaxZ(t *testing.T, args argsForConvertAltitudekeyToMinMaxZ, expectedMin int64, expectedMax int64, wantError bool) {
	gotMinZ, gotMaxZ, err := ConvertAltitudekeyToMinMaxZ(args.altitudekey, args.altitudekeyZoomLevel, args.outputZoomLevel, args.zBaseExponent, args.zBaseOffset)
	if (err != nil) != wantError {
		t.Errorf("ConvertAltitudekeyToMinMaxZ() error = %v, wantErr %v", err, wantError)
		return
	}
	if gotMinZ != expectedMin {
		t.Errorf("ConvertAltitudekeyToMinMaxZ() gotMinZ = %v, minZ %v", gotMinZ, expectedMin)
	}
	if gotMaxZ != expectedMax {
		t.Errorf("ConvertAltitudekeyToMinMaxZ() gotMaxZ = %v, maxZ %v", gotMaxZ, expectedMax)
	}
}
