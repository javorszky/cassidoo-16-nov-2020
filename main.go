package main

func main() {

}

func specialPairs(arr []int) [][]int {
	// create a helper map to track the indices where arr elements appear.
	helper := make(map[int][]int)
	for i, v := range arr {
		helper[v] = append(helper[v], i)
	}

	// create the slice of slices
	pairs := make([][]int, 0)

	// and then let's go over the helper to put together the pairs.
	for _, v := range helper {
		l := len(v)

		// there's only one appearance of the number, can't make a pair out of it.
		if l == 1 {
			continue
		}

		// assemble the pairs.
		for i := 0; i < l; i++ {
			for j := i+1; j < l; j++ {
				pairs = append(pairs, []int{v[i], v[j]})
			}
		}
	}

	return pairs
}

// numPairs returns the number of special pairs.
func numPairs(arr [][]int) int {
	return len(arr)
}
