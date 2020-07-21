package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expect := 6

		if got != expect {
			t.Errorf("got %d and wanted %d, given %v", got, expect, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sums numbers given N number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("got %v expected %v", got, expect)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
