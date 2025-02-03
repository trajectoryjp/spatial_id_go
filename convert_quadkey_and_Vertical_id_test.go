// 空間IDパッケージ
package spatialID

import (
	"reflect"
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		-1, 0,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_OverMaxQuadkeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		36, 0, // 36 = MaxQuadkeyZoomLevel + 1
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_UnderMinAltitudekeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		0, -1,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_OverMaxAltitudekeyZoomLevel(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/85263/65423",

		0, 36, // 36 = MaxAltitudekeyZoomLevel + 1
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_UnderMinZ(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"-1/0/0/0",

		0, 0,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_OverMaxZ(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"36/34359738367/34359738367/34359738367",

		MaxQuadkeyZoomLevel, 24,
	)
}

func TestConvertSpatialIdsToQuadkeysAndVerticalIDs_InvalidY(t *testing.T) {
	testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		"20/32768/test/65423",

		21, 9,
	)
}

func testNewTileXYZBoxFromSpatialIDBox_AllXYZ(
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

type argSetForConvertZToMinMaxAltitudekey struct {
	inputIndex    int64
	inputZoom     int64
	outputZoom    int64
	zBaseExponent int64
	zBaseOffset   int64
}

func assertConvertZToMinMaxAltitudekey(
	t *testing.T,
	expected TileXYZBox,
	spatialIDBox SpatialIDBox,
	tileXYZZBaseExponent int8,
	tileXYZZBaseOffset int64,
	altitudekeyZoomLevel int8,
	) {
	oldZBaseExponent := TileXYZZBaseExponent
	oldZBaseOffset := TileXYZZBaseOffset
	defer func() {
		TileXYZZBaseExponent = oldZBaseExponent
		TileXYZZBaseOffset = oldZBaseOffset
	}()
	TileXYZZBaseExponent = tileXYZZBaseExponent
	TileXYZZBaseOffset = tileXYZZBaseOffset

	tileXYZBox, error := NewTileXYZBoxFromSpatialIDBox(spatialIDBox)
	if error != nil {
		t.Fatal(error)
	}

	error = tileXYZBox.AddZoomLevel(0, altitudekeyZoomLevel-tileXYZBox.GetMin().GetAltitudekeyZoomLevel())
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(tileXYZBox, expected) {
		t.Errorf("expected: %+v, result: %+v", expected, tileXYZBox)
	}
}

func TestConvertZToMinMaxAltitudekey_1(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 27,
				x: 85263,
				y: 65423,
				z: 400,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 27,
				x: 85263,
				y: 65423,
				z: 403,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, 0,

		27,
	)
}

func TestConvertZToMinMaxAltitudekey_2(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 24,
				x: 85263,
				y: 65423,
				z: 50,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 24,
				x: 85263,
				y: 65423,
				z: 50,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, 0,

		24,
	)
}

func TestConvertZToMinMaxAltitudekey_3(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 100,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 100,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, 0,

		25,
	)
}

func TestConvertZToMinMaxAltitudekey_4(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 53,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 53,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, -47,

		25,
	)
}

func TestConvertZToMinMaxAltitudekey_5(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 21,
				x: 85263,
				y: 65423,
				z: 0,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 21,
				x: 85263,
				y: 65423,
				z: 0,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, -272,

		21,
	)
}

func TestConvertZToMinMaxAltitudekey_6(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 0,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 25,
				x: 85263,
				y: 65423,
				z: 0,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		25, -512,

		25,
	)
}

func TestConvertZToMinMaxAltitudekey_7(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 14,
				x: 85263,
				y: 65423,
				z: 1000,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 14,
				x: 85263,
				y: 65423,
				z: 1000,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 28,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 28,
				x: 85263,
				y: 65423,
			},
		},

		25, 2048000,

		14,
	)
}

func TestConvertZToMinMaxAltitudekey_8(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 24,
				x: 85263,
				y: 65423,
				z: 100,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 24,
				x: 85263,
				y: 65423,
				z: 100,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		24, 0,

		24,
	)
}

func TestConvertZToMinMaxAltitudekey_9(t *testing.T) {
	assertConvertZToMinMaxAltitudekey(
		t,

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 27,
				x: 85263,
				y: 65423,
				z: 800,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 25,
				altitudekeyZoomLevel: 27,
				x: 85263,
				y: 65423,
				z: 807,
			},
		},

		SpatialIDBox{
			min: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 25,
				f: 100,
				x: 85263,
				y: 65423,
			},
		},

		24, 0,

		27,
	)
}

type argsForConvertAltitudekeyToMinMaxZ struct {
	altitudekey          int64
	altitudekeyZoomLevel int64
	outputZoomLevel      int64
	zBaseExponent        int64
	zBaseOffset          int64
}

func TestConvertAltitudekeyToMinMaxZ_OffseMustBeConverted(t *testing.T) {
	assertConvertAltitudekeyToMinMaxZ(
		t,

		SpatialIDBox{
			min: SpatialID{
				z: 23,
				f: -2,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 23,
				f: -2,
				x: 85263,
				y: 65423,
			},
		},

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 23,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 23,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
		},

		23, 8,

		23,
	)
}

func TestConvertAltitudekeyToMinMaxZ_ZoomLevelMustBeConverted(t *testing.T) {
	assertConvertAltitudekeyToMinMaxZ(
		t,

		SpatialIDBox{
			min: SpatialID{
				z: 22,
				f: 0,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 22,
				f: 0,
				x: 85263,
				y: 65423,
			},
		},

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 22,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 22,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
		},

		25, -2,

		22,
	)
}

func TestConvertAltitudekeyToMinMaxZ_MinMaxDiffersWhenZBaseOffsetIsNotPowerOf2(t *testing.T) {
	assertConvertAltitudekeyToMinMaxZ(
		t,

		SpatialIDBox{
			min: SpatialID{
				z: 23,
				f: -2,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 23,
				f: -1,
				x: 85263,
				y: 65423,
			},
		},
		
		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 23,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 23,
				altitudekeyZoomLevel: 23,
				x: 85263,
				y: 65423,
				z: 0,
			},
		},

		25, 7,

		23,
	)
}

func TestConvertAltitudekeyToMinMaxZ_MinMaxDiffersWhenZBaseExponentLessThanInputZoomLevelOrOutputZoomLevel(t *testing.T) {
	assertConvertAltitudekeyToMinMaxZ(
		t,

		SpatialIDBox{
			min: SpatialID{
				z: 26,
				f: 6,
				x: 85263,
				y: 65423,
			},
			max: SpatialID{
				z: 26,
				f: 7,
				x: 85263,
				y: 65423,
			},
		},

		TileXYZBox{
			min: TileXYZ{
				quadkeyZoomLevel: 26,
				altitudekeyZoomLevel: 26,
				x: 85263,
				y: 65423,
				z: 3,
			},
			max: TileXYZ{
				quadkeyZoomLevel: 26,
				altitudekeyZoomLevel: 26,
				x: 85263,
				y: 65423,
				z: 3,
			},
		},

		25, -2,

		26,
	)
}

func assertConvertAltitudekeyToMinMaxZ(
	t *testing.T,
	expected SpatialIDBox,
	tileXYZBox TileXYZBox,
	zBaseExponent int8,
	zBaseOffset int64,
	z int8,
) {
	oldZBaseExponent := TileXYZZBaseExponent
	oldZBaseOffset := TileXYZZBaseOffset
	defer func() {
		TileXYZZBaseExponent = oldZBaseExponent
		TileXYZZBaseOffset = oldZBaseOffset
	}()
	TileXYZZBaseExponent = zBaseExponent
	TileXYZZBaseOffset = zBaseOffset

	spatialIDBox, error := NewSpatialIDBoxFromTileXYZBox(tileXYZBox)
	if error != nil {
		t.Fatal(error)
	}

	error = spatialIDBox.AddZ(z-spatialIDBox.GetMin().GetZ())
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(spatialIDBox, expected) {
		t.Errorf("expected: %+v, result: %+v", expected, spatialIDBox)
	}
}
