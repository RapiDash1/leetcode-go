package solutions

func isOpenParanthesis(char string) bool {
	return char == "{" || char == "[" || char == "("
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	mapParanthsisPair := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	var parStack []string

	for _, char := range s {
		if isOpenParanthesis(string(char)) {
			parStack = append(parStack, string(char))
		} else {
			if len(parStack) == 0 {
				return false
			}
			if parStack[len(parStack)-1] != mapParanthsisPair[string(char)] {
				return false
			}
			parStack = parStack[:len(parStack)-1]
		}
	}

	return len(parStack) == 0
}
