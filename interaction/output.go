package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game.")
	fmt.Println("Good Luck")
}

func PrintRoundData(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster by %v\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL ATTACK" {
		fmt.Printf("Player strong attack by %v\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed by %v\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked player by %v\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func GetUserInput(isSpecial bool) {
	fmt.Println("Choose Your Option")
	fmt.Println("-------------------")
	fmt.Println("1) Monster Attack")
	fmt.Println("2) Heal")

	if isSpecial {
		fmt.Println("3) Special Round")
	}
}
func DeclareWinner(winner string) {
	fmt.Println("-------------------")
	asciiGameOver := figure.NewColorFigure("Game Over!", "", "red", true)
	asciiGameOver.Print()
	fmt.Println("-------------------")
	fmt.Printf("%v Won\n", winner)
}
func WriteLogFile(round *[]RoundData) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Not able to write log")
		return
	}
	exPath = filepath.Dir(exPath)
	file, err := os.Create(exPath + "/gamelog.txt")

	if err != nil {
		fmt.Println("Saving log file failed")
		return
	}
	for index, value := range *round {
		logEntry := map[string]string{
			"Round: ":               fmt.Sprint(index + 1),
			"Action: ":              value.Action,
			"Player Attack Value:":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value:":    fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Value:": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health: ":       fmt.Sprint(value.PlayerHealth),
			"Monster Health: ":      fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing failed.")
			continue
		}
	}
	file.Close()
	fmt.Println("Wrote the data")
}
