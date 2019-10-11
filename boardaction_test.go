package main

import (
	"strconv"
	"testing"
)

func TestMaxTile(t *testing.T) {

	// Empty board has a max tile of 0
	var emptyBoard board
	emptyMaxTile := emptyBoard.maxTile()
	if emptyMaxTile != 0 {
		t.Errorf("Expected the maximum tile of an empty board to be 0, but got " + strconv.FormatInt(int64(emptyMaxTile), 10))
	}

	// Non-empty board has the correct max tile
	var nonEmptyBoard board
	nonEmptyBoard[0][0] = 3
	nonEmptyBoard[1][3] = 2048
	nonEmptyBoard[2][1] = 3049
	nonEmptyMaxTile := nonEmptyBoard.maxTile()
	if nonEmptyMaxTile != 3049 {
		t.Errorf("Expected the maximum tile of an empty board to be 3049, but got " + strconv.FormatInt(int64(nonEmptyMaxTile), 10))
	}

}

func TestEmptyTiles(t *testing.T) {

	// All tiles are found if board is empty
	var emptyBoard board
	rEmpty, cEmpty := emptyBoard.emptyTiles()

	if len(rEmpty) != 16 || len(cEmpty) != 16 {
		t.Errorf("Expected all tiles for empty board to be identified as empty but some where not")
	}

	// Finds correct empty tiles
	var nonEmptyBoard board
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			nonEmptyBoard[i][j] = 2048
		}
	}

	i1, j1, i2, j2 := 1, 3, 2, 0
	nonEmptyBoard[i1][j1] = 0
	nonEmptyBoard[i2][j2] = 0
	var boardVal float32

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			boardVal = nonEmptyBoard[i][j]
			if i == i1 && j == j1 || i == i2 && j == j2 {
				if boardVal != 0 {
					t.Errorf("Nonzero value of %d when there should be a zero value at [%d,%d]", int64(boardVal), i, j)
				}
			}
		}
	}

}
