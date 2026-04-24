package solution

func RemoveStars(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '*' {
			n := len(stack)
			if n != 0 {
				stack = stack[:n-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
