package sorting

// merge subarrays into one array
func merge(numbers []int, p, q, r int) {
	nl := q - p + 1
	nr := r - q
	L := make([]int, nl)
	R := make([]int, nr)
	for i := 0; i < nl; i++ {
		L[i] = numbers[p+i]
	}
	for i := 0; i < nr; i++ {
		R[i] = numbers[q+i+1]
	}
	i, j := 0, 0
	k := p
	for i < nl && j < nr {
		if L[i] <= R[j] {
			numbers[k] = L[i]
			i++
		} else {
			numbers[k] = R[j]
			j++
		}
		k++
	}

	for ;i < nl;i++ {
		numbers[k] = L[i]
		k++
	}
	for ;j < nr;j++ {
		numbers[k] = R[j]
		k++
	}
}
// Merge Sort Algorithm
// Sorts the array in O(nlgn)
func MergeSort(list []int, p, r int) {
	if p >= r {
		return
	}
	q := p + (r-p) / 2
	MergeSort(list, p, q)
	MergeSort(list, q+1, r)
	merge(list, p, q, r)
}