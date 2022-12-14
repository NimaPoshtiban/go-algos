package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	result := InsertionSort(list)
	if !sort.SliceIsSorted(result, func(i, j int) bool {
		return result[i] < result[j]
	}) {
		t.Log(list)
		t.Errorf("Slice is not sorted")
	}
}
