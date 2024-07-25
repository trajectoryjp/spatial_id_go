// Package consts 定数パッケージ
package consts

const GeoCrs = 4326            // GeoCrs 空間IDで利用する地理座標系
const OrthCrs = 3857           // OrthCrs 直交座標系のEPSGコード
const SpatialIDDelimiter = "/" // SpatialIDDelimiter 空間IDの区切り文字

const Minima = 1e-10 // Minima 浮動小数点誤差

const ZBaseExponent = 25 // ZBaseExponent 空間IDにおけるボクセルの高さが1mとなるズームレベル

// ZBaseOffsetForNegativeFIndex
//
// 元の最大高さを保ったまま正方向と負方向でfインデックスを半数ずつ導入するためのzBaseOffset
const ZBaseOffsetForNegativeFIndex = 1 << (ZBaseExponent - 1)
