package main

type Game struct {
	playingBoard  Board
	player1       Player
	player2       Player
	currentPlayer Player
	isOver        bool
}

func (g *Game) create() {
	g.playingBoard.create()
	g.player1.create(cross)
	g.player2.create(circle)
	g.currentPlayer = g.player1
	g.isOver = false
}

func (g *Game) switchPlayers() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	}

	g.currentPlayer = g.player1
}

func (g *Game) play() {
	const maximumRounds = cellsPerSide * cellsPerSide
	var roundCounter = 0

	for roundCounter < maximumRounds || !g.isOver {
		// TODO: Implement game logic

		roundCounter++
	}

}
