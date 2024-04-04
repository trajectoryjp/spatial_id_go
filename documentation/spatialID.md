# An Introduction to Spatial ID

Spatial ID (or 空間ID/くうかんID/kūkan ID in Japanese) is a system for managing information in 3D space across time.[^1] Just as [pixels](https://en.wikipedia.org/wiki/Pixel) represent data in two-dimensional space, Spatial ID uses [voxels](https://en.wikipedia.org/wiki/Voxel), or virtual boxes, to map information in three-dimensional space.

## ID Specifications
In the Spatial ID system, each voxel has a unique ID number according to one of two identification specifications.

### 1. Spatial ID: 
In the Spatial ID specification, Zoom Level `z` applies to both x, y, and f

```
{z}/{f}/{x}/{y}
```
where:
- `z` is the Zoom Level
- `x` is the Longitude Index
- `y` is the Latitude Index
- `f` is the Altitude Index

### 2. Extended Spatial ID
In the Extended Spatial ID specification the zoom level is split into two parameters: `hZoom` applies to x and y; `vZoom` applies to f

```
{hZoom}/{x}/{y}/{vZoom}/{f}
```
where:
- `hZoom` is the Horizontal Zoom Level
- `x` is the Longitude Index
- `y` is the Latitude Index
- `vZoom` is the Vertical Zoom Level
- `f` is the Altitude Index

## Zoom Levels

The zoom level determines the resolution of x, y, or f dimensions. The higher the zoom level, the more units exist in a given dimension across the same distance.[^2]

At zoom level 0 the resulting Spatial ID voxel size in meters is 

```
x: 40,075,016.68 
y: 40,075,016.68 
f: 33,554,432.00
```
Mathematically, the zoom level is n, where

```
dimension distance = (zoom level = 0 distance) / (2^n)
```

The Ministry of Economy, Trade and Industry of Japan provides an approximate Spatial ID voxel size table here[^3]



## References

[^1]: Ministry of Economy, Trade and Industry of Japan: 4次元時空間情報基盤 アーキテクチャガイドライン（γ版）. 2024/02/29. Page 7. Access Link: https://www.ipa.go.jp/digital/architecture/Individual-link/nq6ept000000g0fh-att/4dspatio-temporal-guideline-gamma.pdf

[^2]: Ministry of Economy, Trade and Industry of Japan: 4次元時空間情報基盤 アーキテクチャガイドライン（γ版）. 2024/02/29. Page 18. Access Link: https://www.ipa.go.jp/digital/architecture/Individual-link/nq6ept000000g0fh-att/4dspatio-temporal-guideline-gamma.pdf

[^3]: Ministry of Economy, Trade and Industry of Japan: 4次元時空間情報基盤 アーキテクチャガイドライン（γ版）. 2024/02/29. Page 19. Access Link: https://www.ipa.go.jp/digital/architecture/Individual-link/nq6ept000000g0fh-att/4dspatio-temporal-guideline-gamma.pdf