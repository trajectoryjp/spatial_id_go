package main

import (
	"flag"
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
	"strconv"
)

func main() {
	var (
		output      = flag.String("o", "/dev/null", "output path")
		pointOutput = flag.String("p", "/dev/null", "point output path")
	)

	// コマンドライン引数解析
	flag.Parse()

	// 始点
	start0, _ := object.NewPoint(0.0, 0.0, 0.0)
	start1, _ := object.NewPoint(0.0, 0.0, 0.0)
	start2, _ := object.NewPoint(-135.0, 75.7821946114, 4194304.0)
	start3, _ := object.NewPoint(-157.5, 82.1112317102, 2097152.0)
	start4, _ := object.NewPoint(-168.75, 83.863706879, 1048576.0)
	start5, _ := object.NewPoint(-174.375, 84.5151941393, 524288.0)
	start6, _ := object.NewPoint(-177.1875, 84.7962449264, 262144.0)
	start7, _ := object.NewPoint(-178.59375, 84.926801252, 131072.0)
	start8, _ := object.NewPoint(-179.296875, 84.9897248546, 65536.0)
	start9, _ := object.NewPoint(-179.6484375, 85.02061448, 32768.0)
	start10, _ := object.NewPoint(-179.82421875, 85.0359182614, 16384.0)
	start11, _ := object.NewPoint(-179.912109375, 85.0435351431, 8192.0)
	start12, _ := object.NewPoint(-179.9560546875, 85.0473348627, 4096.0)
	start13, _ := object.NewPoint(-179.97802734375, 85.049232546, 2048.0)
	start14, _ := object.NewPoint(-179.989013671875, 85.050180844, 1024.0)
	start15, _ := object.NewPoint(-179.9945068359375, 85.0506548571, 512.0)
	start16, _ := object.NewPoint(-179.99725341796875, 85.0508918298, 256.0)
	start17, _ := object.NewPoint(-179.99862670898438, 85.0510103076, 128.0)
	start18, _ := object.NewPoint(-179.9993133544922, 85.0510695444, 64.0)
	start19, _ := object.NewPoint(-179.9996566772461, 85.0510991622, 32.0)
	start20, _ := object.NewPoint(-179.99982833862305, 85.051113971, 16.0)
	start21, _ := object.NewPoint(-179.99991416931152, 85.0511213754, 8.0)
	start22, _ := object.NewPoint(-179.99995708465576, 85.0511250776, 4.0)
	start23, _ := object.NewPoint(-179.99997854232788, 85.0511269287, 2.0)
	start24, _ := object.NewPoint(-179.99998927116394, 85.0511278542, 1.0)
	start25, _ := object.NewPoint(-179.99999463558197, 85.051128317, 0.5)
	start26, _ := object.NewPoint(-179.99999731779099, 85.0511285484, 0.25)
	start27, _ := object.NewPoint(-179.9999986588955, 85.051128664, 0.125)
	start28, _ := object.NewPoint(-179.99999932944775, 85.0511287219, 0.0625)
	start29, _ := object.NewPoint(-179.99999966472387, 85.0511287508, 0.03125)
	start30, _ := object.NewPoint(-179.99999983236194, 85.0511287653, 0.015625)
	start31, _ := object.NewPoint(-179.99999991618097, 85.0511287725, 0.0078125)
	start32, _ := object.NewPoint(-179.99999995809048, 85.0511287761, 0.00390625)
	start33, _ := object.NewPoint(-179.99999997904524, 85.0511287779, 0.001953125)
	start34, _ := object.NewPoint(-179.99999998952262, 85.0511287788, 0.0009765625)
	start35, _ := object.NewPoint(-179.9999999947613, 85.0511287793, 0.00048828125)

	// 始点のリスト
	startpList := []*object.Point{start0, start1, start2, start3, start4, start5, start6, start7, start8, start9,
		start10, start11, start12, start13, start14, start15, start16, start17, start18, start19,
		start20, start21, start22, start23, start24, start25, start26, start27, start28, start29,
		start30, start31, start32, start33, start34, start35}

	// 終点
	end0, _ := object.NewPoint(179.0, 80.0, 100.0)
	end1, _ := object.NewPoint(179.0, 80.0, 100.0)
	end2, _ := object.NewPoint(45.0, 75.7821946114, 4194304.0)
	end3, _ := object.NewPoint(-67.5, 82.1112317102, 2097152.0)
	end4, _ := object.NewPoint(-123.75, 83.863706879, 1048576.0)
	end5, _ := object.NewPoint(-151.875, 84.5151941393, 524288.0)
	end6, _ := object.NewPoint(-165.9375, 84.7962449264, 262144.0)
	end7, _ := object.NewPoint(-172.96875, 84.926801252, 131072.0)
	end8, _ := object.NewPoint(-176.484375, 84.9897248546, 65536.0)
	end9, _ := object.NewPoint(-178.2421875, 85.02061448, 32768.0)
	end10, _ := object.NewPoint(-179.12109375, 85.0359182614, 16384.0)
	end11, _ := object.NewPoint(-179.560546875, 85.0435351431, 8192.0)
	end12, _ := object.NewPoint(-179.7802734375, 85.0473348627, 4096.0)
	end13, _ := object.NewPoint(-179.89013671875, 85.049232546, 2048.0)
	end14, _ := object.NewPoint(-179.945068359375, 85.050180844, 1024.0)
	end15, _ := object.NewPoint(-179.9725341796875, 85.0506548571, 512.0)
	end16, _ := object.NewPoint(-179.98626708984375, 85.0508918298, 256.0)
	end17, _ := object.NewPoint(-179.99313354492188, 85.0510103076, 128.0)
	end18, _ := object.NewPoint(-179.99656677246094, 85.0510695444, 64.0)
	end19, _ := object.NewPoint(-179.99828338623047, 85.0510991622, 32.0)
	end20, _ := object.NewPoint(-179.99914169311523, 85.051113971, 16.0)
	end21, _ := object.NewPoint(-179.99957084655762, 85.0511213754, 8.0)
	end22, _ := object.NewPoint(-179.9997854232788, 85.0511250776, 4.0)
	end23, _ := object.NewPoint(-179.9998927116394, 85.0511269287, 2.0)
	end24, _ := object.NewPoint(-179.9999463558197, 85.0511278542, 1.0)
	end25, _ := object.NewPoint(-179.99997317790985, 85.051128317, 0.5)
	end26, _ := object.NewPoint(-179.99998658895493, 85.0511285484, 0.25)
	end27, _ := object.NewPoint(-179.99999329447746, 85.051128664, 0.125)
	end28, _ := object.NewPoint(-179.99999664723873, 85.0511287219, 0.0625)
	end29, _ := object.NewPoint(-179.99999832361937, 85.0511287508, 0.03125)
	end30, _ := object.NewPoint(-179.99999916180968, 85.0511287653, 0.015625)
	end31, _ := object.NewPoint(-179.99999958090484, 85.0511287725, 0.0078125)
	end32, _ := object.NewPoint(-179.99999979045242, 85.0511287761, 0.00390625)
	end33, _ := object.NewPoint(-179.9999998952262, 85.0511287779, 0.001953125)
	end34, _ := object.NewPoint(-179.9999999476131, 85.0511287788, 0.0009765625)
	end35, _ := object.NewPoint(-179.99999997380655, 85.0511287793, 0.00048828125)

	// 終点のリスト
	endpList := []*object.Point{end0, end1, end2, end3, end4, end5, end6, end7, end8, end9,
		end10, end11, end12, end13, end14, end15, end16, end17, end18, end19,
		end20, end21, end22, end23, end24, end25, end26, end27, end28, end29,
		end30, end31, end32, end33, end34, end35}

	// 水平精度
	var hZoom int64 = 0
	// 垂直精度
	var vZoom int64 = 0

	for i := 0; i < len(startpList); i++ {

		// 入力の座標出力
		pp, _ := os.Create(*pointOutput)
		defer pp.Close()

		startLon := strconv.FormatFloat(startpList[i].Lon(), 'f', -1, 64)
		startLat := strconv.FormatFloat(startpList[i].Lat(), 'f', -1, 64)
		startAlt := strconv.FormatFloat(startpList[i].Alt(), 'f', -1, 64)
		endLon := strconv.FormatFloat(endpList[i].Lon(), 'f', -1, 64)
		endLat := strconv.FormatFloat(endpList[i].Lat(), 'f', -1, 64)
		endAlt := strconv.FormatFloat(endpList[i].Alt(), 'f', -1, 64)
		pp.WriteString(startLon + "\n")
		pp.WriteString(startLat + "\n")
		pp.WriteString(startAlt + "\n")
		pp.WriteString(endLon + "\n")
		pp.WriteString(endLat + "\n")
		pp.WriteString(endAlt + "\n")

		// 空間ID取得
		spatialIDs, err := shape.GetExtendedSpatialIdsOnLine(startpList[i], endpList[i], hZoom, vZoom)
		if err != nil {
			fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
			os.Exit(1)
		}

		// 結果を出力
		fmt.Printf("始点%d: %v\n", i, startpList[i])
		fmt.Printf("終点%d: %v\n", i, endpList[i])
		fmt.Printf("水平精度: %d 垂直精度: %d\n", hZoom, vZoom)
		fmt.Printf("変換後の空間ID[No%d]: %s\n", i+1, spatialIDs)
		fp, _ := os.Create(*output)
		defer fp.Close()

		for _, spatialID := range spatialIDs {
			fp.WriteString(spatialID + "\n")
		}

		hZoom++
		vZoom++
	}
}
