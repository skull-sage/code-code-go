package adhoc_dp

type Stack struct {
	list []int
}

func (this *Stack) push(x int) {
	this.list = append(this.list, x)
}

func (this Stack) peek() int {
	return this.list[len(this.list)-1]
}

func (this Stack) isNotEmpty() {
	return len(this.list) > 0
}

func (this *Stack) pop() int {
	x := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return x
}

func NewStack() Stack {
	return Stack{list: make([]int, 0)}
}

func calcNextSmaller(nums []int) []int {
	stack := NewStack()
	nextGreater := make([]int, len(nums), len(nums))

	for idx := 0; idx < len(nums); idx++ {

		for stack.isNotEmpty() && nums[stack.peek()] > nums[idx] {
			x := stack.pop()
			// nums[idx] is nearest smaller element to nums[x]
			phi[x] = idx

		}

		stack.push(idx)
	}

	return nextGreater
}

func calcPrevGreater(nums []int) []int {
	stack := NewStack()
	nextGreater := make([]int, len(nums), len(nums))

	for idx := 0; idx < len(nums); idx++ {

		for stack.isNotEmpty() && nums[stack.peek()] > nums[idx] {
			x := stack.pop()
			// nums[idx] is nearest smaller element to nums[x]
			phi[x] = idx

		}

		stack.push(idx)
	}

	return nextGreater
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {

	for idx := 0; idx < len(nums); idx++ {
		jdx := phi[idx]

	}

}
