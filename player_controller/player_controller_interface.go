package player_controller

import (
	"tic-tac-go/game/board"
)

// Typedef for PlayerType

type PlayerType uint8

// Definition of player controllers

const CLIPLAYER = PlayerType(0)
const RANDOMPLAYER = PlayerType(1)

// Intergace for a Generic Player Controller, implementable by CLI, GUI or CPU

type PlayerController interface {
	GetMove() board.Spot
	DisplayBoard()
	DisplayError(err string)
	DisplayWin(winner string, thisPlayer bool)
	DisplayDraw()
	PlayAgain() bool
}

func NewPlayerController(playerType PlayerType, boardToInterface *board.Board, p1Char string, p2Char string) PlayerController {
	switch playerType {
	case CLIPLAYER:
		return NewCLIPlayerController(boardToInterface, p1Char, p2Char)
	case RANDOMPLAYER:
		return NewRandomPLayerController()
	default:
		return NewRandomPLayerController()
	}
}
