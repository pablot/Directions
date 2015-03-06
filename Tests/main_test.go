package tests

import (
	"directions/game"
	"testing"
)

func TestGameShouldStart(t *testing.T) {
	game.Game(0)
}

func TestGenerateDirection(t *testing.T) {
  var value string := generateDirection(1)
  if value != ">" {
    t.Error("Expect \">\", but was: " + value)
  }
}
