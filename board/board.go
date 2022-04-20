package board

import (
	"errors"
	"fmt"
)

// Types used

type Sign int8

type Spot struct {
	Row uint8
	Col uint8
}

type Board [3][3]Sign

type Game struct {
	board      Board
	moveCount  uint8
	conditions [8]int8
}

// Constants

const EMPTY = Sign(0)
const PLAYER_ONE = Sign(1)
const PLAYER_TWO = Sign(-1)

const _ROW = 0
const _COL = 3
const _DIAG = 6
const _ANTIDIAG = 7

// Variables

/* var board BoardType
var moveCount uint8
var conditions [8]int8 */

// Game Contructor

func NewGame() *Game {
	return &Game{
		board: Board{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		moveCount:  0,
		conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// Gets a spot from the board
func (g *Game) GetSpot(spot Spot) Sign {
	return g.board[spot.Row][spot.Col]
}

// Places an X or an O in the spot asked, if it is free

func (g *Game) SetSpot(sign Sign, spot Spot) (Sign, error) {

	// Exit if sign is not valid

	if sign != PLAYER_ONE && sign != PLAYER_TWO && sign != EMPTY {
		fmt.Println("invalid sign")
		return EMPTY, errors.New("invalid sign")
	}

	// Exit if spot is not EMPTY

	if g.GetSpot(spot) != EMPTY {
		fmt.Println("spot taken")
		return EMPTY, errors.New("Spot is already taken")
	}
	g.board[spot.Row][spot.Col] = sign

	// update conditions array and moveCount

	g.moveCount++

	g.conditions[spot.Row+_ROW] += int8(sign)
	g.conditions[spot.Col+_COL] += int8(sign)

	if spot.Row == spot.Col {
		g.conditions[_DIAG] += int8(sign)
	}

	if spot.Row+spot.Col == 2 {
		g.conditions[_ANTIDIAG] += int8(sign)
	}

	// Check if board is winning after current move

	resultSign := EMPTY
	if g.IsWinning() {
		resultSign = sign
	}

	return resultSign, nil
}

// Checks if board is winning using the conditions array

func (g *Game) IsWinning() bool {
	for _, value := range g.conditions {
		if value == 3 || value == -3 {
			return true
		}
	}
	return false
}

// Checks if there is place on the board

func (g *Game) IsFull() bool {
	return g.moveCount == 9
}

func (g *Game) GetBoard() Board {
	result := g.board
	return result
}

// Probably unused

func (g *Game) ResetGame() {
	g.board = [3][3]Sign{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	g.moveCount = 0

	g.conditions = [8]int8{0, 0, 0, 0, 0, 0, 0, 0}
}

// TODO: Move in the Cli display controller

func (g Game) Print() {

	/*fmt.Printf("\n %s | %s | %s \n", g.String(g.board[0][0]), g.String(g.board[0][1]), g.String(g.board[0][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", g.String(g.board[1][0]), g.String(g.board[1][1]), g.String(g.board[1][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n\n", g.String(g.board[2][0]), g.String(g.board[2][1]), g.String(g.board[2][2]))*/

	for _, i := range g.board {
		for _, j := range i {
			fmt.Printf("%d ", j)
		}
		fmt.Println("")
	}

}

// TODO: remove when I have a display controller

func (g *Game) String(sign Sign) string {
	if sign == EMPTY {
		return " "
	} else if sign == PLAYER_ONE {
		return "X"
	} else {
		return "O"
	}
}
