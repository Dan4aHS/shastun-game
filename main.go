package main

import (
	"fmt"
	"shastun-game/chips"
	"shastun-game/game"
	time_rater "shastun-game/time-rater"
	"sync"
	"time"
)

const (
	runnersCount  = 50000
	runChunkCount = 200
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed: %v\n", time.Since(start))
	}()

	totalGames := runChunkCount * runnersCount
	fmt.Printf("Total Games: %.2e\n", float64(totalGames))

	testSum, smartSum := 0, 0
	testChips := []int{1, 2, 3, 4, 11, 12}
	fmt.Printf("TESTING SET: %v\n", testChips)

	tr := time_rater.NewTimeRater(chips.NewChipsFromSlice(testChips))
	fmt.Printf("TIME_RATER: %f\n\n", tr.RateTime())

	partialTestCh := make(chan int, runnersCount)
	partialSmartCh := make(chan int, runnersCount)

	var wg sync.WaitGroup

	for range runnersCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runAsyncGames(testChips, runChunkCount, partialTestCh, partialSmartCh)
		}()
	}

	go func() {
		wg.Wait()
		close(partialTestCh)
		close(partialSmartCh)
	}()

	for testPartial := range partialTestCh {
		testSum += testPartial
	}

	for smartPartial := range partialSmartCh {
		smartSum += smartPartial
	}

	fmt.Printf("test avg: %f\n", float64(testSum)/float64(totalGames))
	fmt.Printf("smart avg: %f\n", float64(smartSum)/float64(totalGames))
}

func runAsyncGames(chips []int, count int, testCh, smartCh chan int) {
	localTestSum, localSmartSum := 0, 0
	for i := 0; i < count; i++ {
		g := game.NewGame(chips)
		testRes, smartRes := g.Play()
		localTestSum += testRes
		localSmartSum += smartRes
	}

	testCh <- localTestSum
	smartCh <- localSmartSum
}
