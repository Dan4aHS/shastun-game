package main

import (
	"fmt"
	"shastun-game/game"
)

func main() {
	runCount := 10000

	testSum, smartSum := 0, 0

	for i := 0; i < runCount; i++ {
		g := game.NewGame()

		testRes, smartRes := g.Play()
		testSum += testRes
		smartSum += smartRes
	}

	fmt.Printf("test avg: %f\n", float64(testSum)/float64(runCount))
	fmt.Printf("smart avg: %f\n", float64(smartSum)/float64(runCount))
}
