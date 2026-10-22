func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix := make([]int, n)
	suffix[n-1] = 1

	for i := n-2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}

	return res
}
