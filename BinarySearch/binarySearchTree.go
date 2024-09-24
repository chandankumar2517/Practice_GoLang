package binarysearch

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {

	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		return
	}

	current := bst.Root

	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = &Node{Value: value}
				return
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &Node{Value: value}
				return
			}
			current = current.Right
		}
	}
}

func (bst *BST) Search(value int) bool {

	if bst.Root == nil {
		return false
	}

	current := bst.Root
	for {
		if value == current.Value {
			return true
		} else if value < current.Value {
			if current.Left == nil {
				return false
			}
			current = current.Left
		} else {
			if current.Right == nil {
				return false
			}
			current = current.Right
		}
	}
}

func (bst *BST) Delete(value int) {
	bst.Root = bst.deleteNode(bst.Root, value)
}

func (bst *BST) deleteNode(node *Node, value int) *Node {

	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = bst.deleteNode(node.Left, value)
		return bst.Root
	} else if value > node.Value {
		node.Right = bst.deleteNode(node.Right, value)
		return bst.Root
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minValueNode := bst.minValue(node.Right)

		node.Value = minValueNode.Value

		node.Right = bst.deleteNode(node.Right, minValueNode.Value)

	}

	return node

}

func (bst *BST) minValue(node *Node) *Node {
	current := node

	if current.Left != nil {
		current = current.Left
	}

	return current
}

func (bst *BST) InOrder() {

	bst.inOrderTraversal(bst.Root)
}

func (bst *BST) inOrderTraversal(node *Node) {
	if node != nil {
		bst.inOrderTraversal(node.Left)  // Visit left subtree
		fmt.Println(node.Value)          // Visit node itself
		bst.inOrderTraversal(node.Right) // Visit right subtree
	}
}
