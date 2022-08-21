package main

import "fmt"

func collectPlayerInput() int {
	fmt.Printf("Select an option: ")
	var playerChoice int
	_, err := fmt.Scan(&playerChoice)

	if err != nil {
		fmt.Println(err)
		playerChoice = collectPlayerInput()
		return playerChoice
	}
	return playerChoice
}

func handlePlayerChoice(choice int) {
	switch choice {
	case 0:
		var g Game
		g.create()
		g.play()
		g.playingBoard.print()
		gameOverScreen(g.state)
		mainMenu()
		return
	case 1:
		fmt.Println("Hope we will play again soon.")
		return
	default:
		fmt.Println("Invalid choice")
		input := collectPlayerInput()
		handlePlayerChoice(input)
		return
	}

}

func mainMenu() {
	fmt.Println("TIC TAC GO")

	options := []string{
		"New Game",
		"Quit",
	}

	for i, text := range options {
		fmt.Printf("[%d] %s\n", i, text)
	}

	input := collectPlayerInput()
	handlePlayerChoice(input)
}

func handleResult(result State) {
	var output string
	switch result {
	case player1Win:
		output = "Player 1 won"
	case player2Win:
		output = "Player 2 won"
	case tie:
		output = "It's a tie"
	default:
		output = "Unexpected error has occurred"
	}

	fmt.Println(output)

}
func gameOverScreen(result State) {
	fmt.Println("Game over")
	handleResult(result)
}
