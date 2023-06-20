package spatial

import (
	"reflect"
	"testing"
)

// TestNewLineFromPoints01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点(1,2,3), 終点：(5,6,7))
//
// + 確認内容
//   - 入力に対応した線オブジェクトが取得できること
func TestNewLineFromPoints01(t *testing.T) {
	start := Point3{1, 2, 3}
	end := Point3{5, 6, 7}
	resultVal := NewLineFromPoints(start, end)

	expectVal := Line3{start, Vector3{end.X - start.X, end.Y - start.Y, end.Z - start.Z}}

	// 戻り値の線オブジェクトと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("線オブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestToPoint01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点(1,2,3), 終点：(5,6,7), パラメータ：10)
//
// + 確認内容
//   - 入力に対応した線に沿った点が取得できること
func TestToPoint01(t *testing.T) {
	var f float64 = 10
	start := Point3{1, 2, 3}
	end := Point3{5, 6, 7}
	l := Line3{start, Vector3{end.X - start.X, end.Y - start.Y, end.Z - start.Z}}
	resultVal := l.ToPoint(f)

	expectVal := Point3{((end.X - start.X) * f) + start.X, ((end.Y - start.Y) * f) + start.Y, ((end.Z - start.Z) * f) + start.Z}

	// 戻り値の点と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestEnd01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点(1,2,3), 終点：(5,6,7))
//
// + 確認内容
//   - 入力に対応した線の終点が取得できること
func TestEnd01(t *testing.T) {
	start := Point3{1, 2, 3}
	end := Point3{5, 6, 7}
	l := Line3{start, Vector3{end.X - start.X, end.Y - start.Y, end.Z - start.Z}}
	resultVal := l.End()

	// 戻り値の終点と期待値の比較
	if !reflect.DeepEqual(resultVal, end) {
		t.Errorf("終点 - 期待値：%v, 取得値：%v", end, resultVal)
	}
	t.Log("テスト終了")
}

// TestStart01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点(1,2,3), 終点：(5,6,7))
//
// + 確認内容
//   - 入力に対応した線の始点が取得できること
func TestStart01(t *testing.T) {
	start := Point3{1, 2, 3}
	end := Point3{5, 6, 7}
	l := Line3{start, Vector3{end.X - start.X, end.Y - start.Y, end.Z - start.Z}}
	resultVal := l.Start()

	// 戻り値の始点と期待値の比較
	if !reflect.DeepEqual(resultVal, start) {
		t.Errorf("始点 - 期待値：%v, 取得値：%v", start, resultVal)
	}

	t.Log("テスト終了")
}
