package list

import (
	"fmt"
)

type Node struct {
	Val      int
	NextNode *Node
}

type List struct {
	Head *Node
}

func (l *List) Insert(val int) {
	if l.Head == nil {
		l.Head = &Node{
			Val:      val,
			NextNode: nil,
		}
		return
	}
	currentNode := l.Head
	lastNode := &Node{
		Val:      val,
		NextNode: nil,
	}
	for {
		if currentNode.NextNode == nil {
			currentNode.NextNode = lastNode
			return
		} else {
			currentNode = currentNode.NextNode
		}
	}
}

func (l *List) Remove(val int) bool {
	if l.Head == nil {
		return false
	}
	currentNode := l.Head.NextNode
	prevNode := l.Head

	if l.Head.Val == val {
		l.Head = l.Head.NextNode
		return true
	}

	for currentNode != nil {
		if currentNode.Val == val {
			prevNode.NextNode = currentNode.NextNode
			return true
		}

		prevNode = prevNode.NextNode
		currentNode = currentNode.NextNode

	}
	return false
}

func (l *List) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.NextNode
	}
}

func (l *List) Reverse() {
	if l.Head == nil {
		return
	}
	var prevNode *Node = nil
	currentNode := l.Head
	var nextNode *Node = nil
	for currentNode != nil {
		nextNode = currentNode.NextNode
		currentNode.NextNode = prevNode
		prevNode = currentNode
		currentNode = nextNode

	}
	l.Head = prevNode
}

func (l *List) Find(val int) *Node {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Val == val {
			return currentNode
		}
		currentNode = currentNode.NextNode
	}
	return nil
}

func (l *List) Len() int {
	currentNode := l.Head
	counter := 0
	for currentNode != nil {
		counter++
		currentNode = currentNode.NextNode
	}
	return counter
}
