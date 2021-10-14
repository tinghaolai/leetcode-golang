package leetcode

func removeDuplicates(nums []int) int {
    if (len(nums) == 0) {
        return 0
    }

    currentIndex := 1
    currentValue := nums[0]
    for _, value := range nums[1:] {
        if (currentValue != value) {
            nums[currentIndex] = value
            currentIndex++
            currentValue = value
        }
    }

    return currentIndex
}
