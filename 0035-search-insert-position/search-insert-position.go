package leetcode

func searchInsert(nums []int, target int) int {
    for i, v := range nums {
        if (target == v || target < v) {
                return i
        }
    }

    return len(nums)
}

