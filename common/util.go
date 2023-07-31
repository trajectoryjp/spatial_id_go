// Package common 共通処理パッケージ。
package common

import (
	"math"

	"github.com/trajectoryjp/spatial_id_go/common/errors"
)

// AlmostEqual 同値確認関数
//
// 引数に入力された数値の浮動小数点誤差以内かを判定する
//
// 引数：
//
//	x： 数値1
//	y： 数値2
//	absTol： 浮動小数点誤差
//
// 戻り値：
//
//	判定結果(True:浮動小数点誤差内、False:浮動小数点誤差範囲外)
func AlmostEqual(x, y, absTol float64) bool {
	return x == y || math.Abs(x-y) <= absTol
}

// Number 数値のインターフェース
type Number interface {
	int | int32 | int64 | float32 | float64
}

// Max 最大値取得関数
//
// 引数に入力された数値スライスの最大値を返却する
//
// 引数：
//
//	numbers：数値スライス
//
// 戻り値：
//
//	数値スライスの最大値
//
// 戻り値（エラー）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力不正：入力スライスが空の場合。
//
// 使用例：
//
//	nums := []int{1, 2, 3}
//	// 最大値取得
//	max := common.Max(nums)
func Max[T Number](numbers []T) (T, error) {
	// スライスが空の場合はエラー
	if len(numbers) == 0 {
		var v T
		return v, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return max, nil
}

// Min 最小値取得関数
//
// 引数に入力された数値スライスの最小値を返却する
//
// 引数：
//
//	numbers：数値スライス
//
// 戻り値：
//
//	数値スライスの最小値
//
// 戻り値（エラー）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力不正：入力スライスが空の場合。
//
// 使用例：
//
//	nums := []int{1, 2, 3}
//	// 最小値取得
//	min := common.Min(nums)
func Min[T Number](numbers []T) (T, error) {
	// スライスが空の場合はエラー
	if len(numbers) == 0 {
		var v T
		return v, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}

	return min, nil
}

// DegreeToRadian ラジアン取得関数
//
// 引数に入力された角度をラジアンに変換して返却する
//
// 引数：
//
//	degree：角度
//
// 戻り値：
//
//	引数に入力した角度をラジアンに変換した結果
func DegreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// RadianToDegree 角度取得関数
//
// 引数に入力されたラジアンを角度に変換して返却する
//
// 引数：
//
//	radian：ラジアン
//
// 戻り値：
//
//	引数に入力したラジアンを角度に変換した結果
func RadianToDegree(radian float64) float64 {
	return radian * (180 / math.Pi)
}

// Union 和集合取得関数
//
// 引数に入力されたスライスの和集合スライスを返却する
//
// 引数：
//
//	l1：スライス1
//	l2：スライス2
//
// 戻り値：
//
//	和集合スライス
func Union[T comparable](l1, l2 []T) []T {
	s := make(map[T]struct{}, len(l1))

	for _, data := range l1 {
		s[data] = struct{}{}
	}

	for _, data := range l2 {
		s[data] = struct{}{}
	}

	// 和集合スライス
	r := make([]T, 0, len(s))

	for key := range s {
		r = append(r, key)
	}

	return r
}

// Difference 差集合取得関数
//
// 引数に入力されたスライスの差集合スライスを返却する
//
// 引数：
//
//	l1：差分を出すスライス1
//	l2：差とするスライス2
//
// 戻り値：
//
//	差集合スライス
func Difference[T comparable](l1, l2 []T) []T {
	s := make(map[T]struct{}, len(l2))

	for _, data := range l2 {
		s[data] = struct{}{}
	}

	// 差集合スライス
	r := make([]T, 0, len(l1))

	for _, data := range l1 {

		if _, ok := s[data]; ok {
			continue
		}
		r = append(r, data)
	}

	return r
}

// Intersect 積集合取得関数
//
// 引数に入力されたスライスの積集合スライスを返却する
//
// 引数：
//
//	l1：スライス1
//	l2：スライス2
//
// 戻り値：
//
//	積集合スライス
func Intersect[T comparable](l1, l2 []T) []T {
	s := make(map[T]struct{}, len(l1))

	for _, data := range l1 {
		s[data] = struct{}{}
	}

	// 積集合スライス
	r := make([]T, 0)

	for _, data := range l2 {
		if _, ok := s[data]; ok {
			r = append(r, data)
		}
	}

	return r
}

// Unique ユニーク化関数
//
// 引数に入力されたスライスのユニークスライスを返却する
//
// 引数：
//
//	target：スライス
//
// 戻り値：
//
//	ユニークスライス
func Unique[T comparable](target []T) []T {
	s := make(map[T]struct{}, len(target))

	for _, data := range target {
		s[data] = struct{}{}
	}

	r := make([]T, 0, len(s))

	for key := range s {
		r = append(r, key)
	}

	return r
}

// Include 内包関数
//
// 引数に入力されたスライスに対象が含まれるかを返却する
//
// 引数：
//
//	slice：スライス
//	target： 対象
//
// 戻り値：
//
//	True：スライスに対象が含まれる場合
//	False：スライスに対象が含まれない場合
func Include[T comparable](slice []T, target T) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

// Combinations 組み合わせ関数
//
// 引数に入力された組み合わせを返却する
//
// 引数：
//
//	n： 組み合わせる数（位数）
//	k： 選択数
//	f： 組み合わせに対する処理
func Combinations(n, k int64, f func([]int64)) {
	pattern := make([]int64, k)
	for i := range pattern {
		pattern[i] = int64(i)
	}

	for {
		f(pattern)

		pos := k - 1
		for {
			if pos == -1 {
				return
			}

			oldNum := pattern[pos]
			if oldNum == n+pos-k {
				pos--
				continue
			}

			pattern[pos]++
			break
		}

		for pos++; pos < k; pos++ {
			pattern[pos] = pattern[pos-1] + 1
		}
	}
}
