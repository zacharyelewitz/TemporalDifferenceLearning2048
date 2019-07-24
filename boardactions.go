package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Define Unique Types
type board [4][4]float32
type tuple [12][12][12][12][12][12]float32

func (b board) maxTile() float32 {

	var mt float32

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] > mt {
				mt = b[i][j]
			}
		}
	}

	return mt

}

func (b board) print() {

	var s string
	var c string

	// Display board to terminal
	for i := 0; i < 4; i++ {
		s = "| "
		for j := 0; j < 4; j++ {
			c = strconv.Itoa(int(b[i][j]))
			s += c
			for k := 0; k < (5 - len(c)); k++ {
				s += " "
			}
		}
		s += "|"
		fmt.Println(s)
		s = ""
	}
}

func (b board) emptyTiles() ([]int, []int) {

	// Empty tile coordinates
	var r []int
	var c []int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] == 0 {
				r = append(r, i)
				c = append(c, j)
			}
		}
	}

	return r, c
}

func (b *board) addTile() {
	r, c := (*b).emptyTiles()

	if len(r) > 0 {
		source := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(source)

		var r1 int

		switch {
		case len(r) > 1:
			r1 = rnd.Intn(len(r))
		default:
			r1 = 0
		}

		r2 := rnd.Intn(100)

		switch {
		case r2 > 9:
			(*b)[r[r1]][c[r1]] = 2
		default:
			(*b)[r[r1]][c[r1]] = 4
		}
	}
}

func startingBoard() board {
	var b board
	b.addTile()
	b.addTile()
	return b
}

func (b board) rotateCopyQuarter() board {

	var r board
	var t float32
	N := 4

	// Rotates clockwise
	for i := 0; i < N/2; i++ {
		for j := 1; j <= N-i-1; j++ {
			t = b[i][j]
			r[i][j] = b[N-1-j][i]
			r[N-1-j][i] = b[N-1-i][N-1-j]
			r[N-1-i][N-1-j] = b[j][N-1-i]
			r[j][N-1-i] = t
		}
	}

	return r

}

func (b *board) rotateOriginalQuarter() {

	// Rotates clockwise
	(*b) = b.rotateCopyQuarter()

}

func (b board) rotateCopy(n int) board {

	// Rotates clockwise
	r := b
	for i := 0; i < n; i++ {
		r = r.rotateCopyQuarter()
	}

	return r
}

func (b *board) rotateOriginal(n int) {
	// Rotates clockwise
	for i := 0; i < n; i++ {
		(*b).rotateOriginalQuarter()
	}
}

func (b *board) swipe(d int) float32 {

	/*
		0 - Left
		1 - Down
		2 - Right
		3 - Up
	*/

	var tempBoard board

	tempBoard = b.rotateCopy(d)
	newPoints := tempBoard.swipeLeft()
	tempBoard.rotateOriginal(4 - d)

	(*b) = tempBoard

	return newPoints

}

func (b *board) swipeLeft() float32 {

	// New points from swipe to add to score
	var newPoints float32

	for r := 0; r < 4; r++ {

		// tile1 and tile2 are adjacent tiles mod empty spaces
		var tile1 float32
		var i int        //
		var nz []float32 //no-zero values
		for _, j := range b[r] {
			if j != 0 {
				nz = append(nz, j)
			}
		}

		nz = append(nz, -1) //-1 indicates end of row

		// Move through non-zero tiles
		for _, tile2 := range nz {
			switch {

			// End of row
			case tile2 == -1:
				if tile1 != 0 {
					(*b)[r][i] = tile1
					i++
				}

			// If left tile is empty
			case tile1 == 0:
				tile1 = tile2

			// Tiles are identical
			case tile1 == tile2:
				combinedTile := tile1 + tile2
				(*b)[r][i] = combinedTile
				newPoints += combinedTile
				i++
				tile1 = 0

			default:
				(*b)[r][i] = tile1
				tile1 = tile2
				i++

			}

		}

		for k := i; k < 4; k++ {
			(*b)[r][k] = 0
		}
	}

	return newPoints

}

func (b board) anyZeros() bool {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] == 0 {
				return true
			}
		}
	}

	return false

}

func (b board) ableToSlideHorizontally() bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == b[i][j+1] {
				return true
			}
		}
	}

	return false
}

func (b board) ableToSlide() bool {

	if b.ableToSlideHorizontally() {
		return true
	}
	r := b.rotateCopyQuarter()

	if r.ableToSlideHorizontally() {
		return true
	}

	return false
}

func (b board) done() bool {
	var n2048 float32 = 2048
	switch {
	case b.maxTile() == n2048:
		return true
	case b.anyZeros():
		return false
	case b.ableToSlide():
		return false
	default:
		return true
	}
}

func (b board) exponentBoard() board {

	var eb board

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] != 0 {
				eb[i][j] = float32(math.Log2(float64(b[i][j])))
			}
		}
	}

	return eb

}

func (b board) transpose() board {
	var t board
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			t[i][j] = b[j][i]
		}
	}

	return t
}

func (b board) getSets() [][][]float32 {

	var boards []board // slice of board variations
	var s1, s2, s3, s4 []float32
	var set [][]float32
	var sets [][][]float32

	e := b.exponentBoard() // board exponents
	t := e.transpose()     // transpose

	// Determine all board variations
	for i := 0; i < 4; i++ {
		boards = append(boards, e.rotateCopy(i))
		boards = append(boards, t.rotateCopy(i))
	}

	for _, d := range boards {
		s1 = []float32{d[0][0], d[1][0], d[2][0],
			d[0][1], d[1][1], d[2][1]}
		s2 = []float32{d[0][1], d[1][1], d[2][1],
			d[0][2], d[1][2], d[2][2]}
		s3 = []float32{d[0][2], d[1][2], d[2][2],
			d[3][2], d[3][1], d[2][1]}
		s4 = []float32{d[0][3], d[1][3], d[2][3],
			d[3][3], d[3][2], d[2][2]}

		set = [][]float32{s1, s2, s3, s4}
		sets = append(sets, set)
	}

	return sets
}

func setIntoTuple(s []float32, V *tuple) float32 {
	val := (*V)[int(s[0])][int(s[1])][int(s[2])][int(s[3])][int(s[4])][int(s[5])]
	return val
}

func f(state board,
	V1 *tuple,
	V2 *tuple,
	V3 *tuple,
	V4 *tuple) float32 {

	allSets := state.getSets()
	var score float32

	for _, sc := range allSets {

		score += setIntoTuple(sc[0], V1)
		score += setIntoTuple(sc[1], V2)
		score += setIntoTuple(sc[2], V3)
		score += setIntoTuple(sc[3], V4)
	}
	return score
}

func matchingBoards(b1 board, b2 board) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b1[i][j] != b2[i][j] {
				return false
			}
		}
	}

	return true

}

func chooseAction(state board,
	V1 *tuple,
	V2 *tuple,
	V3 *tuple,
	V4 *tuple) int {

	var reward float32
	var nextState board
	var score float32
	var values []float32
	var action int
	var endOfGame bool
	nInf := float32(-math.Pow10(333))
	maxVal := float32(-math.Pow10(333))

	for i := 0; i < 4; i++ {
		nextState = state
		reward = nextState.swipe(i)
		switch {
		case endOfGame:
			score = reward
			values = append(values, score)
		case !matchingBoards(state, nextState):
			score = reward + f(nextState, V1, V2, V3, V4)
			values = append(values, score)
		default:
			values = append(values, nInf)
		}

	}

	for i, v := range values {

		if v > maxVal {
			action = i
			maxVal = v
		}
	}

	return action
}

func updateTuple(V *tuple,
	lr float32,
	nReward float32,
	fnnnState float32,
	fnState float32,
	s []float32) {
	update := lr * (nReward + fnnnState - fnState)
	(*V)[int(s[0])][int(s[1])][int(s[2])][int(s[3])][int(s[4])][int(s[5])] = setIntoTuple(s, V) + update
}

func learn(V1 *tuple,
	V2 *tuple,
	V3 *tuple,
	V4 *tuple,
	nState board,
	nnState board,
	lr float32) {
	nAction := chooseAction(nnState, V1, V2, V3, V4)
	nnnState := nnState
	nReward := nnnState.swipe(nAction)

	fnnnState := f(nnnState, V1, V2, V3, V4)
	fnState := f(nState, V1, V2, V3, V4)
	allSets := nState.getSets()

	for _, sc := range allSets {
		updateTuple(V1, lr, nReward, fnnnState, fnState, sc[0])
		updateTuple(V2, lr, nReward, fnnnState, fnState, sc[1])
		updateTuple(V3, lr, nReward, fnnnState, fnState, sc[2])
		updateTuple(V4, lr, nReward, fnnnState, fnState, sc[3])
	}
}

func updateCount(n256 *float32,
	n512 *float32,
	n1024 *float32,
	n2048 *float32,
	maxTile float32) {

	if maxTile >= 256 {
		(*n256)++
		if maxTile >= 512 {
			(*n512)++
			if maxTile >= 1024 {
				(*n1024)++
				if maxTile >= 2048 {
					(*n2048)++
				}
			}
		}
	}
}

func printStatsResetCountsTime(n256 *float32,
	n512 *float32,
	n1024 *float32,
	n2048 *float32,
	q int,
	start *time.Time,
	gameNum int) {

	t := time.Now()
	elapsed := t.Sub(*start)
	d := float32(q)
	fmt.Printf("Game: %d | 256: %.3f | 512: %.3f | 1024: %.3f | 2048: %.3f | Time: %s\n", gameNum, *n256/d, *n512/d, *n1024/d, *n2048/d, elapsed)

	*n256 = 0
	*n512 = 0
	*n1024 = 0
	*n2048 = 0
	*start = time.Now()

}

func playOneGame(V1 *tuple,
	V2 *tuple,
	V3 *tuple,
	V4 *tuple,
	lr float32) float32 {

	var action int
	var n2048 float32 = 2048
	var nState board

	b := startingBoard()
	for {
		if !b.done() && (b.maxTile() < n2048) {

			// Choose action
			action = chooseAction(b, V1, V2, V3, V4)

			// Swipe
			_ = b.swipe(action)

			if b.done() {
				break
			}

			nState = b
			b.addTile()

			if b.done() {
				break
			}

			learn(V1, V2, V3, V4, nState, b, lr)
		} else {
			break
		}
	}

	return b.maxTile()

}
