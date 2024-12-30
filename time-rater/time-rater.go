package time_rater

import (
	"shastun-game/chips"
	"sync"
)

var TimeRaterCache sync.Map

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

	if v, ok := TimeRaterCache.Load(r.chips.String()); ok {
		return v.(float64)
	}

	res := r.complexCalculateTime()
	TimeRaterCache.Store(r.chips.String(), res)

	return res
}

func (r *TimeRater) complexCalculateTime() float64 {
	//var rateWithNoMove float64
	var rateWithOneMove float64
	var rateWithTwoMoves float64

	var noMovesCounter float64
	for i, j := range DiceIterator() {
		switch {
		case r.chips.HasMulti(i, j) && r.chips.Has(i+j):
			newChips1 := chips.NewChipsFromSlice(r.chips.Current())
			newChips1.Remove(i + j)
			removeSumTimeRate := NewTimeRater(newChips1).RateTime()

			newChips2 := chips.NewChipsFromSlice(r.chips.Current())
			newChips2.RemoveMulti(i, j)
			removeTwoTimeRate := NewTimeRater(newChips2).RateTime()

			if removeSumTimeRate < removeTwoTimeRate {
				rateWithTwoMoves += removeSumTimeRate / 36
			} else {
				rateWithTwoMoves += removeTwoTimeRate / 36
			}

		case r.chips.HasMulti(i, j):
			newChips2 := chips.NewChipsFromSlice(r.chips.Current())
			newChips2.RemoveMulti(i, j)
			removeTwoTimeRate := NewTimeRater(newChips2).RateTime()

			rateWithOneMove += removeTwoTimeRate / 36
		case r.chips.Has(i + j):
			newChips1 := chips.NewChipsFromSlice(r.chips.Current())
			newChips1.Remove(i + j)
			removeSumTimeRate := NewTimeRater(newChips1).RateTime()

			rateWithOneMove += removeSumTimeRate / 36
		default:
			noMovesCounter++
		}
	}

	return (1 + rateWithOneMove + rateWithTwoMoves) / (1 - (noMovesCounter / 36))
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
