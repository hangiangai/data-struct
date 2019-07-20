package tree

import "fmt"

type Node struct {
	leftNode  *Node //左节点
	nodeData  int   //节点数据
	rightNode *Node //右节点
}

//生成一个新的节点
func NewNode(value int) *Node {
	return &Node{
		leftNode:  nil,
		nodeData:  value,
		rightNode: nil,
	}
}

//设置左节点
func (node *Node) SetLeftNode(lnode *Node) {
	node.leftNode = lnode
}

//设置右节点
func (node *Node) SetRightNode(rnode *Node) {
	node.rightNode = rnode
}

func (node *Node) TraverseNode() {

	fmt.Println(node.nodeData)
	if node.leftNode != nil {
		node.leftNode.TraverseNode()
	}
	if node.rightNode != nil {
		node.rightNode.TraverseNode()
	}
}

//中序查找
func (node *Node) MidTraverseNode() {
	if node.leftNode != nil {
		node.leftNode.TraverseNode()
	}
	if node.rightNode != nil {
		node.rightNode.TraverseNode()
	}
}

func (node *Node) FindByCycle(v int) *Node {

	for node != nil {
		if v > node.nodeData { //向右子树移动
			node = node.rightNode
		} else if v < node.nodeData { //向左子树移动
			node = node.leftNode
		} else {
			return node
		}
	}
	return nil //不存在返回nil
}

func (node *Node) FindByRecursive(v int) *Node {
	for node != nil { //node.rightNode
		if v > node.nodeData {
			node.rightNode.FindByRecursive(v)
		} else if v < node.nodeData {
			node.leftNode.FindByRecursive(v)
		} else {
			return node
		}
	}
	return nil
}

//查找最小值
func (node *Node) FindMin() *Node {
	for node.leftNode != nil {
		node = node.leftNode
	}
	return node
}

//查找最大值
func (node *Node) FindMax() *Node {
	if node.rightNode == nil { //递归调用 结束条件为右节点为空
		return node
	}
	return node.rightNode.FindMax()
}

//插入数据
func (node *Node) Insert(v int) *Node {

	if v > node.nodeData {
		if node.rightNode == nil {
			node.rightNode = NewNode(v)
			return node.rightNode
		}
		node.rightNode.Insert(v)
	} else if v < node.nodeData {
		if node.leftNode == nil {
			node.leftNode = NewNode(v)
			return node.leftNode
		}
		node.leftNode.Insert(v)
	}
	return node
}

//删除
func (node *Node) Delete(v int) {

}
