# -*- coding: utf-8 -*-

import math
import numpy as np
from skspatial.objects import Point, Points, Line, Vector, Triangle

from SpatialId.common.object.point import Point as SpatialPoint
from SpatialId.common.object.point import Projected_Point
from SpatialId.shape.point import (
    convert_point_list_to_projected_point_list,
    get_point_on_spatial_id
    )

# CRSのデフォルト値のEPSGコード
SPATIAL_ID_CRS = 3857
WGS84_ID_CRS = 4326

# 地理座標のままskspatialのPointsに変換
def spatial_point_to_skspatial_point(points):
    return [Point([point.lon, point.lat, point.alt]) for point in points]

# 地理座標から投影座標へ変換(skspatialのPointsで返却)
def spatial_point_to_proj_skspatial_point(points, crs):
    proj_points = convert_point_list_to_projected_point_list(points, SPATIAL_ID_CRS, crs)
    return [Point([proj_point.x, proj_point.y, proj_point.alt]) for proj_point in proj_points]

# 空間IDから地理座標、投影座標を取得(skspatialのPointsで返却)
def spatial_id_to_skspatial_points(spatial_id, option):
    points = get_point_on_spatial_id(spatial_id, option, WGS84_ID_CRS)
    proj_points = get_point_on_spatial_id(spatial_id, option, SPATIAL_ID_CRS)
    return tuple((
            Points([Point([proj_point.x, proj_point.y, proj_point.alt]) for proj_point in proj_points]),
            Points([Point([point.lon, point.lat, point.alt]) for point in points])
        ))

# 空間IDの三角オブジェクトから投影座標に変換(skspatialのTriangleで返却)
def spatial_triangle_to_proj_skspatial_triangle(triangles, crs):
    results = list()
    for triangle in triangles:
        points = spatial_point_to_proj_skspatial_point([triangle.p1, triangle.p2, triangle.p3], crs)
        results.append(Triangle(points[0], points[1], points[2]))
    return results
