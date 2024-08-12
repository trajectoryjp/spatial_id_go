# 空間ID変換関数

`transform/convert_quadkey_and_Vertical_id.go`内にある各種変換関数の利用方法を記載

## 変換関数一覧

| 変換元         | 変換先         | 対応関数                                                   | 今後の変更予定など備考                                        |
|-------------|-------------|--------------------------------------------------------|----------------------------------------------------|
| 空間ID        | 拡張空間ID      | なし                                                     | `ConvertSpatialIDsToQuadkeysAndVerticalIDs`が部分的に行う |
| 拡張空間ID      | 拡張空間ID      | なし                                                     | `ConvertQuadkeysAndVerticalIDsToSpatialIDs`が部分的に行う |
| 拡張空間ID      | TileXYZ     | なし                                                     | `ConvertExtendedSpatialIDsToTileXYZ()`に実装予定        |
| TileXYZ     | 拡張空間ID      | なし                                                     | `ConvertTileXYZToExtendedSpatialIDs()`に実装予定        |
| 3Dtilekey   | 拡張空間ID      | `ConvertQuadkeysAndAltitudekeysToExtendedSpatialIDs()` | Pull Request#27のもの 廃止する                            |
| 拡張空間ID      | 3Dtilekey   | `ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys()` | 廃止予定                                               |
| 一次元インデックス   | 一次元変換インデックス | `ConvertZToAltitudekey()`                              | `TransformIndexCoordinate()`にリネーム予定                |
| 一次元変換インデックス | 一次元インデックス   | `ConvertAltitudeKeyToZ()`                              | `InverseTransformIndexCoordinate()`にリネーム予定         |
| Key         | 拡張空間ID      | `ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs()`  | 廃止予定                                               |
| Key         | 空間ID        | `ConvertQuadkeysAndVerticalIDsToSpatialIDs()`          | 廃止予定                                               |
| 空間ID        | Key         | `ConvertSpatialIDsToQuadkeysAndVerticalIDs()`          | 廃止予定                                               |
| 拡張空間ID      | Key         | `ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs()`  | 廃止予定                                               |

## ConvertQuadkeysAndAltitudekeysToExtendedSpatialIDs

Quadkey+AltitudeKey空間のボクセルを拡張空間IDに変換する

変換時、QuadKeyはx,yインデックスにデコードされるのみである  
しかしAltitudeKeyは変換によって元の値に近い垂直方向拡張空間インデックスに変換される

これはAltitudeKeyと拡張空間IDの間ではボクセルのインデックス付番や0番の位置の基準を変更されているためである

この基準とデータをまとめて示すために 構造体`FromExtendedSpatialIDToQuadkeyAndAltitudekey`を入力として扱う

```go
[]FromExtendedSpatialIDToQuadkeyAndAltitudekey{
    {
        quadkeyZoom : 20
        innerIDList : [[7432012031,0]] // [[quadkey, alatitudekey],...]
        altitudekeyZoom : 23
        zBaseExponent : 25
        zBaseOffset : -1
    }
}
```

ここで構造体フィールドには次のような値を利用する(fig.1)

- `quadkeyZoom` 水平方向ズームレベル
- `innerIDList` 実際の値: QuadKeyと高度キーの組を配列にしたもの
- `altitudekeyZoom` 垂直方向ズームレベル
- `zBaseExponent` 高度キー1つの実際の高さが1mになるズームレベル
- `zBaseOffset` 高度キー0番に対応する空間ID垂直インデックス
    - ここでは高度キー0番のズームレベルを`zBaseExponent`で扱う

![fig.1](./image/tilekey-standard.png)
<div style="text-align: center">fig1. FromExtendedSpatialIDToQuadkeyAndAltitudekeyで用いる値</div>

出力は拡張空間ID文字列になる  
上記の入力は以下の文字列配列として出力される

```
["20/85263/65423/23/-2"]
```

上記の例では入力と出力は1:1であるが、この変換では1:Nに変換されることがある

その条件は次のどちらか

1. `altitudekeyZoom`が`zBaseExponent`または25(これは空間IDの高度基準における`zBaseExponent`である)より大きい場合
2. `zBaseOffset`が2のべき乗でない場合

また、変換の前後でAltitudeKeyにそのズームレベルで存在しないインデックスが現れた場合エラーとなる  
この際変換は失敗となり、変換後のデータはnilとなる

### 変換例

#### 1. 入力altitudekeyが出力拡張空間ID垂直インデックスに対応する場合

変換前

```
[]FromExtendedSpatialIDToQuadkeyAndAltitudekey{
    {
        quadkeyZoom : 20
        innerIDList : [[7432012031,0]]
        altitudekeyZoom : 23
        zBaseExponent : 25
        zBaseOffset : 8
    }
}
```

変換後

```
extendedSpatialIDs :["20/85263/65423/23/-2"]
```

#### 2. altitudekeyZoomが25より大きい場合

変換前

```
[]FromExtendedSpatialIDToQuadkeyAndAltitudekey{
    {
        quadkeyZoom : 20
        innerIDList : [[7432012031,3]]
        altitudekeyZoom : 26
        zBaseExponent : 25
        zBaseOffset : -2
    }
}
```

変換後

extendedSpatialIDs :["20/85263/65423/26/7", "20/85263/65423/26/7]

#### 3. zBaseOffsetが2のべき乗でない場合

変換前

```
[]FromExtendedSpatialIDToQuadkeyAndAltitudekey{
    {
        quadkeyZoom : 20
        innerIDList : [[7432012031,0]]
        altitudekeyZoom : 23
        zBaseExponent : 25
        zBaseOffset : 7
    }
}
```

変換後

```
extendedSpatialIDs :["20/85263/65423/23/-2", "20/85263/65423/23/-1"]
```

