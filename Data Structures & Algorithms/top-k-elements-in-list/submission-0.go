func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, n+1)

	for key, value := range freq {
		buckets[value] = append(buckets[value], key)
	}

	res := make([]int, 0, k)
	for i := len(buckets)-1; i>=0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}

	return res
}
