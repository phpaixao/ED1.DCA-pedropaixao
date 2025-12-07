package main

import (
	"fmt"
	"errors"
)

type BstNode struct {
	left   *BstNode
	value  int
	height int
	bf     int
	right  *BstNode
}

func (root *BstNode) UpdateProperties() {
	hleft, hright := 0
	if root.left != nil {
		hleft = root.left.height
	}
	if root.right != nil {
		hright = root.right.height
	}
	root.bf = hright - hleft
	if root.left == nil && root.right == nil {
		root.height = 0
	} else {
		if hleft >= hright {
			root.height = hleft + 1
		} else {
			root.height = hright + 1
		}
	}
}

func (root *BstNode) RotRight() *BstNode {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *BstNode) RotLeft() *BstNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (root *BstNode) RebalanceLeftLeft() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftNeutral() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftRight() *BstNode {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *BstNode) RebalanceRightRight() *BstNode {
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightNeutral() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceRightLeft() *BstNode {
	root.right = root.right.RotRight()
	return root.RotLeft()
}

func (root *BstNode) Rebalance() *BstNode {
	if root == nil {
		return nil
	}
	if root.bf < -1 {
		if root.left.bf == -1 {
			root = root.RebalanceLeftLeft()
		} else if root.left.bf == 0 {
			root = root.RebalanceLeftNeutral()
		} else {
			root = root.RebalanceLeftRight()
		}
	} else if root.bf > 1 {
		if root.right.bf == 1 {
			root = root.RebalanceRightRight()
		} else if root.right.bf == 0 {
			root = root.RebalanceRightNeutral()
		} else {
			root = root.RebalanceRightLeft()
		}
	}
	return root
}