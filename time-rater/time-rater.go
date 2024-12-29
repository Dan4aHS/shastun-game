package time_rater

import "shastun-game/chips"

type TimeRater struct {
	chips *chips.Chips
}

func (r *TimeRater) RateTime() float64 {
	if len(r.chips.Current()) == 0 {
		return 0
	}

	return r.calculateChanceToGetNumber(r.chips.Current()[0])
}

func (r *TimeRater) calculateChanceToGetNumber(n int) float64 {
	var goods float64

	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i+j == n {
				goods++
			}

			if i == j && i == n {
				goods++
			}
		}
	}

	return 36 / goods
}

func NewTimeRater(chips *chips.Chips) *TimeRater {
	return &TimeRater{
		chips: chips,
	}
}
