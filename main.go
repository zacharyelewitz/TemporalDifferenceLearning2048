package main

import "time"

// Create a n-tuple temporal different learner to play 2048

func main() {

	// Initialize LUTs
	var V1, V2, V3, V4 tuple

	var n256, n512, n1024, n2048 float32
	ngames := 5000000
	printEvery := 1000
	var learningRate float32 = 0.0025
	var maxTile float32

	start := time.Now()

	for gameNum := 1; gameNum <= ngames; gameNum++ {
		maxTile = playOneGame(&V1, &V2, &V3, &V4, learningRate)
		updateCount(&n256, &n512, &n1024, &n2048, maxTile)
		if gameNum%printEvery == 0 {
			printStatsResetCountsTime(&n256, &n512, &n1024, &n2048, printEvery, &start, gameNum)
		}
	}

}
