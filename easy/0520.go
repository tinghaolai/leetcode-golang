package main

import (
    "fmt"
)

func main() {
    if (detectCapitalUse("AAAAAA")) {
        fmt.Print("true")
    } else {
        fmt.Print("false")
    }
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Detect Capital.
// Memory Usage: 1.9 MB, less than 84.00% of Go online submissions for Detect Capital.
func detectCapitalUse(word string) bool {
    upperCaseLastPos := 0
    upperCaseCount := 0
    for i := 0; i < len(word); i++ {
        if (word[i] >= 65 && word[i] <= 90) {
            upperCaseLastPos = i
            upperCaseCount++
        }
    }

    return upperCaseCount == 0 || upperCaseLastPos == 0 || upperCaseCount == len(word)
}
