package adhoc_dp

type ResultArr [][]int

func addToResult(curr *[]int)

func recurSubset(idx int, curr []int, nums []int, result *ResultArr) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)

	for i := idx + 1; i < len(nums); i++ {

		curr = append(curr, nums[i])
		recurSubset(idx+1, curr, nums, result)
		curr = curr[:len(curr)-1]
	}

}

func subsets(nums []int) [][]int {

}
