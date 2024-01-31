// Package shape 空間IDを任意形状で操作するパッケージ。
//
// 本パッケージは水平、垂直方向精度を一意に設定する'空間ID'及び、
// 水平、垂直方向精度を個別に指定可能な'拡張空間ID'の操作機能を持つ。
//
// 本書における'空間ID'、'拡張空間ID'は以下のフォーマットで構成される。
//
//	空間ID：
//	 [精度]/[高さの位置]/[経度の位置]/[緯度の位置]
//
//	 例)
//	  10/15/20/30
//
//	拡張空間ID：
//	 [水平方向精度]/[経度の位置]/[緯度の位置]/[垂直方向精度]/[高さの位置]
//
//	 例)
//	  10/20/30/10/15
package shape

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/common"
	"github.com/trajectoryjp/spatial_id_go/common/consts"
	"github.com/trajectoryjp/spatial_id_go/common/enum"
	"github.com/trajectoryjp/spatial_id_go/common/errors"
	"github.com/trajectoryjp/spatial_id_go/common/object"

	"github.com/wroge/wgs84"
)

// GetSpatialIdsOnPoints 空間ID取得関数
//
// 地理座標インスタンスのリストから空間IDのリストを取得する。
//
// 引数：
//
//	pointList：地理座標が格納されたインスタンスのリスト。
//	zoom     ：精度レベル。
//
// 戻り値：
//
//	空間IDのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過：精度に 0 ～ 35 の整数値以外が入力されていた場合。
func GetSpatialIdsOnPoints(
	pointList []*object.Point,
	zoom int64,
) ([]string, error) {
	// 拡張空間IDを取得
	ids, err := GetExtendedSpatialIdsOnPoints(pointList, zoom, zoom)

	if err != nil {
		// エラーが発生した場合エラーインスタンスを返却
		return ids, err
	}

	// 拡張空間IDを空間IDのフォーマットに変換
	ids, err = ConvertExtendedSpatialIdsToSpatialIds(ids)

	return ids, err
}

// GetPointOnSpatialId 空間IDの座標取得関数
//
// 単一の空間IDからその空間IDの頂点座標、または中心座標を取得する。
// 頂点の座標は南緯、東経、上空方向は近接する頂点の座標と共有される。
//
// 座標は地理座標で返却される。
//
// 引数：
//
//	spatialId：空間ID。
//	option   ：共通クラスから取得したPointOptionの値。
//
// 戻り値：
//
//	空間IDの各頂点の座標が格納されたインスタンスのリスト、または空間IDの中心点の座標が格納されたインスタンスのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 空間IDフォーマット不正：空間IDのフォーマットに違反する値が"空間ID"に入力されていた場合。
//	 不正なoption入力      ：指定外のoptionを指定した場合。
func GetPointOnSpatialId(
	spatialId string,
	option enum.PointOption,
) ([]*object.Point, error) {
	// 戻り値初期化
	points := []*object.Point{}

	// 空間IDを拡張空間IDのフォーマットに変換
	id, err := ConvertSpatialIdsToExtendedSpatialIds([]string{spatialId})

	// 入力値チェック
	if err != nil {
		// 空間IDのチェック時にエラーが返却された場合、空配列とエラーインスタンスを返却
		return points, err
	}

	// 座標を取得
	points, err = GetPointOnExtendedSpatialId(id[0], option)

	return points, err
}

// GetExtendedSpatialIdsOnPoints 拡張空間ID取得関数
//
// 地理座標インスタンスのリストから拡張空間IDのリストを取得する。
//
// 引数：
//
//	pointList：地理座標が格納されたインスタンスのリスト。
//	hZoom    ：水平方向の精度レベル。
//	vZoom    ：垂直方向の精度レベル。
//
// 戻り値：
//
//	拡張空間IDのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過：水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
func GetExtendedSpatialIdsOnPoints(
	pointList []*object.Point,
	hZoom int64,
	vZoom int64,
) ([]string, error) {

	// 拡張空間IDを格納するスライス
	spatialIds := make([]string, 0, len(pointList))

	// 入力値チェック
	// 水平、垂直方向精度のどちらかが範囲外の場合、空配列とエラーインスタンスを返却
	if !CheckZoom(hZoom) || !CheckZoom(vZoom) {
		return spatialIds, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	} else if common.Include(pointList, nil) {
		return spatialIds, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 先頭から拡張空間IDに変換する
	for _, v := range pointList {
		// 水平方向、垂直方向タイルID格納用スライス
		ids := []string{}

		// 水平方向のタイルを取得する
		ids = append(ids, getHorizontalTileIdOnPoint(v.Lon(), v.Lat(), hZoom))

		// 垂直方向のタイルを取得する
		ids = append(ids, getVerticalTileIdOnAltitude(v.Alt(), vZoom))

		// 変換した拡張空間IDを戻り値に格納
		spatialIds = append(spatialIds, strings.Join(ids, consts.SpatialIDDelimiter))
	}

	return spatialIds, nil
}

// GetPointOnExtendedSpatialId 拡張空間IDの座標取得関数
//
// 単一の拡張空間IDからその空間IDの頂点座標、または中心座標を取得する。
// 頂点の座標は南緯、東経、上空方向は近接する頂点の座標と共有される。
//
// 座標は地理座標で返却される。
//
// 引数：
//
//	extendedSpatialId：拡張空間ID。
//	option           ：共通クラスから取得したPointOptionの値。
//
// 戻り値：
//
//	拡張空間IDの各頂点の座標が格納されたインスタンスのリスト、または拡張空間IDの中心点の座標が格納されたインスタンスのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過              ：水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 拡張空間IDフォーマット不正：拡張空間IDのフォーマットに違反する値が"拡張空間ID"に入力されていた場合。
//	 不正なoption入力          ：指定外のoptionを指定した場合。
func GetPointOnExtendedSpatialId(
	extendedSpatialId string,
	option enum.PointOption,
) ([]*object.Point, error) {

	// 座標を格納するスライス
	outputPoint := []*object.Point{}

	// 拡張空間IDから座標成分を取得する。
	// 拡張空間IDは"[水平精度]/[経度インデックス]/[緯度インデックス]/[垂直精度]/[高さインデックス]"形式とする。
	// 拡張空間IDをx水平方向、y水平方向、垂直方向に分割する。
	id, err := getExtendedSpatialIdAttrs(extendedSpatialId)

	// 入力値チェック
	if err != nil {
		// 拡張空間IDの成分取得時にエラーが返却された場合、空配列とエラーインスタンスを返却
		return outputPoint, err
	}

	// IDから水平方向の位置と精度を取得する。前方2桁は精度。
	hZoom := id[0]
	lonIndex := id[1]
	latIndex := id[2]

	// IDから垂直方向の位置、分解能を取得する。前方2桁は精度。
	vZoom := id[3]
	altIndex := id[4]

	// 入力値チェック
	if !CheckZoom(hZoom) || !CheckZoom(vZoom) {
		// 水平、垂直方向精度のどちらかが範囲外の場合、空配列とエラーインスタンスを返却。
		return outputPoint, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 垂直方向インスタンス取得。
	vPoint := getAltitudeOnVerticalIndexAndZoom(altIndex, vZoom)

	// オプションで呼び出しの分岐。
	if option == enum.Center {
		// 中心点の座標を取得する。
		outputPoint = append(outputPoint, getCenterPointOnVoxelOffset(
			lonIndex, latIndex, hZoom, vPoint))

	} else if option == enum.Vertex {
		// 頂点の座標(計8個)を取得する。
		outputPoint = getVertexOnVoxelOffset(lonIndex, latIndex, hZoom, vPoint)

	} else {
		// サポート対象外のオプションが指定された場合、空配列とエラーインスタンスを返却。
		return outputPoint,
			errors.NewSpatialIdError(errors.OptionFailedErrorCode, "")
	}

	// 拡張空間IDから取得した座標群を返却。
	return outputPoint, nil
}

// ConvertPointListToProjectedPointList 地理座標系リストを投影座標系リストに変換する関数
//
// 地理座標系のリストをユーザ指定の投影座標のリストに変換する。
// 変換前の地理座標のCRSはEPSG:4326固定。
//
// 引数：
//
//	pointList   ：地理座標のデータクラスインスタンスのリスト。
//	projectedCrs：出力の投影座標のCRSのEPSGコード。
//
// 戻り値：
//
//	投影座標のデータクラスインスタンスのリストを返却する。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 不正なEPSG入力：入力されたEPSGコードが存在しない場合。
//	 座標変換に失敗：投影座標への座標変換に失敗した場合。
func ConvertPointListToProjectedPointList(
	pointList []*object.Point,
	projectedCrs int,
) ([]*object.ProjectedPoint, error) {

	// 戻り値格納用変数
	proPointList := make([]*object.ProjectedPoint, 0, len(pointList))

	// 地理座標系、投影座標系をwgs84の座標参照型でインスタンス化
	geoCrsCode := wgs84.EPSG().Code(consts.GeoCrs)
	proCrsCode := wgs84.EPSG().Code(projectedCrs)

	// 入力された地理座標リスト参照
	for _, p := range pointList {

		// 地理座標を投影座標に変換する。
		x, y, _, err := wgs84.SafeTransform(geoCrsCode, proCrsCode)(p.Lon(), p.Lat(), p.Alt())

		if err != nil {
			// 座標変換に失敗した場合エラーインスタンスを返却
			return proPointList,
				errors.NewSpatialIdError(errors.ValueConvertErrorCode, "")
		}

		// 高さは変換前後で不変のため、返還前の値を使用する
		proPointList = append(proPointList, &object.ProjectedPoint{
			X:   x,
			Y:   y,
			Alt: p.Alt(),
		})
	}

	return proPointList, nil
}

// ConvertProjectedPointListToPointList 投影座標系リストを地理座標系リストに変換する関数
//
// ユーザ指定の投影座標系のリストを地理座標のリストに変換する。
// 変換後の地理座標のCRSはEPSG:4326固定。
//
// 引数：
//
//	projectedPointList：投影座標のデータクラスインスタンスのリスト。
//	projectedCrs      ：入力の投影座標のCRSのEPSGコード。
//
// 戻り値：
//
//	地理座標のデータクラスインスタンスのリストを返却する。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 不正なEPSG入力：入力されたEPSGコードが存在しない場合。
//	 座標変換に失敗：地理座標への座標変換に失敗した場合。
func ConvertProjectedPointListToPointList(
	projectedPointList []*object.ProjectedPoint,
	projectedCrs int,
) ([]*object.Point, error) {

	// 戻り値格納用変数
	pointList := make([]*object.Point, 0, len(projectedPointList))

	// 地理座標系、投影座標系をwgs84の座標参照型でインスタンス化
	geoCrsCode := wgs84.EPSG().Code(consts.GeoCrs)
	proCrsCode := wgs84.EPSG().Code(projectedCrs)

	// 入力された投影座標リスト参照
	for _, p := range projectedPointList {

		// 投影座標を地理座標に変換する。
		x, y, _, err := wgs84.SafeTransform(proCrsCode, geoCrsCode)(p.X, p.Y, p.Alt)

		if err != nil {
			// 座標変換に失敗した場合エラーインスタンスを返却
			return pointList,
				errors.NewSpatialIdError(errors.ValueConvertErrorCode, "")
		}

		// 変換後の座標を持つ投影座標用インスタンスを戻り値に格納
		newPoint, _ := object.NewPoint(x, y, p.Alt)
		pointList = append(pointList, newPoint)
	}

	return pointList, nil
}

// CheckZoom 入力精度チェック関数
//
// 入力の精度が0-35の範囲内か判定をする。
//
// 引数：
//
//	zoom：入力された精度
//
// 戻り値：
//
//	入力された精度が0-35の範囲内の場合True、それ以外の場合Falseを返却する。
func CheckZoom(zoom int64) bool {
	// 精度が0-35に収まっている場合はTrueを返却
	return 0 <= zoom && zoom <= 35
}

// getExtendedSpatialIdAttrs 拡張空間IDフォーマットチェック関数
//
// 入力された拡張空間IDの各成分を格納した配列を返却する。
// 以下の条件に当てはまる場合はフォーマット違反となりエラーインスタンスが返却される。
//
//	・拡張空間IDの各成分に数値が入力されていない場合
//	・区切り文字の数が4つで無い場合
//
// 引数：
//
//	extendedSpatialId：拡張空間ID
//
// 戻り値：
//
//	拡張空間IDに含まれる水平、垂直精度、X, Y, Z成分を格納した以下の配列
//	 [水平精度, X成分, Y成分, 垂直精度, Z成分]
//	入力された拡張空間IDのフォーマットが不正な場合、エラーインスタンスを返却。
func getExtendedSpatialIdAttrs(extendedSpatialId string) ([]int64, error) {
	// 戻り値格納用
	result := []int64{}

	// 拡張空間IDをx水平方向、y水平方向、垂直方向に分割する。
	id := strings.Split(extendedSpatialId, consts.SpatialIDDelimiter)

	// 拡張空間IDのフォーマットチェック
	if len(id) != 5 {
		// 拡張空間ID区切り文字の数が不足していた場合エラーインスタンスを返却
		return result, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	for _, s := range id {
		// 空間IDの成分を整数型に変換
		i, err := strconv.ParseInt(s, 10, 64)

		if err == nil {
			// 成分の整数型変換結果を戻り値に格納
			result = append(result, i)
		} else {
			// 変換時にエラーが発生した場合、エラーインスタンスを返却
			return result, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
	}

	return result, nil
}

// getHorizontalTileIdOnPoint 水平方向タイルID取得関数
//
// 経度緯度から水平方向のタイルIDを計算する。
// 水平方向のタイルIDは「[精度]/[x位置]/[y位置]」となる。
//
// 引数：
//
//	lon  ：空間座標系の経度(単位:度)
//	lat  ：空間座標系の緯度(単位:度)
//	hZoom：水平方向の精度
//
// 戻り値：
//
//	水平方向のタイルID
func getHorizontalTileIdOnPoint(lon float64, lat float64, hZoom int64) string {

	// 経度に180が入力されているとタイルインデックス+1の値が出力されるため、補正する
	if lon == 180 {
		lon = -lon
	}

	// 経度方向のインデックスの計算
	lonIndex := math.Floor(math.Pow(2, float64(hZoom)) * ((lon + 180.0) / 360.0))

	// 緯度をラジアンに変換
	latRadian := common.DegreeToRadian(lat)

	// 緯度方向のインデックスの計算
	latIndex := math.Floor(
		math.Pow(2, float64(hZoom)) *
			(1 - math.Log(math.Tan(latRadian)+(1/math.Cos(latRadian)))/math.Pi) / 2)

	// 水平精度、経度方向、緯度方向のインデクスをスライスに格納
	idParams := []string{
		strconv.FormatInt(hZoom, 10),
		strconv.FormatInt(int64(lonIndex), 10),
		strconv.FormatInt(int64(latIndex), 10),
	}

	// 水平精度及び、経度方向と緯度方向のIDを結合して水平方向タイルIDを生成
	tileId := strings.Join(idParams, consts.SpatialIDDelimiter)

	// 水平方向タイルIDを返却
	return tileId
}

// getVerticalTileIdOnAltitude 垂直方向タイルID取得関数
//
// 垂直方向のタイルIDを計算する。
// 垂直方向のタイルIDは「[精度]/[位置]」となる。
//
// 引数：
//
//	alt  ：高さ(単位:m)
//	vZoom：垂直方向の精度
//
// 戻り値：
//
//	垂直方向のタイルID
func getVerticalTileIdOnAltitude(alt float64, vZoom int64) string {

	// 高さ全体の精度あたりの垂直方向の精度
	altResolution := math.Pow(2, 25.0) / math.Pow(2, float64(vZoom))

	// 垂直方向の位置を計算する
	vIndex := math.Floor(alt / altResolution)

	// 垂直精度、高さ方向のインデクスをスライスに格納
	idParams := []string{
		strconv.FormatInt(vZoom, 10),
		strconv.FormatInt(int64(vIndex), 10),
	}

	// 垂直精度及び、高さ方向のIDを結合して垂直方向タイルIDを生成
	tileId := strings.Join(idParams, consts.SpatialIDDelimiter)

	// 垂直方向タイルIDを返却
	return tileId
}

// getAltitudeOnVerticalIndexAndZoom 高さ及び分解能取得関数
//
// 垂直方向の位置と精度から原点に近い高さと分解能を取得する。
// 入力された位置が負の場合は、原点に遠い高さと分解能を取得する。
//
// 引数：
//
//	altIndex：高さの位置
//	vZoom   ：垂直方向精度
//
// 戻り値：
//
//	高さ(単位:m)と分解能が格納されたインスタンス
func getAltitudeOnVerticalIndexAndZoom(
	altIndex int64,
	vZoom int64,
) object.VerticalPoint {

	// 垂直方向位置用の構造体初期化
	vPoint := object.VerticalPoint{}

	// 高さ全体の精度あたりの分解能を取得する
	vPoint.Resolution = math.Pow(2, 25.0) / math.Pow(2, float64(vZoom))

	// 高さを取得
	vPoint.Alt = float64(altIndex) * vPoint.Resolution

	// 垂直方向インスタンスを返却
	return vPoint
}

// getCenterPointOnVoxelOffset ボクセル中心座標取得関数
//
// 緯度、経度、高さから精度ごとの一辺の長さのボクセルの中心点の座標を取得する。
//
// 引数：
//
//	lonIndex：経度方向のインデックス
//	latIndex：緯度方向のインデックス
//	hZoom   ：水平方向の精度
//	vPoint  ：垂直方向インスタンス
//
// 戻り値：
//
//	空間の中心点の座標が格納されたインスタンス
func getCenterPointOnVoxelOffset(
	lonIndex int64,
	latIndex int64,
	hZoom int64,
	vPoint object.VerticalPoint,
) *object.Point {

	// 頂点座標から最大値と最小値を取得する。
	pList := getVertexOnVoxelOffset(lonIndex, latIndex, hZoom, vPoint)

	// 頂点座標格納用。
	lonList := []float64{}
	latList := []float64{}
	altList := []float64{}

	// 頂点座標を配列に格納
	for _, v := range pList {
		lonList = append(lonList, v.Lon())
		latList = append(latList, v.Lat())
		altList = append(altList, v.Alt())
	}

	// 頂点座標をソートし、最大値と最小値を取得する。
	sort.Float64s(lonList)
	sort.Float64s(latList)
	sort.Float64s(altList)

	lonMin := lonList[0]
	lonMax := lonList[len(lonList)-1]
	latMin := latList[0]
	latMax := latList[len(latList)-1]
	altMin := altList[0]
	altMax := altList[len(altList)-1]

	// ボクセルの中心座標を計算
	centerLon := (lonMax + lonMin) / 2
	centerLat := (latMax + latMin) / 2
	centerAlt := (altMax + altMin) / 2
	point, _ := object.NewPoint(centerLon, centerLat, centerAlt)

	// 中心点の座標を返却
	return point
}

// getVertexOnVoxelOffset ボクセル頂点座標取得関数
//
// 緯度、経度、高さから分解能を一辺の長さとしたボクセルの頂点を取得する。
//
// 引数：
//
//	lonIndex：経度方向のインデックス
//	latIndex：緯度方向のインデックス
//	hZoom   ：水平方向の精度
//	vPoint  ：垂直方向インスタンス
//
// 戻り値：
//
//	各頂点の座標が設定されたインスタンスのリスト
func getVertexOnVoxelOffset(
	lonIndex int64,
	latIndex int64,
	hZoom int64,
	vPoint object.VerticalPoint,
) []*object.Point {

	// 返却用のリスト。要素数は頂点の数に設定
	pList := make([]*object.Point, 0, 8)

	// 経度、緯度の判定用の境界値
	hLimit := math.Pow(2, float64(hZoom))

	// 内部計算用の経度方向、緯度方向インデックス格納用
	lonIndexFloat := float64(lonIndex)
	latIndexFloat := float64(latIndex)

	// 緯度の取得
	if (hLimit - 1) <= latIndexFloat {
		latIndexFloat = hLimit - 1
	} else if latIndexFloat < 0.0 {
		latIndexFloat = 0.0
	}

	// タイルの上辺の緯度
	northLat := common.RadianToDegree(
		math.Atan(math.Sinh(math.Pi * (1 - 2*latIndexFloat/hLimit))))

	// タイルの下辺の緯度
	southLat := common.RadianToDegree(
		math.Atan(math.Sinh(math.Pi * (1 - 2*(latIndexFloat+1)/hLimit))))

	// 経度の取得
	if (hLimit-1) <= lonIndexFloat || lonIndexFloat < 0.0 {
		// インデックスの範囲を超えている場合はn周分を無視する
		for lonIndexFloat < 0 {
			lonIndexFloat += hLimit
		}
		lonIndexFloat = math.Mod(lonIndexFloat, hLimit)
	}

	westLon := lonIndexFloat*360/hLimit - 180
	eastLon := (lonIndexFloat+1.0)*360/hLimit - 180

	// ボクセル上面の高さを求める
	vTopAlt := vPoint.Alt + vPoint.Resolution

	// ボクセルの各頂点を定義
	northWestBottom, _ := object.NewPoint(westLon, northLat, vPoint.Alt)
	northEastBottom, _ := object.NewPoint(eastLon, northLat, vPoint.Alt)
	southEastBottom, _ := object.NewPoint(eastLon, southLat, vPoint.Alt)
	southWestBottom, _ := object.NewPoint(westLon, southLat, vPoint.Alt)
	northWestTop, _ := object.NewPoint(westLon, northLat, vTopAlt)
	northEastTop, _ := object.NewPoint(eastLon, northLat, vTopAlt)
	southEastTop, _ := object.NewPoint(eastLon, southLat, vTopAlt)
	southWestTop, _ := object.NewPoint(westLon, southLat, vTopAlt)

	// ボクセルの各頂点を計算して返却用のリストに格納する。
	// ボクセルを上から見下ろし左上から時計回りの順かつ、底面4頂点→上面4頂点の順に格納する。
	pList = append(pList, northWestBottom)
	pList = append(pList, northEastBottom)
	pList = append(pList, southEastBottom)
	pList = append(pList, southWestBottom)
	pList = append(pList, northWestTop)
	pList = append(pList, northEastTop)
	pList = append(pList, southEastTop)
	pList = append(pList, southWestTop)

	// 頂点座標群を返却
	return pList
}

// ConvertSpatialIdsToExtendedSpatialIds 空間IDフォーマット変換関数
//
// 空間IDのフォーマットを拡張空間IDのフォーマットに変換する。
//
// 引数：
//
//	spatialIds：空間IDのリスト。
//
// 戻り値：
//
//	拡張空間IDのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力された空間IDのフォーマットが不正な場合。
func ConvertSpatialIdsToExtendedSpatialIds(spatialIds []string) ([]string, error) {
	// 拡張空間IDを格納するスライス
	resultIds := make([]string, 0, len(spatialIds))

	for _, s := range spatialIds {
		// 空間IDを拡張空間IDのフォーマットに変換して戻り値に格納
		c := strings.Split(s, consts.SpatialIDDelimiter)

		// 空間IDのフォーマットチェック
		if len(c) != 4 {
			// 空間ID区切り文字の数が不足していた場合エラーインスタンスを返却
			return resultIds, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}

		l := []string{c[0], c[2], c[3], c[0], c[1]}
		i := strings.Join(l, consts.SpatialIDDelimiter)
		resultIds = append(resultIds, i)
	}

	return resultIds, nil
}

// ConvertExtendedSpatialIdsToSpatialIds 拡張空間IDフォーマット変換関数
//
// 拡張空間IDのフォーマットを空間IDのフォーマットに変換する。
//
// 引数：
//
//	extendedSpatialIds：拡張空間IDのリスト。
//
// 戻り値：
//
//	空間IDのリスト。
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力された拡張空間IDのフォーマットが不正な場合。
func ConvertExtendedSpatialIdsToSpatialIds(
	extendedSpatialIds []string,
) ([]string, error) {
	// 空間IDを格納するスライス
	resultIds := make([]string, 0, len(extendedSpatialIds))

	for _, s := range extendedSpatialIds {
		// 拡張空間IDを空間IDのフォーマットに変換して戻り値に格納
		c := strings.Split(s, consts.SpatialIDDelimiter)

		// 拡張空間IDのフォーマットチェック
		if len(c) != 5 {
			// 拡張空間ID区切り文字の数が不足していた場合エラーインスタンスを返却
			return resultIds, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}

		l := []string{c[0], c[4], c[1], c[2]}
		i := strings.Join(l, consts.SpatialIDDelimiter)
		resultIds = append(resultIds, i)
	}

	return resultIds, nil
}
