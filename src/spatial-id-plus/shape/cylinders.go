package shape

import (
	// "fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id-plus/shape/physics"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/consts"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/enum"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/errors"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/logger"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/object"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/spatial"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/operated"
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/shape"
)

// convertSpatialProjectionPoint PointオブジェクトからSpatialPointオブジェクトへの変換関数
//
// 引数：
//
//	point： Point3オブジェクト
//
// 戻り値：
//
//	projectedPointオブジェクト
func convertSpatialProjectionPoint(point spatial.Point3) object.ProjectedPoint {
	return object.ProjectedPoint{point.X, point.Y, point.Z}
}

// GetSpatialIDOnAxisIDs 軸IDから拡張空間ID取得
//
// 引数：
//
//	xID: X軸ボクセルID
//	yID: Y軸ボクセルID
//	zID: Z軸ボクセルID
//	hZoom: 水平方向精度
//	vZoom: 垂直方向精度
//
// 戻り値：
//
//	空間ID
func GetSpatialIDOnAxisIDs(xID, yID, zID, hZoom, vZoom int64) string {
	spatialIDs := []string{
		strconv.FormatInt(hZoom, 10),
		strconv.FormatInt(xID, 10),
		strconv.FormatInt(yID, 10),
		strconv.FormatInt(vZoom, 10),
		strconv.FormatInt(zID, 10),
	}

	return strings.Join(spatialIDs, consts.SpatialIDDelimiter)
}

// GetVoxelIDToSpatialID ボクセル成分ID取得
//
// 拡張空間IDからボクセル成分ID取得
//
// 引数：
//
//	spatialID： 取得するボクセルの拡張空間ID
//
// 戻り値：
//
//	(xインデックス, yインデックス, vインデックス)
func GetVoxelIDToSpatialID(spatialID string) []int64 {
	ids := strings.Split(spatialID, consts.SpatialIDDelimiter)
	lonIndex, _ := strconv.ParseInt(ids[1], 10, 64)
	latIndex, _ := strconv.ParseInt(ids[2], 10, 64)
	altIndex, _ := strconv.ParseInt(ids[4], 10, 64)

	return []int64{
		lonIndex,
		latIndex,
		altIndex,
	}
}

// Rectangular 直方体構造体
type Rectangular struct {
	start         spatial.Point3 // 始点
	end           spatial.Point3 // 終点
	radius        float64        // 半径
	hZoom         int64          // 水平精度
	vZoom         int64          // 垂直精度
	height        float64        // 高さ
	allSpatialIDs []string       // 全空間ID
	factor        float64        // Webメルカトル換算係数
}

// NewRectangular 直方体構造体コンストラクタ
//
// 直方体構造体作成
//
// 引数：
//
//	startPoint： 始点
//	endPoint： 終点
//	radius： 半径
//	hZoom： 水平精度
//	vZoom： 垂直精度
//	factor: Webメルカトル換算係数
//
// 戻り値：
//
//	直方体構造体ポインタ
func NewRectangular(
	startPoint spatial.Point3,
	endPoint spatial.Point3,
	radius float64,
	hZoom int64,
	vZoom int64,
	factor float64,
) *Rectangular {

	rect := new(Rectangular)
	rect.start = startPoint
	rect.end = endPoint

	// 【直交座標空間】軸ベクトル
	axis := spatial.NewVectorFromPoints(rect.start, rect.end)
	// 縦の長さ
	rect.height = axis.Norm()

	rect.radius = radius

	// 分解能取得
	rect.hZoom = hZoom
	rect.vZoom = vZoom

	// Webメルカトル換算係数
	rect.factor = factor

	return rect
}

// calcUnitVoxelVector 概算距離用の単位ボクセルベクトルを算出
//
// 引数：
//
//	spatialID： 拡張空間ID
//
// 戻り値：
//
//	単位ボクセルベクトル： spatial.Vector3
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 空間IDフォーマット不正：空間IDのフォーマットに違反する値が"拡張空間ID"に入力されていた場合。
func (r Rectangular) calcUnitVoxelVector(spatialID string) (spatial.Vector3, error) {
	points, err := shape.GetPointOnExtendedSpatialId(
		spatialID,
		enum.Vertex,
	)
	if err != nil {
		return spatial.Vector3{}, err
	}
	orthPoints, _ := shape.ConvertPointListToProjectedPointList(points, consts.OrthCrs)
	// 各XYZ成分座標から単位ボクセルベクトルを算出
	xApexes := make([]float64, 0, 8)
	yApexes := make([]float64, 0, 8)
	zApexes := make([]float64, 0, 8)
	for _, orthPoint := range orthPoints {
		xApexes = append(xApexes, orthPoint.X)
		yApexes = append(yApexes, orthPoint.Y)
		zApexes = append(zApexes, orthPoint.Alt)
	}
	maxX, _ := common.Max(xApexes)
	minX, _ := common.Min(xApexes)
	maxY, _ := common.Max(yApexes)
	minY, _ := common.Min(yApexes)
	maxZ, _ := common.Max(zApexes)
	minZ, _ := common.Min(zApexes)

	logger.Debug("単位ボクセルベクトル(X: %v, Y: %v, Z: %v)",
		math.Floor((maxX - minX) * math.Pow10(6)) / math.Pow10(6),
		math.Floor((maxY - minY) * math.Pow10(6)) / math.Pow10(6),
		math.Floor((maxZ - minZ) * math.Pow10(6)) * r.factor / math.Pow10(6))

	return spatial.Vector3{
		X: math.Floor((maxX-minX)*math.Pow10(6)) / math.Pow10(6),
		Y: math.Floor((maxY-minY)*math.Pow10(6)) / math.Pow10(6),
		Z: math.Floor((maxZ-minZ)*math.Pow10(6)) * r.factor / math.Pow10(6),
	}, nil
}

// Capsule カプセル構造体
type Capsule struct {
	*Rectangular                      // 直方体構造体の埋め込み
	isCapsule         bool            // カプセル判定
	isSphere          bool            // 球判定
	isPrecision       bool            // 衝突判定実施オプション
	object            physics.Physics // 衝突判定オブジェクト
	includeSpatialIDs []string        // 内部判定空間ID
}

// NewCapsule カプセル構造体コンストラクタ
//
// カプセル構造体作成
//
// 引数：
//
//	startPoint： 始点
//	endPoint： 終点
//	radius： 半径
//	hZoom： 水平精度
//	vZoom： 垂直精度
//	isCapsule： カプセル判定
//	isPrecision： 衝突判定実施オプション
//	factor： Webメルカトル換算係数
//
// 戻り値：
//
//	カプセル構造体ポインタ
func NewCapsule(
	startPoint spatial.Point3,
	endPoint spatial.Point3,
	radius float64,
	hZoom int64,
	vZoom int64,
	isCapsule bool,
	isPrecision bool,
	factor float64,
) *Capsule {

	// 球判定
	isSphere := false
	// 物理オブジェクトインターフェース
	var object physics.Physics

	// 球の場合（始点と終点が同一の場合）
	if startPoint.IsClose(endPoint, consts.Minima) {
		// 物理オブジェクトを球物理オブジェクト
		object = physics.NewSpherePhysics(radius*factor, startPoint)
		isSphere = true

		// カプセルの場合
	} else if isCapsule {
		// 物理オブジェクトをカプセル物理オブジェクト
		object = physics.NewCapsulePhysics(radius*factor, startPoint, endPoint)

		// 円柱の場合
	} else {
		// 物理オブジェクトを円柱物理オブジェクト
		object = physics.NewCylinderPhysics(radius*factor, startPoint, endPoint)
	}

	capsule := new(Capsule)
	capsule.Rectangular = NewRectangular(
		startPoint,
		endPoint,
		radius,
		hZoom,
		vZoom,
		factor,
	)
	capsule.isCapsule = isCapsule
	capsule.isSphere = isSphere
	capsule.isPrecision = isPrecision
	capsule.object = object

	return capsule
}

// IsEmpty 空確認
//
// オブジェクトが空であるかを確認
//
// 戻り値：
//
//	True:オブジェクトが空 False:オブジェクトが空でない
func (c Capsule) IsEmpty() bool {
	return c.Rectangular == nil
}

// calcLineSpatialIDs 始点・終点間の軸の空間IDを取得
//
// 始点・終点から軸の空間IDと内部空間用軸の空間IDを取得
//
// 戻り値：
//
//	始点・終点間の軸の空間ID
//	始点・終点間の内部空間用軸の空間ID
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過： 水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
func (c Capsule) calcLineSpatialIDs() ([]string, []string, error) {

	// 【直交座標空間】始点終点間の空間ID取得
	lineSpatialIDs := []string{}
	// 【直交座標空間】内部空間用始点終点間の空間ID取得
	// 	円柱かつ軸の長さが直径より小さい場合は内部空間用始点終点間の空間IDは空のままとする
	insideLineSpatialIDs := []string{}
	var err error

	// Z成分をfactorで補正
	start := spatial.Point3{c.start.X, c.start.Y, c.start.Z / c.factor}
	end := spatial.Point3{c.end.X, c.end.Y, c.end.Z / c.factor}

	// 球の場合
	if c.isSphere {
		// 【直交座標空間⇒緯度経度空間】始点終の座標
		projectesStart := convertSpatialProjectionPoint(start)
		projectedPoints := []*object.ProjectedPoint{
			&projectesStart,
		}
		wgs84Points, _ := shape.ConvertProjectedPointListToPointList(projectedPoints, consts.OrthCrs)
		lineSpatialIDs, err = shape.GetExtendedSpatialIdsOnPoints(wgs84Points, c.hZoom, c.vZoom)
		if err != nil {
			return []string{}, []string{}, err
		}
		// 内部空間用始点終点間の空間IDを更新
		insideLineSpatialIDs = lineSpatialIDs
		logger.Debug("球の場合、接続点の空間ID取得: %s", reflect.ValueOf(lineSpatialIDs))

		// カプセル・円柱の場合
	} else {

		// 【直交座標空間⇒緯度経度空間】始点終の座標
		projectesStart := convertSpatialProjectionPoint(start)
		projectesEnd := convertSpatialProjectionPoint(end)
		projectedPoints := []*object.ProjectedPoint{
			&projectesStart,
			&projectesEnd,
		}
		wgs84Points, _ := shape.ConvertProjectedPointListToPointList(projectedPoints, consts.OrthCrs)
		lineSpatialIDs, err =
			shape.GetExtendedSpatialIdsOnLine(wgs84Points[0], wgs84Points[1], c.hZoom, c.vZoom)
		if err != nil {
			return []string{}, []string{}, err
		}

		logger.Debug("カプセルの場合、始点終点間の空間ID取得: %s", reflect.ValueOf(lineSpatialIDs))

		// カプセルの場合
		if c.isCapsule {
			// 内部空間用始点終点間の空間IDを更新
			insideLineSpatialIDs = lineSpatialIDs

			// 円柱かつ軸の長さが直径より大きい場合
		} else if c.height > 2*c.radius*c.factor {

			// 軸のベクトル
			unitAxis := spatial.NewVectorFromPoints(c.start, c.end).Unit()
			radiusAxis := unitAxis.Scale(c.radius * c.factor)

			// 処理を半径分短くした軸の線に対して行う
			newStart := c.start.Translate(radiusAxis)
			factorStart := spatial.Point3{newStart.X, newStart.Y, newStart.Z / c.factor}
			newEnd := c.end.Translate(radiusAxis.Scale(-1))
			factorEnd := spatial.Point3{newEnd.X, newEnd.Y, newEnd.Z / c.factor}

			// 【直交座標空間⇒緯度経度空間】始点終点の座標
			projectesStart := convertSpatialProjectionPoint(factorStart)
			projectesEnd := convertSpatialProjectionPoint(factorEnd)
			projectedPoints := []*object.ProjectedPoint{
				&projectesStart,
				&projectesEnd,
			}
			wgs84Points, _ := shape.ConvertProjectedPointListToPointList(projectedPoints, consts.OrthCrs)
			// 内部空間用始点終点間の空間IDを更新
			insideLineSpatialIDs, _ =
				shape.GetExtendedSpatialIdsOnLine(wgs84Points[0], wgs84Points[1], c.hZoom, c.vZoom)
			logger.Debug("直径より長い円柱の場合、内部用始点終点間の空間ID取得: %s", reflect.ValueOf(lineSpatialIDs))
		}
	}

	return lineSpatialIDs, insideLineSpatialIDs, nil
}

// calcAllSpatialIDs 全空間IDを取得
//
// 始点・終点間の軸の空間IDから実際の図形を覆う全体の空間IDを取得
//
// 引数：
//
//	lineSpatialIDs： 始点・終点間の軸の空間ID
//	unitVoxel： 単位ボクセル
func (c *Capsule) calcAllSpatialIDs(lineSpatialIDs []string, unitVoxel spatial.Vector3) {
	// オブジェクトに外接する直方体の空間IDを全空間IDとして取得
	xApprNum := math.Ceil(c.radius * c.factor / unitVoxel.X)
	yApprNum := math.Ceil(c.radius * c.factor / unitVoxel.Y)
	zApprNum := math.Ceil(c.radius * c.factor / unitVoxel.Z)

	logger.Debug("X軸シフト数: %f, Y軸シフト数: %f, Z軸シフト数: %f", xApprNum, yApprNum, zApprNum)

	// 【直交空間】オブジェクトの全空間ID簡易取得
	for _, lineSpatialID := range lineSpatialIDs {
		for x := -xApprNum; x <= xApprNum; x++ {
			for y := -yApprNum; y <= yApprNum; y++ {
				for z := -zApprNum; z <= zApprNum; z++ {
					shiftSpatialID :=
						operated.GetShiftingSpatialID(
							lineSpatialID,
							int64(x),
							int64(y),
							int64(z),
						)
					c.allSpatialIDs = append(c.allSpatialIDs, shiftSpatialID)
				}
			}
		}
	}

	c.allSpatialIDs = common.Unique(c.allSpatialIDs)
}

// calcIncludeSpatialIDs 内部空間IDを取得
//
// 始点・終点間の軸の空間IDから実際の図形の内部の空間IDを取得
//
// 引数：
//
//	insideLineSpatialIDs： 始点・終点間の内部空間用軸の空間ID
//	unitVoxel： 単位ボクセル
func (c *Capsule) calcIncludeSpatialIDs(insideLineSpatialIDs []string, unitVoxel spatial.Vector3) {
	// オブジェクトに内接する直方体の空間IDを内部空間IDとして取得
	radius := c.radius / math.Sqrt2
	xApprNum := math.Floor(radius*c.factor/unitVoxel.X) - 1
	yApprNum := math.Floor(radius*c.factor/unitVoxel.Y) - 1
	zApprNum := math.Floor(radius*c.factor/unitVoxel.Z) - 1

	// 【直交空間】オブジェクトの内部空間ID簡易取得
	for _, lineSpatialID := range insideLineSpatialIDs {
		for x := -xApprNum; x <= xApprNum; x++ {
			for y := -yApprNum; y <= yApprNum; y++ {
				for z := -zApprNum; z <= zApprNum; z++ {
					shiftSpatialID := operated.GetShiftingSpatialID(lineSpatialID, int64(x), int64(y), int64(z))
					c.includeSpatialIDs = append(c.includeSpatialIDs, shiftSpatialID)
				}
			}
		}
	}
	c.includeSpatialIDs = common.Unique(c.includeSpatialIDs)
}

// calcCollideSpatialIDs 衝突する空間IDを取得
//
// 全空間IDから内部空間IDを除いた空間IDとオブジェクトで衝突判定を実施し、衝突した空間IDを結果に追加
func (c *Capsule) calcCollideSpatialIDs() {

	// Azul3Dと衝突判定を行う衝突空間ID(全空間IDから内部空間IDを除いた空間ID)
	excludeSpatialIDs := common.Difference(c.allSpatialIDs, c.includeSpatialIDs)

	logger.Debug("Azul3Dと衝突判定を行う空間ID: %s", reflect.ValueOf(excludeSpatialIDs))

	latDict := make(map[int64]spatial.Vector3)
	for _, excludeSpatialID := range excludeSpatialIDs {

		// 【直交座標空間】ボクセルの対角線のベクトルを決定
		lens := spatial.Vector3{}

		lat := GetVoxelIDToSpatialID(excludeSpatialID)[1]
		// 同一緯度のボクセルの対角線のベクトルが取得済みの場合
		if latVoxel, ok := latDict[lat]; ok {
			lens = latVoxel

			// 同一緯度のボクセルの対角線のベクトルが無い場合
		} else {
			lens, _ = c.calcUnitVoxelVector(excludeSpatialID)
			latDict[lat] = lens
		}
		// 【直交座標空間】ボクセルの中心座標
		centers, _ := shape.GetPointOnExtendedSpatialId(excludeSpatialID, enum.Center)
		orthCenters, _ := shape.ConvertPointListToProjectedPointList(centers, consts.OrthCrs)
		orthCenter := spatial.Point3{orthCenters[0].X, orthCenters[0].Y, orthCenters[0].Alt * c.factor}
		logger.Debug("ボクセルの中心座標: %f, %f, %f", orthCenter.X, orthCenter.Y, orthCenter.Z)

		if c.object.IsCollideVoxel(orthCenter, lens) {
			c.includeSpatialIDs = append(c.includeSpatialIDs, excludeSpatialID)
		}
	}

	// 円柱の場合は衝突判定の内部を埋める
	if !c.isCapsule && !c.isSphere {
		// 経度緯度をキーとした辞書初期化
		lotLatMap := map[string][]string{}

		// 空間IDで経度緯度が同一のものを集約
		for _, spatialID := range c.includeSpatialIDs {
			ids := strings.Split(spatialID, consts.SpatialIDDelimiter)
			lotLatMap[ids[1]+ids[2]] = append(lotLatMap[ids[1]+ids[2]], spatialID)
		}

		// 経度緯度が同一の空間ID内の高さ間の空間IDを取得
		for _, spatialIDs := range lotLatMap {
			zIndexes := []int64{}
			baseIndexes := GetVoxelIDToSpatialID(spatialIDs[0])

			// 経度緯度が同一の空間ID内の高さを集約
			for _, checkSpatialID := range spatialIDs {
				ids := GetVoxelIDToSpatialID(checkSpatialID)
				zIndexes = append(zIndexes, ids[2])
			}

			// 高さの最大値・最小値を取得
			minZ, _ := common.Min(zIndexes)
			maxZ, _ := common.Max(zIndexes)
			// 高さの最大値・最小値間の空間IDを結果に追加
			for zIndex := minZ; zIndex <= maxZ; zIndex++ {
				spatialID := GetSpatialIDOnAxisIDs(
					baseIndexes[0],
					baseIndexes[1],
					zIndex,
					c.hZoom,
					c.vZoom,
				)
				c.includeSpatialIDs = append(c.includeSpatialIDs, spatialID)
			}
		}
	}

	// 空間IDの重複を削除
	c.includeSpatialIDs = common.Unique(c.includeSpatialIDs)
}

// CalcValidSpatialIDs 有効な空間ID取得
//
// 始点・終点間の軸の空間IDから実際の図形分の有効な空間IDを取得
//
// 戻り値：
//
//	最終的にオブジェクトと衝突すると判定した空間ID
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過： 水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
func (c *Capsule) CalcValidSpatialIDs() ([]string, error) {

	// 始点・終点間の軸の空間IDを取得
	lineSpatialIDs, insideLineSpatialIDs, err := c.calcLineSpatialIDs()
	if err != nil {
		return []string{}, err
	}

	// 【直交座標空間】単位ボクセル
	baseLat := int64(0)
	if c.hZoom != 0 {
		baseLat = int64(math.Pow(2, float64(c.hZoom-1)))
	}
	baseSpatialID := GetSpatialIDOnAxisIDs(
		0,
		baseLat,
		0,
		c.hZoom,
		c.vZoom,
	)
	unitVoxel, _ := c.calcUnitVoxelVector(baseSpatialID)

	// オブジェクトに外接する直方体の空間IDを全空間IDとして取得
	c.calcAllSpatialIDs(lineSpatialIDs, unitVoxel)

	// 衝突判定実施オプションがfalseの場合は衝突判定をスキップ
	if !c.isPrecision {
		return c.allSpatialIDs, nil
	}

	// オブジェクトに内接する直方体の空間IDを内部空間IDとして取得
	c.calcIncludeSpatialIDs(insideLineSpatialIDs, unitVoxel)

	// 全空間IDから内部空間IDを除いた空間IDとオブジェクトで衝突判定
	c.calcCollideSpatialIDs()

	return c.includeSpatialIDs, nil
}

// IsPrecisionOpts 衝突判定実施オプショナル引数構造体
type IsPrecisionOpts struct {
	IsPrecision bool // 衝突判定実施オプション
}

// 衝突判定実施オプショナル型
type option func(*IsPrecisionOpts)

// IsPrecision isPrecision引数設定関数
//
// 以下の関数のisPrecision引数を設定する。
//   - GetSpatialIdsOnCylinders
//   - GetExtendedSpatialIdsOnCylinders
//
// 引数：
//
//	v: isPrecisionの値
//
// 戻り値：
//
//	衝突判定実施オプショナル型の関数
func IsPrecision(v bool) option {
	return func(p *IsPrecisionOpts) {
		p.IsPrecision = v
	}
}

// GetSpatialIdsOnCylinders 空間ID(円柱)取得
//
// 円柱を複数つなげた経路が通る空間IDを取得する。
// 円柱間の接続面は球状とする。
// ドローンの経路や地中埋設配管が通る経路を空間IDで表現する際に使用する。
//
// 引数：
//
//	center     : 円柱の中心の接続点。Pointを複数指定するリスト。
//	             2つ目の接続点は1つ目の円柱の終点となるが、2つ目の円柱の始点にもなる。
//	radius     : 円柱の半径(単位:m)
//	zoom       : 精度レベル
//	isCapsule  : 始点、終点が球状であるかを示す。True: カプセル / False: 円柱
//	isPrecision: 衝突判定を実施するかのフラグ。True: 実施 / False: 未実施(デフォルトはTrue)
//
// 戻り値：
//
//	円柱を複数つなげた経路が通る空間IDのリスト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力数値不正： 座標の値が不正の場合、もしくは円柱の半径が0以下の場合、エラー
//	 精度閾値超過： 精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 注意: 経度180度をまたがる点をまたがる円柱は空間ID化できない。
func GetSpatialIdsOnCylinders(
	center []*object.Point,
	radius float64,
	zoom int64,
	isCapsule bool,
	isPrecision ...option,
) ([]string, error) {

	// 拡張空間IDを取得
	ids, err := GetExtendedSpatialIdsOnCylinders(center, radius, zoom, zoom, isCapsule, isPrecision...)

	if err != nil {
		// エラーが発生した場合エラーインスタンスを返却
		return ids, err
	}

	// 拡張空間IDを空間IDのフォーマットに変換
	ids, err = shape.ConvertExtendedSpatialIdsToSpatialIds(ids)

	return ids, err
}

// GetExtendedSpatialIdsOnCylinders 拡張空間ID(円柱)取得
//
// 円柱を複数つなげた経路が通る拡張空間IDを取得する。
// 円柱間の接続面は球状とする。
// ドローンの経路や地中埋設配管が通る経路を拡張空間IDで表現する際に使用する。
//
// 引数：
//
//	center     : 円柱の中心の接続点。Pointを複数指定するリスト。
//	             2つ目の接続点は1つ目の円柱の終点となるが、2つ目の円柱の始点にもなる。
//	radius     : 円柱の半径(単位:m)
//	hZoom      : 水平方向の精度レベル
//	vZoom      : 垂直方向の精度レベル
//	isCapsule  : 始点、終点が球状であるかを示す。True: カプセル / False: 円柱
//	isPrecision: 衝突判定を実施するかのフラグ。True: 実施 / False: 未実施(デフォルトはTrue)
//
// 戻り値：
//
//	円柱を複数つなげた経路が通る拡張空間IDのリスト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力数値不正： 座標の値が不正の場合、もしくは円柱の半径が0以下の場合、エラー
//	 精度閾値超過： 水平方向精度、または垂直方向精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 注意: 経度180度をまたがる点をまたがる円柱は拡張空間ID化できない。
func GetExtendedSpatialIdsOnCylinders(
	center []*object.Point,
	radius float64,
	hZoom int64,
	vZoom int64,
	isCapsule bool,
	isPrecision ...option,
) ([]string, error) {

	// デフォルトパラメータを定義
	p := &IsPrecisionOpts{
		IsPrecision: true,
	}

	// ユーザーから渡された値だけ上書き
	for _, opt := range isPrecision {
		opt(p)
	}

	// 空間IDを格納するスライス
	spatialIDs := []string{}

	// 入力値チェック
	// 引数のポインタにnilがある場合
	if common.Include(center, nil) {
		return spatialIDs, errors.NewSpatialIdError(
			errors.InputValueErrorCode, "",
		)

		// 水平、垂直方向精度のどちらかが範囲外の場合、空配列とエラーインスタンスを返却
	} else if !shape.CheckZoom(hZoom) || !shape.CheckZoom(vZoom) {
		return spatialIDs, errors.NewSpatialIdError(
			errors.InputValueErrorCode, "",
		)

		// 半径が0以下の場合は例外を投げる
	} else if radius <= consts.Minima {
		logger.Debug("半径が0以下")
		return spatialIDs, errors.NewSpatialIdError(
			errors.InputValueErrorCode, "",
		)

		// 接続点数が0の場合は空配列を返却
	} else if len(center) == 0 {
		logger.Debug("接続点数が0個")
		return spatialIDs, nil
	}

	// メルカトル距離補正
	radian := common.DegreeToRadian(center[0].Lat())
	factor := 1 / math.Cos(radian)

	logger.Debug("メルカトル係数: %v", factor)

	// 接続点の球
	sphere := new(Capsule)
	// 接続点数
	connectPointNum := 1
	// 入力の接続点数-1の数だけループ
	for i := range center[:len(center)-1] {
		start := center[i]
		end := center[i+1]
		logger.Debug("接続点: %d～%d", connectPointNum, connectPointNum+1)
		logger.Debug("始点(地理座標系): %f, %f, %f", start.Lon(), start.Lat(), start.Alt())
		logger.Debug("終点(地理座標系): %f, %f, %f", end.Lon(), end.Lat(), end.Alt())

		// 同じ座標が連続する場合は次の座標へスキップ
		if start == end {
			logger.Debug("同一座標の接続点が連続しているためスキップ")
			continue
		}

		// 接続点数をインクリメント
		connectPointNum++

		// 【直交座標空間】始点終点の座標
		crsPoints, _ := shape.ConvertPointListToProjectedPointList(
			[]*object.Point{start, end},
			consts.OrthCrs,
		)

		startOrth := spatial.Point3{
			X: crsPoints[0].X,
			Y: crsPoints[0].Y,
			Z: crsPoints[0].Alt * factor,
		}
		endOrth := spatial.Point3{
			X: crsPoints[1].X,
			Y: crsPoints[1].Y,
			Z: crsPoints[1].Alt * factor,
		}
		logger.Debug("始点(投影座標系): %f, %f, %f", startOrth.X, startOrth.Y, startOrth.Z)
		logger.Debug("終点(投影座標系): %f, %f, %f", endOrth.X, endOrth.Y, endOrth.Z)

		if !sphere.IsEmpty() {

			shaveSphereSpatialIDs, _ := sphere.CalcValidSpatialIDs()
			logger.Debug(
				"接続点の空間ID: %s",
				reflect.ValueOf(shaveSphereSpatialIDs),
			)
			spatialIDs = common.Union(shaveSphereSpatialIDs, spatialIDs)
		}

		// 【直交座標空間】始点終点からカプセルの空間ID取得
		capsule := NewCapsule(
			startOrth,
			endOrth,
			radius,
			hZoom,
			vZoom,
			isCapsule,
			p.IsPrecision,
			factor,
		)
		capsuleSpatialIDs, _ := capsule.CalcValidSpatialIDs()
		logger.Debug(
			"接続点間の空間ID: %s",
			reflect.ValueOf(capsuleSpatialIDs),
		)
		// マージ処理
		spatialIDs = common.Union(capsuleSpatialIDs, spatialIDs)

		// 終点もしくはカプセルの場合は接続点の空間IDは取得しない
		if i == len(center[:len(center)-1])-1 || isCapsule {
			continue
		}

		// 【直交座標空間】接続点の球の空間ID取得
		sphere = NewCapsule(
			endOrth,
			endOrth,
			radius,
			hZoom,
			vZoom,
			isCapsule,
			p.IsPrecision,
			factor,
		)
	}

	// 接続点数が1の場合
	if connectPointNum == 1 {
		logger.Debug("接続点数が1個")

		// 【直交座標空間】中心の座標
		orthCenters, _ := shape.ConvertPointListToProjectedPointList(
			[]*object.Point{center[0]}, consts.OrthCrs,
		)
		orthCenter := spatial.Point3{
			X: orthCenters[0].X,
			Y: orthCenters[0].Y,
			Z: orthCenters[0].Alt * factor,
		}
		logger.Debug(
			"中心座標(地理座標系): %f, %f, %f",
			center[0].Lon(), center[0].Lat(), center[0].Alt(),
		)
		logger.Debug(
			"中心座標(投影座標系): %f, %f, %f",
			orthCenter.X, orthCenter.Y, orthCenter.Z,
		)

		// 【直交座標空間】接続点の球の空間ID取得
		sphere = NewCapsule(
			orthCenter,
			orthCenter,
			radius,
			hZoom,
			vZoom,
			isCapsule,
			p.IsPrecision,
			factor,
		)
		spatialIDs, _ = sphere.CalcValidSpatialIDs()
		logger.Debug(
			"球の空間ID: %s",
			reflect.ValueOf(spatialIDs),
		)
	}

	return common.Unique(spatialIDs), nil

}
