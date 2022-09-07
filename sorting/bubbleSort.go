package sorting

// Bubble Sort Algorithm
// Sorts the array in O(n^2)
func BubbleSort(items []int) {
	for i := 0; i < len(items); i++ {
		for j := len(items) - 1; j > 0; j-- {
			if items[j] < items[j-1] {
				temp := items[j]
				items[j] = items[j-1]
				items[j-1] = temp
			}
		}
	}
}
