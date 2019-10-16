package main

import "time"

// Create an n-tuple temporal difference learner to play 2048

func main() {

	// Initialize LUTs
	var V1, V2, V3, V4 tuple

	// Count of games that reach a specific tile
	var n256, n512, n1024, n2048 float32
	ngames := 5000000
	printEvery := 10000
	var learningRate float32 = 0.0025
	var maxTile float32

	start := time.Now()

	for gameNum := 1; gameNum <= ngames; gameNum++ {

		// Maximum Tile from playing a single game
		maxTile = playOneGame(&V1, &V2, &V3, &V4, learningRate)

		// Log the maximum tile
		updateCount(&n256, &n512, &n1024, &n2048, maxTile)

		// Show maximum tile statistics over the past 'printEvery' games
		if gameNum%printEvery == 0 {
			printStatsResetCountsTime(&n256, &n512, &n1024, &n2048, printEvery, &start, gameNum)
		}
	}

}
