package algorithm

type MinStack1 struct {
	stack1 []int
	stack2 []int
}

func Constructor6() MinStack {
	minStack := MinStack{}
	minStack.stack1 = make([]int, 0)
	minStack.stack2 = make([]int, 0)
	return minStack
}

func (this *MinStack) Push1(val int) {
	this.stack1 = append(this.stack1, val)
	if len(this.stack2) > 0 && this.stack2[len(this.stack2)-1] < val {
		this.stack2 = append(this.stack2, this.stack2[len(this.stack2)-1])
	} else {
		this.stack2 = append(this.stack2, val)
	}
}

func (this *MinStack) Pop1() {
	this.stack1 = this.stack1[0 : len(this.stack1)-1]
	this.stack2 = this.stack2[0 : len(this.stack2)-1]
}

func (this *MinStack) Top1() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) GetMin1() int {
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
