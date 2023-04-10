func simpleMatch(text string, pattern string) int {
n := len(text)
m := len(pattern)

for i := 0; i <= n-m; i++ {
	j := 0
	for; j < m; j++ {
		if text[i+j] != pattern[j] {
			break
		}
	}
	if j == m {
		return i
	}
}

return -1
}
