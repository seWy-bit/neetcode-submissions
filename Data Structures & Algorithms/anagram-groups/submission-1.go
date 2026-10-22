func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, word := range strs {
		freq := [26]int{}

		for _, ch := range word {
			freq[ch - 'a']++
		}

		groups[freq] = append(groups[freq], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
