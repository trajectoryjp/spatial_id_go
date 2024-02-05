package integrate

import (
	"reflect"
	"testing"

	"github.com/trajectoryjp/spatial_id_go/v2/common/object"
)

// TestNewUnitDividedSpatialID01 単位分割拡張空間ID構造体のコンストラクタ 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{10, 1024, 2048, 10, 1024}, 水平方向精度差分：1, 垂直方向精度差分：2
//
// + 確認内容
//   - 入力値から単位分割拡張空間ID構造体のポインタを取得できること
func TestNewUnitDividedSpatialID01(t *testing.T) {
	//入力値
	id := "10/1024/2048/10/1024"
	EXSpatialID, _ := object.NewExtendedSpatialID(id)
	var hDiff int64 = 1
	var vDiff int64 = 2
	resultVal := NewUnitDividedSpatialID(EXSpatialID, hDiff, vDiff)

	//期待値
	expectVal := &UnitDividedSpatialID{}
	expectVal.ExtendedSpatialID = EXSpatialID
	expectVal.hDiff = 1
	expectVal.vDiff = 2
	expectVal.unitIDs = map[string]struct{}{
		"11/2048/4096/12/4096": {},
		"11/2048/4096/12/4097": {},
		"11/2048/4096/12/4098": {},
		"11/2048/4096/12/4099": {},
		"11/2048/4097/12/4096": {},
		"11/2048/4097/12/4097": {},
		"11/2048/4097/12/4098": {},
		"11/2048/4097/12/4099": {},
		"11/2049/4096/12/4096": {},
		"11/2049/4096/12/4097": {},
		"11/2049/4096/12/4098": {},
		"11/2049/4096/12/4099": {},
		"11/2049/4097/12/4096": {},
		"11/2049/4097/12/4097": {},
		"11/2049/4097/12/4098": {},
		"11/2049/4097/12/4099": {}}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("単位分割拡張空間ID構造体 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewHighSpatialID01 最適化後拡張空間ID構造体のコンストラクタ 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     単位分割拡張空間ID：{10, 1024, 2048, 10, 1024}, 水平方向精度差分：1, 垂直方向精度差分：2
//
// + 確認内容
//   - 入力値から最適化後拡張空間ID構造体のポインタを取得できること
func TestNewHighSpatialID01(t *testing.T) {
	//入力値
	id := "10/1024/2048/10/1024"
	EXSpatialID, _ := object.NewExtendedSpatialID(id)
	var hDiff int64 = 1
	var vDiff int64 = 2
	UDSpatialID := NewUnitDividedSpatialID(EXSpatialID, hDiff, vDiff)

	resultVal := NewHighSpatialID(UDSpatialID, hDiff, vDiff)

	//期待値
	expectVal := &HighSpatialID{}
	expectVal.ExtendedSpatialID = EXSpatialID.Higher(hDiff, vDiff)
	var threshold int64 = 256
	expectVal.threshold = threshold
	expectVal.lowIDs = []string{"10/1024/2048/10/1024"}
	expectVal.unitIDs = map[string]struct{}{
		"11/2048/4096/12/4096": {},
		"11/2048/4096/12/4097": {},
		"11/2048/4096/12/4098": {},
		"11/2048/4096/12/4099": {},
		"11/2048/4097/12/4096": {},
		"11/2048/4097/12/4097": {},
		"11/2048/4097/12/4098": {},
		"11/2048/4097/12/4099": {},
		"11/2049/4096/12/4096": {},
		"11/2049/4096/12/4097": {},
		"11/2049/4096/12/4098": {},
		"11/2049/4096/12/4099": {},
		"11/2049/4097/12/4096": {},
		"11/2049/4097/12/4097": {},
		"11/2049/4097/12/4098": {},
		"11/2049/4097/12/4099": {}}
	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("最適化後拡張空間ID構造体 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestMerge01 最適化後拡張空間ID構造体の結合 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     単位分割拡張空間ID：{10, 1024, 2048, 10, 1024}, 水平方向精度差分：1, 垂直方向精度差分：2
//
// + 確認内容
//   - 入力値から拡張空間IDが同一の最適化後拡張空間IDの最適化元・単位拡張空間IDのマージを取得できること
func TestMerge01(t *testing.T) {
	//入力値
	id := "10/1024/2048/10/1024"
	EXSpatialID, _ := object.NewExtendedSpatialID(id)
	var hDiff int64 = 1
	var vDiff int64 = 2
	UDSpatialID := NewUnitDividedSpatialID(EXSpatialID, hDiff, vDiff)

	resultVal := NewHighSpatialID(UDSpatialID, hDiff, vDiff)

	resultVal.Merge(resultVal)

	//期待値
	expectVal := &HighSpatialID{}
	expectVal.ExtendedSpatialID = EXSpatialID.Higher(hDiff, vDiff)
	var threshold int64 = 256
	expectVal.threshold = threshold
	expectVal.lowIDs = []string{"10/1024/2048/10/1024", "10/1024/2048/10/1024"}
	expectVal.unitIDs = map[string]struct{}{
		"11/2048/4096/12/4096": {},
		"11/2048/4096/12/4097": {},
		"11/2048/4096/12/4098": {},
		"11/2048/4096/12/4099": {},
		"11/2048/4097/12/4096": {},
		"11/2048/4097/12/4097": {},
		"11/2048/4097/12/4098": {},
		"11/2048/4097/12/4099": {},
		"11/2049/4096/12/4096": {},
		"11/2049/4096/12/4097": {},
		"11/2049/4096/12/4098": {},
		"11/2049/4096/12/4099": {},
		"11/2049/4097/12/4096": {},
		"11/2049/4097/12/4097": {},
		"11/2049/4097/12/4098": {},
		"11/2049/4097/12/4099": {}}
	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDが同一の最適化後拡張空間IDの最適化元・単位拡張空間IDのマージ - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestIsDense01 最適化後拡張空間ID構造体が単位拡張空間ID集合の稠密判定 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     単位分割拡張空間ID：{10, 1024, 2048, 10, 1024}, 水平方向精度差分：1, 垂直方向精度差分：2,
//     単位拡張空間IDの個数の閾値：16
//
// + 確認内容
//   - 入力値から最適化後拡張空間ID構造体が単位拡張空間ID集合の稠密判定の結果を取得できること
func TestIsDense01(t *testing.T) {
	//入力値
	id := "10/1024/2048/10/1024"
	EXSpatialID, _ := object.NewExtendedSpatialID(id)
	var hDiff int64 = 1
	var vDiff int64 = 2

	HSpatialID := &HighSpatialID{}
	HSpatialID.ExtendedSpatialID = EXSpatialID.Higher(hDiff, vDiff)
	var threshold int64 = 16
	HSpatialID.threshold = threshold
	HSpatialID.lowIDs = []string{"10/1024/2048/10/1024"}
	HSpatialID.unitIDs = map[string]struct{}{
		"11/2048/4096/12/4096": {},
		"11/2048/4096/12/4097": {},
		"11/2048/4096/12/4098": {},
		"11/2048/4096/12/4099": {},
		"11/2048/4097/12/4096": {},
		"11/2048/4097/12/4097": {},
		"11/2048/4097/12/4098": {},
		"11/2048/4097/12/4099": {},
		"11/2049/4096/12/4096": {},
		"11/2049/4096/12/4097": {},
		"11/2049/4096/12/4098": {},
		"11/2049/4096/12/4099": {},
		"11/2049/4097/12/4096": {},
		"11/2049/4097/12/4097": {},
		"11/2049/4097/12/4098": {},
		"11/2049/4097/12/4099": {}}

	resultVal := HSpatialID.IsDense()

	//期待値
	expectVal := true
	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("稠密判定 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestIsDense02 最適化後拡張空間ID構造体が単位拡張空間ID集合の稠密判定 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     単位分割拡張空間ID：{10, 1024, 2048, 10, 1024}, 水平方向精度差分：1, 垂直方向精度差分：2,
//     単位拡張空間IDの個数の閾値：256
//
// + 確認内容
//   - 入力値から最適化後拡張空間ID構造体が単位拡張空間ID集合の稠密判定の結果を取得できること
func TestIsDense02(t *testing.T) {
	//入力値
	id := "10/1024/2048/10/1024"
	EXSpatialID, _ := object.NewExtendedSpatialID(id)
	var hDiff int64 = 1
	var vDiff int64 = 2

	HSpatialID := &HighSpatialID{}
	HSpatialID.ExtendedSpatialID = EXSpatialID.Higher(hDiff, vDiff)
	var threshold int64 = 256
	HSpatialID.threshold = threshold
	HSpatialID.lowIDs = []string{"10/1024/2048/10/1024"}
	HSpatialID.unitIDs = map[string]struct{}{
		"11/2048/4096/12/4096": {},
		"11/2048/4096/12/4097": {},
		"11/2048/4096/12/4098": {},
		"11/2048/4096/12/4099": {},
		"11/2048/4097/12/4096": {},
		"11/2048/4097/12/4097": {},
		"11/2048/4097/12/4098": {},
		"11/2048/4097/12/4099": {},
		"11/2049/4096/12/4096": {},
		"11/2049/4096/12/4097": {},
		"11/2049/4096/12/4098": {},
		"11/2049/4096/12/4099": {},
		"11/2049/4097/12/4096": {},
		"11/2049/4097/12/4097": {},
		"11/2049/4097/12/4098": {},
		"11/2049/4097/12/4099": {}}

	resultVal := HSpatialID.IsDense()

	//期待値
	expectVal := false
	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("稠密判定 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds01 空間IDの最適化（マージ） 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"15/0/1024/2048"}, マージ後の精度：11
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeSpatialIds01(t *testing.T) {
	//入力値
	SpatialIDs := []string{"15/0/1024/2048"}
	var zoom int64 = 11
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{"15/0/1024/2048"}

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
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds02 空間IDの最適化（マージ） 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"21/512/1024/2048", "21/512/1024/2049", "21/512/1024/2050", "21/512/1024/2048"},
//     マージ後の精度：11
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeSpatialIds02(t *testing.T) {
	//入力値
	SpatialIDs := []string{"21/512/1024/2048", "21/512/1024/2049", "21/512/1024/2050", "21/512/1024/2048"}
	var zoom int64 = 11
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{"21/512/1024/2048", "21/512/1024/2049", "21/512/1024/2050"}

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
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds03 空間IDの最適化（マージ） 空入力時動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{(空入力))}, マージ後の精度：11
//
// + 確認内容
//   - 空の配列を取得できること
func TestMergeSpatialIds03(t *testing.T) {
	//入力値
	SpatialIDs := []string{}
	var zoom int64 = 11
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{}

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
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds04 空間IDの最適化（マージ） 空間ID内の精度がマージ後精度より低い場合の動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"10/0/1024/2048", "12/0/1024/2048", "9/0/1024/2048"}, マージ後の精度：11
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeSpatialIds04(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/0/1024/2048", "12/0/1024/2048", "9/0/1024/2048"}
	var zoom int64 = 11
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{"10/0/1024/2048", "12/0/1024/2048", "9/0/1024/2048"}

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
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds05 空間IDの最適化（マージ） 空間IDがフォーマット不正の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"15/0/1024/2048/777"}, マージ後の精度：11
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeSpatialIds05(t *testing.T) {
	//入力値
	SpatialIDs := []string{"15/0/1024/2048/777"}
	var zoom int64 = 11
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

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
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds06 空間IDの最適化（マージ） マージ後の精度が不正の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"15/0/1024/2048"}, マージ後の精度：36
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeSpatialIds06(t *testing.T) {
	//入力値
	SpatialIDs := []string{"15/0/1024/2048"}
	var zoom int64 = 36
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

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
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMergeSpatialIds07 空間IDの最適化（マージ） 空間IDフォーマットが不正の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     空間ID文字列配列：{"15|0|1024|2048"}, マージ後の精度：9
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeSpatialIds07(t *testing.T) {
	//入力値
	SpatialIDs := []string{"15|0|1024|2048"}
	var zoom int64 = 9
	resultVal, resultErr := MergeSpatialIds(SpatialIDs, zoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

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
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds01 拡張空間IDの最適化（マージ） 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds01(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"10/1024/2048/10/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds02 拡張空間IDの最適化（マージ） 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024", "10/1024/2048/10/1024", "11/1024/2048/11/1024", "9/1024/2048/9/1024", "11/1024/2048/11/1024"},
//     マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds02(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024", "10/1024/2048/10/1024", "11/1024/2048/11/1024", "9/1024/2048/9/1024", "11/1024/2048/11/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"10/1024/2048/10/1024", "11/1024/2048/11/1024", "9/1024/2048/9/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds03 拡張空間IDの最適化（マージ） 空入力時動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{(空入力)}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 空配列を取得できること
func TestMergeExtendedSpatialIds03(t *testing.T) {
	//入力値
	SpatialIDs := []string{}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds04 拡張空間IDの最適化（マージ） 拡張空間IDの垂直精度のみ低い場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/8/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds04(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/8/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"10/1024/2048/8/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds05 拡張空間IDの最適化（マージ） 拡張空間IDの水平精度のみ低い場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"8/1024/2048/10/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds05(t *testing.T) {
	//入力値
	SpatialIDs := []string{"8/1024/2048/10/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"8/1024/2048/10/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds06 拡張空間IDの最適化（マージ） 拡張空間IDの垂直精度水平精度の両方が低い場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"8/1024/2048/8/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds06(t *testing.T) {
	//入力値
	SpatialIDs := []string{"8/1024/2048/8/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"8/1024/2048/8/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds07 拡張空間IDの最適化（マージ） マージ後の水平方向精度の境界値確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024"}, マージ後の水平方向精度：35, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds07(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024"}
	var hzoom int64 = 35
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"10/1024/2048/10/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds08 拡張空間IDの最適化（マージ） マージ後の水平方向精度の境界値確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024"}, マージ後の水平方向精度：36, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeExtendedSpatialIds08(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024"}
	var hzoom int64 = 36
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds09 拡張空間IDの最適化（マージ） マージ後の垂直方向精度の境界値確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：35
//
// + 確認内容
//   - 入力値からマージ後の空間IDを格納した配列を取得できること
func TestMergeExtendedSpatialIds09(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 35
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{"10/1024/2048/10/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds10 拡張空間IDの最適化（マージ） マージ後の垂直方向精度の境界値確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10/1024/2048/10/1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：36
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeExtendedSpatialIds10(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10/1024/2048/10/1024"}
	var hzoom int64 = 9
	var vzoom int64 = 36
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMergeExtendedSpatialIds11 拡張空間IDの最適化（マージ） 拡張空間IDフォーマット不正
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID文字列配列：{"10|1024|2048|10|1024"}, マージ後の水平方向精度：9, マージ後の垂直方向精度：9
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestMergeExtendedSpatialIds11(t *testing.T) {
	//入力値
	SpatialIDs := []string{"10|1024|2048|10|1024"}
	var hzoom int64 = 9
	var vzoom int64 = 9
	resultVal, resultErr := MergeExtendedSpatialIds(SpatialIDs, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}
