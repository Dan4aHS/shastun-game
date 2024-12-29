package chips

import "testing"

func TestCurrentWithoutMoves(t *testing.T) {
	c := NewChips()
	res := c.Current()

	assertIntSlices(t, res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func TestNewChipsFromSlice(t *testing.T) {
	c := NewChipsFromSlice([]int{1, 2})
	res := c.Current()

	assertIntSlices(t, res, []int{1, 2})
}

func TestRemoveOne(t *testing.T) {
	c := NewChips()

	c.Remove(3)

	assertIntSlices(t, c.Current(), []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func TestRemoveTwo(t *testing.T) {
	c := NewChips()
	c.RemoveMulti(3, 8)
	assertIntSlices(t, c.Current(), []int{1, 2, 4, 5, 6, 7, 9, 10, 11, 12})
}

func TestHasWithoutRemoves(t *testing.T) {
	c := NewChips()
	hv := c.Has(3)
	assertBool(t, hv, true)
}

func TestHasAfterRemove(t *testing.T) {
	c := NewChips()
	c.Remove(4)
	hv := c.Has(4)
	assertBool(t, hv, false)
}

func TestHasMultiple(t *testing.T) {
	c := NewChips()
	hv := c.Has(3)
	assertBool(t, hv, true)

	c.RemoveMulti(3, 8)
	hv = c.HasMulti(3, 8)
	assertBool(t, hv, false)
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

func assertBool(t *testing.T, actual, expected bool) {
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
