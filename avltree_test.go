package avltree

import (
	"testing"
)

func inorder(root *Node) []int {
	res := []int{}
	if root != nil {
		res = append(res, inorder(root.Left)...)
		res = append(res, root.Value.(int))
		res = append(res, inorder(root.Right)...)
	}
	return res
}

func preorder(root *Node) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Value.(int))
		res = append(res, preorder(root.Left)...)
		res = append(res, preorder(root.Right)...)
	}
	return res
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAVLTree(t *testing.T) {
	root := &Node{
		Value:  3,
		Height: 1,
	}
	list := []int{1, 2, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(list); i++ {
		root = root.Insert(list[i])
	}
	in := inorder(root)
	expIn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !equalInts(in, expIn) {
		t.Error("want:", expIn, "but get:", in)
	}
	pre := preorder(root)
	expPre := []int{7, 4, 2, 1, 3, 6, 5, 8, 9}
	if !equalInts(pre, expPre) {
		t.Error("want:", expPre, "but get:", pre)
	}

}
