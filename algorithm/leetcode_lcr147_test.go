package algorithm

import (
	"fmt"
	"testing"
)

type MinStack struct {
	stack1 []int
	stack2 []int
}

/** initialize your data structure here. */
func Constructor3() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	if len(this.stack2) == 0 || this.stack2[len(this.stack2)-1] >= x {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	if this.stack1[len(this.stack1)-1] == this.stack2[len(this.stack2)-1] {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	this.stack1 = this.stack1[:len(this.stack1)-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack2[len(this.stack2)-1]
}

func TestLcr147(t *testing.T) {
	minStack := Constructor3()
	minStack.Push(3)
	minStack.Push(4)
	minStack.Push(1)
	minStack.Push(2)
	minStack.Pop()
	minStack.Pop()
	min := minStack.GetMin()
	fmt.Println(min)
}
