package sorting

// Sorts the given array of Integers
// Best Case: O(n)
// Worst Case: O(n^2)
func InsertionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for ; j >= 0 && list[j] > key; j-- {
			list[j] = list[j+1]
		}
		list[j+1] = key
	}
	return list
}
