package solution

func IsSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	si := 0
	tj := 0
	if (tLen == 0 && sLen == 0) || sLen == 0 {
		return true
	} else if tLen == 0 {
		return false
	}

	for tj < tLen {
		if s[si] == t[tj] {
			si++
			if si == sLen {
				return true
			}
		}
		tj++
	}
	return false
}
