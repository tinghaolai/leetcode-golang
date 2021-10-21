package leetcode

import "strconv"

func fizzBuzz(n int) []string {
    output := []string{}
    for i:= 1; i <= n; i++ {
        switch {
            case i % 15 == 0:
                output = append(output, "FizzBuzz")
            case i % 3 == 0:
                output = append(output, "Fizz")
            case i % 5 == 0:
                output = append(output, "Buzz")
            default:
                output = append(output, strconv.Itoa(i))
        }
    }

    return output
}
