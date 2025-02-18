package spatialID

import (
	radixtree "github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
)

type SpatialIDTree struct {
	positiveTree radixtree.TreeInterface
	negativeTree radixtree.TreeInterface
}

func NewSpatialIDTree(ids []SpatialID) *SpatialIDTree {
	var positiveTree radixtree.TreeInterface
	var negativeTree radixtree.TreeInterface
	for _, id := range ids {
		if id.GetF() >= 0 {
			if positiveTree == nil {
				positiveTree = radixtree.CreateTree(radixtree.Create3DTable())
			}
			treeIndex := radixtree.Indexs{id.GetF(), id.GetX(), id.GetY()}
			positiveTree.Append(treeIndex, radixtree.ZoomSetLevel(id.GetZ()), struct{}{})
		} else {
			if negativeTree == nil {
				negativeTree = radixtree.CreateTree(radixtree.Create3DTable())
			}
			treeIndex := radixtree.Indexs{^id.GetF(), id.GetX(), id.GetY()}
			negativeTree.Append(treeIndex, radixtree.ZoomSetLevel(id.GetZ()), struct{}{})
		}
	}
	return &SpatialIDTree{positiveTree, negativeTree}
}

func (tree *SpatialIDTree) Overlaps(ids []SpatialID) bool {
	for _, id := range ids {
		var isOverlap bool
		if id.GetF() >= 0 {
			if tree.positiveTree != nil {
				treeIndex := radixtree.Indexs{id.GetF(), id.GetX(), id.GetY()}
				isOverlap = tree.positiveTree.IsOverlap(treeIndex, radixtree.ZoomSetLevel(id.GetZ()))
			}
		} else {
			if tree.negativeTree != nil {
				treeIndex := radixtree.Indexs{^id.GetF(), id.GetX(), id.GetY()}
				isOverlap = tree.negativeTree.IsOverlap(treeIndex, radixtree.ZoomSetLevel(id.GetZ()))
			}
		}
		if isOverlap {
			return true
		}
	}
	return false
}
