package insertion_sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	fmt.Println("Testing Insertion Sort")
	unsorted := []int{23, 7, 42, 9, 69}

	got := InsertionSort(unsorted)

	want := []int{7, 9, 23, 42, 69}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
