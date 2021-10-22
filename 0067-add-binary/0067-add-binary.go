package leetcode

import "strconv"

func addBinary(a string, b string) string {
    return algorithm(a, b, 0, "")
}

func algorithm(a string, b string, carry int, result string) string {
    if len(a) == 0 && len(b) == 0 {
        if carry == 1 {
            return "1" + result
        }

        return result
    }

    intA := getCurrentValue(a)
    intB := getCurrentValue(b)
    result = strconv.Itoa(intA ^ intB ^ carry) + result
    if intA + intB + carry > 1 {
        carry = 1
    } else {
        carry = 0
    }

    return algorithm(
        getNextString(a),
        getNextString(b),
        carry,
        result)
}

func getCurrentValue(s string) int {
    if len(s) == 0 {
        return 0
    }

    value, _ := strconv.Atoi(string(s[len(s) - 1]))
    return value
}

func getNextString(s string) string {
    if len(s) == 0 {
        return ""
    }

    return s[:len(s) - 1]
}
