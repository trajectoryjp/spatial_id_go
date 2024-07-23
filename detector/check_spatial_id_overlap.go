package detector

import (
	"fmt"
	"github.com/trajectoryjp/spatial_id_go/v4/integrate"
	"strconv"
	"strings"
)

// CheckSpatialIdsOverlap 2つの空間ID重複の判定関数
//
// 比較対象として入力された2つの空間IDのズームレベルを変換して揃え、重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	spatialId1, spatialId2: 重複判定対象の空間ID。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無が返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときbool値にfalseが返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
func CheckSpatialIdsOverlap(spatialId1 string, spatialId2 string) (bool, error) {
	// ズームレベルの取り出し
	arr1 := strings.Split(spatialId1, "/")
	arr2 := strings.Split(spatialId2, "/")
	if len(arr1) != 4 || len(arr2) != 4 {
		// 不正形式
		return false, fmt.Errorf("invalid format. spatialId1: %v, spatialId2: %v", spatialId1, spatialId2)
	}
	zoom1, _ := strconv.Atoi(arr1[0])
	zoom2, _ := strconv.Atoi(arr2[0])

	// zoomレベルで場合分け
	// zoomレベルが等しいとき、変換なし
	if zoom1 > zoom2 {
		// spatialId2のzoomレベルをzoom1に合わせるよう変換
		convertedSpatialId, err := integrate.ChangeSpatialIdsZoom([]string{spatialId1}, int64(zoom2))
		if err != nil {
			return false, err
		}
		spatialId1 = convertedSpatialId[0]
	} else if zoom1 < zoom2 {
		// spatialId1のzoomレベルをzoom2に合わせるよう変換
		convertedSpatialId, err := integrate.ChangeSpatialIdsZoom([]string{spatialId2}, int64(zoom1))
		if err != nil {
			return false, err
		}
		spatialId2 = convertedSpatialId[0]
	}

	return spatialId1 == spatialId2, nil

}

// CheckSpatialIdsArrayOverlap 2つの空間ID重複の判定関数
//
// 比較対象として入力された2つの空間ID列の重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	spatialIds1, spatialIds2: 重複判定対象の空間ID列。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無が返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときbool値にfalseが返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
func CheckSpatialIdsArrayOverlap(spatialIds1 []string, spatialIds2 []string) (bool, error) {
	// spatialIds1から各要素の取り出し
	for _, spatialId1 := range spatialIds1 {
		// spatialIds2から各要素の取り出し
		for _, spatialId2 := range spatialIds2 {
			// 取り出した要素の比較
			result, err := CheckSpatialIdsOverlap(spatialId1, spatialId2)
			if err != nil {
				// エラー発生時、falseとerrorを返却
				return false, err
			}
			if result {
				// 重複判定時、trueとnilを返却
				return result, nil
			}
		}
	}

	return false, nil
}

// CheckExtendedSpatialIdsOverlap 2つの拡張空間IDの重複の判定関数
//
// 比較対象として入力された2つの拡張空間IDのズームレベルを変換して揃え、重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が拡張空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	extendedSpatialId1, extendedSpatialId2 : 重複判定対象の拡張空間ID。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無がbool値で返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときfalseも返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
func CheckExtendedSpatialIdsOverlap(extendedSpatialId1 string, extendedSpatialId2 string) (bool, error) {
	// ズームレベルの取り出し
	arr1 := strings.Split(extendedSpatialId1, "/")
	arr2 := strings.Split(extendedSpatialId2, "/")
	if len(arr1) != 5 || len(arr2) != 5 {
		// 不正形式
		return false, fmt.Errorf("invalid format. extendedSpatialId1: %v, extendedSpatialId2: %v", extendedSpatialId1, extendedSpatialId2)
	}

	// 変換後の水平方向ズームレベルの決定
	horizontalZoom1, _ := strconv.Atoi(arr1[0])
	horizontalZoom2, _ := strconv.Atoi(arr2[0])
	// zoomレベルで場合分け
	targetHorizontalZoom := horizontalZoom1
	if horizontalZoom1 > horizontalZoom2 {
		targetHorizontalZoom = horizontalZoom2
	}

	// 変換後の垂直方向のズームレベルの設定
	verticalZoom1, _ := strconv.Atoi(arr1[3])
	verticalZoom2, _ := strconv.Atoi(arr2[3])
	targetVerticalZoom := verticalZoom1
	// zoomレベルで場合分け
	if verticalZoom1 > verticalZoom2 {
		targetVerticalZoom = verticalZoom2
	}

	// 変換後のZoomレベルを指定して変換
	extendedSpatialIds1, err := integrate.ChangeExtendedSpatialIdsZoom([]string{extendedSpatialId1}, int64(targetHorizontalZoom), int64(targetVerticalZoom))
	if err != nil {
		// 変換エラー
		return false, err
	}
	extendedSpatialIds2, err := integrate.ChangeExtendedSpatialIdsZoom([]string{extendedSpatialId2}, int64(targetHorizontalZoom), int64(targetVerticalZoom))
	if err != nil {
		// 変換エラー
		return false, err
	}

	return extendedSpatialIds1[0] == extendedSpatialIds2[0], nil

}

// CheckExtendedSpatialIdsArrayOverlap 2つの拡張空間IDの重複の判定関数
//
// 比較対象として入力された2つの拡張空間ID列の重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が拡張空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	extendedSpatialIds1, extendedSpatialIds2 : 重複判定対象の拡張空間ID列。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無がbool値で返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときfalseも返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。

func CheckExtendedSpatialIdsArrayOverlap(extendedSpatialIds1 []string, extendedSpatialIds2 []string) (bool, error) {
	// spatialIds1から各要素の取り出し
	for _, extendedSpatialId1 := range extendedSpatialIds1 {
		// spatialIds2から各要素の取り出し
		for _, extendedSpatialId2 := range extendedSpatialIds2 {
			// 取り出した要素の比較
			result, err := CheckExtendedSpatialIdsOverlap(extendedSpatialId1, extendedSpatialId2)
			if err != nil {
				// エラー発生時、falseとerrorを返却
				return false, err
			}
			if result {
				// 重複判定時、trueとnilを返却
				return result, nil
			}
		}
	}
	return false, nil
}
