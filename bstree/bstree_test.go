package bstree

import (
	"fmt"
	"github.com/linsibolinhong/datastruct/compare"
	"math/rand"
	"testing"
	"time"
)

func Test_BSTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tree := NewBaiscBSTree()
	for i := 0; i < 10; i ++ {
		tree.Insert(compare.IntCompare(rand.Int() % 100))
	}
	sl := tree.Slice()
	fmt.Println(sl)
	tree.Delete(sl[1])
	fmt.Println(tree.Slice())
	fmt.Println(tree.Search(sl[2]))
}
