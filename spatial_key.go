package spatialID

import "math/big"

type SpatialKey struct {
	big.Int
}

func NewSpatialKeyFromSpatialID(spatialID SpatialID) *SpatialKey {
	spatialKey := &SpatialKey{}
	if spatialID.GetF() >= 0 {
		spatialKey.SetBit(&spatialKey.Int, 0, 1)
	}

	for i := 0; i < int(spatialID.GetZ()); i += 1 {
		spatialKey.SetBit(&spatialKey.Int, 3*i+1, uint(spatialID.GetF()) & (1 << i))
		spatialKey.SetBit(&spatialKey.Int, 3*i+2, uint(spatialID.GetX()) & (1 << i))
		spatialKey.SetBit(&spatialKey.Int, 3*i+3, uint(spatialID.GetY()) & (1 << i))
	}
	return spatialKey
}
