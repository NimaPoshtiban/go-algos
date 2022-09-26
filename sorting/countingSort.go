package sorting

// find the lowest number in the array
func minOf(vars ...int) (min int) {
	min = vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return
}

// find the largest number in the array
func maxOf(vars ...int) (max int) {
	max = vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return
}

// CountingSort
// Sorts the given numbers in O(n) using Counting Sort Algorithm
func CountingSort(numbers []int) []int {

	max := maxOf(numbers...)
	min := minOf(numbers...)
	_range := max - min + 1
	count := make([]int, _range)
	output := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		count[numbers[i]-min]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		output[count[numbers[i]-min]-1] = numbers[i]
		count[numbers[i]-min]--
	}
	return output
}

// CountingSortForString
// Sort the array of characters in O(n) using Counting Sort Algorithm
func CountingSortForString(arr []rune) []string {
	length := len(arr)
	count := make([]int, 256)
	output := make([]string, length)

	for i := 0; i < 256; i++ {
		count[i] = 0
	}

	for i := 0; i < length; i++ {
		count[arr[i]]++
	}

	for i := 1; i <= 255; i++ {
		count[i] += count[i-1]
	}

	for i := length - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = string(arr[i])
		count[arr[i]]--
	}
	return output
}
