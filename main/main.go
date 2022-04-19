package main

import (
	"fmt"
	"math/rand"
	"tic-tac-go/game/board"
)

func main() {

	userSign := board.X
	cpuSign := board.O

	playGame(userSign, cpuSign)

}

func playGame(userSign board.Sign, cpuSign board.Sign) {

	fmt.Print("Welcome to the Tic-Tac-Toe game implemented in Go\n\n")

	fmt.Printf("You will be playing as %s\nThe computer will be playing as %s\n\n", board.String(userSign), board.String(cpuSign))

	fmt.Println("You will play by typing the number where you want to place your Sign, then press (ENTER)")
	fmt.Println("Numbers correspond to the following spots on the board:")
	fmt.Print(" 1 | 2 | 3 \n-----------\n 4 | 5 | 6 \n-----------\n 7 | 8 | 9 \n\n")

	fmt.Print("Let's begin\n")

	for !board.IsWinning() {

		if board.IsFull() {
			fmt.Println("Board is full. Draw!")

			if playAgainDialog() {
				fmt.Println("New game!")
				board.ResetGame()
				continue
			} else {
				fmt.Println("Exiting...")
				break
			}
		}

		fmt.Print("Your turn!\n")
		showBoard()

		fmt.Print("Where do you want to place your sign? [1-9]: ")

		userWinning := performMoveWithInput("User", userSign, cpuSign)

		if userWinning != board.EMPTY {
			if playAgainDialog() {
				fmt.Println("New game!")
				board.ResetGame()
				continue
			} else {
				fmt.Println("Exiting...")
				break
			}
		}
		if board.IsFull() {
			fmt.Println("Board is full. Draw!")

			if playAgainDialog() {
				fmt.Println("New game!")
				board.ResetGame()
				continue
			} else {
				fmt.Println("Exiting...")
				break
			}
		}

		fmt.Print("My turn!\n")

		cpuWinning := performMoveWithInput("CPU", userSign, cpuSign)

		if cpuWinning != board.EMPTY {
			if playAgainDialog() {
				fmt.Println("New game!")
				board.ResetGame()
				continue
			} else {
				fmt.Println("Exiting...")
				break
			}
		}
	}
}

func showBoard() {
	currentBoard := board.GetBoard()

	fmt.Printf("\n %s | %s | %s \n", board.String(currentBoard[0][0]), board.String(currentBoard[0][1]), board.String(currentBoard[0][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.String(currentBoard[1][0]), board.String(currentBoard[1][1]), board.String(currentBoard[1][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n\n", board.String(currentBoard[2][0]), board.String(currentBoard[2][1]), board.String(currentBoard[2][2]))
}

func generateCPUMove() uint8 {

	// [0-8]
	position := uint8(rand.Intn(9))
	// fmt.Printf("CPU: Spot %d", position)
	return position
}

func performMoveWithInput(player string, userSign board.Sign, cpuSign board.Sign) board.Sign {
	for {
		if player == "User" {
			var userMove int
			fmt.Scanf("%d", &userMove)
			userMove--

			userWinningSign, userError := board.SetSpot(userSign, board.Spot{Row: uint8(userMove / 3), Col: uint8(userMove % 3)})

			if userWinningSign == userSign {
				fmt.Print("You won!\n")
				return userSign
			} else if userError != nil {
				fmt.Println(userError.Error())
			} else {
				fmt.Printf("You placed a %s in the spot number %d\n\n", board.String(userSign), userMove+1)
				return board.EMPTY
			}
		} else {
			cpuMove := generateCPUMove()

			cpuWinningSign, cpuError := board.SetSpot(cpuSign, board.Spot{Row: uint8(cpuMove / 3), Col: uint8(cpuMove % 3)})
			if cpuWinningSign == cpuSign {
				fmt.Print("The CPU won!\n")
				return cpuSign
			} else if cpuError != nil {
				fmt.Println(cpuError.Error())
			} else {
				fmt.Printf("The CPU placed a %s in the spot number %d\n\n", board.String(cpuSign), cpuMove+1)
				return board.EMPTY
			}
		}
	}
}

func playAgainDialog() bool {
	fmt.Print("Do you want to play again?\n1: Yes\n2: No\nYour choice [1-2]: ")

	var answer int

	fmt.Scanf("%d", &answer)

	return answer == 1
}
