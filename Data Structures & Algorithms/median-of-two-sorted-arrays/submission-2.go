func getKth(a []int, m int, b []int, n int, k int, aStart int, bStart int) int {
    if m > n {
        return getKth(b, n, a, m, k, bStart, aStart)
    }
    if m == 0 {
        return b[bStart+k-1]
    }
    if k == 1 {
        if a[aStart] < b[bStart] {
            return a[aStart]
        }
        return b[bStart]
    }

    i := min(m, k/2)
    j := min(n, k/2)

    if a[aStart+i-1] > b[bStart+j-1] {
        return getKth(a, m, b[bStart+j:], n-j, k-j, aStart, 0)
    }
    return getKth(a[aStart+i:], m-i, b, n, k-i, 0, bStart)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    left := (len(nums1) + len(nums2) + 1) / 2
    right := (len(nums1) + len(nums2) + 2) / 2
    return float64(getKth(nums1, len(nums1), nums2, len(nums2), left, 0, 0) +
           getKth(nums1, len(nums1), nums2, len(nums2), right, 0, 0)) / 2.0
}
