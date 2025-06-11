package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	var prefix string
	for strI, str := range strs {
		if strI == 0 {
			prefix = str
			continue
		}

		var target, compare string
		if len(str) < len(prefix) {
			target = str
			compare = prefix
		} else {
			target = prefix
			compare = str
		}

		var i int
		for i = 0; i < len(target); i++ {
			if target[i] != compare[i] {
				i--
				break
			}
		}
		minLen := min(i+1, len(target))
		prefix = target[:minLen]
	}

	return prefix
}
