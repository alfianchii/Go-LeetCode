package practice

func twoSumQuadratic(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumLinear(nums []int, target int) []int {
	var seen map[int]int = make(map[int]int)

	for i, num := range nums {
		var partner int = target - num
		j, ok := seen[partner]

		if ok {
			return []int{j, i}
		}

		seen[num] = i
	}
	return nil
}
