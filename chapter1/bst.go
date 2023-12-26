package chapter1

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{Val: val, Left: nil, Right: nil}
	}

	if val < n.Val {
		n.Left = n.Left.Insert(val)
	} else if val > n.Val {
		n.Right = n.Right.Insert(val)
	}

	return n
}

func (n *Node) Search(val int) *Node {
	if n == nil || n.Val == val {
		return n
	}

	if val < n.Val {
		return n.Left.Search(val)
	}

	return n.Right.Search(val)
}

func (n *Node) InOrderTraversal() {
	if n != nil {
		n.Left.InOrderTraversal()
		fmt.Print(n.Val, " ")
		n.Right.InOrderTraversal()
	}
}
