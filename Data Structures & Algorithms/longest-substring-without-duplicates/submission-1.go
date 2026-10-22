func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	m := make(map[byte]bool)
	l, res := 0, 0

	for r := 0; r < len(s); r++ {
		for m[s[r]] {
			delete(m, s[l])
			l++
		}

		m[s[r]] = true
		if r - l + 1 > res {
			res = r - l + 1
		}
	}

	return res
}
