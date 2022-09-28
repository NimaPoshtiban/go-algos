package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestRadixSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	sortedList := RadixSort(list)
	if !sort.IntsAreSorted(sortedList) {
		t.Log(list)
		t.Errorf("Slice is not sorted")
	}
	t.Log(list)
}
