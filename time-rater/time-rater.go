package time_rater

import (
	"shastun-game/chips"
)

type TimeRater struct {
	chips *chips.Chips
}

func (r *TimeRater) RateTime() float64 {
	if len(r.chips.Current()) == 0 {
		return 0
	}

	if len(r.chips.Current()) == 1 {
		return r.calculateChanceToGetNumber(r.chips.Current()[0])
	}

	return r.complexCalculateTime()
}

func (r *TimeRater) complexCalculateTime() float64 {
	return 5
}

func (r *TimeRater) calculateChanceToGetNumber(n int) float64 {
	var goods float64

	for i, j := range DiceIterator() {
		if (i+j == n) || (i == j && i == n) {
			goods++
		}
	}

	return 36 / goods
}

func NewTimeRater(chips *chips.Chips) *TimeRater {
	return &TimeRater{
		chips: chips,
	}
}
