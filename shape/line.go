package shape

import (
	"math"

	"github.com/trajectoryjp/spatial_id_go/v3/common"
	"github.com/trajectoryjp/spatial_id_go/v3/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v3/common/object"
	"github.com/trajectoryjp/spatial_id_go/v3/common/spatial"
	"github.com/trajectoryjp/spatial_id_go/v3/operated"
)

const (
	// LonMinima 経度閾値
	LonMinima = 0.00000002
	// LatMinima 緯度閾値
	LatMinima = 0.00000002
	// AltMinima 高さ閾値
	AltMinima = 0.003
	// HightZoomLonMinima 精度が高い場合の経度閾値
	HightZoomLonMinima = 0.000000005
	// HightZoomLatMinima 精度が高い場合の緯度閾値
	HightZoomLatMinima = 0.0000000005
	// HightZoomAltMinima 精度が高い場合の高さ閾値
	HightZoomAltMinima = 0.0005
)

// GetSpatialIdsOnLine 指定範囲の空間ID変換(線分)を取得する。
//
// 始点終点間を結んだ線分上の空間IDを取得する。
//
// 引数：
//
//	start： 始点
//	end： 終点
//	zoom： 精度レベル
//
// 戻り値：
//
//	空間ID集合
//
// 戻り値（例外）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過：精度に 0 ～ 35 の整数値以外が入力されていた場合。
func GetSpatialIdsOnLine(
	start *object.Point,
	end *object.Point,
	zoom int64,
) ([]string, error) {

	// 拡張空間IDを取得
	ids, err := GetExtendedSpatialIdsOnLine(start, end, zoom, zoom)

	if err != nil {
		// エラーが発生した場合エラーインスタンスを返却
		return ids, err
	}

	// 拡張空間IDを空間IDのフォーマットに変換
	ids, err = ConvertExtendedSpatialIdsToSpatialIds(ids)

	return ids, err
}

// GetExtendedSpatialIdsOnLine 指定範囲の拡張空間ID変換(線分)を取得する。
//
// 始点終点間を結んだ線分上の拡張空間IDを取得する。
//
// 引数：
//
//	start： 始点
//	end： 終点
//	hZoom： 垂直方向の精度レベル
//	vZoom： 水平方向の精度レベル
//
// 戻り値：
//
//	拡張空間ID集合
//
// 戻り値（例外）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過：水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
func GetExtendedSpatialIdsOnLine(
	start *object.Point,
	end *object.Point,
	hZoom int64,
	vZoom int64,
) ([]string, error) {

	// nilチェック処理
	if start == nil || end == nil {
		return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 始点・終点の拡張空間IDを格納
	spatialIDs, e := GetExtendedSpatialIdsOnPoints([]*object.Point{start, end}, hZoom, vZoom)
	if e != nil {
		return spatialIDs, e
	}

	// 拡張空間IDをユニーク化
	spatialIDs = common.Unique(spatialIDs)

	// 始点・終点が同じ拡張空間IDかチェック
	if len(spatialIDs) == 1 {
		// 同じ拡張空間IDの場合、 始点・終点の拡張空間IDを返却
		return spatialIDs, nil
	}

	startSpatial := spatial.Point3{X: start.Lon(), Y: start.Lat(), Z: start.Alt()}
	endSpatial := spatial.Point3{X: end.Lon(), Y: end.Lat(), Z: end.Alt()}

	// 中点取得時の閾値初期化
	lonMinima := LonMinima
	latMinima := LatMinima
	altMinima := AltMinima

	// 入力された精度に応じて中点取得時の閾値を設定する
	if hZoom >= 31 {
		lonMinima = HightZoomLonMinima
		latMinima = HightZoomLatMinima
	}

	if vZoom >= 34 {
		altMinima = HightZoomAltMinima
	}

	// 【経度緯度空間】始点・終点の中点を拡張空間IDを再帰的に取得する。
	middleSpatialIds(
		startSpatial, endSpatial, hZoom, vZoom,
		lonMinima, latMinima, altMinima, func(spatialID string) {
			spatialIDs = append(spatialIDs, spatialID)
		})

	return common.Unique(spatialIDs), nil
}

// middleSpatialIds 始点終点の中点の拡張空間IDを再帰的に取得する。
//
// 始点終点の中点の拡張空間IDを再帰的に取得する。
// 始点終点間の距離が引数に入力された閾値以下となった場合、
// または、中点の拡張空間IDが始点終点の周囲に存在するか始点終点と同じIDであった場合、
// それまでに取得した拡張空間IDを全て返却して処理を終了する。
//
// 引数：
//
//	start： 始点
//	end： 終点
//	hZoom： 垂直方向の精度レベル
//	vZoom： 水平方向の精度レベル
//	lonMinima: 中点取得時の経度閾値
//	latMinima: 中点取得時の緯度閾値
//	altMinima: 中点取得時の高さ閾値
//	operate： 中点の拡張空間IDに対する処理関数
func middleSpatialIds(
	start, end spatial.Point3,
	hZoom, vZoom int64,
	lonMinima, latMinima, altMinima float64,
	operate func(string),
) {

	// 中点
	middle := spatial.NewLineFromPoints(start, end).ToPoint(0.5)

	// Pointオブジェクトへ型変換
	startPoint, _ := object.NewPoint(start.X, start.Y, start.Z)
	endPoint, _ := object.NewPoint(end.X, end.Y, end.Z)
	middlePoint, _ := object.NewPoint(middle.X, middle.Y, middle.Z)

	// 拡張空間ID取得
	lineSpatialIDs, _ := GetExtendedSpatialIdsOnPoints(
		[]*object.Point{startPoint, endPoint, middlePoint},
		hZoom,
		vZoom,
	)

	// 中点の拡張空間ID
	middleSpatialID := lineSpatialIDs[2]

	// 中間拡張空間IDへの操作
	operate(middleSpatialID)

	// 始点・終点間のベクトル
	vector := spatial.NewVectorFromPoints(start, end)
	// ベクトルの成分が閾値未満の場合は処理終了
	if math.Abs(vector.X) < lonMinima &&
		math.Abs(vector.Y) < latMinima &&
		math.Abs(vector.Z) < altMinima {
		return
	}

	// 始点の周囲6つの拡張空間ID
	start6SpatialIDs := operated.Get6spatialIdsAdjacentToFaces(lineSpatialIDs[0])
	// 始点自身も含める
	start6SpatialIDs = append(start6SpatialIDs, lineSpatialIDs[0])

	// 終点の周囲6つの拡張空間ID
	end6SpatialIDs := operated.Get6spatialIdsAdjacentToFaces(lineSpatialIDs[1])
	// 終点自身も含める
	end6SpatialIDs = append(end6SpatialIDs, lineSpatialIDs[1])

	// 中点が始点・終点の周囲にある場合
	if common.Include(start6SpatialIDs, middleSpatialID) &&
		common.Include(end6SpatialIDs, middleSpatialID) {
		return

		// 中点が始点の周囲にある場合
	} else if common.Include(start6SpatialIDs, middleSpatialID) {
		// 中点と終点で再帰的に中間拡張空間ID取得
		middleSpatialIds(
			spatial.Point3(middle), end, hZoom, vZoom,
			lonMinima, latMinima, altMinima, operate)

		// 中点が終点の周囲にある場合
	} else if common.Include(end6SpatialIDs, middleSpatialID) {
		// 始点と中点で再帰的に中間拡張空間ID取得
		middleSpatialIds(
			start, spatial.Point3(middle), hZoom, vZoom,
			lonMinima, latMinima, altMinima, operate)

		// 中点が始点・終点の周囲にない場合
	} else {
		// 始点と中点で再帰的に中間拡張空間ID取得
		middleSpatialIds(
			start, spatial.Point3(middle), hZoom, vZoom,
			lonMinima, latMinima, altMinima, operate)
		// 中点と終点で再帰的に中間拡張空間ID取得
		middleSpatialIds(
			spatial.Point3(middle), end, hZoom, vZoom,
			lonMinima, latMinima, altMinima, operate)
	}
}
