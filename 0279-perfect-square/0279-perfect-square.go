package leetcode

func numSquares(n int) int {
    return recursive(n, 1)
}

func recursive(number int, value int) int {
    minValue := 0
    for i:= number; i > 0; i-- {
        if (isPerfectSquare(i)) {
            if (number - i == 0) {
                return value
            }

            if (value + 1 == 4) {
                minValue = 4;
                break
            }

            newValue := recursive(number - i, value + 1)
            if newValue != 0 {
                if (minValue == 0) {
                    minValue = newValue
                } else if (newValue < minValue) {
                    minValue = newValue
                }
            }
        }
    }

    return minValue
}

func isPerfectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}
