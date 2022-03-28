package main

import (
	"github/gauthami/monsterslyer/interaction"
)

func main() {

	startGame()

	winner := "" // "Player" || "Monster"||""

	for winner == "" {
		executeRound()
	}

	endGame()
}

//start exextuteround and end are created

func startGame() {

	interaction.PrintGreeting()
}

func executeRound() {

}

func endGame() {}
