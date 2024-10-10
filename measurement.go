package spatialID

import (
	"math"

	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/mathematics"
)

type Measurement struct {
	// 原点。最小座標です
	Origin coordinates.Geodetic
	// x成分[m]
	XMeter float64
	// y成分[m]
	YMeter float64
	// z成分[m]
	ZMeter float32
	// x成分[°]
	XDegree float64
	// y成分[°]
	YDegree float64
}

func NewMeasurementFromSpatialID(spatialID SpatialID) *Measurement {
	measurement := &Measurement{}

	max := float64(1 << spatialID.GetZ())
	*measurement.Origin.Longitude() = 360.0 * float64(spatialID.GetX()) / max - 180.0
	*measurement.Origin.Latitude() = mathematics.RadianPerDegree * math.Atan(math.Sinh(math.Pi * (1.0 - 2.0*float64(spatialID.GetY())/max)))

	altitudeResolution := float64(1 << (SpatialIDZBaseExponent - spatialID.GetZ()))
	*measurement.Origin.Altitude() = float64(spatialID.GetZ()) * altitudeResolution

	// 垂直方向位置用の構造体初期化
	vPoint := object.VerticalPoint{}

	// 高さ全体の精度あたりの分解能を取得する
	vPoint.Resolution = math.Pow(2, 25.0) / math.Pow(2, float64(vZoom))

	// 高さを取得
	vPoint.Alt = float64(altIndex) * vPoint.Resolution

	// 垂直方向インスタンスを返却
	return vPoint
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

	return Measurement{
		Origin:  spatialID.Origin,
		XMeter:  spatialID.XMeter,
		YMeter:  spatialID.YMeter,
		ZMeter:  spatialID.ZMeter,
		XDegree: spatialID.XDegree,
		YDegree: spatialID.YDegree,
	}
}
