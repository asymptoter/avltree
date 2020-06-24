package avltree

type Node struct {
	Left   *Node
	Right  *Node
	Value  int
	Height int
}

func getHeight(a *Node) int {
	if a == nil || (a.Left == nil && a.Right == nil) {
		return 0
	} else if a.Left == nil {
		return a.Right.Height + 1
	} else if a.Right == nil {
		return a.Left.Height + 1
	}

	res := a.Left.Height
	if res < a.Right.Height {
		res = a.Right.Height
	}
	return res + 1
}

/*
	    a               c
       / \             / \
	  b	  c     ->    a   e
	     / \         / \
	    d   e       b   d
*/
func LeftRotate(a *Node) *Node {
	c := a.Right
	d := a.Right.Left
	a.Right = d
	c.Left = a
	a.Height = getHeight(a)
	c.Height = getHeight(c)
	return c
}

/*
	    a               b
       / \             / \
	  b	  c     ->    d   a
     / \                 / \
	d   e               e   c
*/
func RightRotate(a *Node) *Node {
	b := a.Left
	e := a.Left.Right
	a.Left = e
	b.Right = a
	a.Height = getHeight(a)
	b.Height = getHeight(b)
	return b
}

func balance(root *Node) *Node {
	diffH := diffHeight(root)
	if diffH > 1 {
		if diffHeight(root.Left) < 0 {
			root.Left = LeftRotate(root.Left)
		}
		root = RightRotate(root)
	} else if diffH < -1 {
		if diffHeight(root.Right) > 0 {
			root.Right = RightRotate(root.Right)
		}
		root = LeftRotate(root)
	}
	return root
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{
			Value:  value,
			Height: 1,
		}
	}

	root.Height++
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return balance(root)
}

func Delete(root *Node, value int) *Node {
	cur := root
	for {
		if cur.Value == value {
		}

	}
	return nil
}

func diffHeight(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}

	return leftHeight - rightHeight
}
