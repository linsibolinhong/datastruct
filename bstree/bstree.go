package bstree

import (
	"github.com/linsibolinhong/datastruct/compare"
)

type BSTreeNode struct {
	Value  compare.Comparer
	Left   *BSTreeNode
	Right  *BSTreeNode
	Parent *BSTreeNode `json:"-"`

	Height int
}

func (t *BSTreeNode) Slice() []compare.Comparer {
	ret := []compare.Comparer{}
	if t == nil {
		return  ret
	}

	if t.Left != nil {
		ret = t.Left.Slice()
	}

	ret = append(ret, t.Value)
	ret = append(ret, t.Right.Slice()...)

	return ret
}

func (t *BSTreeNode) InsertNode(node *BSTreeNode) {
	if node == nil {
		return
	}

	node.Left = nil
	node.Right = nil
	node.Height = 0
	node.Parent = nil

	if t.Value.Compare(node.Value) < 0 {
		if t.Right == nil {
			t.Right = node
			node.Parent = t
			if t.Left == nil {
				t.Height = 1
			}
			return
		}

		t.Right.InsertNode(node)
		return
	}

	if t.Left == nil {
		t.Left = node
		node.Parent = t
		if t.Right == nil {
			t.Height = 1
		}
		return
	}

	t.Left.InsertNode(node)
}

func (t *BSTreeNode) updateHeight() {
	if t == nil {
		return
	}

	updated := false
	if t.Right != nil && t.Height < t.Right.Height + 1 {
		t.Height = t.Right.Height + 1
		updated = true
	}

	if t.Left != nil && t.Height < t.Left.Height + 1 {
		t.Height = t.Left.Height + 1
		updated = true
	}

	if updated && t.Parent != nil {
		t.Parent.updateHeight()
	}
}

func (node *BSTreeNode) DeleteNode() *BSTreeNode {
	var newBranch *BSTreeNode
	if node == nil {
		return nil
	}

	if node.Right != nil {
		newBranch = node.Right
		for newBranch.Left != nil {
			newBranch = newBranch.Left
		}

		if newBranch != node.Right {
			newBranch.Parent.Left = nil
			newBranch.Right = node.Right
		}

		newBranch.Left = node.Left
	} else {
		newBranch = node.Left
	}

	if newBranch != nil {
		newBranch.Parent = node.Parent
	}

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = newBranch
		} else {
			node.Parent.Right = newBranch
		}
		return nil
	} else {
		return newBranch
	}
}

func (t *BSTreeNode) Delete(val compare.Comparer) (*BSTreeNode, bool) {
	if t == nil {
		return nil, false
	}

	c := t.Value.Compare(val)
	if c == 0 {
		return t.DeleteNode(), true
	}

	if c > 0 {
		return t.Left.Delete(val)
	}

	return t.Right.Delete(val)
}

type BSTree interface {
	Insert(val compare.Comparer)
	Search(val compare.Comparer) compare.Comparer
	Delete(val compare.Comparer) bool
	Slice() []compare.Comparer
}


type BasicBSTree struct {
	Root *BSTreeNode
}

func NewBaiscBSTree() *BasicBSTree {
	return &BasicBSTree{}
}

func (t *BasicBSTree) Insert(val compare.Comparer) {
	newNode := &BSTreeNode{
		Value:  val,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Height: 0,
	}

	if t.Root == nil {
		t.Root = newNode
	} else {
		t.Root.InsertNode(newNode)
	}
}

func (t *BasicBSTree) Search(val compare.Comparer) compare.Comparer {
	if t.Root == nil {
		return nil
	}

	return nil
}

func (t *BasicBSTree) Delete(val compare.Comparer) bool {
	if t.Root == nil {
		return false
	}

	newBranch, found := t.Root.Delete(val)
	if found && newBranch != nil {
		t.Root = newBranch
	}

	return found
}

func (t *BasicBSTree) Slice() []compare.Comparer {
	return t.Root.Slice()
}

