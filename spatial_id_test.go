package spatialID

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewExtendedSpatialID01 拡張空間ID初期化関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestNewExtendedSpatialID01(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/232837/103222/0",
		&SpatialID{18, 232837, 103222, 0},
		nil,
	)
}

// TestNewExtendedSpatialID02 拡張空間ID初期化関数 異常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestNewExtendedSpatialID02(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/232837/103222",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID03 拡張空間ID再設定関数 入力された拡張空間IDの水平精度がしきい値を超過していた場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："36/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID03(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"36/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID04 拡張空間ID再設定関数 入力された拡張空間IDの水平精度がしきい値より小さい場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："-1/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID04(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"-1/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID05 拡張空間ID再設定関数 正常系動作確認(水平精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："35/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID05(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"35/232837/103222/0",
		&SpatialID{35, 232837, 103222, 0},
		nil,
	)
}

// TestResetExtendedSpatialID06 拡張空間ID再設定関数 正常系動作確認(水平精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："0/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID06(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"0/232837/103222/0",
		&SpatialID{0, 0, 0, 0},
		nil,
	)
}

// TestResetExtendedSpatialID07 拡張空間ID再設定関数 入力された拡張空間IDの垂直精度がしきい値を超過していた場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："25/232837/103222/36/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID07(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"36/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID08 拡張空間ID再設定関数 入力された拡張空間IDの垂直精度がしきい値より小さい場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："25/232837/103222/-1/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID08(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"-1/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID09 拡張空間ID再設定関数 正常系動作確認(垂直精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/35/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID09(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"35/232837/103222/0",
		&SpatialID{35, 232837, 103222, 0},
		nil,
	)
}

// TestResetExtendedSpatialID10 拡張空間ID再設定関数 正常系動作確認(垂直精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/0/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID10(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"0/0/0/-1",
		&SpatialID{0, 0, 0, 0},
		nil,
	)
}

// TestResetExtendedSpatialID11 拡張空間ID再設定関数 入力値に整数以外が存在した場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/2328A7/103222/0/0"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestResetExtendedSpatialID11(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/2328A7/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

func testNewSpatialIDFromString(
	t *testing.T,
	spatialIDString string,
	expectedSpatialID *SpatialID, // TODO: 順番の整理
	expectedError error,
	) {
	spatialID, error := NewSpatialIDFromString(spatialIDString)

	// 始点から終点へのベクトルと期待値の比較
	if !reflect.DeepEqual(spatialID, expectedSpatialID) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expectedSpatialID, spatialID)
	}
	if !reflect.DeepEqual(error, expectedError) {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectedError.Error(), error.Error())
	}
}

// TestSetX01 位置設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetX01(t *testing.T) {
	expected := &SpatialID{1, 0, 1, 0}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetX(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestSetY01 緯度ID設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetY01(t *testing.T) {
	expected := &SpatialID{1, 0, 0, 1}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetY(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestSetZ01 高さID設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetZ01(t *testing.T) {
	expected := &SpatialID{1, 1, 0, 0}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetF(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestX01 経度ID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、経度の値を取得できること。
func TestX01(t *testing.T) {
	expected := int64(3)

	spatialID, error := NewSpatialID(5, 2, 3, 4)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result := spatialID.GetX()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("経度ID - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestY01 緯度ID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、緯度の値を取得できること。
func TestY01(t *testing.T) {
	expected := int64(4)

	spatialID, error := NewSpatialID(5, 2, 3, 4)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result := spatialID.GetY()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("緯度ID - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestZ01 高さID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、高さの値を取得できること。
func TestZ01(t *testing.T) {
	expected := int64(2)

	spatialID, error := NewSpatialID(5, 2, 3, 4)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result := spatialID.GetF()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("高さID - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestHZoom01 水平精度取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、水平精度の値を取得できること。
func TestHZoom01(t *testing.T) {
	expected := int8(5)

	spatialID, error := NewSpatialID(5, 2, 3, 4)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result := spatialID.GetZ()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("精度 - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestID01 空間ID文字列返却関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、拡張空間ID文字列を取得できること。
func TestID01(t *testing.T) {
	expected := "5/2/3/4"

	spatialID, error := NewSpatialID(5, 2, 3, 4)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result := spatialID.String()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間ID文字列 - 期待値：%v, 取得値：%v", expected, result)
	}
}

func TestSpatialIDOverlaps(t *testing.T) {
	testCases := []struct {
		id             SpatialID
		another        SpatialID
		expectedResult bool
	}{
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 0, 0},
			expectedResult: true,
		},
		{
			id:             SpatialID{22, 0, 0, 0},
			another:        SpatialID{23, 1, 1, 1},
			expectedResult: true,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 1, 0, 0},
			expectedResult: false,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 1, 0},
			expectedResult: false,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 0, 1},
			expectedResult: false,
		},
		{
			id:             SpatialID{22, 1, 1, 1},
			another:        SpatialID{23, 1, 1, 1},
			expectedResult: false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		for swapIndex := range 2 {
			actualResult := testCase.id.Overlaps(testCase.another)
			assert.Equal(
				t,
				testCase.expectedResult,
				actualResult,
				fmt.Sprintf("testCaseIndex: %d, swapIndex: %d", testCaseIndex, swapIndex),
			)
		}
		testCase.id, testCase.another = testCase.another, testCase.id
	}
}

func TestMergeSpatialIDs(t *testing.T) {
	testCases := []struct {
		ids            []*SpatialID
		expectedResult []*SpatialID
	}{
		{
			ids: []*SpatialID{
				{23, 1, 2, 4},
			},
			expectedResult: []*SpatialID{
				{23, 1, 2, 4},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 0},
				{24, 1, 1, 1},
			},
			expectedResult: []*SpatialID{
				{23, 0, 0, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
			},
			expectedResult: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
				{23, 1, 1, 1},
				{23, 0, 1, 2},
			},
			expectedResult: []*SpatialID{
				{22, 0, 0, 0},
				{23, 0, 1, 2},
			},
		},
		{
			ids: []*SpatialID{
				{23, -1, 0, 0},
				{23, -1, 0, 1},
				{23, -1, 1, 0},
				{23, -1, 1, 1},
				{23, -2, 0, 0},
				{23, -2, 0, 1},
				{23, -2, 1, 0},
				{23, -2, 1, 1},
			},
			expectedResult: []*SpatialID{
				{22, -1, 0, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{24, 0, 0, 2},
				{24, 0, 0, 3},
				{24, 0, 1, 2},
				{24, 0, 1, 3},
				{24, 1, 0, 2},
				{24, 1, 0, 3},
				{24, 1, 1, 2},
				{24, 1, 1, 3},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
				{23, 1, 1, 1},
			},
			expectedResult: []*SpatialID{
				{22, 0, 0, 0},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		if testCaseIndex != 3 {
			continue
		}
		actualResult := MergeSpatialIDs(testCase.ids)
		assert.ElementsMatch(
			t,
			testCase.expectedResult,
			actualResult,
			fmt.Sprintf("testCaseIndex: %d", testCaseIndex),
		)
	}
}

func BenchmarkMergeSpatialIDs(b *testing.B) {
	type Benchmark struct {
		name string
		ids  []*SpatialID
	}

	benchmarks := []*Benchmark{}
	for _, n := range []int64{16} {
		for _, z := range []int8{23} {
			ids := []*SpatialID{}
			for f := range n {
				for x := range n {
					for y := range n {
						id, _ := NewSpatialID(z, f, x, y)
						ids = append(ids, id)
					}
				}
			}
			benchmarks = append(
				benchmarks,
				&Benchmark{fmt.Sprintf("n=%d,z=%d", n, z), ids},
			)
		}
	}

	for _, benchmark := range benchmarks {
		b.Run(
			benchmark.name,
			func(b *testing.B) {
				MergeSpatialIDs(benchmark.ids)
			},
		)

		startTotalAlloc := getTotalAlloc()
		MergeSpatialIDs(benchmark.ids)
		fmt.Printf("BenchmarkMergeSpatialIDs/%s\t%d B\n", benchmark.name, getTotalAlloc()-startTotalAlloc)
	}
}

func getTotalAlloc() uint64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem.TotalAlloc
}
