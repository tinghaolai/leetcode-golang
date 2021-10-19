package leetcode

func countNumbersWithUniqueDigits(n int) int {
    if n < 1 {
        return 0
    }

    answer := 10
    i := 9
    total := 9
    for i > 0 && n > 1 {
        total *= i
        answer += total
        n--
        i--
    }

    return answer
}
