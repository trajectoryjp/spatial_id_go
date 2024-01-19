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

func TestGetNspatialIdsAroundVoxcels(t *testing.T) {

	ids := []string{"23/7451610/3303539/23/25", "23/7451609/3303518/23/25", "23/7451606/3303477/23/25", "23/7451606/3303473/23/25", "23/7451605/3303455/23/25", "23/7451610/3303536/23/25", "23/7451610/3303531/23/25", "23/7451602/3303409/23/25", "23/7451609/3303515/23/25", "23/7451606/3303469/23/25", "23/7451608/3303505/23/25", "23/7451605/3303461/23/25", "23/7451605/3303456/23/25", "23/7451610/3303544/23/25", "23/7451610/3303541/23/25", "23/7451609/3303528/23/25", "23/7451609/3303524/23/25", "23/7451607/3303493/23/25", "23/7451602/3303418/23/25", "23/7451602/3303408/23/25", "23/7451610/3303543/23/25", "23/7451604/3303451/23/25", "23/7451604/3303436/23/25", "23/7451609/3303523/23/25", "23/7451609/3303526/23/25", "23/7451608/3303511/23/25", "23/7451607/3303491/23/25", "23/7451603/3303421/23/25", "23/7451602/3303410/23/25", "23/7451608/3303512/23/25", "23/7451608/3303509/23/25", "23/7451608/3303498/23/25", "23/7451604/3303435/23/25", "23/7451602/3303413/23/25", "23/7451606/3303480/23/25", "23/7451603/3303424/23/25", "23/7451605/3303453/23/25", "23/7451603/3303427/23/25", "23/7451603/3303428/23/25", "23/7451609/3303527/23/25", "23/7451608/3303499/23/25", "23/7451607/3303496/23/25", "23/7451607/3303489/23/25", "23/7451607/3303490/23/25", "23/7451608/3303507/23/25", "23/7451607/3303494/23/25", "23/7451606/3303466/23/25", "23/7451610/3303532/23/25", "23/7451609/3303521/23/25", "23/7451608/3303497/23/25", "23/7451603/3303429/23/25", "23/7451602/3303417/23/25", "23/7451608/3303503/23/25", "23/7451604/3303447/23/25", "23/7451604/3303445/23/25", "23/7451604/3303440/23/25", "23/7451604/3303449/23/25", "23/7451603/3303435/23/25", "23/7451603/3303432/23/25", "23/7451610/3303535/23/25", "23/7451605/3303460/23/25", "23/7451605/3303454/23/25", "23/7451605/3303452/23/25", "23/7451604/3303448/23/25", "23/7451610/3303540/23/25", "23/7451610/3303538/23/25", "23/7451606/3303475/23/25", "23/7451607/3303484/23/25", "23/7451607/3303485/23/25", "23/7451606/3303470/23/25", "23/7451606/3303468/23/25", "23/7451607/3303492/23/25", "23/7451605/3303462/23/25", "23/7451604/3303441/23/25", "23/7451609/3303514/23/25", "23/7451609/3303513/23/25", "23/7451605/3303466/23/25", "23/7451602/3303414/23/25", "23/7451606/3303467/23/25", "23/7451604/3303443/23/25", "23/7451604/3303438/23/25", "23/7451611/3303544/23/25", "23/7451610/3303534/23/25", "23/7451610/3303529/23/25", "23/7451608/3303504/23/25", "23/7451605/3303459/23/25", "23/7451609/3303519/23/25", "23/7451607/3303482/23/25", "23/7451606/3303471/23/25", "23/7451608/3303502/23/25", "23/7451608/3303501/23/25", "23/7451602/3303411/23/25", "23/7451603/3303434/23/25", "23/7451603/3303425/23/25", "23/7451602/3303416/23/25", "23/7451608/3303510/23/25", "23/7451608/3303500/23/25", "23/7451607/3303487/23/25", "23/7451605/3303458/23/25", "23/7451604/3303439/23/25", "23/7451602/3303419/23/25", "23/7451605/3303451/23/25", "23/7451603/3303433/23/25", "23/7451602/3303415/23/25", "23/7451610/3303533/23/25", "23/7451609/3303525/23/25", "23/7451609/3303520/23/25", "23/7451606/3303482/23/25", "23/7451604/3303450/23/25", "23/7451602/3303412/23/25", "23/7451609/3303522/23/25", "23/7451608/3303513/23/25", "23/7451608/3303506/23/25", "23/7451605/3303464/23/25", "23/7451603/3303422/23/25", "23/7451607/3303486/23/25", "23/7451604/3303446/23/25", "23/7451606/3303474/23/25", "23/7451604/3303437/23/25", "23/7451603/3303423/23/25", "23/7451602/3303420/23/25", "23/7451608/3303508/23/25", "23/7451603/3303426/23/25", "23/7451610/3303528/23/25", "23/7451607/3303495/23/25", "23/7451605/3303463/23/25", "23/7451603/3303420/23/25", "23/7451606/3303476/23/25", "23/7451609/3303516/23/25", "23/7451607/3303488/23/25", "23/7451604/3303442/23/25", "23/7451606/3303472/23/25", "23/7451605/3303465/23/25", "23/7451605/3303457/23/25", "23/7451610/3303542/23/25", "23/7451610/3303530/23/25", "23/7451607/3303497/23/25", "23/7451606/3303478/23/25", "23/7451606/3303479/23/25", "23/7451603/3303431/23/25", "23/7451603/3303430/23/25", "23/7451602/3303407/23/25", "23/7451610/3303537/23/25", "23/7451609/3303517/23/25", "23/7451607/3303483/23/25", "23/7451606/3303481/23/25", "23/7451604/3303444/23/25"}

	nLayer := 5

	start := time.Now()
	resultVal, error := GetNspatialIdsAroundVoxcels(ids, int64(nLayer), int64(nLayer))
	end := time.Since(start)
	if error != nil {
		t.Fatal(error)
	}

	fmt.Println(len(resultVal))
	fmt.Println(end)
}

func TestGetNspatialIdsAroundVoxcel_time(t *testing.T) {

	times := []time.Duration{}
	nIds := []int{}

	var i int64
	for i = 1; i < 50; i += 1 {

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

func TestGetShiftingSpatialIds(t *testing.T) {

	ids := []string{"23/7451603/3303422/23/25", "23/7451603/3303422/23/25", "23/7451605/3303422/23/25"}

	shiftedIds := GetShiftingSpatialIDs(ids, 1, 1, 1)
	if len(ids) == len(shiftedIds) {
		for i, v := range ids {
			log.Printf("id: %v -> %v", v, shiftedIds[i])
		}
	} else {
		log.Println(ids)
		log.Println(shiftedIds)
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
