package main

import "fmt"

// Node is a representation of a single node in tree. (recursive ADT)
type Node struct {
	key   int
	left  *Node
	right *Node
}

// insert is a recursive method for node insertion
func (root *Node) insert(newNode *Node) {

	//if data exists, skip
	if root.key == newNode.key {
		return
	}

	// to right-subtree
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

// inOrderPrint is an in_order traversal (recursively)
func (root *Node) inOrder(fn func(n *Node)) {

	if root.left != nil {
		root.left.inOrder(fn)
	}

	fn(root)

	if root.right != nil {
		root.right.inOrder(fn)
	}

}

// Bst is a tree
// invariant:
//
// given x root, values of all left sub-tree < x, values of all right sub-tree > x
type Bst struct {
	root   *Node
	length int
}

func (bst *Bst) insert(value int) {
	node := &Node{key: value}
	bst.length++

	if bst.root == nil {
		bst.root = node
	} else {
		bst.root.insert(node)
	}
}

func (bst *Bst) getMinimumValueLargerThan(value int) (int, error) {
	if bst.root == nil {
		return 0, fmt.Errorf("Empty tree")
	}
	node, err := bst.find(value)
	if err != nil {
		return 0, err
	}
	if node.key != value {
		return 0, fmt.Errorf("Value %d not found", value)
	}
	minimum, err := bst.min(node)
	if err != nil {
		return 0, err
	}
	if minimum.right != nil {
		rmin, err := bst.min(minimum.right)
		if err != nil {
			return 0, err
		}
		return rmin.key, nil
	}
	if minimum.key == value {
		return 0, fmt.Errorf("No value less than %d", value)
	}
	return minimum.key, nil
}

func (bst *Bst) find(value int) (*Node, error) {
	if bst.root == nil {
		return nil, fmt.Errorf("Empty tree")
	}
	v := bst._find(bst.root, value)
	if v.key != value {
		return nil, fmt.Errorf("Value %d not found", value)
	}
	return v, nil
}

func (bst *Bst) min(node *Node) (*Node, error) {
	if node == nil {
		return nil, fmt.Errorf("nil input node")
	}
	if node.left == nil {
		return node, nil
	}
	return bst.min(node.left)
}

func (bst *Bst) _find(node *Node, value int) *Node {
	if node.key == value {
		return node
	}
	if node.key < value {
		return bst._find(node.right, value)
	}
	return bst._find(node.left, value)

}

func (bst *Bst) inorderPrint() {
	bst.root.inOrder(func(node *Node) {
		fmt.Printf("%p: %v\n", &node, node)
	})
}

func main() {

	tree := new(Bst)
	tree.insert(1)
	tree.insert(0)
	tree.insert(20)
	tree.insert(4)
	tree.insert(5)
	tree.insert(6)

	tree.inorderPrint()
	fmt.Println(tree.getMinimumValueLargerThan(0))
}
