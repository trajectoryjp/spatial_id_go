package spatialID

import (
	radixtree "github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
)

type SpatialIDDetector interface {
	IsOverlap(ids SpatialIDs) bool
}

type SpatialIDGreedyDetector struct {
	ids SpatialIDs
}

func NewSpatialIDGreedyDetector(ids SpatialIDs) SpatialIDDetector {
	return &SpatialIDGreedyDetector{ids}
}

func (detector *SpatialIDGreedyDetector) IsOverlap(targetIDs SpatialIDs) bool {
	return detector.ids.Overlaps(targetIDs)
}

type SpatialIDTreeDetector struct {
	positiveTree radixtree.TreeInterface
	negativeTree radixtree.TreeInterface
}

func NewSpatialIDTreeDetector(ids SpatialIDs) SpatialIDDetector {
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
	return &SpatialIDTreeDetector{positiveTree, negativeTree}
}

func (tree *SpatialIDTreeDetector) IsOverlap(targetIds SpatialIDs) bool {
	for _, id := range targetIds {
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
