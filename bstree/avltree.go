package bstree

import (
	"encoding/json"
	"fmt"
	"github.com/linsibolinhong/datastruct/compare"
	"github.com/linsibolinhong/datastruct/list"
)

type AVLTreeNode struct {
	Left  *AVLTreeNode
	Right *AVLTreeNode

	Value compare.Comparer

	Changed bool `json:"-"`
	Height  int
}

func rightRotate(t *AVLTreeNode) *AVLTreeNode {
	if t == nil {
		return nil
	}

	if t.Left == nil {
		fmt.Println("no left child, cannot right rotate")
		return t
	}

	ret := t.Left

	t.Left = ret.Right

	ret.Right = t
	ret.Changed = true
	ret.Right.Changed = true
	ret.CalcHeight()

	return ret
}

func leftRotate(t *AVLTreeNode) *AVLTreeNode {
	if t == nil {
		return nil
	}

	if t.Right == nil {
		fmt.Println("no right child, cannot left rotate")
		return t
	}

	ret := t.Right

	t.Right = ret.Left
	ret.Left = t

	ret.Changed = true
	ret.Left.Changed = true
	ret.CalcHeight()

	return ret
}

func ll_balance(t *AVLTreeNode) *AVLTreeNode {
	return rightRotate(t)
}

func rr_balance(t *AVLTreeNode) *AVLTreeNode {
	return leftRotate(t)
}

func rl_balance(t *AVLTreeNode) *AVLTreeNode {
	t.Right = rightRotate(t.Right)
	return leftRotate(t)
}

func (t *AVLTreeNode) Print() {
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}

func lr_balance(t *AVLTreeNode) *AVLTreeNode {
	t.Left = leftRotate(t.Left)
	return rightRotate(t)
}

func getHeight(t *AVLTreeNode) int {
	if t == nil {
		return -1
	}
	return t.Height
}

func rebalance(t *AVLTreeNode) *AVLTreeNode {
	if t == nil {
		return nil
	}

	t.CalcHeight()

	//t.Print()

	lh := -1
	rh := -1
	if t.Left != nil {
		lh = t.Left.Height
	}
	if t.Right != nil {
		rh = t.Right.Height
	}

	balance := lh - rh
	newBranch := t

	if balance > 1 {
		if getHeight(t.Left.Left) >= getHeight(t.Left.Right) {
			newBranch = ll_balance(t)
		} else {
			newBranch = lr_balance(t)
		}
	}

	if balance < -1 {
		if getHeight(t.Right.Right) >= getHeight(t.Right.Left) {
			newBranch = rr_balance(t)
		} else {
			newBranch = rl_balance(t)
		}
	}

	if newBranch != t {
		//newBranch.Print()
	}

	return newBranch
}

func (t *AVLTreeNode) CalcHeight() int {
	if t == nil {
		return 0
	}

	if !t.Changed {
		return t.Height
	}

	t.Height = 0
	if t.Left != nil {
		lh := t.Left.CalcHeight()
		if lh+1 > t.Height {
			t.Height = lh + 1
		}
	}

	if t.Right != nil {
		rh := t.Right.CalcHeight()
		if rh+1 > t.Height {
			t.Height = rh + 1
		}
	}

	t.Changed = false
	return t.Height
}

func (t *AVLTreeNode) InsertNode(node, parent *AVLTreeNode, root **AVLTreeNode) {
	if node == nil {
		return
	}

	defer func() {
		newBranch := rebalance(t)
		if parent == nil {
			*root = newBranch
		} else {
			if parent.Left == t {
				parent.Left = newBranch
			} else {
				parent.Right = newBranch
			}
			parent.Changed = true
			parent.CalcHeight()
		}
	}()

	if t.Value.Compare(node.Value) <= 0 {
		if t.Right == nil {
			t.Right = node
			t.Changed = true
			t.CalcHeight()
			return
		}
		t.Right.InsertNode(node, t, root)
		return
	}

	if t.Left == nil {
		t.Left = node
		t.Changed = true
		t.CalcHeight()
		return
	}

	t.Left.InsertNode(node, t, root)
}

func (t *AVLTreeNode) DeleteNode(val compare.Comparer, parent *AVLTreeNode, root **AVLTreeNode) (found bool) {
	if t == nil {
		return false
	}

	newBranch := t
	defer func() {
		if !found {
			return
		}

		if newBranch != nil {
			newBranch.CalcHeight()
		}
		newBranch = rebalance(newBranch)
		if parent == nil {
			*root = newBranch
		} else {
			if parent.Left == t {
				parent.Left = newBranch
			} else {
				parent.Right = newBranch
			}
			parent.Changed = true
			parent.CalcHeight()
		}
	}()

	cmp := t.Value.Compare(val)
	if cmp == 0 {
		if t.Left == nil && t.Right == nil {
			newBranch = nil
			return true
		}

		if t.Left == nil {
			newBranch = t.Right
			return true
		}

		if t.Right == nil {
			newBranch = t.Left
			return true
		}

		if t.Right.Left == nil {
			newBranch = t.Right
			newBranch.Left = t.Left
			newBranch.Changed = true
			return true
		}

		l := list.New()
		cur := t.Right
		for cur.Left != nil {
			l.PushBack(cur)
			cur = cur.Left
		}

		newBranch = cur
		newBranch.Left = t.Left

		tmpNode := newBranch.Right
		newBranch.Right = t.Right

		for l.Len() > 0 {
			pathNode := l.Back().Value.(*AVLTreeNode)
			l.Remove(l.Back())
			pathNode.Left = tmpNode

			tmpNode = pathNode
			tmpNode.Changed = true
			tmpNode.CalcHeight()
			tmpNode = rebalance(tmpNode)
		}

		newBranch.Right = tmpNode
		newBranch.Changed = true

		return true
	}

	if cmp < 0 {
		if t.Right != nil {
			return t.Right.DeleteNode(val, t, root)
		}
		return false
	}

	if t.Left != nil {
		return t.Left.DeleteNode(val, t, root)
	}

	return false
}

func (t *AVLTreeNode) Slice() []compare.Comparer {
	if t == nil {
		return []compare.Comparer{}
	}

	ret := []compare.Comparer{}
	if t.Left != nil {
		ret = t.Left.Slice()
	}
	ret = append(ret, t.Value)
	if t.Right != nil {
		ret = append(ret, t.Right.Slice()...)
	}
	return ret
}

type AVLTree struct {
	Root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(val compare.Comparer) {
	newNode := &AVLTreeNode{Value: val}
	if t.Root == nil {
		t.Root = newNode
		return
	}
	t.Root.InsertNode(newNode, nil, &t.Root)
}

func (t *AVLTree) Slice() []compare.Comparer {
	if t.Root == nil {
		return []compare.Comparer{}
	}
	return t.Root.Slice()
}

func (t *AVLTree) Delete(val compare.Comparer) bool {
	if t.Root == nil {
		return false
	}

	return t.Root.DeleteNode(val, nil, &t.Root)
}

func (t *AVLTreeNode) Find(val compare.Comparer) *AVLTreeNode {
	if t == nil {
		return nil
	}

	c := t.Value.Compare(val)
	if c == 0 {
		return t
	}

	if c > 0 {
		return t.Left.Find(val)
	}

	return t.Right.Find(val)
}

func (t *AVLTree) Search(val compare.Comparer) compare.Comparer  {
	if t.Root == nil {
		return nil
	}

	node := t.Root.Find(val)
	if node != nil {
		return node.Value
	}
	return nil
}