package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

// Insert inserts a Node into the tree
func (tree *BinarySearchTree) Insert(data int) *BinarySearchTree {
	if tree.Root == nil {
		tree.Root = &Node{Value: data, Left: nil, Right: nil}
	} else {
		tree.Root.insert(data)
	}
	return tree
}

// Remove removes a single node from the binary tree using a
// recursive
func (tree *BinarySearchTree) Remove(data int) {
	tree.Root.remove(data, tree.Root)
}

func (tree *BinarySearchTree) Find(data int) *Node {
	return tree.Root.find(data)
}

// Insert recursively inserts a new Node into the tree.
func (n *Node) insert(data int) {
	if n == nil {
		return
	}

	if data == n.Value {
		return
	}

	if data < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: data}
			return
		}
		n.Left.insert(data)
	}

	if data > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: data}
			return
		}
		n.Right.insert(data)
	}
	return
}

func (n *Node) remove(data int, parent *Node) {
	if data < parent.Value {
		n.Left.remove(data, n)
	} else if data > parent.Value {
		n.Right.remove(data, n)
	} else {
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return
		}

		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return
		}
	}
}

func (n *Node) find(data int) *Node {
	// traverse the tree checking each Nodes value
	// against the passed data parameter
	if n == nil {
		return nil
	}
	if data == n.Value {
		return n
	}
	if data < n.Value {
		n.Left.find(data)
	}
	if data > n.Value {
		n.Right.find(data)
	}
	return nil
}

func (n *Node) replaceNode(parent, replacement *Node) {
	if n == nil {
		return
	}
	if n == parent.Left {
		parent.Left = replacement
		return
	}
	parent.Right = replacement
}
