package main

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

}
