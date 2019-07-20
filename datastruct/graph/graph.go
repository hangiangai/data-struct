package graph

type Graph struct {
	g []*Node //存储链表的第一个节点
}

//创建一个图
func New() *Graph {
	return &Graph{
		g: make([]*Node, 0, 0),
	}
}

//添加节点
func (gh *Graph) Add(node *Node) {
	gh.g = append(gh.g, node)
}
