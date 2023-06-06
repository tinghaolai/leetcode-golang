package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	heads := make(map[int]bool)
	haveRun := make(map[int]bool)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		if _, ok := haveRun[from]; !ok {
			heads[from] = true
		}

		if _, ok := heads[to]; ok {
			delete(heads, to)
		}

		haveRun[from] = true
		haveRun[to] = true
	}

	result := make([]int, 0)

	for key := range heads {
		result = append(result, key)
	}

	return result
}

func main() {
	result := findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}})
	println(result)
}
