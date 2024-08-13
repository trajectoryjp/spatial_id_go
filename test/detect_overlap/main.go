package main

import (
	"fmt"
	"github.com/trajectoryjp/spatial_id_go/v4/detector"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

const maxCases = 10

func main() {
	// profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	// 入力値
	spatialIdCount := maxCases
	fmt.Println("inputCount,id1Len,id2Len,truePoint,duration(ns),detection result")
	_ = pprof.StartCPUProfile(f)
	// caseCount < spatialIdCountまで: caseCountで重複する配列データ生成
	// caseCount => spatialIdCount: 重複なしデータ生成
	for caseCount := 0; caseCount < spatialIdCount+1; caseCount++ {
		spatialIds1, spatialIds2 := makeTargets(caseCount, spatialIdCount, []string{}, []string{})
		start := time.Now()
		result, _ := detector.CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)
		end := time.Now()
		duration := end.Sub(start)
		fmt.Printf("%v,%v,%v,%v,%v,%v\n", spatialIdCount, len(spatialIds1), len(spatialIds2), caseCount, duration.Nanoseconds(), result)
	}
	pprof.StopCPUProfile()
}

func makeTargets(truePoint int, spatialIdCount int, spatialIds1 []string, spatialIds2 []string) ([]string, []string) {
	for i := 0; i < spatialIdCount; i++ {
		if i == truePoint {
			spatialIds1 = append(spatialIds1, "13/0/7274/3225")
			spatialIds2 = append(spatialIds2, "16/0/58198/25804")
		} else {
			spatialIds1 = append(spatialIds1, "13/0/7275/3226")
			spatialIds2 = append(spatialIds2, "16/0/58198/25804")
		}
	}
	return spatialIds1, spatialIds2
}
