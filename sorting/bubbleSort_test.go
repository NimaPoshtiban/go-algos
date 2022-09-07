package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var list = make([]int, 5)

func TestBubbleSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		list[i] = rand.Intn(100)
	}
	BubbleSort(list)
	if !sort.SliceIsSorted(list, func(i, j int) bool {
		return list[i] < list[j]
	}) {
		t.Log(list)
		t.Errorf("Slice is not sorted")
	}
}
