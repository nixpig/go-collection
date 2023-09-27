package insertion_sort

import (
	"bytes"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	unsorted := []byte{23, 7, 42, 9, 69}

	got := InsertionSort(unsorted)

	want := []byte{7, 9, 23, 42, 69}

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
