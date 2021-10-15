package leetcode

func searchInsert(nums []int, target int) int {
    for i, v := range nums {
        switch {
            case target == v || target < v:
                return i
        }
    }

    return len(nums)
}

