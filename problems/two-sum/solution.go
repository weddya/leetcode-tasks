package two_sum

func TwoSum(nums []int, target int) []int {
	diffs := make(map[int]int, len(nums))

	diff := 0
	for k, v := range nums {
		diff = target - v
		if d, ok := diffs[v]; ok {
			return []int{d, k}
		}

		diffs[diff] = k
	}

	return []int{}
}
