package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var numbers = make([]int, 5)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	MergeSort(numbers, 0, 4)
	if !sort.SliceIsSorted(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	}) {
		t.Log(numbers)
		t.Errorf("Slice is not sorted")
	}

	t.Log(numbers)
}
