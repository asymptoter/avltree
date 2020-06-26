package avltree

type AVLTree interface {
	Less(a, b interface{}) bool
}

type Node struct {
	Left   *Node
	Right  *Node
	Value  interface{}
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

func (r *Node) balance() *Node {
	diffH := r.diffHeight()
	if diffH > 1 {
		if r.Left.diffHeight() < 0 {
			r.Left = LeftRotate(r.Left)
		}
		r = RightRotate(r)
	} else if diffH < -1 {
		if r.Right.diffHeight() > 0 {
			r.Right = RightRotate(r.Right)
		}
		r = LeftRotate(r)
	}
	return r
}

func Less(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func (r *Node) Insert(value int) *Node {
	if r == nil {
		return &Node{
			Value:  value,
			Height: 1,
		}
	}

	r.Height++
	if Less(value, r.Value) {
		r.Left = r.Left.Insert(value)
	} else {
		r.Right = r.Right.Insert(value)
	}

	return r.balance()
}

func (r *Node) Delete(value int) *Node {
	cur := r
	for {
		if cur.Value == value {
		}

	}
	return nil
}

func (r *Node) diffHeight() int {
	if r == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0
	if r.Left != nil {
		leftHeight = r.Left.Height
	}
	if r.Right != nil {
		rightHeight = r.Right.Height
	}

	return leftHeight - rightHeight
}
