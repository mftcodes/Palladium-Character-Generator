package roller

import (
	"testing"
)

// Test that random number is between or including values specified.
func TestRandom(t *testing.T) {
	min := 1
	max := 2
	rando := random(min, max)
	if !(rando >= min || rando <= min) {
		t.Errorf(`Not good: %d is not between %d & %d`, rando, min, max)
	}
}

// Test roll is between known possible values for human.
func TestRollAttributesHuman(t *testing.T) {
	isHuman := true
	sides := 6
	numDice := 3
	bonus := 0

	min := 3
	max := 30

	roll := RollAttributes(isHuman, sides, numDice, bonus)
	if !(roll >= min || roll <= min) {
		t.Errorf(`Bad Human: %d is not between %d & %d`, roll, min, max)
	}
}

// Test roll is between known possible values for non-human.
func TestRollAttributesNonHuman(t *testing.T) {
	isHuman := true
	sides := 6
	numDice := 2
	bonus := 0

	min := 2
	max := 18

	roll := RollAttributes(isHuman, sides, numDice, bonus)
	if !(roll >= min || roll <= min) {
		t.Errorf(`Bad Human: %d is not between %d & %d`, roll, min, max)
	}
}
