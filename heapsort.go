package algorithms

func heapify(array []int, N, i int, cmp func(i, j int) bool) {
	candidate := i

	l, r := candidate*2+1, candidate*2+2

	if l < N && cmp(array[candidate], array[l]) {
		candidate = l
	}

	if r < N && cmp(array[candidate], array[r]) {
		candidate = r
	}

	if candidate != i {
		array[candidate], array[i] = array[i], array[candidate]
		heapify(array, N, candidate, cmp)
	}
}

func heapSort(array []int, cmp func(i, j int) bool) []int {
	N := len(array)
	for i := int(N/2) - 1; i >= 0; i-- {
		heapify(array, N, i, cmp)
	}

	for i := N - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0, cmp)
	}
	return array
}

// HeapSortInt takes an integer array and returns a new sorted copy of the array.
// The order of sorting is determined by the `reverse` boolean value:
// - If reverse is false, the array is sorted in ascending order
// - If reverse is true, the array is sorted in descending order
// The original array remains unmodified.
func HeapSortInt(array []int, reverse bool) []int {
	var arrayCopy = make([]int, len(array))
	copy(arrayCopy, array)
	var cmpInt = func(i, j int) bool {
		if !reverse {
			return i < j
		}
		return i > j
	}
	return heapSort(arrayCopy, cmpInt)
}
