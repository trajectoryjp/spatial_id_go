package detector

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/integrate"
	"github.com/trajectoryjp/spatial_id_go/v4/transform"
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
	ids1 := []string{spatialId1}
	ids2 := []string{spatialId2}
	return CheckSpatialIdsArrayOverlap(ids1, ids2)
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
// ただし高度は16,777,216未満 〜 -16,777,216m以上に制限される(この範囲で高度変換を行う)
//
// 戻り値:
//
//	bool:
//		重複の有無が返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。このときbool値はfalseで返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
//	 		空間ID高度範囲外：高度範囲外の空間IDが入力されていた場合。
func CheckSpatialIdsArrayOverlap(spatialIds1 []string, spatialIds2 []string) (bool, error) {
	// f,x,yで3次元分の2分木
	tr := tree.CreateTree(tree.Create3DTable())
	// spatialIds1から各要素をradix treeに格納
	for indexSpatialId1, spatialId1 := range spatialIds1 {
		zoom1, f1, x1, y1, err := getSpatialIdAttrs(spatialId1)
		if err != nil {
			return false, fmt.Errorf("%w @spatialId1[%v]", err, indexSpatialId1)
		}
		// 高度インデックスをオフセット変換のみ実行して自然数にする
		// minAltitudeKey == maxAltitudeKeyになるため結果は片方のみ利用する
		convertedFIndex, _, errAltConversion := transform.ConvertZToMinMaxAltitudekey(int64(f1), int64(zoom1), int64(zoom1), consts.ZOriginValue, consts.ZBaseOffsetForNegativeFIndex)
		if convertedFIndex < 0 {
			return false, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("input f-index %v is out of altitude range @spatialId1[%v] = %v", f1, indexSpatialId1, spatialId1))
		}
		if errAltConversion != nil {
			return false, fmt.Errorf("%w @spatialId1[%v] = %v", errAltConversion, indexSpatialId1, spatialId1)
		}
		index1 := tree.Indexs{convertedFIndex, int64(x1), int64(y1)}
		tr.Append(index1, tree.ZoomSetLevel(zoom1), spatialId1)
	}
	// spatialIds2から各要素の取り出し
	for indexSpatialId2, spatialId2 := range spatialIds2 {
		zoom2, f2, x2, y2, err := getSpatialIdAttrs(spatialId2)
		if err != nil {
			return false, fmt.Errorf("%w @spatialId2[%v]", err, indexSpatialId2)
		}
		// 取り出した要素の比較
		// 高度インデックスをオフセット変換のみ実行して自然数にする
		// minAltitudeKey == maxAltitudeKeyになるため結果は片方のみ利用する
		convertedFIndex2, _, errAltConversion := transform.ConvertZToMinMaxAltitudekey(int64(f2), int64(zoom2), int64(zoom2), consts.ZOriginValue, consts.ZBaseOffsetForNegativeFIndex)
		if convertedFIndex2 < 0 {
			return false, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("input f-index %v is out of altitude range @spatialId2[%v] = %v", f2, indexSpatialId2, spatialId2))
		}
		if errAltConversion != nil {
			return false, fmt.Errorf("%w @spatialId2[%v] = %v", errAltConversion, indexSpatialId2, spatialId2)
		}
		result := tr.IsOverlap(tree.Indexs{convertedFIndex2, int64(x2), int64(y2)}, tree.ZoomSetLevel(zoom2))
		if result {
			// 重複判定時、trueとnilを返却
			return result, nil
		}
	}

	return false, nil
}

// getSpatialIdAttrs 空間IDフォーマットチェック関数
//
// 入力された空間IDの各成分を格納した配列を返却する。
// 以下の条件に当てはまる場合はフォーマット違反となりエラーインスタンスが返却される。
//
//	・空間IDの各成分に数値が入力されていない場合
//	・区切り文字の数が3つで無い場合
//
// 引数：
//
//	spatialId：空間ID
//
// 戻り値：
//
//	空間IDに含まれる精度(ズームレベル), F, X, Y成分を格納した以下のtuple
//	 (精度(ズームレベル), F成分, X成分, Y成分)
//	入力された拡張空間IDのフォーマットが不正な場合、戻り値を0にしエラーインスタンスを返却。
func getSpatialIdAttrs(spatialId string) (int, int, int, int, error) {
	// 分割
	spatialIdAttributes := strings.Split(spatialId, consts.SpatialIDDelimiter)
	if len(spatialIdAttributes) != 4 {
		// 不正形式(要素数)
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	zoom, err := strconv.Atoi(spatialIdAttributes[0])
	// 不正形式(数値)
	if err != nil {
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	f, err := strconv.Atoi(spatialIdAttributes[1])
	// 不正形式(数値)
	if err != nil {
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	x, err := strconv.Atoi(spatialIdAttributes[2])
	// 不正形式(数値)
	if err != nil {
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	y, err := strconv.Atoi(spatialIdAttributes[3])
	// 不正形式(数値)
	if err != nil {
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	return zoom, f, x, y, nil
}

// SpatialIdOverlapDetector 方式吸収のための重複検知インターフェース
type SpatialIdOverlapDetector interface {
	IsOverlap(spatialIds []string) (bool, error)
}

// SpatialIdGreedyOverlapDetector 愚直方式での重複検知構造体
type SpatialIdGreedyOverlapDetector struct {
	detectedIds []PartedSpatialId
}

// NewSpatialIdGreedyOverlapDetector 被検知IDの空間ID列spatialIdsを受け取り、SpatialIdOverlapDetectorを返す
// 空間IDのフォーマットが正しくない場合エラーを返す
func NewSpatialIdGreedyOverlapDetector(spatialIds []string) (SpatialIdOverlapDetector, error) {
	ids := []PartedSpatialId{}
	for index, spatialId := range spatialIds {
		id, err := NewPartedSpatialId(spatialId)
		if err != nil {
			return nil, fmt.Errorf("%w @spatialIds[%v]", err, index)
		}
		ids = append(ids, id)
	}
	return &SpatialIdGreedyOverlapDetector{ids}, nil
}

// IsOverlap 重複検知を行う
// 検知IDの空間ID列spatialIdsを受け取り、重複検知結果をboolで返す
// 空間IDのフォーマットが正しくない場合エラーを返す
func (detector *SpatialIdGreedyOverlapDetector) IsOverlap(spatialIds []string) (bool, error) {
	for index, spatialId := range spatialIds {
		detectId, err := NewPartedSpatialId(spatialId)
		if err != nil {
			return false, fmt.Errorf("%w @spatialIds[%v]", err, index)
		}
		for _, detectedId := range detector.detectedIds {
			if detectId.IsOverlap(detectedId) {
				return true, nil
			}
		}
	}
	return false, nil
}

// SpatialIdGreedyOverlapDetector 多次元基数木方式での重複検知構造体
type SpatialIdTreeOverlapDetector struct {
	positiveTree tree.TreeInterface
	negativeTree tree.TreeInterface
}

// NewSpatialIdTreeOverlapDetector 被検知IDの空間ID列spatialIdsを受け取り、SpatialIdOverlapDetectorを返す
// 空間IDのフォーマットが正しくない場合エラーを返す
func NewSpatialIdTreeOverlapDetector(spatialIds []string) (SpatialIdOverlapDetector, error) {
	// f,x,yで3次元分の2分木
	var positiveTree tree.TreeInterface
	var negativeTree tree.TreeInterface

	for spatialIdIndex, spatialId := range spatialIds {
		zoom, f, x, y, err := getSpatialIdAttrs(spatialId)
		if err != nil {
			return nil, fmt.Errorf("%w @spatialIds[%v]", err, spatialIdIndex)
		}

		if f >= 0 {
			if positiveTree == nil {
				positiveTree = tree.CreateTree(tree.Create3DTable())
			}
			treeIndex := tree.Indexs{int64(f), int64(x), int64(y)}
			positiveTree.Append(treeIndex, tree.ZoomSetLevel(zoom), spatialId)
		} else {
			if negativeTree == nil {
				negativeTree = tree.CreateTree(tree.Create3DTable())
			}
			treeIndex := tree.Indexs{int64(^f), int64(x), int64(y)}
			negativeTree.Append(treeIndex, tree.ZoomSetLevel(zoom), spatialId)
		}
	}
	return &SpatialIdTreeOverlapDetector{positiveTree, negativeTree}, nil
}

// IsOverlap 検知IDの空間ID列spatialIdsを受け取り、重複検知を行う
// 空間IDのフォーマットが正しくない場合エラーを返す
func (detector *SpatialIdTreeOverlapDetector) IsOverlap(spatialIds []string) (bool, error) {
	for spatialIdIndex, spatialId := range spatialIds {
		zoomLevel, f, x, y, err := getSpatialIdAttrs(spatialId)
		if err != nil {
			return false, fmt.Errorf("%w @spatialIds[%v]", err, spatialIdIndex)
		}

		var isOverlap bool
		if f >= 0 {
			if detector.positiveTree != nil {
				treeIndex := tree.Indexs{int64(f), int64(x), int64(y)}
				isOverlap = detector.positiveTree.IsOverlap(treeIndex, tree.ZoomSetLevel(zoomLevel))
			}
		} else {
			if detector.negativeTree != nil {
				treeIndex := tree.Indexs{int64(^f), int64(x), int64(y)}
				isOverlap = detector.negativeTree.IsOverlap(treeIndex, tree.ZoomSetLevel(zoomLevel))
			}
		}
		if isOverlap {
			return true, nil
		}
	}
	return false, nil
}

// PartedSpatialId 文字列の空間IDを分解した構造体
type PartedSpatialId struct {
	zoomLevel int
	indices   [3]int
}

// NewPartedSpatialId 文字列の空間IDであるspatialIdを受け取り、PartedSpatialIdを返す
// 空間IDのフォーマットが正しくない場合エラーを返す
func NewPartedSpatialId(spatialId string) (PartedSpatialId, error) {
	zoomLevel, f, x, y, err := getSpatialIdAttrs(spatialId)
	if err != nil {
		return PartedSpatialId{}, err
	}
	indices := [3]int{f, x, y}
	return PartedSpatialId{zoomLevel, indices}, nil
}

// DownZoomLevel ズームレベルをtoZoomLevelに下げる
// 現在のズームレベルがtoZoomLevel以下の場合は何もしない
func (id PartedSpatialId) DownZoomLevel(toZoomLevel int) PartedSpatialId {
	if id.zoomLevel > toZoomLevel {
		for i := range 3 {
			id.indices[i] >>= id.zoomLevel - toZoomLevel
		}
		id.zoomLevel = toZoomLevel
	}
	return id
}

// IsOverlap 空間IDであるid1およびid2の重複検知を行う
func (id1 PartedSpatialId) IsOverlap(id2 PartedSpatialId) bool {
	if id1.zoomLevel < id2.zoomLevel {
		id2 = id2.DownZoomLevel(id1.zoomLevel)
	} else {
		id1 = id1.DownZoomLevel(id2.zoomLevel)
	}
	return id1 == id2
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
