package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	r := NewNode(nil, nil)
	r.SetData("root node")

	a := NewNode(nil, nil)
	a.SetData("left node")

	b := NewNode(nil, nil)
	b.SetData("right node")

	r.le = a
	r.ri = b

	fmt.Printf("%v\n", r)
}