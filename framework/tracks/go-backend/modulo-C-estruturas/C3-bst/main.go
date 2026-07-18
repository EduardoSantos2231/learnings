package main

import "fmt"

func main() {
	rootNode := &root{
		root: &treeNode{
			value: 15,
			left: &treeNode{
				value: 8,
				left: &treeNode{
					value: 5,
				},
				right: &treeNode{
					value: 11,
				},
			},
			right: &treeNode{
				value: 24,
				left: &treeNode{
					value: 19,
				},
				right: &treeNode{
					value: 28,
				},
			},
		},
	}
	rootNode.insert(2)
	rootNode.insert(4)
	rootNode.insert(10)
	ordered := rootNode.inOrder()
	fmt.Println(ordered)
	rootNode.delete(10)
	ordered = rootNode.inOrder()
	fmt.Println(ordered)
}
