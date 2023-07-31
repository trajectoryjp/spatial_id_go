// 空間IDパッケージ
package transform

import (
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/trajectoryjp/spatial_id_go/common/errors"
	"github.com/trajectoryjp/spatial_id_go/common/object"
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

	quadkeyAndVerticalIDsValueE := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(31, [][2]int64{{29031296, 1}, {29031296, 0}}, 10, 500, 0)
	quadkeyAndVerticalIDsValueE = append(quadkeyAndVerticalIDsValueE, newQuadkeyAndVerticalID)

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

	quadkeyAndVerticalIDsValueE := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	newQuadkeyAndVerticalID = object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(31, [][2]int64{{29031296, 1}, {29031296, 0}}, 10, 500, 0)
	quadkeyAndVerticalIDsValueE = append(quadkeyAndVerticalIDsValueE, newQuadkeyAndVerticalID)

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
