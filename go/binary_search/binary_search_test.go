package binary_search_test

import "testing"

func TestBinarySearch_UpperHalf(t *testing.T) {
	arr := []int{1, 7, 23, 42, 75, 116, 794, 800, 1099}

	got := BinarySearch(arr, 800)
	want := 7

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBinarySearch_LowerHalf(t *testing.T) {
	arr := []int{1, 7, 23, 42, 75, 116, 794, 800, 1099}

	got := BinarySearch(arr, 23)
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBinarySearch_MidPoint(t *testing.T) {
	arr := []int{1, 7, 23, 42, 75, 116, 794, 800, 1099}

	got := BinarySearch(arr, 75)
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBinarySearch_OutOfBounds_Upper(t *testing.T) {
	arr := []int{1, 7, 23, 42, 75, 116, 794, 800, 1099}

	got := BinarySearch(arr, 1200)
	want := -1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBinarySearch_OutOfBounds_Lower(t *testing.T) {
	arr := []int{1, 7, 23, 42, 75, 116, 794, 800, 1099}

	got := BinarySearch(arr, -23)
	want := -1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
