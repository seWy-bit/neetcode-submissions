func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0
	med1, med2 := 0, 0

	for count := 0; count < (len1+len2)/2+1; count++ {
		med2 = med1
		if i < len1 && j < len2 {
			if nums1[i] > nums2[j] {
				med1 = nums2[j]
				j++
			} else {
				med1 = nums1[i]
				i++
			}
		} else if i < len1 {
			med1 = nums1[i]
			i++
		} else {
			med1 = nums2[j]
			j++
		}
	}

	if (len1+len2) % 2 == 1 {
		return float64(med1)
	}

	return float64(med1 + med2) / 2.0
}
