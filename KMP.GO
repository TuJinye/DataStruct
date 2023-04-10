func kmpMatch(text string, pattern string) int {
    n := len(text)
    m := len(pattern)

    if m == 0 {
        return 0
    }

    // 计算模式串的 next 数组
    next := make([]int, m)
    j := 0
    for i := 1; i < m; i++ {
        for j > 0 && pattern[j] != pattern[i] {
            j = next[j-1]
        }
        if pattern[j] == pattern[i] {
            j++
        }
        next[i] = j
    }

    // 在文本串中查找模式串
    j = 0
    for i := 0; i < n; i++ {
        for j > 0 && text[i] != pattern[j] {
            j = next[j-1]
        }
        if text[i] == pattern[j] {
            j++
        }
        if j == m {
            return i - m + 1
        }
    }

    return -1
}
