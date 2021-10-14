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

func removeDuplicates_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func removeDuplicates_2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length-1; i++ {
		if nums[i] == lastNum {

			break
		}
		if nums[i+1] == nums[i] {
			removeElement(nums, i+1, nums[i])
		}
	}
	return i + 1
}

func removeElement(nums []int, start, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}
