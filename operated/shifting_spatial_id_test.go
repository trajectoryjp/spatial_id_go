package operated

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"testing"
	"time"
)

// TestGet6spatialIdsAdjacentToFaces01 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (空間ID:16/468/95/20/3)
//
// + 確認内容
//   - 入力の空間IDの面に直接接している6個の空間IDが返却されること
func TestGet6spatialIdsAdjacentToFaces01(t *testing.T) {
	resultVal := Get6spatialIdsAdjacentToFaces("16/468/95/20/3")

	expectVal := []string{"16/467/95/20/3", "16/469/95/20/3", "16/468/94/20/3", "16/468/96/20/3", "16/468/95/20/2", "16/468/95/20/4"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	t.Log("テスト終了")
}

// TestGet8spatialIdsAroundHorizontal01 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (空間ID:16/468/95/20/3)
//
// + 確認内容
//   - 入力の空間IDの水平方向の一周分の8個の空間IDが返却されること
func TestGet8spatialIdsAroundHorizontal01(t *testing.T) {
	resultVal := Get8spatialIdsAroundHorizontal("16/468/95/20/3")

	expectVal := []string{"16/467/95/20/3", "16/469/95/20/3", "16/468/94/20/3", "16/468/96/20/3", "16/469/94/20/3", "16/467/94/20/3", "16/467/96/20/3", "16/469/96/20/3"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	t.Log("テスト終了")
}

func TestGet24spatialIdsAroundHorizontal(t *testing.T) {
	resultVal := Get24spatialIdsAroundHorizontal("16/468/95/20/3")

	if float64(len(resultVal)) != 24 {
		t.Fatalf("N responses is not expected")
	} else {
		log.Printf("\nNumber of responses matches expected\n")
	}

	fmt.Println(resultVal)
}

// TestGet26spatialIdsAroundVoxel01 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (空間ID:16/468/95/20/3)
//
// + 確認内容
//   - 入力の空間IDを囲う26個の空間IDが返却されること
func TestGet26spatialIdsAroundVoxel01(t *testing.T) {
	resultVal := Get26spatialIdsAroundVoxel("16/468/95/20/3")

	expectVal := []string{"16/468/95/20/4", "16/469/95/20/4", "16/467/95/20/4", "16/468/96/20/4", "16/468/94/20/4", "16/469/94/20/4", "16/467/94/20/4", "16/469/96/20/4", "16/467/96/20/4",
		"16/469/95/20/3", "16/467/95/20/3", "16/468/96/20/3", "16/468/94/20/3", "16/469/94/20/3", "16/467/94/20/3", "16/469/96/20/3", "16/467/96/20/3",
		"16/468/95/20/2", "16/469/95/20/2", "16/467/95/20/2", "16/468/96/20/2", "16/468/94/20/2", "16/469/94/20/2", "16/467/94/20/2", "16/469/96/20/2", "16/467/96/20/2"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID01 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (空間ID:16/468/95/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID01(t *testing.T) {
	resultVal := GetShiftingSpatialID("16/468/95/20/3", 2, 3, 4)

	expectVal := "16/470/98/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID02 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が(2^精度-1)を超えている場合
//     (空間ID:5/30/29/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID02(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/30/29/20/3", 2, 3, 4)
	expectVal := "5/0/0/20/7"
	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID03 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が(2^精度-1)と等しい場合
//     (空間ID:5/29/28/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID03(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/29/28/20/3", 2, 3, 4)

	expectVal := "5/31/31/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID04 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が(2^精度-1)未満の場合
//     (空間ID:5/28/27/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID04(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/28/27/20/3", 2, 3, 4)

	expectVal := "5/30/30/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID05 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が0より大きい場合
//     (空間ID:5/-1/-2/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID05(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/-1/-2/20/3", 2, 3, 4)

	expectVal := "5/1/1/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID06 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が0と等しい場合
//     (空間ID:5/-2/-3/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID06(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/-2/-3/20/3", 2, 3, 4)

	expectVal := "5/0/0/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetShiftingSpatialID07 xおよびy方向インデックスの境界値確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     移動先の位置が0より小さい場合
//     (空間ID:5/-3/-4/20/3, 空間IDを経度方向に動かす数値:2, 空間IDを緯度方向に動かす数値:3, 空間IDを高さ方向に動かす数値:4)
//
// + 確認内容
//   - 入力の空間IDから指定の数値分移動した場合の空間IDが返却されること
func TestGetShiftingSpatialID07(t *testing.T) {
	resultVal := GetShiftingSpatialID("5/-3/-4/20/3", 2, 3, 4)

	expectVal := "5/31/31/20/7"

	//戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

func TestGet124spatialIdsAroundVoxcel(t *testing.T) {
	resultVal := Get124spatialIdsAroundVoxcel("16/468/95/20/3")

	if float64(len(resultVal)) != 124 {
		t.Fatalf("N responses is not expected")
	} else {
		log.Printf("\nNumber of responses matches expected\n")
	}
	fmt.Println(resultVal)

}

func TestGetNspatialIdsAroundVoxcel(t *testing.T) {

	nLayer := 2

	start := time.Now()
	resultVal, error := GetNspatialIdsAroundVoxcel("16/468/95/20/3", int64(nLayer), int64(nLayer))
	end := time.Since(start)
	if error != nil {
		t.Fatal(error)
	}

	expandParam := float64(nLayer * 2)
	if float64(len(resultVal)) != (math.Pow(float64(expandParam+1), 3) - 1) {
		t.Fatalf("N responses is not expected")
	} else {
		log.Printf("\nNumber of responses matches expected\n")
	}

	fmt.Println(len(resultVal))
	fmt.Println(end)

}

func TestGetNspatialIdsAroundVoxcel_time(t *testing.T) {

	times := []time.Duration{}
	nIds := []int{}

	var i int64
	for i = 1; i < 100; i += 1 {

		expandParam := float64(i * 2)

		start := time.Now()
		result, error := GetNspatialIdsAroundVoxcel("16/468/95/20/3", i, i)
		if error != nil {
			t.Fatal(error)
		}
		end := time.Since(start)

		if float64(len(result)) != (math.Pow(float64(expandParam+1), 3) - 1) {
			t.Fatalf("N responses is not expected")
		}

		times = append(times, end)
		nIds = append(nIds, len(result))

	}

	for i, time := range times {
		fmt.Printf("nLayer: %v\t nIds:%v\t %v\n", i+1, nIds[i], time)

	}
}

// stringスライスの中に指定文字列を含むか判定する
//
// 引数：
//
//	slice： stringスライス
//	target： 検索文字列
//
// 戻り値：
//
//	含む場合：true
//	含まない場合：false
func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
