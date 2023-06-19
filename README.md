# 空間IDライブラリ

## 概要
- 任意の座標を空間IDに変換するライブラリです。
- spatial-idとspatial-id-plusの2つのパッケージがあります。spatial-idは外部ライブラリの事前インストールなしで利用が可能です。spatial-id-plusを利用するためには別途、外部ライブラリのインストールが必要です(後述)。それぞれのパッケージで提供する機能は下記です。
  - spatial-idの提供機能は以下です。
    - 任意の座標から空間IDを取得する機能
    - 空間IDを任意の精度に拡大・縮小する機能
    - 任意の空間IDの周辺の空間IDを取得する機能
    - 任意の形状から空間IDを取得する機能
    - 空間IDをquadkeyと2分木におけるbit形式のIDに変換する機能
  - spatial-id-plusの提供機能は以下です。
    - 任意の座標と座標を結ぶ線を中心軸とした円柱状の空間IDを取得する機能
- 各機能の詳細についてはdocフォルダ配下にAPI仕様書があります。
- 空間ID仕様については以下のリンクを参照して下さい。
[Digital Architecture Design Center 3次元空間情報基盤アーキテクチャ検討会 会議資料](https://www.ipa.go.jp/dadc/architecture/pdf/pj_report_3dspatialinfo_doc-appendix_202212_1.pdf)

# ライブラリの構成
ライブラリは下記のパッケージ構成となっています。
- src/spatial-id
  - common
    - const
    - enum
    - errors
    - logger
    - object
    - spatial
  - integrate
  - operated
  - shape
  - transform
- src/spatial-id-plus
  - shape


# ライブラリのimport方法
importは下記のように記載します。
下記は、2つのパッケージを同時にimportをする例になります。

```
import (
	"fmt"
	plus "<リポジトリURL>/src/spatial-id-plus/shape"
	"<リポジトリURL>/src/spatial-id/shape"
)
```


# 事前インストールが必要な外部ライブラリ
spatial-idのみを使用する場合は、不要です。
spatial-id-plusで使用している外部ライブラリにAzul3Dがあります。
Azul3Dの動作の前提としてODEライブラリが必要になるため、事前にインストールが必要です。

インストール手順は下記です。

## ODEのインストール手順
ODEはC++の物理エンジンです。

[公式サイト](http://www.ode.org/)

Azul3DではODEをWrapして衝突判定に用いています。そのため、Azul3Dの前提ライブラリとしてインストールします。

1. ODEのソースを取得します。
[最新版のソース](https://bitbucket.org/odedevs/ode/downloads/ode-0.16.2.tar.gz)
1. ファイルを解凍して配置します。
1. 配置先をカレントにして下記コマンドでインストールします。
```
$ cd ode-0.16.2
$ ./configure --enable-double-precision --enable-shared
$ make
$ sudo make install
```
 - トラブルシューティング
Azul3Dのパッケージをimportしたプログラムの実行時に下記のメッセージが出た場合
```
error while loading shared libraries: libode.so.8: cannot open shared object file: No such file or directory
```
1. 「/etc/ld.so.conf」を編集し、「/usr/local/lib」をファイル末尾に追加します。
2. 下記、コマンドを実行します。
```
$ sudo /sbin/ldconfig
```

## 注意事項
* ライブラリの入力可能な緯度の最大、最小値は「±85.0511287798」とします。
* 精度レベルの指定範囲は、0から35とします。
* 経度の限界値は±180ですが、180と-180は同じ個所を指すこととZFXY形式のインデックスの考え方により、180はライブラリ内部では-180として扱われます。(180の入力は可能とします。)

## 外部ライブラリ
- 外部ライブラリ
    - WGS84
        - バージョン:1.1.6
        - 確認日:2023/3/8
        - 用途:座標変換に使用します。
    - azul3D  
        - バージョン:バージョンなし
        - 確認日:2023/3/8
        - 用途:円柱と空間ボクセルの衝突確認に使用します。spatial-id-plusを呼び出す場合にのみ利用されます。
    - ODE
        - バージョン:0.16.2
        - 確認日:2023/3/8
        - 用途:円柱と空間ボクセルの衝突確認に使用します。spatial-id-plusを呼び出す場合にのみ利用されます。
