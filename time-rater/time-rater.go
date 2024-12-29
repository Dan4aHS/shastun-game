package time_rater

import "shastun-game/chips"

type TimeRater struct {
	chips *chips.Chips
}

func (r *TimeRater) RateTime() int {
	return 0
}

func NewTimeRater(chips *chips.Chips) *TimeRater {
	return &TimeRater{
		chips: chips,
	}
}
