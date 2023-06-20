// Package spatial 空間座標操作パッケージ
package spatial


// Line3 線の構造体
type Line3 struct {
	Point Point3      // 始点
	Direction Vector3 // 方向ベクトル
}

// NewLineFromPoints 2点から線オブジェクト作成
//
// 2点から線オブジェクト作成
//
// 引数：
//  start： 始点
//  end： 終点
//
// 戻り値：
//  2点間の線オブジェクト
func NewLineFromPoints(start Point3, end Point3) Line3 {
	// 方向ベクトル
	vec := NewVectorFromPoints(start, end)

	return Line3{start, vec}
}

// ToPoint 線に沿った点を返却
//
// パラメータtを使用して、線に沿った点を返却
//
// 引数：
//  t： パラメータ
//
// 戻り値：
//  線に沿った点
func (l Line3) ToPoint(t float64) Point3 {
	// 線に沿った分平行移動するベクトル
	vec := l.Direction.Scale(t)

	return l.Point.Translate(vec)
}

// End 終点取得
//
// 線の終点を返却する
//
// 戻り値：
//  線の終点
func (l Line3) End() Point3 {
	return l.Point.Translate(l.Direction)
}

// Start 始点取得
//
// 線の始点を返却する
//
// 戻り値：
//  線の始点
func (l Line3) Start() Point3 {
	return l.Point
}
