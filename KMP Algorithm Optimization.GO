func kmpMatchOptimized(text string, pattern string) int {
    n := len(text)
    m := len(pattern)

    if m == 0 {
        return 0
    }

    // 计算模式串的 next 数组
    next := make([]int, m)
    next[0] = -1

    j := -1
    for i := 1; i < m; i++ {
        for j >= 0 && pattern[j+1] != pattern[i] {
            j = next[j]
        }
        if pattern[j+1] == pattern[i] {
            j++
        }
        if j+1 < i && pattern[j+1] == pattern[next[j]+1] {
            next[i] = next[j]
        } else {
            next[i] = j
        }
    }

    // 在文本串中查找模式串
    j := -1
    for i := 0; i < n; i++ {
        for j >= 0 && pattern[j+1] != text[i] {
            j = next[j]
        }
        if pattern[j+1] == text[i] {
            j++
        }
        if j == m-1 {
            return i-m+1
        }
    }

    return -1
}
