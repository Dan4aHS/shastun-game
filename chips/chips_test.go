package chips

import "testing"

func TestCurrentChipsWithoutMoves(t *testing.T) {
	c := NewChips()
	res := c.Current()

	assertIntSlices(t, res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func assertIntSlices(t *testing.T, actual, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}

	expectedMap := make(map[int]int)
	for v := range expected {
		expectedMap[v]++
	}

	for v := range actual {
		expectedMap[v]--
	}

	for _, v := range expectedMap {
		if v != 0 {
			t.Errorf("actual %v != expected %v", actual, expected)
		}
	}
}
