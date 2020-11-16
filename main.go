package main

func main() {

}

func specialPairs(arr []int) [][]int {
	// create a helper map to track the indices where arr elements appear.
	helper := make(map[int][]int)

	// create the slice of slices to keep track of the pairs.
	pairs := make([][]int, 0)

	for i, v := range arr {
		helper[v] = append(helper[v], i)
		if len(helper[v]) > 1 {
			for _, k := range helper[v][:len(helper[v])-1] {
				pairs = append(pairs, []int{k, i})
			}
		}
	}

	return pairs
}

// numPairs returns the number of special pairs.
func numPairs(arr [][]int) int {
	return len(arr)
}
