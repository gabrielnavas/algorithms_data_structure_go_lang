package main

import "fmt"

type Tree struct {
	LeftNode  *Tree
	num       int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree == nil {
		tree = &Tree{nil, m, nil}
		return
	}
	if tree.LeftNode == nil {
		tree.LeftNode = &Tree{nil, m, nil}
	} else if tree.RightNode == nil {
		tree.RightNode = &Tree{nil, m, nil}
	} else if tree.LeftNode != nil {
		tree.LeftNode.insert(m)
	} else {
		tree.RightNode.insert(m)
	}
}

func (tree *Tree) show() {
	if tree != nil {
		fmt.Println(tree.num)
		tree.LeftNode.show()
		tree.RightNode.show()
	}
}

func main() {
	tree := Tree{nil, 2, nil}
	for _, n := range []int{5, 8, 9, 10, 15, 20, 33, 77, 100, 150, 162} {
		tree.insert(n)
	}
	tree.show()
	fmt.Println("this example is O(log n) -> Logarithmic complexity.")
}
