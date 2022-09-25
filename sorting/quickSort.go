package sorting

// partition the array around the pivot
func partition(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if list[j] <= pivot {
			i++
			temp := list[i]
			list[i] = list[j]
			list[j] = temp
		}
	}
	temp := list[i+1]
	list[i+1] = list[high]
	list[high] = temp

	return i + 1
}

// QuickSort Algorithm
// Sorts the array in O(nlogn)
// Worst case: O(n^2)
func QuickSort(list []int, low, high int) {
	if low < high {
		q := partition(list, low, high)
		QuickSort(list, low, q-1)
		QuickSort(list, q+1, high)
	}
}
