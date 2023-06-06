import argparse
from skspatial.objects import Point, Vector, Line
from spatialid.plot_utils import Plot3D
from spatialid import utils
from SpatialId.common.object.point import Point as SpatialPoint

from SpatialId.common.exception import SpatialIdError

def main():

    parser = argparse.ArgumentParser()
    parser.add_argument('-i', required=True, dest="input", type=open)
    parser.add_argument('-p', required=True, dest="input_point", type=open)

    args = parser.parse_args()

    # デフォルト角度
    plt = Plot3D()
    ax = plt.ax

    points = [float(point.strip()) for point in args.input_point.readlines()]
    start_point = SpatialPoint(points[0], points[1], points[2])
    end_point = SpatialPoint(points[3], points[4], points[5])

    proj_points = utils.spatial_point_to_proj_skspatial_point([start_point, end_point], 4326)

    s_p = proj_points[0]
    e_p = proj_points[1]
    vec = Vector.from_points(s_p, e_p)

    vec.plot_3d(ax, point=s_p, c='r', alpha=1.0)
    s_p.plot_3d(ax, c='r', alpha=1.0)
    e_p.plot_3d(ax, c='g', alpha=1.0)

    spatial_ids = [spatial_id.strip() for spatial_id in args.input.readlines()]

    # 空間IDをプロット
    plt.plot_voxel_on_spatial_ids(spatial_ids)
    # ファイル保存
    plt.savefig('output/plot_out.pdf')

if __name__ == "__main__":
    try:
        main()
    except SpatialIdError as e:
        print(e)
        raise e
