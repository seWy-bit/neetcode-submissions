func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	totalLen := len(merged)
	if totalLen % 2 == 0 {
		return float64(merged[totalLen/2-1] + merged[totalLen/2]) / 2.0
	}

	return float64(merged[totalLen/2])
}
