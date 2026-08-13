type MinStack struct {
	value  []int
	min    []int // keep track of min val sorted  
}

func Constructor() MinStack {
	return MinStack{
		value: []int{},
		min  :  []int{},
	}
}

func (this *MinStack) Push(val int) {
	// never changes if not for this
	if len(this.value) == 0 {
		this.min = append(this.min, val)
	}
	this.value = append(this.value, val)
	// only min vals keep getting added: [2, 1, 0, -2,...]
	if val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	// if stack if empty
	if len(this.value) == 0 {
		return
	}
	popVal := this.value[len(this.value)-1]
	if popVal == this.min[len(this.min)-1] {
		// decrement the min stack by one to get to next el
		this.min = this.min[:len(this.min)-1]
	}
	this.value = this.value[:len(this.value)-1] 
}

func (this *MinStack) Top() int {
	if len(this.value) == 0 {
		return 0
	}
	return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
