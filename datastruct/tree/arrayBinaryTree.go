package tree

import "fmt"

type ArrayBinaryTree struct {
	data []interface{}
}

func NewArrayBinaryTree(value []interface{}) *ArrayBinaryTree {
	return &ArrayBinaryTree{
		data: value,
	}
}

func (a *ArrayBinaryTree) Traverse(index int) {
	//数组没有赋值或长度为零
	if a.data == nil || len(a.data) == 0 {
		return
	}
	//先输出自己的值
	fmt.Println(a.data[index])
	//处理左子树
	if 2*index+1 < len(a.data) {
		a.Traverse(2*index + 1)
	}
	//处理右子树
	if 2*index+2 < len(a.data) {
		a.Traverse(2*index + 2)
	}
}
