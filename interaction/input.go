package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttack bool) string {

	for {
		selectedOption, _ := cleanData("Please Select an option")
		if selectedOption == "1" {
			return "ATTACK"
		} else if selectedOption == "2" {
			return "HEAL"
		} else if selectedOption == "3" && specialAttack {
			return "SPECIAL ATTACK"
		}
		fmt.Println("Invalid Entry")

	}

}

func cleanData(inputPrompt string) (string, error) {
	fmt.Println(inputPrompt)
	playerChoice, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	return playerChoice, nil
}
