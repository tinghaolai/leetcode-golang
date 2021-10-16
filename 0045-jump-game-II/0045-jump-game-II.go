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

func jump_1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
