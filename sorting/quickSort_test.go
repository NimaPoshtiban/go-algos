package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	QuickSort(list, 0, len(list)-1)
	if !sort.SliceIsSorted(list, func(i, j int) bool {
		return list[i] < list[j]
	}) {
		t.Log(list)
		t.Errorf("Slice is not sorted")
	}
}
