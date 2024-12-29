package time_rater

import (
	"shastun-game/chips"
	"testing"
)

func TestTimeRaterOnZeroChips(t *testing.T) {
	ch := chips.NewChips()
	ch.RemoveMulti(1, 2)
	ch.RemoveMulti(3, 4)
	ch.RemoveMulti(5, 6)
	ch.RemoveMulti(7, 8)
	ch.RemoveMulti(9, 10)
	ch.RemoveMulti(11, 12)
	tr := NewTimeRater(ch)

	res := tr.RateTime()

	assertInt(t, 0, res)
}

func TestTimeRaterOnOnlyTwelveLeft(t *testing.T) {
	ch := chips.NewChips()
	ch.RemoveMulti(1, 2)
	ch.RemoveMulti(3, 4)
	ch.RemoveMulti(5, 6)
	ch.RemoveMulti(7, 8)
	ch.RemoveMulti(9, 10)
	ch.Remove(11)
	tr := NewTimeRater(ch)

	res := tr.RateTime()
	assertInt(t, 0, res)
}

func assertInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected: %d, actual: %d", expected, actual)
	}
}
