// Package enum オプション定数管理用パッケージ
package enum

// PointOption 座標取得オプション用の型
type PointOption int

// MergeOption 空間IDマージオプション用の型
type MergeOption int

// 点群APIで入力可能な座標取得のオプション
const (
	Vertex PointOption = iota // 空間IDの頂点座標を取得(0)
	Center                    // 空間IDの中心座標を取得(1)
)
