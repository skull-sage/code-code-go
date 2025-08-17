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

func (this Stack) isNotEmpty() bool {
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

func calcRightSmaller(nums []int) []int {
	stack := NewStack()
	rightSmaller := make([]int, len(nums), len(nums))

	for idx := 0; idx < len(nums); idx++ {

		rightSmaller[idx] = -1
		for stack.isNotEmpty() && nums[stack.peek()] >= nums[idx] {
			top := stack.pop()
			// nums[idx] is nearest smaller element to nums[x]
			rightSmaller[top] = idx

		}

		stack.push(idx)
	}

	return rightSmaller
}

func calcLeftSmaller(nums []int) []int {
	stack := NewStack()
	leftSmaller := make([]int, len(nums), len(nums))

	for idx := len(nums) - 1; idx >= 0; idx-- {
		leftSmaller[idx] = -1
		for stack.isNotEmpty() && nums[stack.peek()] >= nums[idx] {
			top := stack.pop()
			// nums[idx] is nearest smaller element to nums[x]
			leftSmaller[top] = idx

		}

		stack.push(idx)
	}

	return leftSmaller
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {

	leftSmaller := calcLeftSmaller(nums)
	rightSmaller := calcRightSmaller(nums)

	for idx := 0; idx < len(nums); idx++ {
		leftIdx := leftSmaller[idx]
		rightIdx := rightSmaller[idx]
		if leftIdx != -1 && idx-leftIdx <= indexDiff {
			if abs(nums[idx]-nums[leftIdx]) <= valueDiff {
				return true
			}
		}
		if rightIdx != -1 && rightIdx-idx <= indexDiff {
			if abs(nums[idx]-nums[rightIdx]) <= valueDiff {
				return true
			}
		}
	}

	return false

}
