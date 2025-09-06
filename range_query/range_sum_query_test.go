package range_query

type SegmentNode struct {
	start    int
	end      int
	rangeVal int

	left  *SegmentNode
	right *SegmentNode
}

func buildTree(numArr *[]int, start, end int) *SegmentNode {

	if start == end {
		return &SegmentNode{start, end, (*numArr)[start], nil, nil}
	}

	mid := start + (end-start)/2

	left := buildTree(numArr, start, mid)
	right := buildTree(numArr, mid, end)
	rangeVal := left.rangeVal + right.rangeVal

	return &SegmentNode{start, end, rangeVal, left, right}

}

func updateTree(node *SegmentNode, idx, newVal int) {
	if node.start == node.end { // for leaf.left == left.right == nil
		node.rangeVal = newVal
		return
	}

	if idx <= node.left.end {
		updateTree(node.left, idx, newVal)
	} else {
		updateTree(node.right, idx, newVal)
	}

	node.rangeVal = node.left.rangeVal + node.right.rangeVal
}

func queryTail(node *SegmentNode, r int) int {
	if r == node.end {
		return node.rangeVal
	}

	if r <= node.left.end {
		return queryTail(node.left, r)
	} else {
		return node.left.rangeVal + queryTail(node.right, r)
	}

}

func queryRange(node *SegmentNode, l, r int) int {
	return queryTail(node, r) - queryTail(node, l)
}

// func queryRange(node *SegmentNode, l, r int) int {
// 	// out of range
// 	if node == nil || node.start > r || node.end < l {
// 		// we have reach to a node where there is no overlapp
// 		return 0
// 	}

// 	// found match
// 	if node.start == l && node.end == r {
// 		return node.rangeVal
// 	}

// 	if r <= node.left.end {
// 		return queryRange(node.left, l, r)
// 	} else if l >= node.right.start {
// 		return queryRange(node.right, l, r)
// 	} else {
// 		mid := l + (r-l)/2
// 		leftVal := queryRange(node.left, l, mid)
// 		rightVal := queryRange(node.right, mid, r)
// 		return leftVal + rightVal
// 	}

// }

type NumArray struct {
	numArr []int
	root   *SegmentNode
}

func Constructor(nums []int) NumArray {

	root := buildTree(&nums, 0, len(nums)-1)
	return NumArray{numArr: nums, root: root}
}

func (this *NumArray) Update(idx int, val int) {
	this.numArr[idx] = val
	updateTree(this.root, idx, val)

}

func (this *NumArray) SumRange(left int, right int) int {
	return queryRange(this.root, left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
