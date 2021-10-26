package leetcode

func countRangeSum_self_1(nums []int, lower int, upper int) int {
    sums := make([]int, len(nums))
    total := 0
    for i, v := range nums {
        for j := 0; j <= i; j++ {
            sums[j] += v
            if (sums[j] >= lower && sums[j] <= upper) {
                total++
            }
        }
    }

    return total
}

