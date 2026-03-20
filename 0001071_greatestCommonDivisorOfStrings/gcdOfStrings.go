package _001071_greatestCommonDivisorOfStrings

func gcdOfStrings(str1 string, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	n := min(n1, n2)

	answer := ""

	for i := 1; i <= n; i++ {
		if n1%i != 0 || n2%i != 0 {
			continue
		}
		print("i: ")
		println(i)
		for j := 1; j <= n; j += i {
			print("str: ")
			println(str1[n1-i:])
			print("str1: ")
			println(str1[n1-j : n1-i+1])
			print("str2: ")
			println(str2[n2-j : n2-i+1])
			print("answer: ")
			println(answer)
			println("")
			if j >= n1 || str1[n1-i:] != str1[n1-j:n1-i+1] {
				break
			} else if j >= n2 || str1[n1-i:] != str2[n2-j:n2-i+1] {
				break
			}
		}
		answer = str1[n-i:]
	}

	return answer
}
