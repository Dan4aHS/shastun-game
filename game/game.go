package game

import (
	"shastun-game/chips"
	"shastun-game/random"
	time_rater "shastun-game/time-rater"
)

type Game struct {
	testChips  *chips.Chips
	smartChips *chips.Chips

	testMoves  int
	smartMoves int
}

func NewGame() *Game {
	return &Game{
		testChips:  chips.NewChips(),
		smartChips: chips.NewChips(),
		testMoves:  0,
		smartMoves: 0,
	}
}

func (g *Game) Play() (int, int) {
	for len(g.testChips.Current()) > 0 || len(g.smartChips.Current()) > 0 {

		cube1, cube2 := random.ThrowCubes()

		if len(g.testChips.Current()) > 0 {
			g.TestMove(cube1, cube2)
			g.testMoves++
		}
		if len(g.smartChips.Current()) > 0 {
			g.SmartMove(cube1, cube2)
			g.smartMoves++
		}

		//s1 := g.testChips.Current()
		//s2 := g.smartChips.Current()
		//
		//sort.Ints(s1)
		//sort.Ints(s2)
		//
		//fmt.Println("AFTER MOVE", s1, s2, cube1, cube2)
		//fmt.Println()
	}

	return g.testMoves, g.smartMoves
}

func (g *Game) TestMove(first, second int) {
	switch {
	case g.testChips.Has(first + second):
		g.testChips.Remove(first + second)
	case g.testChips.HasMulti(first, second):
		g.testChips.RemoveMulti(first, second)

	}
}

func (g *Game) SmartMove(first, second int) {
	switch {
	case g.smartChips.HasMulti(first, second) && g.smartChips.Has(first+second):
		newChipsForSum := chips.NewChipsFromSlice(g.smartChips.Current())
		newChipsForSum.Remove(first + second)
		trForSum := time_rater.NewTimeRater(newChipsForSum)
		rateSum := trForSum.RateTime()

		newChipsForTwo := chips.NewChipsFromSlice(g.smartChips.Current())
		newChipsForTwo.RemoveMulti(first, second)
		trForTwo := time_rater.NewTimeRater(newChipsForTwo)
		rateTwo := trForTwo.RateTime()

		if rateSum < rateTwo {
			g.smartChips.Remove(first + second)
		} else {
			g.smartChips.RemoveMulti(first, second)
		}
	case g.smartChips.Has(first + second):
		g.smartChips.Remove(first + second)
	case g.smartChips.HasMulti(first, second):
		g.smartChips.RemoveMulti(first, second)
	}
}
