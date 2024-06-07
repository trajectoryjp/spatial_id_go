// Package integrate 空間IDパッケージ
package integrate

import (
	"math"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v4/common"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/shape"
)

// ChangeSpatialIdsZoom 空間IDの精度変換関数
//
// 変換対象として入力された空間IDを、指定された精度の空間IDに変換する。
// 変換対象には複数の空間IDを指定可能、精度が異なる空間IDが混在している場合も入力を許容する。
//
// 精度を上げる場合、入力された空間IDボクセルを変換後の精度で分割し、
// 入力された空間IDに内包される空間IDを返却する。
//
// 精度を下げる場合、入力された空間IDボクセルを変換後の精度となるよう拡大し、
// 入力された空間IDを内包している空間IDを返却する。
//
// 引数の入力値がAPIの仕様に違反していた場合、精度変換失敗となりエラーインスタンスが返却される。
//
// 引数：
//
//	spatialIds：精度変換対象の空間ID。複数の空間IDを指定可能、空間ID毎に精度が異なっている入力も許容。
//	zoom      ：変換後の精度。空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値：
//
//	精度変換後の全空間IDを格納した配列が返却される。IDの重複は解消された形で返却される。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 空間IDフォーマット不正：空間IDのフォーマットに違反する値が"変換対象空間ID"に入力されていた場合。
//
// 補足事項：
//
//	複数空間ID変換時の入出力対応について：
//	 本APIでは、入力された全空間IDの精度変換結果をまとめて返却している。
//	 そのため、複数の空間IDを入力とした場合、精度変換前後の空間IDの対応を取ることはできない。
//	 精度変換前後の空間IDの対応を保持したい場合は、空間IDを1つずつ変換する必要がある。
//
//	精度値を上げる場合について：
//	 空間IDにおける精度を上げる場合、精度が1上がるごとに8のべき乗で空間ID数が増加する。
//	 そのため、低い精度から高い精度に変換する際、精度差が大きすぎると変換後の空間ID数は大幅に増大する。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func ChangeSpatialIdsZoom(spatialIds []string, zoom int64) ([]string, error) {
	// 変換後の空間ID格納用のスライス
	resultIDList := []string{}

	// 拡張空間ID変換結果
	extendedSpatialIds, convertErr :=
		shape.ConvertSpatialIdsToExtendedSpatialIds(spatialIds)

	if convertErr != nil {
		// 変換に失敗した場合エラーインスタンスを返却
		return resultIDList, convertErr
	}

	// 拡張空間IDの精度変換を実行
	changeIDList, changeZoomErr := ChangeExtendedSpatialIdsZoom(
		extendedSpatialIds,
		zoom,
		zoom,
	)

	if changeZoomErr != nil {
		// 変換に失敗した場合エラーインスタンスを返却
		return resultIDList, changeZoomErr
	}

	// 拡張空間IDを空間IDに変換
	resultIDList, _ = shape.ConvertExtendedSpatialIdsToSpatialIds(changeIDList)

	return resultIDList, nil
}

// ChangeExtendedSpatialIdsZoom 拡張空間IDの精度変換関数
//
// 変換対象として入力された拡張空間IDを、指定された水平/垂直方向精度の拡張空間IDに変換する。
// 変換対象には複数の拡張空間IDを指定可能、精度が異なる拡張空間IDが混在している場合も入力を許容する。
//
// 精度を上げる場合、入力された拡張空間IDボクセルを変換後の精度で分割し、
// 入力された拡張空間IDに内包される拡張空間IDを返却する。
//
// 精度を下げる場合、入力された拡張空間IDボクセルを変換後の精度となるよう拡大し、
// 入力された拡張空間IDを内包している拡張空間IDを返却する。
//
// 水平/垂直方向精度について、一方の精度が上がり一方の精度が下がる場合、
// 精度が上がった方向はボクセルを分割、精度が下がった方向はボクセルの拡大がされた拡張空間IDを返却する。
//
// 引数の入力値がAPIの仕様に違反していた場合、精度変換失敗となりエラーインスタンスが返却される。
//
// 引数：
//
//	extendedSpatialIds：精度変換対象の拡張空間ID。複数の拡張空間IDを指定可能、拡張空間ID毎に精度が異なっている入力も許容。
//	hZoom             ：変換後の水平方向精度。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//	vZoom             ：変換後の垂直方向精度。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値：
//
//	精度変換後の全拡張空間IDを格納した配列が返却される。IDの重複は解消された形で返却される。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過              ：水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 拡張空間IDフォーマット不正：拡張空間IDのフォーマットに違反する値が"変換対象の拡張空間ID"に入力されていた場合。
//
// 補足事項：
//
//	複数変換時の入出力対応について：
//	 本APIでは、入力された全拡張空間IDの精度変換結果をまとめて返却している。
//	 そのため、複数の拡張空間IDを入力とした場合、精度変換前後の拡張空間IDの対応を取ることはできない。
//	 精度変換前後の拡張空間IDの対応を保持したい場合は、変換対象を1つずつ変換する必要がある。
//
//	精度値を上げる場合について：
//	 拡張空間IDにおける精度を上げる場合、水平精度が1上がるごとに4のべき乗、垂直方向精度が1上がるごとに2のべき乗で拡張空間ID数が増加する。
//	 そのため、低い精度から高い精度に変換する際、精度差が大きすぎると変換後の拡張空間ID数は大幅に増大する。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func ChangeExtendedSpatialIdsZoom(
	extendedSpatialIds []string,
	hZoom int64,
	vZoom int64,
) ([]string, error) {
	// 変換後の拡張空間ID格納用のスライス
	resultIDList := []string{}

	// 入力値チェック
	// 水平、垂直方向精度のどちらかが範囲外の場合、空配列とエラーインスタンスを返却
	if !shape.CheckZoom(hZoom) || !shape.CheckZoom(vZoom) {
		return resultIDList, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 空間IDインスタンス初期化
	s := &object.ExtendedSpatialID{}

	for _, id := range extendedSpatialIds {
		// 拡張空間ID再設定
		if err := s.ResetExtendedSpatialID(id); err != nil {
			// 拡張空間IDの再設定時にエラーが返却された場合、エラーインスタンスを返却
			return resultIDList, err
		}

		// 拡張空間IDの成分取得
		components := s.FieldParams()

		// 水平方向成分の精度変換結果取得
		hComponents := HorizontalZoom(
			components[0],
			components[1],
			components[2],
			hZoom,
		)

		// 垂直方向成分の精度変換結果取得
		vComponents := VerticalZoom(
			components[3],
			components[4],
			vZoom,
		)

		for _, h := range hComponents {
			for _, v := range vComponents {
				// 水平、垂直方向の精度変換結果から拡張空間IDを生成
				resultID := strings.Join([]string{h, v}, "/")
				resultIDList = append(resultIDList, resultID)
			}
		}
	}

	return common.Unique(resultIDList), nil
}

// HorizontalZoom 拡張空間IDの水平方向の精度変換関数
//
// 変換対象として入力された拡張空間IDの水平方向成分を、指定された精度に変換する。
// 精度を上げる場合、入力された拡張空間IDボクセルを変換後の精度で分割し、
// 入力された拡張空間IDに内包される拡張空間IDを返却する。
//
// 精度を下げる場合、入力された拡張空間IDボクセルを変換後の精度となるよう拡大し、
// 入力された拡張空間IDを内包している拡張空間IDを返却する。
//
// 引数：
//
//	inputZoom ：精度変換対象の拡張空間IDの水平方向精度。
//	xIndex    ：精度変換対象の拡張空間IDのxIndex成分。
//	yIndex    ：精度変換対象の拡張空間IDのyIndex成分。
//	outputZoom：変換後の水平方向精度。
//
// 戻り値：
//
//	精度変換後の全拡張空間IDの水平方向成分を格納したスライスが返却される。
//
// 補足事項：
//
//	精度値を上げる場合について：
//	 拡張空間IDにおける精度を上げる場合、水平精度が1上がるごとに4のべき乗で拡張空間ID数が増加する。
//	 そのため、低い精度から高い精度に変換する際、精度差が大きすぎると変換後の拡張空間ID数は大幅に増大する。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func HorizontalZoom(
	inputZoom int64,
	xIndex int64,
	yIndex int64,
	outputZoom int64,
) []string {
	// 変換後の拡張空間ID格納用のスライス
	horizontalIDs := []string{}

	// 精度の差異を取得
	hZoomDiff := outputZoom - inputZoom

	// 変換後の x, y, z 成分の最小値を初期化
	minXparam := xIndex
	minYparam := yIndex

	// 変換後の x, y, z 成分の最大値を初期化
	maxXparam := xIndex
	maxYparam := yIndex

	// 水平方向の拡張空間ID分割数を定義
	hVoxelNum := int64(math.Pow(2, math.Abs(float64(hZoomDiff))))

	// 精度の増減によって処理を分岐
	if hZoomDiff > 0 {
		// 水平精度が上がった場合
		// 変換後の x, y 成分の最小値を定義
		minXparam = xIndex * hVoxelNum
		minYparam = yIndex * hVoxelNum

		// 変換後の x, y 成分の最大値を定義
		maxXparam = minXparam + hVoxelNum - 1
		maxYparam = minYparam + hVoxelNum - 1
	} else if hZoomDiff < 0 {
		// 水平精度が下がった場合
		// 変換後の x, y 成分の最小値を定義
		minXparam = xIndex / hVoxelNum
		minYparam = yIndex / hVoxelNum

		// 変換後の x, y 成分の最大値を定義
		maxXparam = minXparam
		maxYparam = minYparam
	}

	// 変換後の拡張空間IDを定義
	for y := minYparam; y <= maxYparam; y++ {
		for x := minXparam; x <= maxXparam; x++ {
			horizontalIDs = append(
				horizontalIDs,
				strconv.FormatInt(outputZoom, 10)+
					"/"+strconv.FormatInt(x, 10)+
					"/"+strconv.FormatInt(y, 10))
		}
	}
	return horizontalIDs
}

// VerticalZoom 拡張空間IDの垂直方向の精度変換関数
//
// 変換対象として入力された拡張空間IDの高さ方向成分を、指定された精度に変換する。
// 精度を上げる場合、入力された拡張空間IDボクセルを変換後の精度で分割し、
// 入力された拡張空間IDに内包される拡張空間IDを返却する。
//
// 精度を下げる場合、入力された拡張空間IDボクセルを変換後の精度となるよう拡大し、
// 入力された拡張空間IDを内包している拡張空間IDを返却する。
//
// 引数：
//
//	inputZoom ：精度変換対象の拡張空間IDの垂直方向精度成分。
//	vIndex    ：精度変換対象の拡張空間IDのvIndex成分。
//	outputZoom：変換後の垂直方向精度。
//
// 戻り値：
//
//	精度変換後の全拡張空間IDの垂直方向成分を格納したスライスが返却される。
//
// 補足事項：
//
//	精度値を上げる場合について：
//	 拡張空間IDにおける精度を上げる場合、垂直方向精度が1上がるごとに2のべき乗で拡張空間ID数が増加する。
//	 そのため、低い精度から高い精度に変換する際、精度差が大きすぎると変換後の拡張空間ID数は大幅に増大する。
//	 動作環境によってはメモリ不足となる可能性があるため、注意すること。
func VerticalZoom(inputZoom int64, vIndex int64, outputZoom int64) []string {
	verticalIDs := []string{}

	// 精度の差異を取得
	vZoomDiff := outputZoom - inputZoom

	minVparam := vIndex
	maxVparam := vIndex

	// 垂直方向の拡張空間ID分割数を定義
	vVoxelNum := int64(math.Pow(2, math.Abs(float64(vZoomDiff))))

	if vZoomDiff > 0 {
		// 垂直精度が上がった場合
		// 変換後の v 成分の最小値を定義
		minVparam = vIndex * vVoxelNum

		// 変換後の v 成分の最大値を定義
		maxVparam = minVparam + vVoxelNum - 1
	} else if vZoomDiff < 0 {
		// 垂直精度が下がった場合
		// 変換後の v 成分の最小値を定義
		minVparam = vIndex / vVoxelNum

		// 変換後の z 成分の最大値を定義
		maxVparam = minVparam
	}
	// 変換後の拡張空間IDを定義
	for v := minVparam; v <= maxVparam; v++ {
		// 生成した拡張空間IDを戻り値に格納
		verticalIDs = append(
			verticalIDs,
			strconv.FormatInt(outputZoom, 10)+"/"+strconv.FormatInt(v, 10))
	}
	return verticalIDs
}
