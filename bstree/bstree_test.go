package bstree

import (
	"fmt"
	"github.com/linsibolinhong/datastruct/compare"
	"math/rand"
	"testing"
)

func Test_BSTree(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	tree := NewRBTree()
	for i := 0; i < 10000; i ++ {
		k := rand.Int() % 1000
		tree.Insert(compare.IntCompare(k))
		//tree.Root.Print()
	}

	tree.Root.Print()
	s := tree.Slice()
	for i := 0; i < 110; i++ {
		k := s[rand.Int() % len(s)]
		tree.Delete(k)
		fmt.Println(tree.Root.IsRBTree())
	}

	//tree.Delete(sl[1])
	//tree.Root.Print()

	//fmt.Println(tree.Slice())
	//fmt.Println(tree.Root.IsRBTree())



}
