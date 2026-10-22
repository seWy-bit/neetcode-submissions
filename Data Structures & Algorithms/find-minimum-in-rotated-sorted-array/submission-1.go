func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	
	l, r := 0, len(nums)-1
	res := 1001

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < res {
				res = nums[l]
			}
			break
		}
		 m := l + (r-l)/2
		 if nums[m] < res {
			res = nums[m]
		 }
		 if nums[l] <= nums[m] {
			l = m + 1
		 }
		 if nums[m] < nums[r] {
			r = m - 1
		 }
	}

	return res
}
