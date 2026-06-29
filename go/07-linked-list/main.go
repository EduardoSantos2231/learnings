package main

import (
	"linked-list/list"
)

func main() {
	node := list.Node{
		Val:      1,
		NextNode: nil,
	}
	list := list.List{
		Head: &node,
	}
	list.Insert(14)
	list.Insert(15)
	list.Insert(16)
	list.Print()
	list.Remove(15)
	list.Print()
}
