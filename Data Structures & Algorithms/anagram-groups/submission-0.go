func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		key := getKey(word)

		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func getKey(word string) string {
	word2 := []byte(word)

	sort.Slice(word2, func(i, j int) bool {return word2[i]<word2[j]})

	sortWord := string(word2)

	return sortWord
}
