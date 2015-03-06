package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameShouldStart(t *testing.T) {
	Game(0)
}

func TestGenerateDirectionLeft(t *testing.T) {
	var value, _ = generateDirection(0)

	assert.Equal(t, "<", value, "they should be equal")
}

func TestGenerateDirectionRight(t *testing.T) {
	var value, _ = generateDirection(1)

	assert.Equal(t, ">", value, "they should be equal")
}

func TestGenerateDirectionUp(t *testing.T) {
	var value, _ = generateDirection(2)

	assert.Equal(t, "^", value, "they should be equal")
}

func TestGenerateDirectionDown(t *testing.T) {
	var value, _ = generateDirection(3)

	assert.Equal(t, "v", value, "they should be equal")
}

func TestGenerateDirectionShoudError(t *testing.T) {
	var _, err = generateDirection(4)

	assert.Nil(t, err)
}

func Test_game_ShouldReturnTrueIfUserWin(t *testing.T) {
	genDirectionsStub := func(quant int) string { return "<<<" }
	getRspGameStub := func() string { return "<<<" }
	countingToXStub := func(quant int) {}

	assert.False(t, game(3, genDirectionsStub, getRspGameStub, countingToXStub))
}

func Test_game_ShouldReturnFalseIfUserLose(t *testing.T) {
	genDirectionsStub := func(quant int) string { return "<<<" }
	getRspGameStub := func() string { return ">>>" }
	countingToXStub := func(quant int) {}

	assert.True(t, game(3, genDirectionsStub, getRspGameStub, countingToXStub))
}
