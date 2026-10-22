type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""

	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for j < len(encoded) && encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])
		i = end
	}
	
	return result
}
