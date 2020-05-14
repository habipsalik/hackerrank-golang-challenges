package main

import (
	"errors"
	"sort"
	"testing"
)

func minimumIndex(seq []int) (int, error) {
	if len(seq) == 0 {
		return -1, errors.New("Cannot get the minimum value index from an empty sequence")
	}
	var min int
	for i := range seq {
		if seq[i] < seq[min] {
			min = i
		}
	}
	return min, nil
}
func dataEmptyGetArray() []int {
	var seq []int
	return seq
}

func dataUniqueValuesGetArray() []int {
	seq := []int{1, 2, 3}
	return seq
}

func dataUniqueValuesGetExpectedResult() int {
	seq := dataUniqueValuesGetArray()
	min, _ := minimumIndex(seq)
	return min
}

func dataExactlyTwoDifferentMinimumsGetArray() []int {
	seq := []int{1, 1, 2, 3}
	return seq
}

func dataExactlyTwoDifferentMinimumsGetExpectedResult() int {
	seq := dataExactlyTwoDifferentMinimumsGetArray()
	min, _ := minimumIndex(seq)
	return min
}

func TestWithEmptyArray(t *testing.T) {
	seq := dataEmptyGetArray()
	result, _ := minimumIndex(seq)
	if result != -1 {
		t.Fail()
	}
}

func TestWithUniqueValues(t *testing.T) {
	seq := dataUniqueValuesGetArray()
	if len(seq) < 2 {
		t.Fail()
	}

	expectedResult := dataUniqueValuesGetExpectedResult()
	result, _ := minimumIndex(seq)
	if result != expectedResult {
		t.Fail()
	}
}

func TestWithExactyTwoDifferentMinimums(t *testing.T) {
	seq := dataExactlyTwoDifferentMinimumsGetArray()
	if len(seq) < 2 {
		t.Fail()
	}

	sort.Sort(sort.IntSlice(seq))
	if seq[0] != seq[1] && (len(seq) != 2 || seq[1] >= seq[2]) {
		t.Fail()
	}

	expectedResult := dataExactlyTwoDifferentMinimumsGetExpectedResult()
	result, _ := minimumIndex(seq)
	if result != expectedResult {
		t.Fail()
	}
}
