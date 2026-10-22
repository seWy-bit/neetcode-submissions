func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1
	res := make([]int, 0, 2)

	for l < r {
		if numbers[l] + numbers[r] > target {
			r--
		}
		if numbers[l] + numbers[r] < target {
			l++
		}
		if numbers[l] + numbers[r] == target {
			res = append(res, l+1)
			res = append(res, r+1)
			break
		}
	}

	return res
}
