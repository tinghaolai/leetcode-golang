package leetcode

func buildArray(nums []int) []int {
    answer := []int{}
    for _, v := range nums {
        answer = append(answer, nums[v])
    }

    return answer
}
