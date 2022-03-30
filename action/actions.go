package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixMicro())
var randGenerator = rand.New(randSource)
var currentPlayerHealth = PLAYER_HEALTH
var currentMonsterHealth = PLAYER_HEALTH

func AttackMoster(isSpecialAttack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_VALUE
	maxAttackValue := PLAYER_ATTACK_MAX_VALUE
	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_VALUE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_VALUE
	}

	attackValue := generateRandomNumber(minAttackValue, maxAttackValue)

	currentMonsterHealth -= attackValue

}
func PlayerHeal() {
	healValue := generateRandomNumber(PLAYER_HEAL_MIN_VALUE, PLAYER_ATTACK_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttackValue := MONSTER_ATTACK_MIN_VALUE
	maxAttackValue := MONSTER_ATTACK_MAX_VALUE

	attackValue := generateRandomNumber(minAttackValue, maxAttackValue)

	currentPlayerHealth -= attackValue
}

func generateRandomNumber(min int, max int) int {
	return randGenerator.Intn(max-min) + max
}
