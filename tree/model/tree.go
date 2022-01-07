package model

import "fmt"

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(h Human) {
	if t.Root == nil {
		t.Root = &Node{Key: h}
		return
	}

	t.Root.Insert(h)
}

func (t *Tree) Find(key Human) (*Node, error) {
	if t.Root == nil {
		return nil, fmt.Errorf("root not found. The tree is empty")
	}

	return t.Root.find(key)
}

func (t *Tree) Remove(key Human) error {
	if t.Root == nil {
		return fmt.Errorf("root not found. The tree is empty")
	}

	return remove(&t.Root, key)
}

type Node struct {
	Key   Human
	Left  *Node
	Right *Node
}

type Human struct {
	FirstName string
	LastName  string
	Age       uint
	Iq        uint
}

func (n *Node) Insert(h Human) {
	if h.Age < n.Key.Age {
		if n.Left == nil {
			n.Left = &Node{Key: h}
		} else {
			n.Left.Insert(h)
		}
		return
	}

	if n.Right == nil {
		n.Right = &Node{Key: h}
	} else {
		n.Right.Insert(h)
	}
}

func (n *Node) find(key Human) (*Node, error) {
	if n == nil {
		return nil, fmt.Errorf("node does not exist")
	}
	if n.Key.Age > key.Age {
		return n.Left.find(key)
	}
	if n.Key.Age < key.Age {
		return n.Right.find(key)
	}

	return n, nil
}

func remove(node **Node, key Human) error {
	n := *node
	switch {
	case n == nil:
		return fmt.Errorf("node does not exist")
	case n.Key.Age > key.Age:
		return remove(&n.Left, key)
	case n.Key.Age < key.Age:
		return remove(&n.Right, key)
	}
	if n.Left == nil && n.Right == nil {
		*node = nil
		return nil
	}
	if n.Left != nil && n.Right != nil {
		minParent := n
		minRight := n.Right

		for minRight.Left != nil {
			minParent = minRight
			minRight = minRight.Left
		}

		if minParent != n {
			minParent.Left = minRight.Right
		} else {
			minParent.Right = minRight.Right
		}

		tmp := *minRight
		tmp.Left = n.Left
		tmp.Right = n.Right
		*node = &tmp
		return nil
	}

	if n.Left != nil {
		*node = (*node).Left
	} else {
		*node = (*node).Right
	}
	return nil
}
