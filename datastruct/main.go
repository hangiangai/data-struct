package main

import (
	"fmt"
	"goworkerspace/datastruct/tree"
)

func main() {
	//创建一颗空树
	t := tree.NewTree()

	//创建根节点
	root := tree.NewNode(10)
	//为二叉树设置根节点
	t.SetRootNode(root)

	rootl := tree.NewNode(8)
	rootr := tree.NewNode(12)

	root.SetLeftNode(rootl)
	root.SetRightNode(rootr)

	lnodel := tree.NewNode(6)
	lnoder := tree.NewNode(9)

	rootl.SetLeftNode(lnodel)
	rootl.SetRightNode(lnoder)

	rnodel := tree.NewNode(11)
	rnoder := tree.NewNode(15)

	rootr.SetLeftNode(rnodel)
	rootr.SetRightNode(rnoder)

	fmt.Println(t.FindByCycle(15))

	fmt.Println(t.Insert(13))

	fmt.Println(t.Insert(16))

	t.FrontTraverseTree()

	fmt.Println("=======================")

	arr := []int{12, 3, 4, 5, 56, 6, 7}

	arr1 := make([]int, 0, 0)

	for _, val := range arr {

	}

}
