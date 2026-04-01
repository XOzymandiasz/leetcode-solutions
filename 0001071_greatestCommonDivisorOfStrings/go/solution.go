package greatestCommonDivisors

func GcdOfStrings(str1 string, str2 string) string {
	k1 := len(str1)
	k2 := len(str2)
	if str1+str2 != str2+str1 {
		return ""
	}
	for k2 != 0 {
		k1, k2 = k2, k1%k2
	}
	return str1[:k1]
}
