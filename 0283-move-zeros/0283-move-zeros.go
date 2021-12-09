package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
     		return
     	}

	nonzeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
            if i != nonzeroIndex {
                nums[i], nums[nonzeroIndex] = nums[nonzeroIndex], nums[i]
            }

            nonzeroIndex++
		}
	}
}
