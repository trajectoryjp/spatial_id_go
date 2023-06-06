# -*- coding: utf-8 -*-
"""プロットユーティリティモジュール
"""

import numpy as np
import matplotlib.pyplot as plt
from skspatial.objects import Points, Point, Vector, Line
from spatialid import utils
from mpl_toolkits.mplot3d.art3d import Poly3DCollection, Line3DCollection
from SpatialId.common.object.point import Projected_Point
from SpatialId.common.object.enum import Point_Option

# 浮動小数点誤差
__MINIMA__ = 1e-10

class Plot3D():
    """3Dプロットクラス

    3Dプロットを行う汎用処理クラス

    Attributes:
        fig(matplotlib.figure.Figure): プロットコンテナ
        ax(matplotlib.axes._subplots.Axes3DSubplot): 3D軸オブジェクト
    """

    def __init__(self, elev=None, azim=None):
        """コンストラクタ

        Args:
            elev(int): z平面の仰角(Noneの場合、デフォルトの仰角となる)
            azim(int): x, y平面の仰角(Noneの場合、デフォルトの仰角となる)
        """

#        plt.figaspect(1)
        self.fig = plt.figure(figsize=plt.figaspect(1))
        self.ax = self.fig.add_subplot(111, projection='3d')
        self.ax.view_init(elev=elev, azim=azim)

    def plot_voxel_on_spatial_id(self, spatial_id, facecolors='red', linewidths=0.3, edgecolors='black', alpha=0.3):
        points, _ = utils.spatial_id_to_skspatial_points(spatial_id, Point_Option.VERTEX)
        min_x = min(points, key=lambda x: x[0])[0]
        max_x = max(points, key=lambda x: x[0])[0]
        min_y = min(points, key=lambda x: x[1])[1]
        max_y = max(points, key=lambda x: x[1])[1]
        min_z = min(points, key=lambda x: x[2])[2]
        max_z = max(points, key=lambda x: x[2])[2]
        Z = Points([
            [min_x, min_y, min_z],
            [min_x, max_y, min_z],
            [min_x, min_y, max_z],
            [min_x, max_y, max_z],
            [max_x, min_y, min_z],
            [max_x, max_y, min_z],
            [max_x, min_y, max_z],
            [max_x, max_y, max_z],
        ])
        verts = [[Z[0],Z[1],Z[3],Z[2]],
                 [Z[0],Z[1],Z[5],Z[4]],
                 [Z[0],Z[2],Z[6],Z[4]],
                 [Z[4],Z[5],Z[7],Z[6]],
                 [Z[1],Z[3],Z[7],Z[5]],
                 [Z[2],Z[3],Z[7],Z[6]]]
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors='cyan', linewidths=1, edgecolors='black', alpha=.1))
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors='red', linewidths=0.1, edgecolors='black', alpha=.1))
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors='red', linewidths=0.0, edgecolors='black', alpha=.1))
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors='red', linewidths=0.1, edgecolors='black', alpha=0.05))
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors='cyan', linewidths=0.1, edgecolors='black', alpha=0.05))
        self.ax.add_collection3d(Poly3DCollection(verts, facecolors=facecolors, linewidths=linewidths, edgecolors=edgecolors, alpha=alpha))


    def plot_voxel_on_spatial_ids(self, spatial_ids, facecolors='red', linewidths=0.3, edgecolors='black', alpha=0.3):
        h_x_min = None
        h_x_max = None
        h_y_min = None
        h_y_max = None
        v_z_min = None
        v_z_max = None
        for spatial_id in spatial_ids:
#            print(spatial_id)
            proj_points, gis_points = utils.spatial_id_to_skspatial_points(spatial_id, Point_Option.VERTEX)
#            print(gis_points)
#            print(proj_points)
            x_max = max(proj_points, key=lambda x: x[0])
            x_min = min(proj_points, key=lambda x: x[0])
            y_max = max(proj_points, key=lambda x: x[1])
            y_min = min(proj_points, key=lambda x: x[1])
            z_max = max(proj_points, key=lambda x: x[1])
            z_min = min(proj_points, key=lambda x: x[1])

            if h_x_min is None:
                h_x_max = x_max[0]
                h_x_min = x_min[0]
                h_y_max = y_max[1]
                h_y_min = y_min[1]
                v_z_max = z_max[2]
                v_z_min = z_min[2]
            else:
                h_x_max = max([x_max[0], h_x_max])
                h_x_min = min([x_min[0], h_x_min])
                h_y_max = max([y_max[1], h_y_max])
                h_y_min = min([y_min[1], h_y_min])
                v_z_max = max([z_max[2], v_z_max])
                v_z_min = min([z_min[2], v_z_min])

            proj_points.plot_3d(self.ax, c='b', alpha=0.0)
            self.plot_voxel_on_spatial_id(spatial_id, facecolors, linewidths, edgecolors, alpha)

        h_size = max([h_x_max - h_x_min, h_y_max - h_y_min])
        v_size = v_z_max - v_z_min
        max_size = max([h_size, v_size])
        Point([h_x_min + max_size, h_y_min + max_size, v_z_min + max_size]).plot_3d(self.ax, c='b', alpha=0.0)

    def plot_triangle(self, triangle, facecolors='red', linewidths=0.3, edgecolors='black', alpha=0.3):
        verts = [[triangle.point_a, triangle.point_b, triangle.point_c]]
        #self.ax.add_collection3d(Poly3DCollection(verts, facecolors=facecolors, linewidths=linewidths, edgecolors=edgecolors, alpha=alpha))

        # 始点終点プロット
        triangle.point_a.plot_3d(self.ax, c=facecolors, alpha=alpha)
        triangle.point_b.plot_3d(self.ax, c=facecolors, alpha=alpha)
        triangle.point_c.plot_3d(self.ax, c=facecolors, alpha=alpha)

        # 線プロット
        vec_a = Vector.from_points(triangle.point_a, triangle.point_b)
        vec_a.plot_3d(self.ax, point=triangle.point_a, c=facecolors, alpha=alpha)
        vec_b = Vector.from_points(triangle.point_b, triangle.point_c)
        vec_b.plot_3d(self.ax, point=triangle.point_b, c=facecolors, alpha=alpha)
        vec_c = Vector.from_points(triangle.point_c, triangle.point_a)
        vec_c.plot_3d(self.ax, point=triangle.point_c, c=facecolors, alpha=alpha)


    def plot_cylinder(self, cylinder, facecolors='blue', linewidths=0.3, edgecolors='black', alpha=0.3):
        cylinder.plot_3d(self.ax, c=facecolors, alpha=alpha)


    def savefig(self, path):
        """プロット出力
        Args:
            path(str): プロット出力先ファイルパス
        """

        plt.savefig(path)


class Rectangular3DPlot:
    """ 始点・終点の２面が正方形の直方体の空間ID

    :var _start_point: 直方体の始点。
    :var _end_point: 直方体の終点。
    :var _radius: 正方形の半径(単位:m)
    :var _all_cross_points: 全交点
    :var _all_spatial_ids: 全空間ID
    :var _height: 図形の長さ
    :var _unit_axis_vector: 軸方向の単位ベクトル
    :var _radius_orth_vector: 軸に対する半径長さの直交ベクトル
    :var _contact_point: 直交ベクトルと図形の交点
    :var _radius_normal_vector: 軸方向のベクトルと直交ベクトルの外積
    """
    def __init__(
        self, start_point: Point, end_point: Point, radius: float,
    ) -> None:
        """ コンストラクタ

        :param start_point: 直方体の始点。
        :type  start_point: skspatial.objects.Point
        :param end_point: 直方体の終点。
        :type  end_point: skspatial.objects.Point
        :param radius: 正方形の半径(単位:m)
        :type  radius: float
        """
        # 【直交座標空間】
        self._start_point = start_point
        self._end_point = end_point
        self._radius = radius

        # 全交点リスト
        self._all_cross_points = list()

        # 空間ID
        self._all_spatial_ids = set()

    def _init_axis(self) -> None:
        # 【直交座標空間】軸ベクトル
        axis_vector = Vector.from_points(self._start_point, self._end_point)
        # 縦の長さ
        self._height = axis_vector.norm()
        x_vec, y_vec, z_vec = axis_vector

        # 開始点の直行ベクトルを一つ求める
        # 直行ベクトル
        start_orth_vector = None
        if abs(y_vec) < __MINIMA__:
            start_orth_vector = Vector([0, 1, 0])
        elif (abs(z_vec) < __MINIMA__):
            start_orth_vector = Vector([0, 0, 1])
        else:
            start_orth_vector = Vector([0, -z_vec, y_vec])

        # 直交ベクトルを機体半径の大きさにする。
        self._radius_orth_vector = self._radius * start_orth_vector.unit()

        # 単位軸ベクトル
        self._unit_axis_vector = axis_vector.unit()

        # 円との接点
        self._contact_point = self._radius_orth_vector + self._start_point

        # 軸ベクトルと直交ベクトルの法線ベクトル(大きさは機体半径)
        self._radius_normal_vector = \
            self._radius * self._radius_orth_vector.cross(
                self._unit_axis_vector).unit()

    def _calc_rectangular_apex(self, ax) -> None:
        # 【直交座標空間】直方体の頂点
        apex_points = [0] * 8
        apex_points[0] = self._contact_point + self._radius_normal_vector
        apex_points[1] = self._contact_point - self._radius_normal_vector
        apex_points[2] = apex_points[0] - 2 * self._radius_orth_vector
        apex_points[3] = apex_points[1] - 2 * self._radius_orth_vector
        apex_points[4] = apex_points[0] + self._height * self._unit_axis_vector
        apex_points[5] = apex_points[1] + self._height * self._unit_axis_vector
        apex_points[6] = apex_points[2] + self._height * self._unit_axis_vector
        apex_points[7] = apex_points[3] + self._height * self._unit_axis_vector

        # 【直交座標空間】直方体の辺
        rect_lines = [
            Line.from_points(apex_points[0], apex_points[1]),
            Line.from_points(apex_points[0], apex_points[4]),
            Line.from_points(apex_points[1], apex_points[5]),
            Line.from_points(apex_points[4], apex_points[5]),
            Line.from_points(apex_points[2], apex_points[3]),
            Line.from_points(apex_points[2], apex_points[6]),
            Line.from_points(apex_points[6], apex_points[7]),
            Line.from_points(apex_points[3], apex_points[7]),
            Line.from_points(apex_points[0], apex_points[2]),
            Line.from_points(apex_points[1], apex_points[3]),
            Line.from_points(apex_points[4], apex_points[6]),
            Line.from_points(apex_points[5], apex_points[7])
        ]
        for line in rect_lines:
            line.plot_3d(ax, c='b', alpha=1.0)

    def plot_3d(self, ax):
        # 軸の初期化
        self._init_axis()
        # 直方体の境界面に対する交点取得
        self._calc_rectangular_apex(ax)

