package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestCountingSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	result := CountingSort(list)
	if !sort.SliceIsSorted(result, func(i, j int) bool {
		return result[i] < result[j]
	}) {
		t.Log(result)
		t.Errorf("Slice is not sorted")
	}
}

func TestCoutingSortForString(t *testing.T) {
	letters := make([]rune, 5)
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < len(letters); i++ {
		letters[i] = rune(rand.Intn(256))
	}
	result := CountingSortForString(letters)

	if !sort.StringsAreSorted(result) {
		t.Log(result)
		t.Errorf("Slice is not sorted")
	}
}
