package integrate

import (
	"math"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v4/common"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/shape"
)

// UnitDividedSpatialID 単位分割拡張空間ID構造体
// 拡張空間IDを最小の空間IDである単位拡張空間IDへ分割した情報を保持する構造体
type UnitDividedSpatialID struct {
	*object.ExtendedSpatialID                     // 元の拡張空間ID
	hDiff                     int64               // 水平精度差分
	vDiff                     int64               // 垂直精度差分
	unitIDs                   map[string]struct{} // 単位拡張空間ID集合
}

// NewUnitDividedSpatialID 単位分割拡張空間ID構造体のコンストラクタ
//
// UnitDividedSpatialIDオブジェクトを返却する。
//
// 引数：
//
//	 s    ：拡張空間ID
//		hDiff：水平方向精度差分
//		vDiff：垂直方向精度差分
//
// 戻り値：
//
//	UnitDividedSpatialIDオブジェクトのポインタ
func NewUnitDividedSpatialID(
	s *object.ExtendedSpatialID,
	hDiff, vDiff int64,
) *UnitDividedSpatialID {

	// 単位拡張空間IDの個数
	hDiffIndex := int64(math.Pow(2, float64(hDiff)))
	vDiffIndex := int64(math.Pow(2, float64(vDiff)))
	// 単位空間ID集合
	unitIDs := map[string]struct{}{}
	for x := s.X() * hDiffIndex; x < (s.X()+1)*hDiffIndex; x++ {
		for y := s.Y() * hDiffIndex; y < (s.Y()+1)*hDiffIndex; y++ {
			for z := s.Z() * vDiffIndex; z < (s.Z()+1)*vDiffIndex; z++ {
				spatialStr :=
					strings.Join([]string{
						strconv.FormatInt(hDiff+s.HZoom(), 10),
						strconv.FormatInt(x, 10),
						strconv.FormatInt(y, 10),
						strconv.FormatInt(vDiff+s.VZoom(), 10),
						strconv.FormatInt(z, 10)},
						consts.SpatialIDDelimiter)
				unitIDs[spatialStr] = struct{}{}
			}
		}
	}

	return &UnitDividedSpatialID{
		ExtendedSpatialID: s,
		hDiff:             hDiff,
		vDiff:             vDiff,
		unitIDs:           unitIDs,
	}
}

// HighSpatialID 最適化後拡張空間ID構造体
// 拡張空間IDのマージ候補である精度が粗い最適化後の拡張空間IDの構造体
// マージ元の拡張空間IDは最適化元拡張空間として保持する。
type HighSpatialID struct {
	*object.ExtendedSpatialID                     // 最適化後の拡張空間ID
	threshold                 int64               // 単位拡張空間IDの個数の閾値
	lowIDs                    []string            // 最適化元拡張空間ID配列
	unitIDs                   map[string]struct{} // 単位拡張空間ID集合
}

// NewHighSpatialID 最適化後拡張空間ID構造体のコンストラクタ
//
// HighSpatialIDオブジェクトを返却する。
//
// 引数：
//
//	 u    ：単位分割拡張空間ID
//		hDiff：単位分割拡張空間IDからの水平方向精度差分
//		vDiff：単位分割拡張空間IDからの垂直方向精度差分
//
// 戻り値：
//
//	HighSpatialIDオブジェクトのポインタ
func NewHighSpatialID(u *UnitDividedSpatialID, hDiff, vDiff int64) *HighSpatialID {

	// 最適化後拡張空間IDを求める
	highID := u.Higher(hDiff, vDiff)
	// 単位拡張空間IDの個数
	hDiffIndex := int64(math.Pow(2, float64(hDiff+u.hDiff)))
	vDiffIndex := int64(math.Pow(2, float64(vDiff+u.vDiff)))
	// 単位拡張空間IDの個数の閾値
	threshold := hDiffIndex * hDiffIndex * vDiffIndex

	// 最適化元拡張空間ID配列
	lowIDs := []string{u.ID()}
	// 単位拡張空間ID集合
	unitIDs := u.unitIDs

	return &HighSpatialID{
		ExtendedSpatialID: highID,
		threshold:         threshold,
		lowIDs:            lowIDs,
		unitIDs:           unitIDs,
	}
}

// Merge 最適化後拡張空間ID構造体の結合
//
// 拡張空間IDが同一の最適化後拡張空間IDの最適化元・単位拡張空間IDをマージする
//
// 引数：
//
//	s：マージ対象の最適化後拡張空間ID
func (r *HighSpatialID) Merge(s *HighSpatialID) {
	// 最適化元拡張空間ID配列を結合
	r.lowIDs = append(r.lowIDs, s.lowIDs...)
	// 単位拡張空間ID集合を結合
	for k := range s.unitIDs {
		r.unitIDs[k] = struct{}{}
	}
}

// IsDense 最適化後拡張空間ID構造体が単位拡張空間ID集合で稠密であるかの判定
//
// 最適化後拡張空間ID構造体が単位拡張空間ID集合で稠密であるかの判定結果を返却する
//
// 戻り値：
//
//	True：最適化後拡張空間ID構造体が単位拡張空間ID集合で稠密
//	False：最適化後拡張空間ID構造体が単位拡張空間ID集合で稠密でない
func (r HighSpatialID) IsDense() bool {
	return int64(len(r.unitIDs)) == r.threshold
}

// MergeSpatialIds 空間IDの最適化（マージ）関数
//
// 入力された空間ID配列をより大きな空間IDにマージし、最適化した結果を返却する。
// マージ後の空間IDの精度は、引数で指定された精度となる。
//
// 入力された空間ID配列の内、以下の条件を満たす空間IDがマージ対象となる。
//
//	・空間IDの精度が、マージ後の精度以上。
//
// 入力された空間ID配列がマージされる条件は以下となる。
//
//	・マージ後の精度の空間IDボクセルを完全に満たす空間ID群が、入力された空間ID配列に含まれている場合
//	   ボクセル内を満たす空間ID群を、マージ後の精度の空間IDにマージする。
//
// マージ条件に合致しなかった空間ID、マージ対象外の空間IDはそのまま返却される。
// マージ後に重複する空間IDが存在した場合は、重複を解消したうえで返却する。
//
// 引数：
//
//	spatialIds：マージ対象の空間ID文字列配列
//	zoom      ：マージ後の精度。空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値：
//
//	マージ後の空間IDを格納した配列が返却される。IDの重複は解消された形で返却される。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 空間IDフォーマット不正：空間IDのフォーマットに違反する値が"空間ID配列"に入力されていた場合。
//
// 補足事項：
//
//	多数の空間IDのマージを行う場合について：
//	 空間IDのマージ処理では内部で空間IDを配列内の一番小さい空間IDに合わせて精度を上げて分割する処理を行っている。
//	 空間IDにおける精度を上げる場合、精度が1上がるごとに8のべき乗で空間ID数が増加する。
//	 そのため、以下のような場合に分割後の空間ID数は大幅に増大する。
//	  * 入力の空間ID配列全体での精度の幅が大きい場合。
//	  * 空間IDが大量に入力された場合。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func MergeSpatialIds(spatialIds []string, zoom int64) ([]string, error) {
	// マージ後の空間ID格納用のスライス
	resultIDs := []string{}

	// 拡張空間ID変換結果
	extendedSpatialIds, convertErr :=
		shape.ConvertSpatialIdsToExtendedSpatialIds(spatialIds)

	if convertErr != nil {
		// 変換に失敗した場合エラーインスタンスを返却
		return resultIDs, convertErr
	}

	// 拡張空間IDのマージを実行
	mergeIDs, mergeZoomErr := MergeExtendedSpatialIds(
		extendedSpatialIds,
		zoom,
		zoom,
	)

	if mergeZoomErr != nil {
		// マージに失敗した場合エラーインスタンスを返却
		return resultIDs, mergeZoomErr
	}

	// 拡張空間IDを空間IDに変換
	resultIDs, _ = shape.ConvertExtendedSpatialIdsToSpatialIds(mergeIDs)

	return resultIDs, nil
}

// MergeExtendedSpatialIds 拡張空間IDの最適化（マージ）関数
//
// 入力された拡張空間ID配列をより大きな拡張空間IDにマージし、最適化した結果を返却する。
// マージ後の拡張空間IDの精度は、引数で指定された水平方向精度、垂直方向精度となる。
//
// 入力された拡張空間ID配列の内、以下の条件を満たす拡張空間IDがマージ対象となる。
//
//	・拡張空間IDの水平方向精度が、マージ後の水平方向精度以上。
//	・拡張空間IDの垂直方向精度が、マージ後の垂直方向精度以上。
//
// 入力された拡張空間ID配列がマージされる条件は以下となる。
//
//	・マージ後の精度の拡張空間IDボクセルを完全に満たす拡張空間ID群が、入力された拡張空間ID配列に含まれている場合
//	   ボクセル内を満たす拡張空間ID群を、マージ後の精度の拡張空間IDにマージする。
//
// マージ条件に合致しなかった拡張空間ID、マージ対象外の拡張空間IDはそのまま返却される。
// マージ後に重複する拡張空間IDが存在した場合は、重複を解消したうえで返却する。
//
// 引数：
//
//	extendedSpatialIds：マージ対象の拡張空間ID文字列配列
//	hZoom             ：マージ後の水平方向精度。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//	vZoom             ：マージ後の垂直方向精度。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値：
//
//	マージ後の拡張空間IDを格納した配列が返却される。IDの重複は解消された形で返却される。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過              ：精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 拡張空間IDフォーマット不正：拡張空間IDのフォーマットに違反する値が"拡張空間ID配列"に入力されていた場合。
//
// 補足事項：
//
//	多数の拡張空間IDのマージを行う場合について：
//	 拡張空間IDのマージ処理では内部で拡張空間IDを配列内の一番小さい拡張空間IDに合わせて精度を上げて分割する処理を行っている。
//	 拡張空間IDにおける精度を上げる場合、水平精度が1上がるごとに4のべき乗、垂直方向精度が1上がるごとに2のべき乗で拡張空間ID数が増加する。
//	 そのため、以下のような場合に分割後の拡張空間ID数は大幅に増大する。
//	  * 入力の拡張空間ID配列全体での精度の幅が大きい場合。
//	  * 入力の拡張空間ID内の水平精度と垂直精度に大きい差分がある場合。
//	  * 拡張空間IDが大量に入力された場合。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func MergeExtendedSpatialIds(extendedSpatialIds []string, hZoom, vZoom int64) ([]string, error) {
	// 変換後の拡張空間ID格納用のスライス
	resultIDs := []string{}
	// 拡張空間ID配列
	spatialIDs := []*object.ExtendedSpatialID{}

	// 入力値チェック
	// 水平、垂直方向精度のどちらかが範囲外の場合、空配列とエラーインスタンスを返却
	if !shape.CheckZoom(hZoom) || !shape.CheckZoom(vZoom) {
		return resultIDs, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 最大水平精度
	maxHZoom := int64(0)
	// 最大垂直精度
	maxVZoom := int64(0)
	// 拡張空間ID配列を作成しつつ、最大水平/垂直精度を算出
	for _, extendedSpatialID := range extendedSpatialIds {
		spatialID, err := object.NewExtendedSpatialID(extendedSpatialID)

		if err != nil {
			// 拡張空間ID初期化時にエラーが発生した場合、フォーマットチェックエラーとしてエラーインスタンスを返却
			return resultIDs, err
		}

		// 入力精度より細かい精度もしくは等しい精度の場合はマージ対象の空間IDとする。
		if spatialID.HZoom() >= hZoom && spatialID.VZoom() >= vZoom {
			spatialIDs = append(spatialIDs, spatialID)

			// 入力精度より粗い精度の場合は処理を行わず、結果空間IDに格納
		} else {
			resultIDs = append(resultIDs, spatialID.ID())
		}

		// 最大水平/垂直精度を更新
		if spatialID.HZoom() > maxHZoom {
			maxHZoom = spatialID.HZoom()
		}
		if spatialID.VZoom() > maxVZoom {
			maxVZoom = spatialID.VZoom()
		}
	}

	unitSpatialIDs := []*UnitDividedSpatialID{}
	for _, spatialID := range spatialIDs {
		// 最大水平/垂直精度を単位分割の単位になるようにして、単位分割空間ID作成
		unitSpatialID := NewUnitDividedSpatialID(
			spatialID,
			maxHZoom-spatialID.HZoom(),
			maxVZoom-spatialID.VZoom(),
		)
		unitSpatialIDs = append(unitSpatialIDs, unitSpatialID)
	}

	// 最適化後空間ID辞書
	highSpatialIDs := map[string]*HighSpatialID{}
	for _, unitSpatialID := range unitSpatialIDs {

		// 最適化後空間ID
		highSpatialID := NewHighSpatialID(
			unitSpatialID,
			unitSpatialID.HZoom()-hZoom,
			unitSpatialID.VZoom()-vZoom,
		)

		// 最適化後空間IDが辞書に登録済みの場合は最適化後空間ID構造体を結合
		if val, ok := highSpatialIDs[highSpatialID.ID()]; ok {
			val.Merge(highSpatialID)

			// 未登録の場合は辞書に登録
		} else {
			highSpatialIDs[highSpatialID.ID()] = highSpatialID
		}
	}

	// 最適化後空間IDごとに処理
	for _, highSpatialID := range highSpatialIDs {
		// 単位分割数が単位空間ID集合の数に等しい場合
		if highSpatialID.IsDense() {
			// 最適化後空間IDを結果空間ID配列に格納
			resultIDs = append(resultIDs, highSpatialID.ID())

			// 単位分割数が単位空間ID集合の数と異なる場合
		} else {
			// 最適化元空間ID配列を結果空間ID配列に格納
			resultIDs = append(resultIDs, highSpatialID.lowIDs...)
		}
	}

	// 重複する空間IDを削除
	resultIDs = common.Unique(resultIDs)

	return resultIDs, nil
}
