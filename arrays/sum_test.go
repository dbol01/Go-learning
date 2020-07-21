package sum

import "testing"

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
