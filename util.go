// Package common 共通処理パッケージ。
package spatialID


// computes arithmatic shift of index and shift parameters. When index = 1, similar to returning 2^shift. When index > 1,
// similar to returning index*2^shift.
func CalculateArithmeticShift(index int64, shift int64) int64 {

	// determine if shift is non-negative
	if shift >= 0 {
		return index << shift
	} else {
		return index >> -shift
	}

}
