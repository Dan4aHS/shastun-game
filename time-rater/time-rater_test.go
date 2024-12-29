package time_rater

import (
	"shastun-game/chips"
	"testing"
)

func TestTimeRaterOnZeroChips(t *testing.T) {
	ch := chips.NewChipsFromSlice([]int{})
	tr := NewTimeRater(ch)

	res := tr.RateTime()

	assertFloat(t, 0, res)
}

func TestTimeRaterOnOnlyOneChipLeft(t *testing.T) {
	tests := []struct {
		chip int
		time float64
	}{
		{1, 36},
		{2, 18},
		{3, 12},
		{4, 9},
		{5, 7.2},
		{6, 6},
		{7, 6},
		{8, 7.2},
		{9, 9},
		{10, 12},
		{11, 18},
		{12, 36},
	}

	for _, tt := range tests {
		ch := chips.NewChipsFromSlice([]int{tt.chip})
		tr := NewTimeRater(ch)

		res := tr.RateTime()
		assertFloat(t, tt.time, res)
	}
}

func TestTimeRaterOnTwoChipsLeft(t *testing.T) {
	tests := []struct {
		chips []int
		time  float64
	}{
		{
			chips: []int{1, 12},
			time:  54,
		},
	}

	for _, tt := range tests {
		ch := chips.NewChipsFromSlice(tt.chips)
		tr := NewTimeRater(ch)

		res := tr.RateTime()
		assertFloat(t, tt.time, res)
	}
}

func assertInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected: %d, actual: %d", expected, actual)
	}
}

func assertFloat(t *testing.T, expected, actual float64) {
	if expected != actual {
		t.Errorf("Expected: %f, actual: %f", expected, actual)
	}
}
