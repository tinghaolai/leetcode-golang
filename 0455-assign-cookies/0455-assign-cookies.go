package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    childIndex := 0
    for _, cookieSize := range s {
        if cookieSize >= g[childIndex] {
            childIndex ++
            if childIndex == len(g) {
                return childIndex
            }
        }
    }

    return childIndex
}
