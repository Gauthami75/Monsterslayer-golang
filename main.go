package main

import (
	"fmt"
	"github/gauthami/monsterslyer/action"
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

var countRounds = 0

//start exextuteround and end are created

func startGame() {

	interaction.PrintGreeting()
}

func executeRound() string {
	countRounds++
	isSpecialRound := countRounds%3 == 0
	interaction.GetUserInput(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	fmt.Println(userChoice)
	if userChoice == "ATTACK" {
		action.AttackMoster(false)
	} else if userChoice == "HEAL" {
		action.PlayerHeal()
	} else if userChoice == "SPECIAL ATTACK" {
		action.AttackMoster(true)
	}
	action.AttackPlayer()
	return ""
}

func endGame() {}
