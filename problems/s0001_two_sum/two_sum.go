package problems

func twoSum(nums []int, target int) []int {
	for idx, num := range nums {
		if idx+1 < len(nums) {
			nextIdx := idx + 1
			nextItem := nums[nextIdx]

			result := num + nextItem
			if result == target {
				return []int{idx, nextIdx}
			}
		}
	}
	return nil
}
