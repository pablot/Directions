package game

import "testing"

func TestGameShouldStart(t *testing.T) {
	Game(0)
}

func TestGenerateDirectionLeft(t *testing.T) {
	var value, _ = generateDirection(0)
	if value != "<" {
		t.Error("Expect \"<\", but was " + value)
	}
}

func TestGenerateDirectionRight(t *testing.T) {
	var value, _ = generateDirection(1)
	if value != ">" {
		t.Error("Expect \">\", but was " + value)
	}
}

func TestGenerateDirectionUp(t *testing.T) {
	var value, _ = generateDirection(2)
	if value != "^" {
		t.Error("Expect \"^\", but was " + value)
	}
}

func TestGenerateDirectionDown(t *testing.T) {
	var value, _ = generateDirection(3)
	if value != "v" {
		t.Error("Expect \"v\", but was " + value)
	}
}

func TestGenerateDirectionShoudError(t *testing.T) {
	var _, err = generateDirection(4)
	if err == nil {
		t.Error("Expect err, but error was nil")
	}
}

func Test_game_ShouldReturnTrueIfUserWin(t *testing.T) {
	genDirectionsStub := func(quant int) string { return "<<<" }
	getRspGameStub := func() string { return "<<<" }
	countingToXStub := func(quant int) {}
	if game(3, genDirectionsStub, getRspGameStub, countingToXStub) == false {
		t.Error("Expect true but was false")
	}
}

func Test_game_ShouldReturnFalseIfUserLose(t *testing.T) {
	genDirectionsStub := func(quant int) string { return "<<<" }
	getRspGameStub := func() string { return ">>>" }
	countingToXStub := func(quant int) {}
	if game(3, genDirectionsStub, getRspGameStub, countingToXStub) {
		t.Error("Expect false but was true")
	}
}
