package match

func WdMatch(s, str string) string {
	j := 0
	for i := 0; i < len(s); i++ {
		found := false
		for j < len(str) {
			if s[i] == str[j] {
				found = true
				j++
				break
			}
			j++
		}
		if !found {
			return "Your string don't match..."
		}
	}
	return s
}
