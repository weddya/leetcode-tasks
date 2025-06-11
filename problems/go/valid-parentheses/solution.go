package valid_parentheses

func isValid(s string) bool {
	closingBrackets := map[string]string{")": "(", "]": "[", "}": "{"}

	stack := ""
	for _, ch := range s {
		if _, ok := closingBrackets[string(ch)]; ok {
			last := stack[max(0, len(stack)-1):]
			stack = stack[:max(0, len(stack)-1)]
			if last != closingBrackets[string(ch)] {
				return false
			}

			continue
		}

		stack += string(ch)
	}

	return len(stack) == 0
}
