package tree

type Binarytree struct {
	RootNode *Node //根节点
}

//创建空二叉树
func NewTree() *Binarytree {
	return &Binarytree{}
}

//设置二叉树的根节点
func (bt *Binarytree) SetRootNode(node *Node) {
	bt.RootNode = node
}

//前置遍历二叉树
func (bt *Binarytree) FrontTraverseTree() {
	bt.RootNode.TraverseNode()
}

//中序遍历
func (bt *Binarytree) MidTraverseTree() {
	bt.RootNode.MidTraverseNode()
}

//通过指定的值找节点(递归)
func (bt *Binarytree) FindByRecursive(value int) *Node {
	if bt.RootNode != nil {
		return bt.RootNode.FindByRecursive(value)
	}
	return nil
}

//通过指定的值找节点(循环)
func (bt *Binarytree) FindByCycle(value int) *Node {
	if bt.RootNode != nil { //根节点不为空
		return bt.RootNode.FindByCycle(value)
	}
	return nil
}

//查找最小值
func (bt *Binarytree) FindMin() *Node {
	if bt.RootNode != nil {
		return bt.RootNode.FindMin()
	}
	return nil
}

//差找最大值
func (bt *Binarytree) FindMax() *Node {
	if bt.RootNode != nil {
		return bt.RootNode.FindMax()
	}
	return nil
}

//插入一个值
func (bt *Binarytree) Insert(value int) *Node {
	if bt.RootNode != nil {
		return bt.RootNode.Insert(value)
	}
	//当树为空树 初始化跟节点
	bt.RootNode = NewNode(value)
	return bt.RootNode
}
