

package main

import (
    "fmt"
)

func main() {
    result := reverseStr("abcdefg", 8)
    fmt.Print(result)
}

//Runtime: 7 ms, faster than 13.85% of Go online submissions for Reverse String II.
//Memory Usage: 3.8 MB, less than 32.31% of Go online submissions for Reverse String II.
func reverseStr(s string, k int) string {
    maxIndex := len(s) - 1
    runes := []rune(s)
    for index := 0; index <= len(runes) / k; index++ {
        if (index % 2 == 1) {
            continue;
        }

        for i, j := index * k, min(maxIndex, (index + 1) * k - 1); i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
    }

    return string(runes)
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
