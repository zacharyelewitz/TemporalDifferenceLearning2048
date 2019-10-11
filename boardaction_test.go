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
			} else {
				if boardVal == 0 {
					t.Errorf("Zero value when value should be 2048 at [%d,%d]", i, j)
				}
			}

		}
	}
}

func TestAddTile(t *testing.T) {

	// Adding a tile means there is one more non-zero tile on the board
	var b board
	r, _ := b.emptyTiles()
	nonZeroCount := 0

	for i := 0; i < 16; i++ {
		b.addTile()
		r, _ = b.emptyTiles()
		nonZeroCount++
		if 16-len(r) != nonZeroCount {
			t.Errorf("%d too many tiles added when addTile method is applied", 16-len(r)-nonZeroCount)
		}
	}

	// Added tiles are always a 2 or a 4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] != 2 && b[i][j] != 4 {
				t.Errorf("Tile at [%d,%d] has value %d (not a 2 or a 4)", i, j, int32(b[i][j]))
			}
		}
	}
}

func TestStartingBoard(t *testing.T) {

	// Starting board exists and has two tiles
	b := startingBoard()
	r, _ := b.emptyTiles()
	if len(r) != 14 {
		t.Errorf("Starting board should only have 2 tiles, actually has %d", 16-len(r))
	}

}
