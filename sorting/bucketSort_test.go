package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBucketSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	list := make([]float64, 5)
	for i := 0; i < len(list); i++ {
		list[i] = float64(rand.Intn(100)) + 0.5
	}
	result := BucketSort(list, 3)
	if !sort.SliceIsSorted(result, func(i, j int) bool {
		return result[i] < result[j]
	}) {
		t.Log(result)
		t.Errorf("Slice is not sorted")
	}
}
