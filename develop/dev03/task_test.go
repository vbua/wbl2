package main

import (
	"reflect"
	"testing"
)

func TestReverseSliceOfStrings(t *testing.T) {
	got := []string{"cccc", "zzzz", "wwwww", "aaaaa"}
	want := []string{"aaaaa", "wwwww", "zzzz", "cccc"}
	if !reflect.DeepEqual(reverseSliceOfStrings(got), want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestRemoveDuplicateStrFromSlice(t *testing.T) {
	got := []string{"cccc", "zzzz", "zzzz", "wwwww", "aaaaa"}
	want := []string{"cccc", "zzzz", "wwwww", "aaaaa"}
	if !reflect.DeepEqual(removeDuplicateStrFromSlice(got), want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestSortStrings(t *testing.T) {
	got := []string{"cccc", "zzzz", "zzzz", "wwwww", "aaaaa"}
	want := []string{"aaaaa", "cccc", "wwwww", "zzzz", "zzzz"}
	k = 1
	if !reflect.DeepEqual(sortStrings(got), want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestSortNum(t *testing.T) {
	got := []string{"7cccc", "2zzzz", "1zzzz", "5wwwww", "3aaaaa"}
	want := []string{"1zzzz", "2zzzz", "3aaaaa", "5wwwww", "7cccc"}
	if !reflect.DeepEqual(SortNum(got), want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
