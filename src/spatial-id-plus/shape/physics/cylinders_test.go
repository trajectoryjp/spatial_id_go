package shape

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"spatial-id/common"
	"spatial-id/common/consts"
	"spatial-id/common/enum"
	"spatial-id/common/object"
	"spatial-id/common/spatial"
	"spatial-id/shape/physics"
	"testing"
)

// TestConvertSpatialProjectionPoint01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - Point3型オブジェクト：{1.0, 2.0, 3.0}
//
// + 確認内容
//   - convertSpatialProjectionPoint Point3型オブジェクトからprojectedPoint型オブジェクトへ変換されること
func TestConvertSpatialProjectionPoint01(t *testing.T) {
	// Point3型オブジェクト
	p := spatial.Point3{1.0, 2.0, 3.0}

	expectVal := object.ProjectedPoint{1.0, 2.0, 3.0}
	resultVal := convertSpatialProjectionPoint(p)

	//戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("オブジェクト - 期待値%v, 取得値%v", expectVal, resultVal)
	}

	// 戻り値と期待値の型の比較
	// 型が変換されていない場合Errorをログに出力
	if reflect.DeepEqual(reflect.TypeOf(p), reflect.TypeOf(resultVal)) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v", reflect.TypeOf(expectVal), reflect.TypeOf(p))
	}

	t.Log("テスト終了")
}

// TestConvertSpatialProjectionPoint02 異常系動作確認(引数に空の構造体オブジェクトを指定)
//
// 試験詳細：
// + 試験データ
//   - Point3型オブジェクト：{}
//
// + 確認内容
//   - convertSpatialProjectionPoint Point3型オブジェクトからprojectedPoint型オブジェクトへ変換されること
func TestConvertSpatialProjectionPoint02(t *testing.T) {
	// Point3型オブジェクト
	p := spatial.Point3{}

	expectVal := object.ProjectedPoint{}
	resultVal := convertSpatialProjectionPoint(p)

	//戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("オブジェクト - 期待値%v, 取得値%v", expectVal, resultVal)
	}

	// 戻り値と期待値の型の比較
	// 型が変換されていない場合Errorをログに出力
	if reflect.DeepEqual(reflect.TypeOf(p), reflect.TypeOf(resultVal)) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v", reflect.TypeOf(expectVal), reflect.TypeOf(p))
	}

	t.Log("テスト終了")
}

// TestGetSpatialIDOnAxisIDs01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//  xID: 10
//  yID: 20
//  zID: 30
//  hZoom: 15
//  vZoom: 25
//
// + 確認内容
//   - 軸IDから拡張空間IDを取得すること
func TestGetSpatialIDOnAxisIDs01(t *testing.T) {
	//入力パラメータ
	var xID, yID, zID, hZoom, vZoom int64 = 10, 20, 30, 15, 25

	// 期待値
	expectVal := "15/10/20/25/30"

	// テスト対象呼び出し
	resultVal := GetSpatialIDOnAxisIDs(xID, yID, zID, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetVoxelIDToSpatialID01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   拡張空間ID： 25/200/29803148/13212522
//
// + 確認内容
//   - 拡張空間IDに対応するボクセル成分IDを取得すること
func TestGetVoxelIDToSpatialID01(t *testing.T) {
	//入力パラメータ
	spatialID := "25/200/29803148/25/0"

	// 期待値
	expectVal := []int64{200, 29803148, 0}

	// テスト対象呼び出し
	resultVal := GetVoxelIDToSpatialID(spatialID)

	// 戻り値のボクセル成分IDと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値のボクセル成分IDが期待値と異なる場合Errorをログに出力
		t.Errorf("ボクセル成分ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewRectangular01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   始点： (1,2,3)
//   終点： (5,6,7)
//   半径： 2.0
//   水平精度： 2
//   垂直精度： 5
//   Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 入力値に対応する直方体構造体ポインタが返されること
func TestNewRectangular01(t *testing.T) {
	//入力パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{5, 6, 7}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	factor := 2.0

	// 期待値
	expectVal := new(Rectangular)
	expectVal.start = start
	expectVal.end = end
	expectVal.radius = radius
	expectVal.hZoom = hZoom
	expectVal.vZoom = vZoom
	expectVal.factor = factor
	// 【直交座標空間】軸ベクトル
	axis := spatial.NewVectorFromPoints(expectVal.start, expectVal.end)
	// 縦の長さ
	expectVal.height = axis.Norm()

	// テスト対象呼び出し
	resultVal := NewRectangular(start, end, radius, hZoom, vZoom, factor)

	// 戻り値の直方体構造体ポインタと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の直方体構造体ポインタが期待値と異なる場合Errorをログに出力
		t.Errorf("直方体構造体 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestcalcUnitVoxelVector01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   パターン1
//    拡張空間ID：25/200/29803148/25/0
//
// + 確認内容
//   - 入力値に対応する概算距離用の単位ボクセルベクトルが返却されること
func TestCalcUnitVoxelVector01(t *testing.T) {
	// 拡張空間ID
	spatialID := "25/200/29803148/25/0"
	//初期化用入力パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{4, 5, 6}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	factor := 1.15473441108545

	// 入力パラメータ初期化
	rect := NewRectangular(start, end, radius, hZoom, vZoom, factor)

	points, _ := GetPointOnExtendedSpatialId(
		spatialID,
		enum.Vertex,
	)

	// 期待値
	orthPoints, _ := ConvertPointListToProjectedPointList(points, consts.OrthCrs)

	// 各座標リストの最大値と最小値を取得(PrintしてX,Y,Altの最大となる確認して直接代入)
	// &{-2.0037269477075852e+07 -1.5557242698819885e+07 0}
	// &{-2.0037268282747287e+07 -1.5557242698819885e+07 0}
	// &{-2.0037268282747287e+07 -1.5557243893136082e+07 0} Min要素
	// &{-2.0037269477075852e+07 -1.5557243893136082e+07 0}
	// &{-2.0037269477075852e+07 -1.5557242698819885e+07 1}
	// &{-2.0037268282747287e+07 -1.5557242698819885e+07 1} Max要素
	// &{-2.0037268282747287e+07 -1.5557243893136082e+07 1}
	// &{-2.0037269477075852e+07 -1.5557243893136082e+07 1}
	fmt.Printf("%v\n", orthPoints[0])
	fmt.Printf("%v\n", orthPoints[1])
	fmt.Printf("%v\n", orthPoints[2])
	fmt.Printf("%v\n", orthPoints[3])
	fmt.Printf("%v\n", orthPoints[4])
	fmt.Printf("%v\n", orthPoints[5])
	fmt.Printf("%v\n", orthPoints[6])
	fmt.Printf("%v\n", orthPoints[7])
	maxX := orthPoints[5].X
	minX := orthPoints[3].X
	maxY := orthPoints[5].Y
	minY := orthPoints[3].Y
	maxZ := orthPoints[5].Alt
	minZ := orthPoints[3].Alt

	vect1 := spatial.Vector3{
		X: (maxX - minX),
		Y: (maxY - minY),
		Z: (maxZ - minZ) * rect.factor,
	}

	// テスト対象呼び出し
	vect2, _ := rect.calcUnitVoxelVector(spatialID)

	// 小数点6桁以下を切り捨てているため、6桁までで比較
	if !common.AlmostEqual(vect1.X, vect2.X, 1e-6) {
		// 戻り値の単位ボクセルベクトルが期待値と異なる場合Errorをログに出力
		t.Errorf("単位ボクセルベクトル - 期待値：%v, 取得値：%v", vect1.X, vect2.X)
	}
	if !common.AlmostEqual(vect1.Y, vect2.Y, 1e-6) {
		// 戻り値の単位ボクセルベクトルが期待値と異なる場合Errorをログに出力
		t.Errorf("単位ボクセルベクトル - 期待値：%v, 取得値：%v", vect1.Y, vect2.Y)
	}
	if !common.AlmostEqual(vect1.Z, vect2.Z, 1e-6) {
		// 戻り値の単位ボクセルベクトルが期待値と異なる場合Errorをログに出力
		t.Errorf("単位ボクセルベクトル - 期待値：%v, 取得値：%v", vect1.Z, vect2.Z)
	}

	t.Log("テスト終了")
}

// TestCalcUnitVoxelVector02 異常系動作確認(不正な拡張空間ID)
//
// 試験詳細：
// + 試験データ
//   パターン2
//    拡張空間ID：25/200/29803148/25
//
// + 確認内容
//   - 不正なフォーマットの拡張空間IDを指定した場合にエラーが発生すること
func TestCalcUnitVoxelVector02(t *testing.T) {
	// 拡張空間ID
	spatialID := "25/200/29803148/25"
	//初期化用入力パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{4, 5, 6}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	factor := 1.15473441108545

	// 入力パラメータ初期化
	rect := NewRectangular(start, end, radius, hZoom, vZoom, factor)
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	_, resultErr := rect.calcUnitVoxelVector(spatialID)

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestNewCapsule01 正常系動作確認(球の場合)
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： (1,2,3)
//    終点： (1,2,3)
//    半径： 2.0
//    水平精度： 2
//    垂直精度： 5
//    カプセル判定： false
//    衝突判定実施オプション：true
//   Webメルカトル換算係数： 1.15473441108545
//
// + 確認内容
//   - 入力に応じた構造体ポインタを取得できること
//   - 戻り値の型がsphereであること
func TestNewCapsule01(t *testing.T) {
	// 初期化用パラメータ
	start := spatial.Point3{-1, -2, -3}
	end := spatial.Point3{-1, -2, -3}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	isCapsule := false
	isPrecision := true
	factor := 1.15473441108545

	// 期待されるカプセル構造体(球)
	var object physics.Physics
	object = physics.NewSpherePhysics(radius, start)
	isSphere := true
	expectVal := new(Capsule)
	expectVal.Rectangular = NewRectangular(start, end, radius, hZoom, vZoom, factor)
	expectVal.isCapsule = isCapsule
	expectVal.isSphere = isSphere
	expectVal.isPrecision = isPrecision
	expectVal.object = object

	// テスト対象呼び出し
	resultVal := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 戻り値の構造体ポインタと期待値の比較
	// 型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectVal), reflect.TypeOf(resultVal)) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectVal), reflect.TypeOf(resultVal))
	}
	// 直方体構造体
	if !reflect.DeepEqual(expectVal.Rectangular, resultVal.Rectangular) {
		t.Errorf("直方体構造体 - 期待値：%v, 取得値：%v",
			expectVal.Rectangular, resultVal.Rectangular)
	}
	// カプセル判定
	if !reflect.DeepEqual(expectVal.isCapsule, resultVal.isCapsule) {
		t.Errorf("カプセル判定 - 期待値：%v, 取得値：%v",
			expectVal.isCapsule, resultVal.isCapsule)
	}
	// 球判定
	if !reflect.DeepEqual(expectVal.isSphere, resultVal.isSphere) {
		t.Errorf("球判定 - 期待値：%v, 取得値：%v",
			expectVal.isSphere, resultVal.isSphere)
	}
	// 衝突判定実施オプション
	if !reflect.DeepEqual(expectVal.isPrecision, resultVal.isPrecision) {
		t.Errorf("衝突判定実施オプション - 期待値：%v, 取得値：%v",
			expectVal.isPrecision, resultVal.isPrecision)
	}
	// 衝突判定オブジェクトはインスタンスごとに値が変わるため、
	// objectについてはNewSpherePhysics()で試験することとする
	// if !reflect.DeepEqual(expectVal.object, resultVal.object) {
	// 	t.Errorf("衝突判定オブジェクト - 期待値：%v, 取得値：%v",
	// 		expectVal.object, resultVal.object)
	// }
	// 内部判定空間ID
	if !reflect.DeepEqual(expectVal.includeSpatialIDs, resultVal.includeSpatialIDs) {
		t.Errorf("内部判定空間ID - 期待値：%v, 取得値：%v",
			expectVal.includeSpatialIDs, resultVal.includeSpatialIDs)
	}
	t.Log("テスト終了")
}

// TestNewCapsule02 正常系動作確認(カプセルの場合)
//
// 試験詳細：
// + 試験データ
//   パターン2
//    始点： (1,2,3)
//    終点： (4,5,6)
//    半径： 2.0
//    水平精度： 2
//    垂直精度： 5
//    カプセル判定： true
//    衝突判定実施オプション：false
//    Webメルカトル換算係数： 1.15473441108545
//
// + 確認内容
//   - 入力に応じたカプセル構造体ポインタを取得できること
//   - 衝突判定オブジェクトは入力値に関係なくインスタンス化のたびに値が変更されるため、比較の対象外とする
func TestNewCapsule02(t *testing.T) {
	// 初期化用パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{4, 5, 6}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	isSphere := false
	isCapsule := true
	isPrecision := false
	factor := 1.15473441108545

	// 期待されるカプセル構造体(カプセル)
	var object physics.Physics
	object = physics.NewCapsulePhysics(radius, start, end)
	expectVal := new(Capsule)
	expectVal.Rectangular = NewRectangular(start, end, radius, hZoom, vZoom, factor)
	expectVal.isCapsule = isCapsule
	expectVal.isSphere = isSphere
	expectVal.isPrecision = isPrecision
	expectVal.object = object

	// テスト対象呼び出し
	resultVal := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 戻り値の構造体ポインタと期待値の比較
	// 型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectVal), reflect.TypeOf(resultVal)) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectVal), reflect.TypeOf(resultVal))
	}
	// 直方体構造体
	if !reflect.DeepEqual(expectVal.Rectangular, resultVal.Rectangular) {
		t.Errorf("直方体構造体 - 期待値：%v, 取得値：%v",
			expectVal.Rectangular, resultVal.Rectangular)
	}
	// カプセル判定
	if !reflect.DeepEqual(expectVal.isCapsule, resultVal.isCapsule) {
		t.Errorf("カプセル判定 - 期待値：%v, 取得値：%v",
			expectVal.isCapsule, resultVal.isCapsule)
	}
	// 球判定
	if !reflect.DeepEqual(expectVal.isSphere, resultVal.isSphere) {
		t.Errorf("球判定 - 期待値：%v, 取得値：%v",
			expectVal.isSphere, resultVal.isSphere)
	}
	// 衝突判定実施オプション
	if !reflect.DeepEqual(expectVal.isPrecision, resultVal.isPrecision) {
		t.Errorf("衝突判定実施オプション - 期待値：%v, 取得値：%v",
			expectVal.isPrecision, resultVal.isPrecision)
	}
	// 衝突判定オブジェクトはインスタンスごとに値が変わるため、
	// objectについてはNewCapsulePhysics()で試験することとする
	// if !reflect.DeepEqual(expectVal.object, resultVal.object) {
	// 	t.Errorf("衝突判定オブジェクト - 期待値：%v, 取得値：%v",
	// 		expectVal.object, resultVal.object)
	// }
	// 内部判定空間ID
	if !reflect.DeepEqual(expectVal.includeSpatialIDs, resultVal.includeSpatialIDs) {
		t.Errorf("内部判定空間ID - 期待値：%v, 取得値：%v",
			expectVal.includeSpatialIDs, resultVal.includeSpatialIDs)
	}

	t.Log("テスト終了")
}

// TestNewCapsule03 正常系動作確認(円柱の場合)
//
// 試験詳細：
// + 試験データ
//   パターン3
//    始点： (1,2,3)
//    終点： (4,5,6)
//    半径： 2.0
//    水平精度： 2
//    垂直精度： 5
//    カプセル判定： false
//    衝突判定実施オプション：false
//    Webメルカトル換算係数： 1.15473441108545
//
// + 確認内容
//   - 入力に応じた構造体ポインタを取得できること
//   - 衝突判定オブジェクトは入力値に関係なくインスタンス化のたびに値が変更されるため、比較の対象外とする
func TestNewCapsule03(t *testing.T) {
	// 初期化用パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{4, 5, 6}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	isCapsule := false
	isPrecision := false
	factor := 1.15473441108545

	// 期待されるカプセル構造体(円柱)
	var object physics.Physics
	object = physics.NewCylinderPhysics(radius, start, end)
	expectVal := new(Capsule)
	expectVal.Rectangular = NewRectangular(start, end, radius, hZoom, vZoom, factor)
	expectVal.isCapsule = isCapsule
	expectVal.isSphere = false
	expectVal.isPrecision = isPrecision
	expectVal.object = object

	// テスト対象呼び出し
	resultVal := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 戻り値の構造体ポインタと期待値の比較
	// 型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectVal), reflect.TypeOf(resultVal)) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectVal), reflect.TypeOf(resultVal))
	}
	// 直方体構造体
	if !reflect.DeepEqual(expectVal.Rectangular, resultVal.Rectangular) {
		t.Errorf("直方体構造体 - 期待値：%v, 取得値：%v",
			expectVal.Rectangular, resultVal.Rectangular)
	}
	// カプセル判定
	if !reflect.DeepEqual(expectVal.isCapsule, resultVal.isCapsule) {
		t.Errorf("カプセル判定 - 期待値：%v, 取得値：%v",
			expectVal.isCapsule, resultVal.isCapsule)
	}
	// 球判定
	if !reflect.DeepEqual(expectVal.isSphere, resultVal.isSphere) {
		t.Errorf("球判定 - 期待値：%v, 取得値：%v",
			expectVal.isSphere, resultVal.isSphere)
	}
	// 衝突判定実施オプション
	if !reflect.DeepEqual(expectVal.isPrecision, resultVal.isPrecision) {
		t.Errorf("衝突判定実施オプション - 期待値：%v, 取得値：%v",
			expectVal.isPrecision, resultVal.isPrecision)
	}
	// 衝突判定オブジェクトはインスタンスごとに値が変わるため、
	// objectについてはNewCylinderPhysics()で試験することとする
	// if !reflect.DeepEqual(expectVal.object, resultVal.object) {
	// 	t.Errorf("衝突判定オブジェクト - 期待値：%v, 取得値：%v",
	// 		expectVal.object, resultVal.object)
	// }
	// 内部判定空間ID
	if !reflect.DeepEqual(expectVal.includeSpatialIDs, resultVal.includeSpatialIDs) {
		t.Errorf("内部判定空間ID - 期待値：%v, 取得値：%v",
			expectVal.includeSpatialIDs, resultVal.includeSpatialIDs)
	}

	t.Log("テスト終了")
}

// TestIsEmpty01 正常系動作確認(オブジェクトが空でない)
//
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： (1,2,3)
//    終点： (4,5,6)
//    半径： 2.0
//    水平精度： 2
//    垂直精度： 5
//    カプセル判定： false
//    衝突判定実施オプション：false
//
// + 確認内容
//   - オブジェクトの状態に対応した真偽値(false)が返ってくること
func TestIsEmpty01(t *testing.T) {
	// 初期化用パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{4, 5, 6}
	radius := 2.0
	hZoom := int64(2)
	vZoom := int64(5)
	isCapsule := false
	isPrecision := false
	factor := 2.0
	// カプセルオブジェクト初期化
	cap := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := cap.IsEmpty()

	//戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("オブジェクトの比較結果 - 期待値%v, 取得値%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestIsEmpty02 正常系動作確認(オブジェクトが空)
//
// 試験詳細：
// + 試験データ
//   パターン2
//     空のカプセルオブジェクト
//
// + 確認内容
//   - オブジェクトの状態に対応した真偽値(true)が返ってくること
func TestIsEmpty02(t *testing.T) {
	// 空のカプセルオブジェクト初期化
	cap := new(Capsule)

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := cap.IsEmpty()

	//戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("オブジェクトの比較結果 - 期待値%v, 取得値%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs01 正常系動作確認(球の場合)
//
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： 任意
//    終点： 始点と同一
//    半径： 2.0
//    水平精度： 26
//    垂直精度： 27
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 1.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs01(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(26)
	vZoom := int64(27)
	isCapsule := false
	isPrecision := true
	factor := 1.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点座標
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnPoints([]*object.Point{start}, hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs02 正常系動作確認(球の場合: Webメルカトル係数補正が1以外)
//
// 試験詳細：
// + 試験データ
//   パターン2
//    始点： 任意
//    終点： 始点と同一
//    半径： 2.0
//    水平精度： 27
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs02(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, -35.685371, 11.0)
	end, _ := object.NewPoint(139.753098, -35.685371, 11.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(27)
	vZoom := int64(26)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点座標、補正有
	start.SetAlt(start.Alt() / factor)
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnPoints([]*object.Point{start}, hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs03 異常系動作確認(球の場合)
//
// 試験詳細：
// + 試験データ
//   パターン3
//    始点： 任意
//    終点： 始点と同一
//    半径： 2.0
//    水平精度： 36(不正値)
//    垂直精度： 25
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 戻り値の空間IDのリストが二つとも空であること
//   - 戻り値としてエラー内容が返却されること
func TestCalcLineSpatialIDs03(t *testing.T) {
	// 期待値
	expectErr := "InputValueError,入力チェックエラー"

	//入力パラメータ
	start, _ := object.NewPoint(-139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(-139.753098, 35.685371, 11.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(36)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// テスト対象呼び出し
	result1, result2, resultErr := cap.calcLineSpatialIDs()

	// 空間IDが返却されていないこと
	if len(result1) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(result1))
	}

	// 空間IDが返却されていないこと
	if len(result2) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の内部空間用軸の空間ID: %v", len(result2))
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs04 正常系動作確認(カプセルの場合)
//
// 試験詳細：
// + 試験データ
//   パターン4
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 26
//    垂直精度： 27
//    カプセル判定： true
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 1.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs04(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(26)
	vZoom := int64(27)
	isCapsule := true
	isPrecision := true
	factor := 1.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点終点から取得できる線分APIの空間ID結果
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expecetedLineSpatialIDs)
	sort.Strings(lineSpatialIDs)
	sort.Strings(insideLineSpatialIDs)

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs05 正常系動作確認(カプセルの場合: Webメルカトル係数補正が1以外)
//
// 試験詳細：
// + 試験データ
//   パターン5
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 27
//    垂直精度： 26
//    カプセル判定： true
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs05(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(27)
	vZoom := int64(26)
	isCapsule := true
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点終点から取得できる線分APIの空間ID結果
	end.SetAlt(end.Alt() / factor)
	start.SetAlt(start.Alt() / factor)
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expecetedLineSpatialIDs)
	sort.Strings(lineSpatialIDs)
	sort.Strings(insideLineSpatialIDs)

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs06 異常系動作確認(カプセルの場合)
//
// 試験詳細：
// + 試験データ
//   パターン6
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 36(不正値)
//    垂直精度： 25
//    カプセル判定： true
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 戻り値の空間IDのリストが二つとも空であること
//   - 戻り値としてエラー内容が返却されること
func TestCalcLineSpatialIDs06(t *testing.T) {
	// 期待値
	expectErr := "InputValueError,入力チェックエラー"

	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(36)
	vZoom := int64(25)
	isCapsule := true
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// テスト対象呼び出し
	result1, result2, resultErr := cap.calcLineSpatialIDs()

	// 空間IDが返却されていないこと
	if len(result1) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(result1))
	}

	// 空間IDが返却されていないこと
	if len(result2) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の内部空間用軸の空間ID: %v", len(result2))
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs07 正常系動作確認(円柱の場合: 軸の長さが直径より大きい場合)
//
// 試験詳細：
// + 試験データ
//   パターン7
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 26
//    垂直精度： 27
//    カプセル判定： false
//    衝突判定実施オプション： false
//    Webメルカトル換算係数： 1.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs07(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(26)
	vZoom := int64(27)
	isCapsule := false
	isPrecision := true
	factor := 1.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点終点から取得できる線分APIの空間ID結果
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)
	// 期待値: 始点終点を半径分短くした軸の線に対して行う線分APIの空間ID結果
	// 軸のベクトルを短くする
	startRadiusAxis := spatial.NewVectorFromPoints(startOrth, endOrth).Unit().Scale(radius)
	endRadiusAxis := spatial.NewVectorFromPoints(startOrth, endOrth).Unit().Scale(-1 * radius)
	newStart := startOrth.Translate(startRadiusAxis)
	factorStart := spatial.Point3{newStart.X, newStart.Y, newStart.Z}
	newEnd := endOrth.Translate(endRadiusAxis)
	factorEnd := spatial.Point3{newEnd.X, newEnd.Y, newEnd.Z}
	// 【直交座標空間⇒緯度経度空間】始点終点の座標
	projectesStart := convertSpatialProjectionPoint(factorStart)
	projectesEnd := convertSpatialProjectionPoint(factorEnd)
	projectedPoints := []*object.ProjectedPoint{
		&projectesStart,
		&projectesEnd,
	}
	wgs84Points, _ := ConvertProjectedPointListToPointList(projectedPoints, consts.OrthCrs)
	expecetedInsideLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(wgs84Points[0], wgs84Points[1], hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expecetedLineSpatialIDs)
	sort.Strings(expecetedInsideLineSpatialIDs)
	sort.Strings(lineSpatialIDs)
	sort.Strings(insideLineSpatialIDs)

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedInsideLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs08 正常系動作確認(円柱の場合: 軸の長さが直径より大きい場合かつWebメルカトル補正有)
//
// 試験詳細：
// + 試験データ
//   パターン8
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 27
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： false
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 球に応じた始点・終点間の軸の空間IDを返却すること
//   - 球に応じた始点・終点間の内部空間用軸の空間IDを返却すること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs08(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(27)
	vZoom := int64(26)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点終点から取得できる線分APIの空間ID結果
	end.SetAlt(end.Alt() / factor)
	start.SetAlt(start.Alt() / factor)
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)
	// 期待値: 始点終点を半径分短くした軸の線に対して行う線分APIの空間ID結果
	// 軸のベクトルを短くする
	startRadiusAxis := spatial.NewVectorFromPoints(startOrth, endOrth).Unit().Scale(radius * factor)
	endRadiusAxis := spatial.NewVectorFromPoints(startOrth, endOrth).Unit().Scale(-1 * radius * factor)
	newStart := startOrth.Translate(startRadiusAxis)
	factorStart := spatial.Point3{newStart.X, newStart.Y, newStart.Z / factor}
	newEnd := endOrth.Translate(endRadiusAxis)
	factorEnd := spatial.Point3{newEnd.X, newEnd.Y, newEnd.Z / factor}
	// 【直交座標空間⇒緯度経度空間】始点終点の座標
	projectesStart := convertSpatialProjectionPoint(factorStart)
	projectesEnd := convertSpatialProjectionPoint(factorEnd)
	projectedPoints := []*object.ProjectedPoint{
		&projectesStart,
		&projectesEnd,
	}
	wgs84Points, _ := ConvertProjectedPointListToPointList(projectedPoints, consts.OrthCrs)
	expecetedInsideLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(wgs84Points[0], wgs84Points[1], hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expecetedLineSpatialIDs)
	sort.Strings(expecetedInsideLineSpatialIDs)
	sort.Strings(lineSpatialIDs)
	sort.Strings(insideLineSpatialIDs)

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//戻り値と期待値の始点・終点間の内部空間用軸の空間IDを比較
	if !reflect.DeepEqual(expecetedInsideLineSpatialIDs, insideLineSpatialIDs) {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs09 異常系動作確認(円柱の場合)
//
// 試験詳細：
// + 試験データ
//   パターン9
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 36(不正値)
//    垂直精度： 25
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 戻り値の空間IDのリストが二つとも空であること
//   - 戻り値としてエラー内容が返却されること
func TestCalcLineSpatialIDs09(t *testing.T) {
	// 期待値
	expectErr := "InputValueError,入力チェックエラー"

	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753198, 35.685471, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(36)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// テスト対象呼び出し
	result1, result2, resultErr := cap.calcLineSpatialIDs()

	// 空間IDが返却されていないこと
	if len(result1) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(result1))
	}

	// 空間IDが返却されていないこと
	if len(result2) > 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("始点・終点間の内部空間用軸の空間ID: %v", len(result2))
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestCalcLineSpatialIDs10 正常系動作確認(円柱の場合: 軸の長さが直径より小さい場合)
//
// 試験詳細：
// + 試験データ
//   パターン10
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 26
//    垂直精度： 27
//    カプセル判定： false
//    衝突判定実施オプション： false
//    Webメルカトル換算係数： 1.0
//
// + 確認内容
//   - 円柱に応じた始点・終点間の軸の空間IDを返却すること
//   - 始点・終点間の内部空間用軸の空間IDが空リストであること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcLineSpatialIDs10(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(26)
	vZoom := int64(27)
	isCapsule := false
	isPrecision := true
	factor := 1.0

	cap := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 期待値: 始点終点から取得できる線分APIの空間ID結果
	expecetedLineSpatialIDs, _ :=
		GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)

	// テスト対象呼び出し
	lineSpatialIDs, insideLineSpatialIDs, err := cap.calcLineSpatialIDs()

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expecetedLineSpatialIDs)
	sort.Strings(lineSpatialIDs)
	sort.Strings(insideLineSpatialIDs)

	//戻り値と期待値の始点・終点間の軸の空間IDを比較
	if !reflect.DeepEqual(expecetedLineSpatialIDs, lineSpatialIDs) {
		t.Errorf("始点・終点間の軸の空間ID - 期待値%v, 取得値%v",
			expecetedLineSpatialIDs, lineSpatialIDs)
	}

	//始点・終点間の内部空間用軸の空間IDが空リスト
	if len(insideLineSpatialIDs) > 0 {
		t.Errorf("始点・終点間の内部空間用軸の空間ID - 取得値%v",
			insideLineSpatialIDs)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が27であること
func TestCalcAllSpatialIDs01(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDの数が27個であること
	if len(capsule.allSpatialIDs) != 27 {
		// 27個でない場合Errorをログに出力
		t.Errorf("全体の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs02 正常系動作確認(始点・終点間の軸の空間IDが0)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 0個
//    単位ボクセル成分が全て: 1
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全空間IDの数が0であること
func TestCalcAllSpatialIDs02(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 任意の空間ID１つ
	lineSpatialIDs := []string{}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 0 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs03 正常系動作確認(全方向の移動数が1以下となるが切り上げで1となること確認)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 2
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1(切り上げ)であるため、全空間IDの数が27であること
func TestCalcAllSpatialIDs03(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 2,
		Y: 2,
		Z: 2,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 27 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs04 正常系動作確認(全方向の移動数が1で空間ID数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 2個(X方向に隣に並んでいるケース)
//    単位ボクセル成分が全て: 2
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が36であること
func TestCalcAllSpatialIDs04(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID1 := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	baseSpatialID2 := GetSpatialIDOnAxisIDs(
		11,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID1, baseSpatialID2}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 36 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs05 正常系動作確認(全方向の移動数が1で空間ID数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 2個(XYZ方向の斜めに並んでいるケース)
//    単位ボクセル成分が全て: 2
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が46であること
func TestCalcAllSpatialIDs05(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID1 := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	baseSpatialID2 := GetSpatialIDOnAxisIDs(
		11,
		11,
		11,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID1, baseSpatialID2}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 46 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs06 正常系動作確認(全成分の移動数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 2
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が125であること
func TestCalcAllSpatialIDs06(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 125 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs07 正常系動作確認(全成分の移動数が2, 係数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 1
//    係数: 2
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が125であること
func TestCalcAllSpatialIDs07(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 125 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcAllSpatialIDs08 正常系動作確認(移動数が成分ごとに異なる)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセルX成分: 1
//    単位ボクセルX成分: 2
//    単位ボクセルX成分: 3
//    半径: 1
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、全空間IDの数が45であること
func TestCalcAllSpatialIDs08(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 1.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	// テスト対象呼び出し
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 全体の空間IDが空でないこと
	if len(capsule.allSpatialIDs) != 45 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("始点・終点間の軸の空間IDの個数: %v", len(capsule.allSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 3
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、内部空間IDの数が27であること
//   - 内部空間IDのリストが期待値と完全に一致すること
func TestCalcIncludeSpatialIDs01(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 3.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// 期待値
	expectLineSpatialIDs := []string{}
	for x := 9; x <= 11; x++ {
		for y := 9; y <= 11; y++ {
			for z := 9; z <= 11; z++ {
				expectLineSpatialID := fmt.Sprintf("25/%v/%v/25/%v", x, y, z)
				expectLineSpatialIDs = append(expectLineSpatialIDs, expectLineSpatialID)
			}
		}
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(capsule.includeSpatialIDs)
	sort.Strings(expectLineSpatialIDs)

	// 全体の空間IDの数が27であること
	if len(capsule.includeSpatialIDs) != 27 {
		// 27個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	// 空間IDが期待値と一致していること
	if !reflect.DeepEqual(expectLineSpatialIDs, capsule.includeSpatialIDs) {
		t.Errorf("衝突判定済みの空間ID - 期待値%v, 取得値%v",
			expectLineSpatialIDs, capsule.includeSpatialIDs)
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs02 正常系動作確認(始点・終点間の軸の空間IDが0)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 0個
//    単位ボクセル成分が全て: 1
//    半径: 3
//    係数: 1
//
// + 確認内容
//   - 内部空間IDの数が0であること
func TestCalcIncludeSpatialIDs02(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 3.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 空間IDが0個
	lineSpatialIDs := []string{}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が0個であること
	if len(capsule.includeSpatialIDs) != 0 {
		// 0個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs03 正常系動作確認
// (全方向の移動数が「-1」される前の状態で2以上となるが、切り下げで2となること確認)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 2
//    半径: 8
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が1であるため、内部空間IDの数が27であること
func TestCalcIncludeSpatialIDs03(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 8.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 2,
		Y: 2,
		Z: 2,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が27であること
	if len(capsule.includeSpatialIDs) != 27 {
		// 27個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs04 正常系動作確認(全方向の移動数が1で空間ID数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 2個(X方向に隣に並んでいるケース)
//    単位ボクセル成分が全て: 2
//    半径: 3
//    係数: 2
//
// + 確認内容
//   - 全方向の移動数が1であるため、内部空間IDの数が36であること
func TestCalcIncludeSpatialIDs04(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 3.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID1 := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	baseSpatialID2 := GetSpatialIDOnAxisIDs(
		11,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID2つ
	lineSpatialIDs := []string{baseSpatialID1, baseSpatialID2}
	unitVoxel := spatial.Vector3{
		X: 2,
		Y: 2,
		Z: 2,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が36であること
	if len(capsule.includeSpatialIDs) != 36 {
		// 36個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs05 正常系動作確認(全方向の移動数が1で空間ID数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 2個(XYZ方向の斜めに並んでいるケース)
//    単位ボクセル成分が全て: 2
//    半径: 4
//    係数: 2
//
// + 確認内容
//   - 内部空間IDの数が46であること
func TestCalcIncludeSpatialIDs05(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 4.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID1 := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	baseSpatialID2 := GetSpatialIDOnAxisIDs(
		11,
		11,
		11,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID1, baseSpatialID2}
	unitVoxel := spatial.Vector3{
		X: 2,
		Y: 2,
		Z: 2,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が46であること
	if len(capsule.includeSpatialIDs) != 46 {
		// 46個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs06 正常系動作確認(全成分の移動数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 9
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が2であるため、内部空間IDの数が125であること
func TestCalcIncludeSpatialIDs06(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 5.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が125であること
	if len(capsule.includeSpatialIDs) != 125 {
		// 125個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs07 正常系動作確認(全成分の移動数が2, 係数が2)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセル成分が全て: 1
//    半径: 4
//    係数: 2
//
// + 確認内容
//   - 全方向の移動数が2であるため、内部空間IDの数が125であること
func TestCalcIncludeSpatialIDs07(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 2.5
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 1,
		Z: 1,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数がvであること
	if len(capsule.includeSpatialIDs) != 125 {
		// 125個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs08 正常系動作確認(移動数が成分ごとに異なる)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセルX成分: 1
//    単位ボクセルX成分: 2
//    単位ボクセルX成分: 3
//    半径: 3
//    係数: 1
//
// + 確認内容
//   - 内部空間IDの数が315であること
func TestCalcIncludeSpatialIDs08(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 3.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 4.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が315であること
	if len(capsule.includeSpatialIDs) != 315 {
		// 315個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcIncludeSpatialIDs09 正常系動作確認(移動数が0)
// (全方向の移動数が「-1」される前の状態で1以上となるが、切り下げで1となること確認)
//
// 試験詳細：
// + 試験データ
//    始点・終点間の軸の空間ID: 1個
//    単位ボクセルX成分: 1
//    単位ボクセルX成分: 2
//    単位ボクセルX成分: 3
//    半径: 2
//    係数: 1
//
// + 確認内容
//   - 全方向の移動数が0であるため、内部空間IDの数が0であること
func TestCalcIncludeSpatialIDs09(t *testing.T) {
	//カプセルオブジェクト作成(任意値)
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 1.0
	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	//  始点・終点間の軸の空間
	baseSpatialID := GetSpatialIDOnAxisIDs(
		10,
		10,
		10,
		hZoom,
		vZoom,
	)
	// 任意の空間ID１つ
	lineSpatialIDs := []string{baseSpatialID}
	unitVoxel := spatial.Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	// テスト対象呼び出し
	capsule.calcIncludeSpatialIDs(lineSpatialIDs, unitVoxel)

	// 内部空間IDの数が0であること
	if len(capsule.includeSpatialIDs) != 0 {
		// 0個でない場合Errorをログに出力
		t.Errorf("内部空間IDの個数: %v", len(capsule.includeSpatialIDs))
	}

	t.Log("テスト終了")
}

// TestCalcCollideSpatialIDs01 正常系動作確認(カプセルの場合)
//
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： (139.92271122072384, 35.5610740346, -0.8500000000029104)
//    終点： (139.92259973802746, 35.5608653809, -0.8500000000029104)
//    半径： 2.0
//    水平精度： 26
//    垂直精度： 26
//    カプセル判定： true
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 1.2292622540181326
//
// + 確認内容
//   - 衝突した空間IDを取得すること
//   - 内部空間IDリストの要素が重複していないこと
func TestCalcCollideSpatialIDs01(t *testing.T) {

	// カプセルの全空間ID
	allSpatialIds := []string{
		"26/59637911/26453548/26/5",
		"26/59637911/26453548/26/4",
		"26/59637911/26453548/26/3",
		"26/59637911/26453548/26/2",
		"26/59637911/26453548/26/1",
		"26/59637911/26453548/26/0",
		"26/59637911/26453548/26/0",
		"26/59637911/26453548/26/-1",
		"26/59637911/26453548/26/-2",
		"26/59637911/26453548/26/-3",
		"26/59637911/26453548/26/-4",
		"26/59637911/26453548/26/-5",
		"26/59637911/26453548/26/-6",
		"26/59637911/26453548/26/-7",
		"26/59637911/26453548/26/-8",
		"26/59637911/26453548/26/-9",
	}

	// 期待値(カプセルの内部空間ID)
	expecetIncludeSpatialIds := []string{
		"26/59637911/26453548/26/3",
		"26/59637911/26453548/26/2",
		"26/59637911/26453548/26/1",
		"26/59637911/26453548/26/0",
		"26/59637911/26453548/26/-1",
		"26/59637911/26453548/26/-2",
		"26/59637911/26453548/26/-3",
		"26/59637911/26453548/26/-4",
		"26/59637911/26453548/26/-5",
		"26/59637911/26453548/26/-6",
		"26/59637911/26453548/26/-7",
	}
	sort.Strings(expecetIncludeSpatialIds)

	//入力パラメータ
	radius := 3.15
	hZoom := int64(26)
	vZoom := int64(26)
	isCapsule := true
	isPrecision := true
	p1, _ := object.NewPoint(139.92271122072384, 35.5610740346, -0.8500000000029104)
	p2, _ := object.NewPoint(139.92259973802746, 35.5608653809, -0.8500000000029104)
	center := []*object.Point{p1, p2}
	radian := common.DegreeToRadian(center[0].Lat())
	factor := 1 / math.Cos(radian)

	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{p1, p2},
		consts.OrthCrs,
	)

	startOrth := spatial.Point3{
		X: crsPoints[0].X,
		Y: crsPoints[0].Y,
		Z: crsPoints[0].Alt * factor,
	}
	endOrth := spatial.Point3{
		X: crsPoints[1].X,
		Y: crsPoints[1].Y,
		Z: crsPoints[1].Alt * factor,
	}

	// カプセルオブジェクト
	capsule := NewCapsule(
		startOrth,
		endOrth,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		isPrecision,
		factor,
	)

	// 全空間IDを設定
	for _, x := range allSpatialIds {
		capsule.allSpatialIDs = append(capsule.allSpatialIDs, x)
	}

	// テスト対象呼び出し
	capsule.calcCollideSpatialIDs()

	// 内部空間IDが期待値と一致すること
	sort.Strings(capsule.includeSpatialIDs)
	if !reflect.DeepEqual(expecetIncludeSpatialIds, capsule.includeSpatialIDs) {
		t.Errorf("内部空間ID - 期待値%v, 取得値%v",
			expecetIncludeSpatialIds, capsule.includeSpatialIDs)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range capsule.includeSpatialIDs {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在。重複ID： %v", x)
		}
	}

	t.Log("テスト終了")
}

// TestCalcCollideSpatialIDs02 正常系動作確認(円柱の場合)
//
// 試験詳細：
// + 試験データ
//   パターン2
//    始点： (139.92271122072384, 35.5610740346, -0.8500000000029104)
//    終点： (139.92259973802746, 35.5608653809, -0.8500000000029104)
//    半径： 3.15
//    水平精度： 26
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 1.2292622540181326
//
// + 確認内容
//   - 衝突した空間IDを取得すること
//   - 内部空間IDリストの要素が重複していないこと
func TestCalcCollideSpatialIDs02(t *testing.T) {
	// 円柱の全空間ID
	allSpatialIds := []string{
		"26/59637892/26453591/26/-7",
		"26/59637892/26453591/26/-8",
		"26/59637892/26453591/26/1",
		"26/59637892/26453591/26/1",
		"26/59637892/26453591/26/2",
		"26/59637891/26453592/26/2",
		"26/59637891/26453592/26/2",
		"26/59637891/26453592/26/-6",
	}

	// 期待値
	expecetIncludeSpatialIds := []string{
		"26/59637891/26453592/26/-1",
		"26/59637891/26453592/26/-2",
		"26/59637891/26453592/26/-3",
		"26/59637891/26453592/26/-4",
		"26/59637891/26453592/26/-5",
		"26/59637891/26453592/26/-6",
		"26/59637891/26453592/26/0",
		"26/59637891/26453592/26/1",
		"26/59637891/26453592/26/2",
		"26/59637892/26453591/26/-1",
		"26/59637892/26453591/26/-2",
		"26/59637892/26453591/26/-3",
		"26/59637892/26453591/26/-4",
		"26/59637892/26453591/26/-5",
		"26/59637892/26453591/26/-6",
		"26/59637892/26453591/26/-7",
		"26/59637892/26453591/26/0",
		"26/59637892/26453591/26/1",
		"26/59637892/26453591/26/2",
	}
	sort.Strings(expecetIncludeSpatialIds)

	//入力パラメータ
	p1, _ := object.NewPoint(139.92271122072384, 35.5610740346, -0.8500000000029104)
	p2, _ := object.NewPoint(139.92259973802746, 35.5608653809, -0.8500000000029104)
	center := []*object.Point{p1, p2}
	radian := common.DegreeToRadian(center[0].Lat())
	factor := 1 / math.Cos(radian)

	radius := 3.15
	hZoom := int64(26)
	vZoom := int64(26)
	isCapsule := false
	isPrecision := true

	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{p1, p2},
		consts.OrthCrs,
	)

	startOrth := spatial.Point3{
		X: crsPoints[0].X,
		Y: crsPoints[0].Y,
		Z: crsPoints[0].Alt * factor,
	}
	endOrth := spatial.Point3{
		X: crsPoints[1].X,
		Y: crsPoints[1].Y,
		Z: crsPoints[1].Alt * factor,
	}

	// 円柱オブジェクト
	capsule := NewCapsule(
		startOrth,
		endOrth,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		isPrecision,
		factor,
	)

	// 全空間IDを設定
	for _, x := range allSpatialIds {
		capsule.allSpatialIDs = append(capsule.allSpatialIDs, x)
	}

	// テスト対象呼び出し
	capsule.calcCollideSpatialIDs()

	// 内部空間IDが期待値と一致すること
	sort.Strings(capsule.includeSpatialIDs)
	if !reflect.DeepEqual(expecetIncludeSpatialIds, capsule.includeSpatialIDs) {
		t.Errorf("内部空間ID - 期待値%v, 取得値%v",
			expecetIncludeSpatialIds, capsule.includeSpatialIDs)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range capsule.includeSpatialIDs {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在。重複ID： %v", x)
		}
	}

	t.Log("テスト終了")
}

// TestCalcCollideSpatialIDs03 正常系動作確認(球の場合)
//
// 試験詳細：
// + 試験データ
//   パターン3
//    始点： (139.92271122072384, 35.5610740346, -0.8500000000029104)
//    終点： (139.92271122072384, 35.5610740346, -0.8500000000029104)
//    半径： 2.0
//    水平精度： 23
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： false
//    Webメルカトル換算係数： 1.2292622540181326
//
// + 確認内容
//   - 衝突した空間IDを取得すること
//   - 内部空間IDリストの要素が重複していないこと
func TestCalcCollideSpatialIDs03(t *testing.T) {
	// 球の全空間ID
	allSpatialIds := []string{
		"26/59637911/26453548/26/5",
		"26/59637911/26453548/26/4",
		"26/59637911/26453548/26/3",
		"26/59637911/26453548/26/2",
		"26/59637911/26453548/26/2",
		"26/59637911/26453548/26/1",
		"26/59637911/26453548/26/0",
		"26/59637911/26453548/26/-1",
		"26/59637911/26453548/26/-2",
		"26/59637911/26453548/26/-3",
		"26/59637911/26453548/26/-4",
		"26/59637911/26453548/26/-5",
		"26/59637911/26453548/26/-6",
		"26/59637911/26453548/26/-7",
		"26/59637911/26453548/26/-8",
		"26/59637911/26453548/26/-9",
		"26/59637911/26453548/26/-9",
	}

	// 期待値(球の内部空間ID)
	expecetIncludeSpatialIds := []string{
		"26/59637911/26453548/26/3",
		"26/59637911/26453548/26/2",
		"26/59637911/26453548/26/1",
		"26/59637911/26453548/26/0",
		"26/59637911/26453548/26/-1",
		"26/59637911/26453548/26/-2",
		"26/59637911/26453548/26/-3",
		"26/59637911/26453548/26/-4",
		"26/59637911/26453548/26/-5",
		"26/59637911/26453548/26/-6",
		"26/59637911/26453548/26/-7",
	}
	sort.Strings(expecetIncludeSpatialIds)

	//入力パラメータ
	radius := 3.15
	hZoom := int64(26)
	vZoom := int64(26)
	isCapsule := true
	isPrecision := true
	p1, _ := object.NewPoint(139.92271122072384, 35.5610740346, -0.8500000000029104)
	p2, _ := object.NewPoint(139.92271122072384, 35.5610740346, -0.8500000000029104)
	center := []*object.Point{p1, p2}
	radian := common.DegreeToRadian(center[0].Lat())
	factor := 1 / math.Cos(radian)

	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{p1, p2},
		consts.OrthCrs,
	)

	startOrth := spatial.Point3{
		X: crsPoints[0].X,
		Y: crsPoints[0].Y,
		Z: crsPoints[0].Alt * factor,
	}
	endOrth := spatial.Point3{
		X: crsPoints[1].X,
		Y: crsPoints[1].Y,
		Z: crsPoints[1].Alt * factor,
	}

	// 球オブジェクト
	capsule := NewCapsule(
		startOrth,
		endOrth,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		isPrecision,
		factor,
	)

	// 全空間IDを設定
	for _, x := range allSpatialIds {
		capsule.allSpatialIDs = append(capsule.allSpatialIDs, x)
	}

	// テスト対象呼び出し
	capsule.calcCollideSpatialIDs()

	// 内部空間IDが期待値と一致すること
	sort.Strings(capsule.includeSpatialIDs)
	if !reflect.DeepEqual(expecetIncludeSpatialIds, capsule.includeSpatialIDs) {
		t.Errorf("内部空間ID - 期待値%v, 取得値%v",
			expecetIncludeSpatialIds, capsule.includeSpatialIDs)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range capsule.includeSpatialIDs {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在。重複ID： %v", x)
		}
	}

	t.Log("テスト終了")
}

// TestCalcValidSpatialIDs01 正常系動作確認(衝突判定あり)
//
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 27
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 取得した結果が衝突判定済みの空間IDであること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcValidSpatialIDs01(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(27)
	vZoom := int64(26)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	capsule := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 比較用の実際の図形を覆う全体の空間ID
	lineSpatialIDs, _, _ := capsule.calcLineSpatialIDs()
	baseLat := int64(math.Pow(2, float64(hZoom-1)))
	baseSpatialID := GetSpatialIDOnAxisIDs(
		0,
		baseLat,
		0,
		hZoom,
		vZoom,
	)
	unitVoxel, _ := capsule.calcUnitVoxelVector(baseSpatialID)
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// テスト対象呼出し
	resultVal, err := capsule.CalcValidSpatialIDs()

	// 取得した結果が衝突判定済みの空間IDであること
	if !reflect.DeepEqual(capsule.includeSpatialIDs, resultVal) {
		t.Errorf("衝突判定済みの空間ID - 期待値%v, 取得値%v",
			capsule.includeSpatialIDs, resultVal)
	}
	// 取得した結果が全空間IDでないこと
	if reflect.DeepEqual(capsule.allSpatialIDs, resultVal) {
		t.Errorf("全空間ID - 期待値%v, 取得値%v",
			capsule.includeSpatialIDs, resultVal)
	}

	// 空間IDが取得できていること
	if len(resultVal) == 0 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("実際の図形分の有効な空間IDの個数: %v", len(resultVal))
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcValidSpatialIDs02 正常系動作確認(衝突判定なし)
//
// 試験詳細：
// + 試験データ
//   パターン1
//    始点： 任意
//    終点： 任意
//    半径： 2.0
//    水平精度： 27
//    垂直精度： 26
//    カプセル判定： false
//    衝突判定実施オプション： false
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 取得した結果が全空間IDであること
//   - 戻り値としてエラー内容が返却されないこと
func TestCalcValidSpatialIDs02(t *testing.T) {
	//入力パラメータ
	start, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	end, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	// 【直交座標空間】始点終点の座標
	crsPoints, _ := ConvertPointListToProjectedPointList(
		[]*object.Point{start, end},
		consts.OrthCrs,
	)
	//初期化用入力パラメータ
	startOrth := spatial.Point3{crsPoints[0].X, crsPoints[0].Y, crsPoints[0].Alt}
	endOrth := spatial.Point3{crsPoints[1].X, crsPoints[1].Y, crsPoints[1].Alt}
	radius := 2.0
	hZoom := int64(27)
	vZoom := int64(26)
	isCapsule := false
	isPrecision := false
	factor := 2.0

	capsule := NewCapsule(startOrth, endOrth, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// 比較用の実際の図形を覆う全体の空間ID
	lineSpatialIDs, _, _ := capsule.calcLineSpatialIDs()
	baseLat := int64(math.Pow(2, float64(hZoom-1)))
	baseSpatialID := GetSpatialIDOnAxisIDs(
		0,
		baseLat,
		0,
		hZoom,
		vZoom,
	)
	unitVoxel, _ := capsule.calcUnitVoxelVector(baseSpatialID)
	capsule.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// テスト対象呼出し
	resultVal, err := capsule.CalcValidSpatialIDs()

	// 取得した結果が衝突判定済みの空間IDでないこと
	if reflect.DeepEqual(capsule.includeSpatialIDs, resultVal) {
		t.Errorf("衝突判定済みの空間ID - 期待値%v, 取得値%v",
			capsule.includeSpatialIDs, resultVal)
	}
	// 取得した結果が全空間IDであること
	if !reflect.DeepEqual(capsule.allSpatialIDs, resultVal) {
		t.Errorf("全空間ID - 期待値%v, 取得値%v",
			capsule.includeSpatialIDs, resultVal)
	}

	// 空間IDが取得できていること
	if len(resultVal) == 0 {
		// 内部空間IDのリストが空の場合Errorをログに出力
		t.Errorf("実際の図形分の有効な空間IDの個数: %v", len(resultVal))
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestCalcValidSpatialIDs03 異常系動作確認(水平精度が不正)
//
// 試験詳細：
// + 試験データ
//   パターン2
//    始点： (1,2,3)
//    終点： (-20,-32,-30)
//    半径： 2.0
//    水平精度： -1(不正値)
//    垂直精度： 25
//    カプセル判定： false
//    衝突判定実施オプション： true
//    Webメルカトル換算係数： 2.0
//
// + 確認内容
//   - 戻り値の空間IDが空であること
//   - 戻り値としてエラー内容が返却されること
func TestCalcValidSpatialIDs03(t *testing.T) {
	// 期待値
	expectErr := "InputValueError,入力チェックエラー"

	//入力パラメータ
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{-20, -32, -30}
	radius := 2.0
	hZoom := int64(-1)
	vZoom := int64(25)
	isCapsule := false
	isPrecision := true
	factor := 2.0

	capsule := NewCapsule(start, end, radius, hZoom, vZoom, isCapsule, isPrecision, factor)

	// テスト対象呼出し
	resultVal, resultErr := capsule.CalcValidSpatialIDs()

	// 内部空間IDが取得できていること
	if len(resultVal) > 0 {
		// 内部空間IDのリストが空でない場合Errorをログに出力
		t.Errorf("実際の図形分の有効な空間IDの個数: %v", len(resultVal))
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestIsPrecision01 正常系動作確認(衝突判定フラグ：true)
//
// 試験詳細：
// + 試験データ
//    パターン1
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値関数がshape.option型であること
//   - 戻り値関数にIsPrecisionOpts構造体ポインタを渡した際に、
//     IsPrecisionOpts構造体の衝突判定実施オプションが更新されること
func TestIsPrecision01(t *testing.T) {
	// 衝突判定フラグ(shape.IsPrecisionOpts)をfalseで初期化
	val := IsPrecisionOpts{false}

	// テスト対象呼出し
	resultVal := IsPrecision(true)
	resultVal(&val)

	// 期待値
	expectedType := "shape.option"
	expectedValue := true

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(expectedType, reflect.TypeOf(resultVal).String()) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v",
			expectedType, reflect.TypeOf(resultVal))
	}

	// IsPrecisionOpts構造体の衝突判定実施オプションがtrueに更新されていること
	if val.IsPrecision != expectedValue {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("衝突判定実施オプション - 期待値：%v, 取得値：%v", expectedValue, val.IsPrecision)
	}

	t.Log("テスト終了")
}

// TestIsPrecision02 正常系動作確認(衝突判定フラグ：false)
//
// 試験詳細：
// + 試験データ
//    パターン2
//     衝突判定フラグ：false
//
// + 確認内容
// 　- 戻り値関数がshape.option型であること
//   - 戻り値関数にIsPrecisionOpts構造体ポインタを渡した際に、
//     IsPrecisionOpts構造体の衝突判定実施オプションが更新されること
func TestIsPrecision02(t *testing.T) {

	// 衝突判定フラグ(shape.IsPrecisionOpts)をtrueで初期化
	val := IsPrecisionOpts{true}

	// テスト対象呼出し
	resultVal := IsPrecision(false)
	resultVal(&val)

	// 期待値
	expectedType := "shape.option"
	expectedValue := false

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(expectedType, reflect.TypeOf(resultVal).String()) {
		t.Errorf("オブジェクトの型 - 期待値：%v, 取得値：%v",
			expectedType, reflect.TypeOf(resultVal))
	}

	// IsPrecisionOpts構造体の衝突判定実施オプションがfalseに更新されていること
	if val.IsPrecision != expectedValue {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("衝突判定実施オプション - 期待値：%v, 取得値：%v", expectedValue, val.IsPrecision)
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnCylinders01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//    パターン1
//	   円柱の中心の接続点：Pointオブジェクト
//	   円柱の半径：2.0
//	   精度レベル：25
//	   始点、終点の球状判定：true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 円柱を複数つなげた経路が通る空間IDのリストが正しく返却されること
func TestGetSpatialIdsOnCylinders01(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	pList := []*object.Point{p1}
	radius := 2.0
	zoom := int64(25)
	isCapsule := true

	// 期待値
	// 拡張空間IDを取得
	ids, _ := GetExtendedSpatialIdsOnCylinders(
		pList,
		radius,
		zoom,
		zoom,
		isCapsule,
		IsPrecision(true),
	)
	// 拡張空間IDを空間IDのフォーマットに変換
	ids, _ = ConvertExtendedSpatialIdsToSpatialIds(ids)
	expectVal := ids

	// テスト対象呼出し
	resultVal, _ := GetSpatialIdsOnCylinders(
		pList,
		radius,
		zoom,
		isCapsule,
		IsPrecision(true),
	)

	// 戻り値の一致確認用に空間IDリストを昇順ソート
	sort.Strings(expectVal)
	sort.Strings(resultVal)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnCylinders02 異常系動作確認(入力値不正)
//
// 試験詳細：
// + 試験データ
//    パターン2
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：-2.0(不正値)
//	   精度レベル：25
//     始点、終点の球状判定：true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値にエラー内容が含まれていること
func TestGetSpatialIdsOnCylinders02(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	pList := []*object.Point{p1}
	radius := -2.0
	zoom := int64(25)
	isCapsule := true

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼出し
	_, resultErr := GetSpatialIdsOnCylinders(
		pList,
		radius,
		zoom,
		isCapsule,
		IsPrecision(true),
	)

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnCylinders03 異常系動作確認(精度閾値超過)
//
// 試験詳細：
// + 試験データ
//    パターン3
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   精度レベル：36(不正値)
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値にエラー内容が含まれていること
func TestGetSpatialIdsOnCylinders03(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	radius := 2.0
	zoom := int64(36)
	isCapsule := true

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼出し
	_, resultErr := GetSpatialIdsOnCylinders(
		pList,
		radius,
		zoom,
		isCapsule,
		IsPrecision(true),
	)

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnCylinders04 正常系動作確認(接続点のリストが空)
//
// 試験詳細：
// + 試験データ
//    パターン4
//	   円柱の中心の接続点：Pointオブジェクト(空)
//     円柱の半径：2.0
//	   精度レベル：25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDのリストが空であること
func TestGetSpatialIdsOnCylinders04(t *testing.T) {

	//初期化用入力パラメータ
	pList := []*object.Point{}
	radius := 2.0
	zoom := int64(25)
	isCapsule := true

	// 期待値
	expectVal := []string{}
	expectErr := new(error)

	// テスト対象呼出し
	resultVal, resultErr := GetSpatialIdsOnCylinders(
		pList,
		radius,
		zoom,
		isCapsule,
		IsPrecision(true),
	)

	// 戻り値のエラー内容の比較
	if !reflect.DeepEqual(*expectErr, resultErr) {
		// エラーが返された場合Errorをログに出力
		t.Errorf("エラー内容 - 期待値：%v, 取得値：%v",
			*expectErr, resultErr)
	}

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders01 正常系動作確認(接続点数が1)
//
// 試験詳細：
// + 試験データ
//    パターン1
//	   円柱の中心の接続点：Pointオブジェクト(接続点数が1つ)
//     円柱の半径：2.0
//	   水平方向の精度レベル：25
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定：false
//     衝突判定フラグ：true
//
// + 確認内容
//   - 円柱を複数つなげた経路が通る空間IDのリストが返却されること(結果の妥当性は結合試験で確認)
//   - 戻り値にエラー内容が含まれていないこと
// 　- 戻り値の空間IDが空でないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders01(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	center := []*object.Point{p1}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在するかを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders02 異常系動作確認(引数のポインタにnilが存在)
//
// 試験詳細：
// + 試験データ
//    パターン2
//	   円柱の中心の接続点：Pointオブジェクトとnil
//     円柱の半径：2.0
//	   水平方向の精度レベル：25
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
//   - 円柱を複数つなげた経路が通る空間IDのリストが返却されること(結果の妥当性は結合試験で確認)
//   - 戻り値にエラー内容が含まれていないこと
// 　- 戻り値の空間IDが空でないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders02(t *testing.T) {

	//入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1, nil}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// 期待値
	expectVal := []string{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}
	// 重複するIDが存在するかを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders03 異常系動作確認(水平方向の精度が下限)
//
// 試験詳細：
// + 試験データ
//    パターン3
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル：0
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders03(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1}
	radius := 2.0
	hZoom := int64(0)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if resultErr != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", resultErr)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders04 異常系動作確認(水平方向の精度が下限-1)
//
// 試験詳細：
// + 試験データ
//    パターン4
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル -1
//	   垂直方向の精度レベル 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空であること
// 　- 戻り値にエラー内容が含まれていること
func TestGetExtendedSpatialIdsOnCylinders04(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	p2, _ := object.NewPoint(19.753098, 35.371, 12.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(-1)
	vZoom := int64(25)
	isCapsule := true

	// 期待値
	expectErr := "InputValueError,入力チェックエラー"
	expectVal := []string{}

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders05 正常系動作確認(水平方向の精度が上限)
//
// 試験詳細：
// + 試験データ
//    パターン5
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：0.01
//	   水平方向の精度レベル 35
//	   垂直方向の精度レベル 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders05(t *testing.T) {
	//入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1}
	radius := 0.01
	hZoom := int64(35)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if resultErr != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", resultErr)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders06 異常系動作確認(水平方向の精度が上限+1)
//
// 試験詳細：
// + 試験データ
//    パターン6
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル 36
//	   垂直方向の精度レベル 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空であること
// 　- 戻り値にエラー内容が含まれていること
func TestGetExtendedSpatialIdsOnCylinders06(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	p2, _ := object.NewPoint(19.753098, 35.371, 12.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(36)
	vZoom := int64(25)
	isCapsule := true

	// 期待値
	expectErr := "InputValueError,入力チェックエラー"
	expectVal := []string{}

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders07 異常系動作確認(垂直方向の精度が下限)
//
// 試験詳細：
// + 試験データ
//    パターン7
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル 25
//	   垂直方向の精度レベル 0
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders07(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(0)
	isCapsule := true

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if resultErr != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", resultErr)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders08 異常系動作確認(垂直方向の精度が下限-1)
//
// 試験詳細：
// + 試験データ
//    パターン8
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル 25
//	   垂直方向の精度レベル -1
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空であること
// 　- 戻り値にエラー内容が含まれていること
func TestGetExtendedSpatialIdsOnCylinders08(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	p2, _ := object.NewPoint(139.75308, 35.685381, 13.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(-1)
	isCapsule := true

	// 期待値
	expectErr := "InputValueError,入力チェックエラー"
	expectVal := []string{}

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders09 異常系動作確認(垂直方向の精度が上限)
//
// 試験詳細：
// + 試験データ
//    パターン9
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：0.02
//	   水平方向の精度レベル 25
//	   垂直方向の精度レベル 35
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders09(t *testing.T) {

	//入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1}
	radius := 0.02
	hZoom := int64(25)
	vZoom := int64(35)
	isCapsule := true

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if resultErr != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", resultErr)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders10 異常系動作確認(垂直方向の精度が上限+1)
//
// 試験詳細：
// + 試験データ
//    パターン6
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル 25
//	   垂直方向の精度レベル 36
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空であること
// 　- 戻り値にエラー内容が含まれていること
func TestGetExtendedSpatialIdsOnCylinders10(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	p2, _ := object.NewPoint(139.753099, 35.685381, 12.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(36)
	isCapsule := true

	// 期待値
	expectErr := "InputValueError,入力チェックエラー"
	expectVal := []string{}

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders11 異常系動作確認(半径が下限値)
//
// 試験詳細：
// + 試験データ
//    パターン11
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：0.0(不正値)
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値にエラー内容が含まれていること
// 　- 戻り値の空間IDが空であること
func TestGetExtendedSpatialIdsOnCylinders011(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	center := []*object.Point{p1}
	radius := 0.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// 期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders12 正常系動作確認(半径が下限値+0.1)
//
// 試験詳細：
// + 試験データ
//    パターン12
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：0.1
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
//   - 戻り値にエラー内容が含まれていないこと
// 　- 戻り値の空間IDが空でないこと
//   - 戻り値の空間IDリストの要素が重複していないことする
func TestGetExtendedSpatialIdsOnCylinders12(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 12.0)
	p2, _ := object.NewPoint(139.753099, 35.685381, 12.0)
	center := []*object.Point{p1, p2}
	radius := 0.1
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, resultErr := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if resultErr != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", resultErr)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders13 正常系動作確認(接続点数が0個)
//
// 試験詳細：
// + 試験データ
//    パターン13
//	   円柱の中心の接続点：空のPointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値にエラー内容が含まれていること
// 　- 戻り値の空間IDが空であること
func TestGetExtendedSpatialIdsOnCylinders13(t *testing.T) {

	//入力パラメータ
	center := []*object.Point{}

	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// 期待値
	expectVal := []string{}

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDの比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders14 正常系動作確認(同一座標の接続点が存在)
//
// 試験詳細：
// + 試験データ
//    パターン7
//	   円柱の中心の接続点：Pointオブジェクト(同一座標の接続点が存在)(接続点数が2以上)
//     円柱の半径：2.0
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： false
//     衝突判定フラグ：false
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 同一座標の場合は接続点の空間IDの取得がスキップされること
//   - 終点の場合はの場合は接続点の空間IDの取得がスキップされること
//   - 球オブジェクトが空でない分岐を通ること
//     (カバレッジを確認する)
func TestGetExtendedSpatialIdsOnCylinders14(t *testing.T) {

	//入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	p2, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	p3, _ := object.NewPoint(139.753099, 35.685371, 11.0)
	p4, _ := object.NewPoint(139.753097, 35.685371, 11.0)
	center := []*object.Point{p1, p1, p2, p3, p4}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := false

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(false),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders15 正常系動作確認(始点、終点が球状(カプセル))
//
// 試験詳細：
// + 試験データ
//    パターン15
//	   円柱の中心の接続点：Pointオブジェクト(3点以上)
//     円柱の半径：0.001
//	   水平方向の精度レベル： 15
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 終点もしくはカプセルの場合は接続点の空間IDの取得がスキップされること
//     (カバレッジを確認する)
func TestGetExtendedSpatialIdsOnCylinders15(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	p2, _ := object.NewPoint(134.753098, 34.685371, 10.0)
	p3, _ := object.NewPoint(134.753097, 34.685373, 10.0)
	center := []*object.Point{p1, p2, p3}
	radius := 0.001
	hZoom := int64(15)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders16 正常系動作確認(経度方向にまっすぐなオブジェクト)
//
// 試験詳細：
// + 試験データ
//    パターン16
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders16(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753078, 35.685371, 11.0)
	p2, _ := object.NewPoint(139.753088, 35.685371, 11.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders17 正常系動作確認(緯度方向にまっすぐなオブジェクト)
//
// 試験詳細：
// + 試験データ
//    パターン17
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders17(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	p2, _ := object.NewPoint(139.753098, 35.685381, 11.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders18 正常系動作確認(高さ方向にまっすぐなオブジェクト)
//
// 試験詳細：
// + 試験データ
//    パターン18
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル： 25
//	   垂直方向の精度レベル： 25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders18(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 11.0)
	p2, _ := object.NewPoint(139.753098, 35.685371, 14.394)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders19 正常系動作確認(経度、緯度、高さ方向全てにまっすぐでないオブジェクト)
//
// 試験詳細：
// + 試験データ
//    パターン19
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル：25
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：false
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders19(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(138.753098, 34.685371, 11.0)
	p2, _ := object.NewPoint(138.753099, 34.685381, 13.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(false),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders20 正常系動作確認(緯度がマイナスの値)
//
// 試験詳細：
// + 試験データ
//    パターン20
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：0.01
//	   水平方向の精度レベル：25
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
//   - 戻り値の空間IDが空でないこと
//   - 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders20(t *testing.T) {

	//入力パラメータ
	p1, _ := object.NewPoint(138.753098, -34.685371, 11.0)
	p2, _ := object.NewPoint(138.753099, -34.685481, 13.0)
	center := []*object.Point{p1, p2}
	radius := 0.01
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnCylinders21 正常系動作確認(緯度がプラスとマイナスの値)
//
// 試験詳細：
// + 試験データ
//    パターン21
//	   円柱の中心の接続点：Pointオブジェクト
//     円柱の半径：2.0
//	   水平方向の精度レベル：25
//	   垂直方向の精度レベル：25
//     始点、終点の球状判定： true(カプセル)
//     衝突判定フラグ：true
//
// + 確認内容
// 　- 戻り値の空間IDが空でないこと
// 　- 戻り値にエラー内容が含まれていないこと
//   - 戻り値の空間IDリストの要素が重複していないこと
func TestGetExtendedSpatialIdsOnCylinders21(t *testing.T) {

	//初期化用入力パラメータ
	p1, _ := object.NewPoint(138.753098, -0.00001234, 11.0)
	p2, _ := object.NewPoint(138.753099, 0.00001234, 11.0)
	center := []*object.Point{p1, p2}
	radius := 2.0
	hZoom := int64(25)
	vZoom := int64(25)
	isCapsule := true

	// テスト対象呼出し
	resultVal, err := GetExtendedSpatialIdsOnCylinders(
		center,
		radius,
		hZoom,
		vZoom,
		isCapsule,
		IsPrecision(true),
	)

	// 空間IDが返却されていること(要素数が0以上)
	if len(resultVal) <= 0 {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間IDの取得失敗: %v", resultVal)
	}

	// エラーが返された場合はErrorをログに出力
	if err != nil {
		t.Errorf("error - 期待値：nil, 取得値：%v", err)
	}

	// 重複するIDが存在しないことを確認
	var uniqueSlice []string
	for _, x := range resultVal {
		//  False：スライスに対象が含まれない場合
		if common.Include(uniqueSlice, x) == false {
			uniqueSlice = append(uniqueSlice, x)
		} else {
			//  True：スライスに対象が含まれる場合はエラー
			t.Errorf("重複するIDが存在")
		}
	}

	t.Log("テスト終了")
}
