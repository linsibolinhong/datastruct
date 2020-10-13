package bstree

import (
	"encoding/json"
	"fmt"
	"github.com/linsibolinhong/datastruct/compare"
)

/*
 * RBTree
 * 1. node has two color, black and red
 * 2. root node is black
 * 3. leaf node (nil node) is black
 * 4. red node's children is black
 * 5. any node as a source, all path to leaf node contains same number of black node
 */

type NodeColor int

const (
	RedNode   NodeColor = 0
	BlackNode NodeColor = 1
)

type Positon int

const (
	LeftPos  Positon = 0
	RightPos Positon = 1
)

type RBNode struct {
	Left  *RBNode
	Right *RBNode
	Value compare.Comparer
	Color NodeColor // 0 is red, 1 is black
}

func rightRotateT(t, p *RBNode, root **RBNode) *RBNode {
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

	if p != nil {
		if p.Left == t {
			p.Left = ret
		} else {
			p.Right = ret
		}
	} else {
		*root = ret
		ret.Color = BlackNode
	}
	return ret
}

func leftRotateT(t, p *RBNode, root **RBNode) *RBNode {
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

	if p != nil {
		if p.Left == t {
			p.Left = ret
		} else {
			p.Right = ret
		}
	} else {
		*root = ret
		ret.Color = BlackNode
	}

	return ret
}

func (r *RBNode) IsRBTree() (int, bool) {
	if r == nil {
		return 0, true
	}

	lc, lb := 0, true
	rc, rb := 0, true

	if r.Left != nil {
		lc, lb = r.Left.IsRBTree()
		if !lb {
			return 0, false
		}
		co := r.Value.Compare(r.Left.Value)
		if co < 0 {
			return 0, false
		}
	}

	if r.Right != nil {
		rc, rb = r.Right.IsRBTree()
		co := r.Value.Compare(r.Right.Value)
		if co > 0 {
			return 0, false
		}
	}

	if !lb || !rb || lc != rc {
		return 0, false
	}

	return lc + int(r.Color), true
}

func (t *RBTree) RB_Balance(p []*RBNode) {
	var parent *RBNode
	var gparent *RBNode
	var ggparent *RBNode
	var uncle *RBNode
	var son *RBNode

	l := len(p)
	if l == 1 {
		p[0].Color = BlackNode
		return
	}
	if l < 2 {
		fmt.Println("errorrrr lenth")
		return
	}

	son = p[l-1]
	parent = p[l-2]
	if son.Color == BlackNode {
		fmt.Println("son color is black")
		return
	}

	var sonPos Positon
	if parent.Left == son {
		sonPos = LeftPos
	} else {
		sonPos = RightPos
	}


	//case 1: parent is black
	if parent.Color == BlackNode {
		return
	}

	gparent = p[l-3]
	if l >= 4 {
		ggparent = p[l-4]
	}
	var pPos Positon
	if gparent.Left == parent {
		uncle = gparent.Right
		pPos = LeftPos
	} else {
		uncle = gparent.Left
		pPos = RightPos
	}

	// case 2:
	// 1. parent is red
	// 2. uncle is red
	// infer grandparent is black
	if uncle != nil && uncle.Color == RedNode {
		gparent.Color = RedNode
		uncle.Color = BlackNode
		parent.Color = BlackNode
		t.RB_Balance(p[:l-2])
		return
	}

	// case 3:
	// 1. parent is red
	// 2. uncle is black
	// infer grandparent is black

	if pPos == LeftPos {
		if sonPos == LeftPos {
			// case 3.1
			// 1. son is left & parent is left
			rightRotateT(gparent, ggparent, &t.Root)
			gparent.Color = RedNode
			parent.Color = BlackNode
			return
		} else {
			// case 3.2
			// 1. son is right & parent is left
			leftRotateT(parent, gparent, &t.Root)
			rightRotateT(gparent, ggparent, &t.Root)
			son.Color = BlackNode
			gparent.Color = RedNode
		}
	} else {
		if sonPos == RightPos {
			// case 3.3
			// 1. son is right & parent is right
			leftRotateT(gparent, ggparent, &t.Root)
			gparent.Color = RedNode
			parent.Color = BlackNode
			return
		} else {
			// case 3.4
			// 1. son is left & parent is right
			rightRotateT(parent, gparent, &t.Root)
			leftRotateT(gparent, ggparent, &t.Root)
			son.Color = BlackNode
			gparent.Color = RedNode
		}
	}
}

func (t *RBTree) InsertNode(r, node *RBNode, path []*RBNode) {
	if path == nil {
		path = []*RBNode{}
	}
	path = append(path, r)
	if r.Value.Compare(node.Value) <= 0 {
		if r.Right == nil {
			r.Right = node
			path = append(path, node)
			t.RB_Balance(path)
			return
		}

		t.InsertNode(r.Right, node, path)
		return
	}

	if r.Left == nil {
		r.Left = node
		path = append(path, node)
		t.RB_Balance(path)
		return
	}

	t.InsertNode(r.Left, node, path)
}

func (t *RBTree) DeleteNode(path []*RBNode, del bool) {
	l := len(path)
	if l == 0 {
		return
	}

	var p, b, r, pp *RBNode
	var bl, br *RBNode
	r = path[l-1]
	if l > 1 {
		p = path[l-2]
		if p.Left == r {
			b = p.Right
		} else {
			b = p.Left
		}
	}

	if l >= 3 {
		pp = path[l-3]
	}

	if p == nil {
		t.Root = nil
		return
	}

	// case 1: r is red, delete directory
	if r.Color == RedNode {
		if p.Left == r {
			p.Left = nil
		} else {
			p.Right = nil
		}
		return
	}

	// 根据红黑树性质5，r的颜色为黑色时，其兄弟节点必然存在
	// case 2: r is left
	bl = b.Left
	br = b.Right
	if p.Left == r {
		if b.Color == BlackNode {
			if del {
				p.Left = nil
			}

			if br != nil && br.Color == RedNode {
				leftRotateT(p, pp, &t.Root)
				b.Color = p.Color
				br.Color = BlackNode
				p.Color = BlackNode
				return
			}

			if bl != nil && bl.Color == RedNode {
				rightRotateT(b, p, &t.Root)
				b.Color = RedNode
				bl.Color = BlackNode
				t.DeleteNode(path, false)

				return
			}

			// default brother's children is all black
			b.Color = RedNode
			if p.Color == RedNode {
				p.Color = BlackNode
			} else {
				t.DeleteNode(path[:l-1], false)
			}
			return
		} else {
			leftRotateT(p, pp, &t.Root)
			b.Color = BlackNode
			p.Color = RedNode
			path = append(path[:l-2], b, p, r)
			t.DeleteNode(path, del)
		}
	} else {
		if b.Color == BlackNode {
			if del {
				p.Right = nil
			}
			if bl != nil && bl.Color == RedNode {
				rightRotateT(p, pp, &t.Root)
				b.Color = p.Color
				bl.Color = BlackNode
				p.Color = BlackNode
				return
			}

			if br != nil && br.Color == RedNode {
				leftRotateT(b, p, &t.Root)
				b.Color = RedNode
				br.Color = BlackNode
				t.DeleteNode(path, false)

				return
			}

			// default brother's children is all black
			b.Color = RedNode
			if p.Color == RedNode {
				p.Color = BlackNode
			} else {
				t.DeleteNode(path[:l-1], false)
			}
			return
		} else {
			rightRotateT(p, pp, &t.Root)
			b.Color = BlackNode
			p.Color = RedNode
			path = append(path[:l-2], b, p, r)
			t.DeleteNode(path, del)
		}
	}

}

func (t *RBTree) SwapNode(node *RBNode, val compare.Comparer, path []*RBNode) bool {
	if node == nil {
		return false
	}

	if path == nil {
		path = []*RBNode{}
	}
	path = append(path, node)
	c := node.Value.Compare(val)
	if c == 0 {
		if node.Right == nil && node.Left != nil {
			node.Value = node.Left.Value
			path = append(path, node.Left)
		}

		if node.Right != nil {
			cur := node.Right
			for cur.Left != nil {
				path = append(path, cur)
				cur = cur.Left
			}
			path = append(path, cur)
			node.Value = cur.Value

			if cur.Right != nil {
				cur.Value = cur.Right.Value
				cur.Right = nil
				return true
			}
		}

		t.DeleteNode(path, true)
		return true
	}

	if c < 0 {
		return t.SwapNode(node.Right, val, path)
	} else {
		return t.SwapNode(node.Left, val, path)
	}
}

type RBTree struct {
	Root *RBNode
}

func (t *RBTree) Insert(val compare.Comparer) {
	node := &RBNode{
		Left:  nil,
		Right: nil,
		Value: val,
		Color: RedNode,
	}

	if t.Root == nil {
		node.Color = BlackNode
		t.Root = node
		return
	}

	t.InsertNode(t.Root, node, nil)
}

func (t *RBNode) Slice() []compare.Comparer {
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

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) Slice() []compare.Comparer {
	if t.Root == nil {
		return []compare.Comparer{}
	}
	return t.Root.Slice()
}

func (t *RBNode) Print() {
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}

func (t *RBTree) Delete(val compare.Comparer) bool {
	return t.SwapNode(t.Root, val, nil)
}