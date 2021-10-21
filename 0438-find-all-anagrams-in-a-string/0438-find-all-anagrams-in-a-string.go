package leetcode

import "strings"

func findAnagrams(s string, p string) []int {
    found := []int{}
    currentStringFound := make(map[string]int)
    bias := len(p) - 1
    for i := 0; i < bias && i < len(s); i ++ {
        char := string(s[i])
        if strings.Contains(p, char) {
            currentStringFound[char] ++
        }
    }

    for i:= bias; i < len(s); i++ {
        char := string(s[i])
        if strings.Contains(p, char) {
            currentStringFound[char]++
        }

        if len(currentStringFound) == len(p) {
            found = append(found, i - bias)
        }

        lastIndex := string(s[i - bias])
        currentStringFound[lastIndex] --
        if (currentStringFound[lastIndex] <= 0) {
            delete(currentStringFound, lastIndex)
        }
    }

    return found
}
