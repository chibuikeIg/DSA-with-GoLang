package main

import "fmt"

type Node struct {
	value int
	right *Node
	left  *Node
}

type BinarySearchTree struct {
	root   *Node
	length int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) insert(value int) {

	if bst.length == 0 {
		bst.root = &Node{value: value}
	}

	bst.addByTraverse(value)

	bst.length += 1

}

func (bst *BinarySearchTree) lookup(value int) *Node {

	currentNode := bst.root

	counter := 0

	for counter != bst.length {

		if currentNode.value < value {
			currentNode = currentNode.right
		} else if currentNode.value > value {
			currentNode = currentNode.left
		} else if currentNode.value == value {
			break
		}

		counter++

	}

	return currentNode

}

func (bst *BinarySearchTree) addByTraverse(value int) {

	currentNode := bst.root

	counter := 0

	for counter != bst.length {

		if currentNode.value < value && currentNode.right == nil {

			currentNode.right = &Node{value: value}
			break

		} else if currentNode.value < value && currentNode.right != nil {

			currentNode = currentNode.right

		}

		if currentNode.value > value && currentNode.left == nil {

			currentNode.left = &Node{value: value}
			break

		} else if currentNode.value > value && currentNode.left != nil {

			currentNode = currentNode.left

		}

		counter++

	}

}

func main() {

	newBST := NewBinarySearchTree()

	newBST.insert(3)

	newBST.insert(5)

	newBST.insert(2)

	newBST.insert(4)

	fmt.Println(newBST.root.right.left.value)

	fmt.Println(newBST.lookup(5).left)

}
