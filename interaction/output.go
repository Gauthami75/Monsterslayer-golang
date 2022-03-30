package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game.")
	fmt.Println("Good Luck")
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
