package leetcode

func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }

    currentIndex := 0
    step := 1
    for (currentIndex + nums[currentIndex] < len(nums) - 1) {
        var nextIndex int
        step ++
        max := 0
        for i:= currentIndex + 1; i <= currentIndex + nums[currentIndex] && i < len(nums); i++ {
            jumpLength := i + nums[i]
            if jumpLength > max {
                max = jumpLength
                nextIndex = i
            }
        }

        currentIndex = nextIndex
    }

    return step
}
