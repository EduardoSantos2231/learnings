package main


type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

type root struct {
	root *treeNode
}

func (r *root) insert(value int) {
	if r.treeIsEmpty() {
		r.root = &treeNode{
			value: value,
		}
		return
	}

	r.root = advanceToInsert(r.root, value)
}

func advanceToInsert(node *treeNode, value int) *treeNode {
	if node == nil {
		return &treeNode{
			value: value,
		}
	}

	if value < node.value {
		node.left = advanceToInsert(node.left, value)
	} else {
		node.right = advanceToInsert(node.right, value)
	}

	return node
}

func (r *root) delete(value int) {

}

func findNodeToRemove(node *treeNode, val int) *treeNode {

}

func (r *root) search(value int) *treeNode {

}

func searchAndAdvance(node *treeNode, val int) *treeNode {
	if node == nil {
		return nil
	}
	if node.value == val {
		return node
	}
	if node.value > val {
		return searchAndAdvance(node.left, val)
	}
	return searchAndAdvance(node.right, val)
}

func (r *root) inOrder() []int {
	if r.treeIsEmpty() {
		return []int{}
	}
	result := []int{}
	traverseInOrder(r.root, &result)
	return result
}

func traverseInOrder(node *treeNode, result *[]int) {
	if node == nil {
		return
	}
	traverseInOrder(node.left, result)
	*result = append(*result, node.value)
	traverseInOrder(node.right, result)
}

// o valor mínimo é encontrado quando percorremos a BST toda pela esquerda (traverse - atravessar)
func (r *root) min() *treeNode {
	if r.treeIsEmpty() {
		return nil
	}

	return traverseLeft(r.root)
}

// função recursiva como o caso base (finalmente chegamos ao final e à esquerda do nó não existe mais nenhum elemento), percorrendo até a pilha se desmontar (chegamos no caso base)
func traverseLeft(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return traverseLeft(node.left)
}

// o valor máximo é encontrado percorrendo a BST por toda a direita (traverse - atravessar)
func (r *root) max() *treeNode {
	if r.treeIsEmpty() {
		return nil
	}

	return traverseRight(r.root)
}

// função recursiva com caso base (finalment chegamos ao fim e o à direita do nó não existe nada), percorrendo até a pilha começar a se desmontar (chegamos no caos base - valor máximo achado)
func traverseRight(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}
	return traverseRight(node.right)
}

func (r *root) treeIsEmpty() bool {
	return r.root == nil
}
