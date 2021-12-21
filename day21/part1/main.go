package main

import "fmt"

type Player struct {
	pos   int
	score int
}

var player1 *Player
var player2 *Player
var dieRolls int = 0
var diceValue int = 1

func setUpBoard() {
	//test
	player1 = &Player{pos: 4, score: 0}
	player2 = &Player{pos: 8, score: 0}

	//real
	player1 = &Player{pos: 2, score: 0}
	player2 = &Player{pos: 1, score: 0}
}

func roll() int {
	v := diceValue + diceValue + 1 + diceValue + 2
	diceValue += 3
	if diceValue > 100 {
		diceValue -= 100
	}
	dieRolls += 3
	return v
}

func (p *Player) move(spaces int) {
	p.pos = p.pos + spaces
	for p.pos > 10 {
		p.pos -= 10
	}
	p.score += p.pos
}

func main() {
	setUpBoard()
	currentPlayer := player1
	for player1.score < 1000 && player2.score < 1000 {
		r := roll()
		currentPlayer.move(r)
		fmt.Println(*currentPlayer)

		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}

	loser := player1
	if player2.score < player1.score {
		loser = player2
	}

	fmt.Println("Rolls ", dieRolls, " Loser score ", loser.score, " math ", dieRolls*loser.score)
}
