package algorithm

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := ""
	// 用闭包也能实现，不需要传指针了。还得传指针？？？
	var preOrder func(node *TreeNode, res *string)
	preOrder = func(node *TreeNode, res *string) {
		if node == nil {
			return
		}
		*res += strconv.Itoa(node.Val) + ","
		preOrder(node.Left, res)
		preOrder(node.Right, res)
	}

	preOrder(root, &res)
	return res[:len(res)-1]
}

//func preOrder(node *TreeNode, res *string) {
//	if node == nil {
//		return
//	}
//	*res += strconv.Itoa(node.Val) + ","
//	preOrder(node.Left, res)
//	preOrder(node.Right, res)
//}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	arrStr := strings.Split(data, ",")
	arrInt := make([]int, 0)
	for _, val := range arrStr {
		v, _ := strconv.Atoi(val)
		arrInt = append(arrInt, v)
	}

	var strc func(lower, upper int) *TreeNode
	strc = func(lower, upper int) *TreeNode {
		// 这是二叉搜索树才能这样处理
		if len(arrInt) == 0 || arrInt[0] < lower || arrInt[0] > upper {
			return nil
		}
		val := arrInt[0]
		node := TreeNode{val, nil, nil}
		arrInt = arrInt[1:]
		node.Left = strc(lower, val)
		node.Right = strc(val, upper)
		return &node
	}

	root := strc(math.MinInt, math.MaxInt)
	return root
}

func Test449(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	left := TreeNode{1, nil, nil}
	right := TreeNode{3, nil, nil}
	root.Left = &left
	root.Right = &right
	codec := Codec{}
	str := codec.serialize(root)
	fmt.Println(str)
	resRoot := codec.deserialize(str)
	fmt.Println(resRoot.Val)
}

//---------------------------------------------------------------------
// 第二种解法, 层次遍历。提交通过
type Codec2 struct {
}

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize2(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([]string, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
		} else {
			res = append(res, "X")
		}
		if node != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return strings.Join(res, ",")
}

func (this *Codec) deserialize2(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	arr := strings.Split(data, ",")
	queue := make([]*TreeNode, 0)
	rootVal, _ := strconv.Atoi(arr[0])
	arr = arr[1:]
	root := &TreeNode{rootVal, nil, nil}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left := arr[0]
		right := arr[1]
		arr = arr[2:]
		if left != "X" {
			val, _ := strconv.Atoi(left)
			leftNode := TreeNode{val, nil, nil}
			node.Left = &leftNode
			queue = append(queue, &leftNode)
		}
		if right != "X" {
			val, _ := strconv.Atoi(right)
			rightNode := TreeNode{val, nil, nil}
			node.Right = &rightNode
			queue = append(queue, &rightNode)
		}
	}

	return root
}

func Test4991(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	left := TreeNode{1, nil, nil}
	right := TreeNode{3, nil, nil}
	root.Left = &left
	root.Right = &right
	codec := Codec{}
	str := codec.serialize(root)
	fmt.Println(str)
	resRoot := codec.deserialize(str)
	fmt.Println(resRoot.Val)
}

//---------------------------------------------------------------------

// string和int数字类型啥的一样, 数组也是这样的
// 变成了123
func test1(res *string) {
	res1 := "123"
	*res = res1
}

// 没变
func test2(res *string) {
	res1 := "123"
	res = &res1
}

func Test1(t *testing.T) {
	res := "12"
	test2(&res)
	fmt.Println(res)
}
