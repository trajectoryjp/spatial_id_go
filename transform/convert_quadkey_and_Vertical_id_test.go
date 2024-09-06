// 空間IDパッケージ
package transform

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"

	"github.com/stretchr/testify/assert"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
)

func TestConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(t *testing.T) {
	// データの作成
	quadkeyAndVerticalIDList := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID := object.NewQuadkeyAndVerticalID(6, 2914, 7, 74, 500, 0) // 231202 6/24/53/5/0
	quadkeyAndVerticalIDList = append(quadkeyAndVerticalIDList, newQuadkeyAndVerticalID)
	newQuadkeyAndVerticalID = object.NewQuadkeyAndVerticalID(6, 2882, 25, 0, 0, 0) // 231002 "7/48/58/2/0","7/49/58/2/0","7/48/59/2/0","7/49/59/2/0"
	quadkeyAndVerticalIDList = append(quadkeyAndVerticalIDList, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDList2 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewQuadkeyAndVerticalID(9, 451739, 25, 0, 0, 0) // 123210212 "9/338/229/2/0"
	quadkeyAndVerticalIDList2 = append(quadkeyAndVerticalIDList2, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDListE := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE := object.NewQuadkeyAndVerticalID(6, 2882, 25, 0, -500, 0) // 231002
	quadkeyAndVerticalIDListE = append(quadkeyAndVerticalIDListE, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE2 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(6, 2882, 2, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE2 = append(quadkeyAndVerticalIDListE2, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE3 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(32, 2882, 2, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE3 = append(quadkeyAndVerticalIDListE3, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE4 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(31, 2882, 36, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE4 = append(quadkeyAndVerticalIDListE4, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE5 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(31, 4611686018427388065, 25, 0, 500, 0) // 231002
	quadkeyAndVerticalIDListE5 = append(quadkeyAndVerticalIDListE5, newQuadkeyAndVerticalIDE)

	datas := []struct {
		quadkeyAndVerticalIDs []*object.QuadkeyAndVerticalID
		ToHZoom               int64
		ToVZoom               int64
		result                []string
		pattern               int64
		e                     error
	}{
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToHZoom: 6, ToVZoom: 20, pattern: 0, result: []string{"6/24/53/20/9", "6/24/49/20/0"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToHZoom: 7, ToVZoom: 2, pattern: 0, result: []string{"7/48/107/2/0", "7/49/107/2/0", "7/48/98/2/0", "7/49/98/2/0", "7/48/99/2/0", "7/49/99/2/0", "7/48/106/2/0", "7/49/106/2/0"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToHZoom: 5, ToVZoom: 8, pattern: 0, result: []string{"5/12/26/8/0", "5/12/24/8/0"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList2, ToHZoom: 9, ToVZoom: 2, pattern: 0, result: []string{"9/338/229/2/0"}},
		// 異常系(精度エラー)
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE, ToHZoom: 0, ToVZoom: 35, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE, ToHZoom: 5, ToVZoom: 36, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE2, ToHZoom: 5, ToVZoom: 35, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE3, ToHZoom: 5, ToVZoom: 35, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE4, ToHZoom: 5, ToVZoom: 35, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE5, ToHZoom: 5, ToVZoom: 35, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
	}

	for _, p := range datas {
		result, e := ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(p.quadkeyAndVerticalIDs, p.ToHZoom, p.ToVZoom)
		sort.Strings(result)
		sort.Strings(p.result)
		if p.pattern == 0 && !reflect.DeepEqual(result, p.result) {
			t.Log(t.Name())
			t.Errorf("ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(%+v,%d,%d) == %s, result: %s", p.quadkeyAndVerticalIDs, p.ToHZoom, p.ToVZoom, p.result, result)
			return
		}
		if p.pattern == 1 && e != p.e {
			t.Log(t.Name())
			t.Errorf("ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(%+v,%d,%d) == %+v, result: %+v", p.quadkeyAndVerticalIDs, p.ToHZoom, p.ToVZoom, e, p.e)
		}
	}

}

func TestConvertQuadkeysAndVerticalIDsToSpatialIDs(t *testing.T) {
	// データの作成
	quadkeyAndVerticalIDList := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID := object.NewQuadkeyAndVerticalID(6, 2914, 7, 74, 500, 0) // 231202 6/24/53/5/0
	quadkeyAndVerticalIDList = append(quadkeyAndVerticalIDList, newQuadkeyAndVerticalID)
	newQuadkeyAndVerticalID = object.NewQuadkeyAndVerticalID(6, 2882, 25, 0, 0, 0) // 231002 "7/48/58/2/0","7/49/58/2/0","7/48/59/2/0","7/49/59/2/0"
	quadkeyAndVerticalIDList = append(quadkeyAndVerticalIDList, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDList2 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewQuadkeyAndVerticalID(9, 451739, 25, 0, 0, 0) // 123210212 "9/338/229/2/0"
	quadkeyAndVerticalIDList2 = append(quadkeyAndVerticalIDList2, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDListE := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE := object.NewQuadkeyAndVerticalID(6, 2882, 25, 0, -500, 0) // 231002
	quadkeyAndVerticalIDListE = append(quadkeyAndVerticalIDListE, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE2 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(6, 2882, 2, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE2 = append(quadkeyAndVerticalIDListE2, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE3 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(32, 2882, 2, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE3 = append(quadkeyAndVerticalIDListE3, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE4 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(31, 2882, 36, 10, 500, 0) // 231002
	quadkeyAndVerticalIDListE4 = append(quadkeyAndVerticalIDListE4, newQuadkeyAndVerticalIDE)

	quadkeyAndVerticalIDListE5 := []*object.QuadkeyAndVerticalID{}
	newQuadkeyAndVerticalIDE = object.NewQuadkeyAndVerticalID(31, 4611686018427388065, 25, 0, 500, 0) // 231002
	quadkeyAndVerticalIDListE5 = append(quadkeyAndVerticalIDListE5, newQuadkeyAndVerticalIDE)

	datas := []struct {
		quadkeyAndVerticalIDs []*object.QuadkeyAndVerticalID
		ToZoom                int64
		result                []string
		pattern               int64
		e                     error
	}{
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToZoom: 6, pattern: 0, result: []string{"6/0/24/53", "6/0/24/49"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToZoom: 7, pattern: 0, result: []string{"7/0/48/107", "7/0/49/107", "7/0/48/98", "7/0/49/98", "7/0/48/99", "7/0/49/99", "7/0/48/106", "7/0/49/106"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList, ToZoom: 5, pattern: 0, result: []string{"5/0/12/26", "5/0/12/24"}},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDList2, ToZoom: 9, pattern: 0, result: []string{"9/0/338/229"}},
		// 異常系(精度エラー)
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE, ToZoom: 0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE, ToZoom: 5, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE2, ToZoom: 5, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE3, ToZoom: 5, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE4, ToZoom: 5, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{quadkeyAndVerticalIDs: quadkeyAndVerticalIDListE5, ToZoom: 5, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
	}

	for _, p := range datas {
		result, e := ConvertQuadkeysAndVerticalIDsToSpatialIDs(p.quadkeyAndVerticalIDs, p.ToZoom)
		sort.Strings(result)
		sort.Strings(p.result)
		if p.pattern == 0 && !reflect.DeepEqual(result, p.result) {
			t.Log(t.Name())
			t.Errorf("ConvertQuadkeysAndVerticalIDsToSpatialIDs(%+v,%d) == %s, result: %s", p.quadkeyAndVerticalIDs, p.ToZoom, p.result, result)
			return
		}
		if p.pattern == 1 && e != p.e {
			t.Log(t.Name())
			t.Errorf("ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(%+v,%d) == %+v, result: %+v", p.quadkeyAndVerticalIDs, p.ToZoom, e, p.e)
		}
	}

}

func TestConvertExtendedSpatialIdsToQuadkeysAndVerticalIDs(t *testing.T) {
	// 結果確認用の構造体を作成する
	//"20/85263/65423"→ 00012322332320003333 →7432012031 21:29728048124,29728048125,29728048126,29728048127,
	//horizontalID: "20/45621/43566", result: 3448507833},         //"00003031203000312321"
	//horizontalID: "26/4562451/2343566", result: 26508024119725}, //"00012001233201113020012231"
	//horizontalID: "26/1/2", result: 9},                          //"00000000000000000000000021"
	//horizontalID: "26/2/1", result: 6},                          //"00000000000000000000000012"
	//horizontalID: "5/4562451/2343566", result: 429},             //"12231"

	quadkeyAndVerticalIDs := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID := object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(21, [][2]int64{{29728048124, 58}, {29728048124, 57}, {29728048125, 58}, {29728048125, 57}, {29728048126, 58}, {29728048126, 57}, {29728048127, 58}, {29728048127, 57}}, 10, 500, 0)
	quadkeyAndVerticalIDs = append(quadkeyAndVerticalIDs, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsSpatialIDs := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(19, [][2]int64{{1858003007, 56}}, 26, 0, 0)
	quadkeyAndVerticalIDsSpatialIDs = append(quadkeyAndVerticalIDsSpatialIDs, newQuadkeyAndVerticalID)

	quadkeyAndVerticalIDsHBorders1 := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(1, [][2]int64{{0, 56}}, 26, 0, 0)
	quadkeyAndVerticalIDsHBorders1 = append(quadkeyAndVerticalIDsHBorders1, newQuadkeyAndVerticalID)
	quadkeyAndVerticalIDsHBorders31 := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(31, [][2]int64{{29031296, 1}, {29031296, 0}}, 10, 500, 0)
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
		// 正常
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 21, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, pattern: 0},
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 19, ToVZoom: 26, maxHeight: 0, minHeight: 0.0, result: quadkeyAndVerticalIDsSpatialIDs, pattern: 0},

		// 水平精度個数確認 低精度は1、高精度は精度差^4
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 24, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 512, pattern: 2},
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 2, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 2, pattern: 2},
		// 水平精度境界値
		{spatialIds: []string{"20/85263/65423/26/0"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsHBorders1, resultLength: 2, pattern: 2},
		{spatialIds: []string{"35/85263/65423/26/0"}, ToHZoom: 31, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsHBorders31, pattern: 0},

		// 垂直精度境界値
		{spatialIds: []string{"20/85263/65423/26/0"}, ToHZoom: 21, ToVZoom: 0, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 4, pattern: 3},
		{spatialIds: []string{"20/85263/65423/26/0"}, ToHZoom: 21, ToVZoom: 1, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 4, pattern: 3},

		// 異常系(精度エラー)
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 0, ToVZoom: 10, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 20, ToVZoom: -1, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"35/85263/65423/26/56"}, ToHZoom: 32, ToVZoom: 10, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"20/85263/65423/35/56"}, ToHZoom: 20, ToVZoom: 36, maxHeight: 0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"36/85263/65423/26/56"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		{spatialIds: []string{"20/85263/65423/36/56"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},

		// 異常系(高度エラー)
		{spatialIds: []string{"20/85263/65423/26/56"}, ToHZoom: 1, ToVZoom: 10, maxHeight: -500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, "")},
		// 異常系(入力エラー)
		{spatialIds: []string{"20/test/65423/26/56"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500.0, minHeight: 0.0, pattern: 1, e: errors.NewSpatialIdError(errors.InputValueErrorCode, err.Error())},
	}
	for _, p := range datas {

		result, e := ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight)
		if p.pattern == 0 && !reflect.DeepEqual(result, p.result) {
			t.Log(t.Name())
			t.Errorf("ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, p.result[0], result[0])
		}
		if p.pattern == 1 && e != p.e {
			t.Log(t.Name())
			t.Errorf("ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, e, p.e)
		}
		if p.pattern == 2 && p.resultLength != len(result[0].InnerIDList()) {
			t.Log(t.Name())
			t.Errorf("ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, len(result[0].InnerIDList()), p.resultLength)
		}
		if p.pattern == 3 && p.resultLength != len(result[0].InnerIDList()) {
			t.Log(t.Name())
			t.Errorf("ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(%s,%d,%d,%f,%f) == %+v, result: %+v", p.spatialIds, p.ToHZoom, p.ToVZoom, p.maxHeight, p.minHeight, len(result[0].InnerIDList()), p.resultLength)
		}

	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_1(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 56}},
			26,
			25,
			0,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/26/56"},
		20,
		26,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_2(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			21,
			[][2]int64{{29728048124, 56}, {29728048125, 56}, {29728048126, 56}, {29728048127, 56}},
			26,
			25,
			0,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/26/56"},
		21,
		26,
		25,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_3(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 7}},
			12,
			14, // 2^14
			0,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/26/56"},
		20,
		12,
		14,
		0,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_4(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 54}},
			12,
			14,
			188,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/26/56"},
		20,
		12,
		14,
		188,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_5(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			21,
			[][2]int64{{29728048124, 12}, {29728048124, 13}, {29728048125, 12}, {29728048125, 13}, {29728048126, 12}, {29728048126, 13}, {29728048127, 12}, {29728048127, 13}},
			15,
			14,
			-50,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/56"},
		21,
		15,
		14,
		-50,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func newTileXYZ(t *testing.T, hZoom int64, x int64, y int64, vZoom int64, z int64) *object.TileXYZ {
	t.Helper()
	xyz, err := object.NewTileXYZ(hZoom, x, y, vZoom, z)
	if err != nil {
		t.Fatal(err)
	}
	return xyz
}

func newExtendedSpatialID(t *testing.T, id string) *object.ExtendedSpatialID {
	t.Helper()
	extendedSpatialId, err := object.NewExtendedSpatialID(id)
	if err != nil {
		t.Fatal(err)
	}
	return extendedSpatialId
}

func TestConvertTileXYZsToSpatialIDs(t *testing.T) {
	type argSet struct {
		tile          []*object.TileXYZ
		zBaseExponent int64
		zBaseOffset   int64
		outputVZoom   int64
	}
	testCases := []struct {
		expected []string
		request  argSet
	}{
		{
			// z/f/x/y
			[]string{"23/-2/85263/65423"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					23,
					85263,
					65423,
					23,
					0,
				)},
				23,
				8,
				23,
			},
		},
		{
			[]string{"25/1/170526/130846", "25/1/170526/130847", "25/1/170527/130846", "25/1/170527/130847"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					24,
					85263,
					65423,
					25,
					3,
				)},
				25,
				2,
				25,
			},
		},
		{
			[]string{"4/0/63/23", "4/1/63/23"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					4,
					63,
					23,
					3,
					3,
				)},
				3,
				2,
				3,
			},
		},
		{
			[]string{"23/-2/85263/65423", "23/-1/85263/65423"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					23,
					85263,
					65423,
					23,
					0,
				)},
				25,
				7,
				23,
			},
		},
		{
			[]string{"26/6/85263/65423", "26/7/85263/65423"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					26,
					85263,
					65423,
					26,
					3,
				)},
				25,
				-2,
				26,
			},
		},
		{
			[]string{
				"23/0/170526/130846",
				"23/1/170526/130846",
				"23/2/170526/130846",
				"23/0/170526/130847",
				"23/1/170526/130847",
				"23/2/170526/130847",
				"23/0/170527/130846",
				"23/1/170527/130846",
				"23/2/170527/130846",
				"23/0/170527/130847",
				"23/1/170527/130847",
				"23/2/170527/130847",
			},
			argSet{
				[]*object.TileXYZ{
					newTileXYZ(
						t,
						22,
						85263,
						65423,
						23,
						0,
					),
					newTileXYZ(
						t,
						22,
						85263,
						65423,
						23,
						1,
					)},
				25,
				-1,
				23,
			},
		},
	}
	for _, testCase := range testCases {
		result, err := ConvertTileXYZsToSpatialIDs(testCase.request.tile, testCase.request.zBaseExponent, testCase.request.zBaseOffset, testCase.request.outputVZoom)
		if err != nil {
			t.Fatal(err)
		}
		if assert.ElementsMatch(t, testCase.expected, result) == false {
			// t.Error(result):
			t.Errorf("expected: %v result: %v", testCase.expected, result)
		} else {
			t.Log("Success", result)
		}
	}

}

func ExampleConvertTileXYZsToSpatialIDs() {
	inputData := []struct {
		hZoom int64
		x     int64
		y     int64
		vZoom int64
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
	outputData, err := ConvertTileXYZsToSpatialIDs(
		inputXYZ,
		25,
		-1,
		23,
	)
	if err != nil {
		panic(err)
	}
	for _, out := range outputData {
		fmt.Println(out)
	}
	// Unordered output:
	// 23/0/170526/130846
	// 23/0/170527/130846
	// 23/0/170526/130847
	// 23/0/170527/130847
	// 23/1/170526/130846
	// 23/1/170526/130847
	// 23/1/170527/130846
	// 23/1/170527/130847
	// 23/2/170526/130846
	// 23/2/170526/130847
	// 23/2/170527/130846
	// 23/2/170527/130847
}

func TestConvertTileXYZsToExtendedSpatialIDs(t *testing.T) {
	type argSet struct {
		tile          []*object.TileXYZ
		zBaseExponent int64
		zBaseOffset   int64
		outputVZoom   int64
	}
	testCases := []struct {
		expected []string
		request  argSet
	}{
		{
			// hZoom/x/y/vZoom/z
			[]string{"20/85263/65423/23/-2"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					23,
					0,
				)},
				23,
				8,
				23,
			},
		},
		{
			[]string{"20/85263/65423/25/1"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					25,
					3,
				)},
				25,
				2,
				25,
			},
		},
		{
			[]string{"20/85263/65423/3/0"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					3,
					3,
				)},
				3,
				2,
				3,
			},
		},
		{
			[]string{"20/85263/65423/23/-2"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					23,
					0,
				)},
				25,
				8,
				23,
			},
		},
		{
			[]string{"20/85263/65423/23/-2", "20/85263/65423/23/-1"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					23,
					0,
				)},
				25,
				7,
				23,
			},
		},
		{
			[]string{"20/85263/65423/26/6", "20/85263/65423/26/7"},
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					26,
					3,
				)},
				25,
				-2,
				26,
			},
		},
		{
			[]string{"20/85263/65423/23/0", "20/85263/65423/23/1", "20/85263/65423/23/2"},
			argSet{
				[]*object.TileXYZ{
					newTileXYZ(
						t,
						20,
						85263,
						65423,
						23,
						0,
					),
					newTileXYZ(
						t,
						20,
						85263,
						65423,
						23,
						1,
					)},
				25,
				-1,
				23,
			},
		},
		{
			[]string{"20/85263/65423/23/0", "20/85263/65423/23/1", "20/85263/65423/23/2", "20/85264/65424/23/0", "20/85264/65424/23/1", "20/85264/65424/23/2"},
			argSet{
				[]*object.TileXYZ{
					newTileXYZ(
						t,
						20,
						85263,
						65423,
						23,
						0,
					),
					newTileXYZ(
						t,
						20,
						85263,
						65423,
						23,
						1,
					),
					newTileXYZ(
						t,
						20,
						85264,
						65424,
						23,
						0,
					),
					newTileXYZ(
						t,
						20,
						85264,
						65424,
						23,
						1,
					)},
				25,
				-1,
				23,
			},
		},
	}
	for _, testCase := range testCases {
		expectedData := []object.ExtendedSpatialID{}
		for i := 0; i < len(testCase.expected); i++ {
			extendedSpatialId := newExtendedSpatialID(t, testCase.expected[i])
			expectedData = append(expectedData, *extendedSpatialId)
		}
		result, err := ConvertTileXYZsToExtendedSpatialIDs(testCase.request.tile, testCase.request.zBaseExponent, testCase.request.zBaseOffset, testCase.request.outputVZoom)
		if err != nil {
			t.Fatal(err)
		}
		if assert.ElementsMatch(t, expectedData, result) == false {
			// t.Error(result):
			t.Errorf("expected: %v result: %v", expectedData, result)
		} else {
			t.Log("Success", result)
		}
	}

}

func TestErrorConvertTileXYZsToExtendedSpatialIDs(t *testing.T) {
	type argSet struct {
		tile          []*object.TileXYZ
		zBaseExponent int64
		zBaseOffset   int64
		outputVZoom   int64
	}
	testCases := []struct {
		expected string
		request  argSet
	}{
		{
			"input index does not exist",
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					25,
					(1<<consts.ZOriginValue)+1,
				)},
				25,
				8,
				25,
			},
		},
		{
			"output index does not exist",
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					25,
					(1<<consts.ZOriginValue)-1,
				)},
				25,
				-1,
				25,
			},
		},
		{
			"extendSpatialID zoom level must be in 0-35",
			argSet{
				[]*object.TileXYZ{newTileXYZ(
					t,
					20,
					85263,
					65423,
					23,
					0,
				)},
				23,
				8,
				230,
			},
		},
	}
	for _, testCase := range testCases {
		_, err := ConvertTileXYZsToExtendedSpatialIDs(testCase.request.tile, testCase.request.zBaseExponent, testCase.request.zBaseOffset, testCase.request.outputVZoom)
		if err == nil {
			t.Fatal(err)
		}
		if strings.Contains(err.Error(), testCase.expected) == false {
			t.Errorf("expected: %v result: %v", testCase.expected, err)
		} else {
			t.Log("Success", err)
		}
	}
}

func ExampleConvertTileXYZsToExtendedSpatialIDs() {
	inputData := []struct {
		hZoom int64
		x     int64
		y     int64
		vZoom int64
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
		{
			22,
			85264,
			65424,
			23,
			0,
		},
		{
			22,
			85264,
			65424,
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
	outputData, err := ConvertTileXYZsToExtendedSpatialIDs(
		inputXYZ,
		25,
		-1,
		23,
	)
	if err != nil {
		panic(err)
	}
	for _, out := range outputData {
		fmt.Println(out.ID())
	}
	// Unordered output:
	// 22/85263/65423/23/0
	// 22/85264/65424/23/0
	// 22/85263/65423/23/1
	// 22/85264/65424/23/1
	// 22/85263/65423/23/2
	// 22/85264/65424/23/2
}

func TestConvertExtendedSpatialIDsToSpatialIDs(t *testing.T) {
	testCases := []struct {
		expected []string
		id       *object.ExtendedSpatialID
	}{
		{
			// 水平精度の方が低い場合
			[]string{"7/0/48/106", "7/0/48/107", "7/0/49/106", "7/0/49/107"},
			newExtendedSpatialID(t, "6/24/53/7/0"),
		},
		{
			// 水平精度の方が低い場合
			[]string{"7/0/48/98", "7/0/48/99", "7/0/49/98", "7/0/49/99"},
			newExtendedSpatialID(t, "6/24/49/7/0"),
		},
		{
			// 垂直精度の方が低い場合
			[]string{"7/48/24/53", "7/49/24/53"},
			newExtendedSpatialID(t, "7/24/53/6/24"),
		},
		{
			// 水平精度、垂直精度に差がない場合
			[]string{"6/0/24/49"},
			newExtendedSpatialID(t, "6/24/49/6/0"),
		},
	}
	for _, testCase := range testCases {
		result := ConvertExtendedSpatialIDToSpatialIDs(testCase.id)
		if !assert.ElementsMatch(t, testCase.expected, result) {
			t.Errorf("expected: %v result: %v", testCase.expected, result)
		}
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example1(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 3}},
			3,
			3,
			2,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/1"},
		20,
		3,
		3,
		2,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example2(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 2}},
			2,
			3,
			2,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/3"},
		20,
		2,
		3,
		2,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example3(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 2}, {7432012031, 3}},
			3,
			3,
			2,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/24/0"},
		20,
		3,
		3,
		2,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example4(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 0}, {7432012031, 1}, {7432012031, 2}, {7432012031, 3}},
			25,
			23,
			-1,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/1"},
		20,
		25,
		23,
		-1,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example5(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			20,
			[][2]int64{{7432012031, 0}},
			23,
			25,
			-1,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/1"},
		20,
		23,
		25,
		-1,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
}

func TestConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys_Example6(t *testing.T) {
	expected := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{
		object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey( // 例6
			20,
			[][2]int64{{7432012031, 0}},
			23,
			25,
			-1,
		),
	}

	result, error := ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(
		[]string{"20/85263/65423/25/4"},
		20,
		23,
		25,
		-1,
	)
	if error != nil {
		t.Fatal(error)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result)
	}
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
		// 正常
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 21, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, pattern: 0},     //all1
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 21, ToVZoom: 11, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsUpup, pattern: 0}, //all1
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 19, ToVZoom: 9, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsDwdw, pattern: 0},  //all1
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 19, ToVZoom: 11, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsDwup, pattern: 0}, //all1

		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 21, ToVZoom: 21, maxHeight: 0, minHeight: 0.0, result: quadkeyAndVerticalIDsSpatialIDs, pattern: 0},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 21, ToVZoom: 19, maxHeight: 0, minHeight: 0.0, result: quadkeyAndVerticalIDsSpatialIDsUpdw, pattern: 0},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 19, ToVZoom: 21, maxHeight: 0, minHeight: 0.0, result: quadkeyAndVerticalIDsSpatialIDsDwup, pattern: 0},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 19, ToVZoom: 19, maxHeight: 0, minHeight: 0.0, result: quadkeyAndVerticalIDsSpatialIDsDwdw, pattern: 0},

		// 水平精度個数確認 低精度は1、高精度は精度差^4
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 24, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 256, pattern: 2},
		{spatialIds: []string{"20/56/85263/65423"}, ToHZoom: 2, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 1, pattern: 2},
		// 水平精度境界値
		{spatialIds: []string{"20/0/85263/65423"}, ToHZoom: 1, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsHBorders1, resultLength: 66, pattern: 2},
		{spatialIds: []string{"35/0/85263/65423"}, ToHZoom: 31, ToVZoom: 10, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDsHBorders31, pattern: 0},

		// 垂直精度境界値
		{spatialIds: []string{"20/0/85263/65423"}, ToHZoom: 21, ToVZoom: 0, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 4, pattern: 3},
		{spatialIds: []string{"20/0/85263/65423"}, ToHZoom: 21, ToVZoom: 1, maxHeight: 500, minHeight: 0.0, result: quadkeyAndVerticalIDs, resultLength: 4, pattern: 3},

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
