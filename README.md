# 空間IDモジュール
[![Go Doc](https://pkg.go.dev/badge/github.com/trajectoryjp/spatial_id_go)](https://pkg.go.dev/github.com/trajectoryjp/spatial_id_go)
[![Go Report Card](https://goreportcard.com/badge/github.com/trajectoryjp/spatial_id_go)](https://goreportcard.com/report/github.com/trajectoryjp/spatial_id_go)


## 概要
任意の座標を空間IDに変換するモジュールです。
* 外部ライブラリの事前インストールなしで利用が可能です。
* 提供機能は以下の通りです。
  * 任意の座標から空間IDを取得する機能
  * 空間IDを任意の精度に拡大・縮小する機能
  * 任意の空間IDの周辺の空間IDを取得する機能
  * 任意の形状から空間IDを取得する機能
  * 空間IDをquadkeyと2分木におけるbit形式のIDに変換する機能
* 空間ID仕様については[Digital Architecture Design Center 3次元空間情報基盤アーキテクチャ検討会 会議資料](https://www.ipa.go.jp/digital/architecture/Individual-link/ps6vr7000000qmcv-att/pj_report_3dspatialinfo_doc-appendix_202212_1.pdf)を参照して下さい。


## 注意事項
* 入力可能な緯度の最大、最小値は「±85.0511287798」とします。
* 精度レベルの指定範囲は、0から35とします。
* 経度の限界値は±180ですが、180と-180は同じ個所を指すこととZFXY形式のインデックスの考え方により、180はモジュール内部では-180として扱われます。(180の入力は可能とします。)
