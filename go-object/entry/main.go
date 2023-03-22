package main

import (
	"fmt"
	"github.com/menah3m/go-tutorial/go-object/tree"
)

/*
   @Auth: menah3m
   @Desc:
*/

type myNode struct {
	node *tree.Node
}

func main() {
	root := tree.Node{
		Value: 3,
	}
	root.Left = &tree.Node{}
	root.Left.Value = 4
	root.Right = &tree.
		Node{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = tree.CreateNode(2)
	root.Left.Right = tree.CreateNode(2)

	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }
	// fmt.Println(nodes)
	// root.print()
	// root.setValue(100)
	// root.print()
	// root.left.right.setValue(3)
	// root.left.right.print()
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node  value:", maxNode)

}

func (m *myNode) postOrder() {
	if &m == nil || m.node == nil {
		return
	}
	left := myNode{m.node.Left}
	left.postOrder()
	right := myNode{m.node.Right}
	right.postOrder()
	m.node.Print()
}

func (m myNode) preOrder() {
	if &m == nil || m.node == nil {
		return
	}
	m.node.Print()
	left := myNode{m.node.Left}
	left.preOrder()
	right := myNode{m.node.Right}
	right.preOrder()
}
