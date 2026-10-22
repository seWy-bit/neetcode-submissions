func maxArea(heights []int) int {
	if len(heights) == 2 {
		return min(heights[0], heights[1])
	}

	l, r := 0, len(heights)-1
	res := 0
	for l < r {
		curr := (r - l) * min(heights[l], heights[r])
		if curr > res {
			res = curr
		}

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}

	return x2
}
