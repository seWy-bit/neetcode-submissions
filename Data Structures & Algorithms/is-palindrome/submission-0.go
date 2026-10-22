func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	runes := []rune(s)
	start, end := 0, len(runes)-1
	
	for start < end {
		for start < end && !isLetter(runes[start]) {
			start++
		}

		for end > start && !isLetter(runes[end]) {
			end--
		}

		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}
