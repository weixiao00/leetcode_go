package algorithm

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Codec1 struct{}

func Constructor1() Codec1 {
	return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	s := ""
	dfs1(root, &s)
	return s
}

func dfs1(root *TreeNode, s *string) {

	if root == nil {
		*s += "null,"
	} else {
		*s += strconv.Itoa(root.Val) + ","
		dfs1(root.Left, s)
		dfs1(root.Right, s)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	var de func() *TreeNode

	// 闭包可以使用方法里全局变量
	data1 := strings.Split(data, ",")
	de = func() *TreeNode {
		s := data1[0]
		if s == "null" {
			data1 = data1[1:]
			return nil
		}
		i, _ := strconv.Atoi(s)
		node := &TreeNode{i, nil, nil}
		data1 = data1[1:]
		node.Left = de()
		node.Right = de()
		return node
	}

	return de()
}

func TestLcr048(t *testing.T) {
	//1,2,3,null,null,4,5
	root := TreeNode{1, nil, nil}
	left := TreeNode{2, nil, nil}
	right := TreeNode{3, nil, nil}
	left4 := TreeNode{4, nil, nil}
	right5 := TreeNode{5, nil, nil}
	root.Left = &left
	root.Right = &right
	right.Left = &left4
	right.Right = &right5
	codec1 := Codec1{}
	s := codec1.serialize(&root)
	fmt.Println(s)
	retRoot := codec1.deserialize(s)
	fmt.Println(retRoot.Val)
}
