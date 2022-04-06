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
		winner = executeRound()
	}

	endGame(winner)
}

var countRounds = 0
var gameRound = []interaction.RoundData{}

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

	var playerAttackDmg int
	var monsterAttackDmg int
	var PlayerHealValue int

	if userChoice == "ATTACK" {
		playerAttackDmg = action.AttackMoster(false)
	} else if userChoice == "HEAL" {
		PlayerHealValue = action.PlayerHeal()
	} else if userChoice == "SPECIAL ATTACK" {
		playerAttackDmg = action.AttackMoster(true)
	}
	monsterAttackDmg = action.AttackPlayer()
	playerHealth, monsterHealth := action.GetHealthAmount()

	newData := interaction.RoundData{
		Action:           userChoice,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  PlayerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}
	interaction.PrintRoundData(&newData)
	gameRound = append(gameRound, newData)
	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRound)
}
