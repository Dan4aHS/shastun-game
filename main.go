package main

import (
	"fmt"
	"shastun-game/chips"
	"shastun-game/game"
	time_rater "shastun-game/time-rater"
)

func main() {
	runCount := 100000

	testSum, smartSum := 0, 0
	testChips := []int{1, 2, 3, 4, 11, 12}
	fmt.Printf("TESTING SET: %v\n", testChips)

	tr := time_rater.NewTimeRater(chips.NewChipsFromSlice(testChips))
	fmt.Printf("TIME_RATER: %f\n\n", tr.RateTime())

	for i := 0; i < runCount; i++ {
		g := game.NewGame(testChips)

		testRes, smartRes := g.Play()
		testSum += testRes
		smartSum += smartRes
	}

	fmt.Printf("test avg: %f\n", float64(testSum)/float64(runCount))
	fmt.Printf("smart avg: %f\n", float64(smartSum)/float64(runCount))
}
