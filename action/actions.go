package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentPlayerHealth = PLAYER_HEALTH
var currentMonsterHealth = PLAYER_HEALTH

func AttackMoster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_VALUE
	maxAttackValue := PLAYER_ATTACK_MAX_VALUE
	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_VALUE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_VALUE
	}

	attackValue := generateRandomNumber(minAttackValue, maxAttackValue)

	currentMonsterHealth -= attackValue

	return attackValue
}
func PlayerHeal() int {
	healValue := generateRandomNumber(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}
	//return currentPlayerHealth
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_VALUE
	maxAttackValue := MONSTER_ATTACK_MAX_VALUE

	attackValue := generateRandomNumber(minAttackValue, maxAttackValue)

	currentPlayerHealth -= attackValue

	return attackValue
}
func GetHealthAmount() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}
func generateRandomNumber(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
