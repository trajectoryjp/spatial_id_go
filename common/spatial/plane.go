// Package spatial 空間座標操作パッケージ
package spatial

// Plane 面の構造体
type Plane struct {
	Point Point3 // 面上の一点
	Normal Vector3 // 法線ベクトル
}
