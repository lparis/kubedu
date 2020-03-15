package hangman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessWrong(t *testing.T) {
	g := NewGame("hello")

	g.Guess('z')

	assert.Equal(t, []rune{'z'}, g.Guesses)
	assert.Equal(t, Active, g.Tally.Status)
	assert.Equal(t, "_____", string(g.Tally.Letters))
	assert.Equal(t, 6, g.Tally.TurnsLeft)
}

func TestGuessGood(t *testing.T) {
	g := NewGame("hello")

	g.Guess('l')
	assert.Equal(t, []rune{'l'}, g.Guesses)
	assert.Equal(t, Active, g.Tally.Status)
	assert.Equal(t, "__ll_", string(g.Tally.Letters))
	assert.Equal(t, 7, g.Tally.TurnsLeft)
}

func TestLoss(t *testing.T) {
	g := NewGame("hello")

	for _, l := range []rune("abcdfgi") {
		g.Guess(l)
	}

	assert.Equal(t, []rune("abcdfgi"), g.Guesses)
	assert.Equal(t, Lost, g.Tally.Status)
	assert.Equal(t, "_____", string(g.Tally.Letters))
	assert.Equal(t, 0, g.Tally.TurnsLeft)
}

func TestWin(t *testing.T) {
	g := NewGame("hello")

	for _, l := range []rune("helo") {
		g.Guess(l)
	}

	assert.Equal(t, []rune("helo"), g.Guesses)
	assert.Equal(t, Won, g.Tally.Status)
	assert.Equal(t, "hello", string(g.Tally.Letters))
	assert.Equal(t, 7, g.Tally.TurnsLeft)
}

func TestAlreadyGuessed(t *testing.T) {
	g := NewGame("hello")

	for _, l := range []rune("hh") {
		g.Guess(l)
	}

	assert.Equal(t, []rune("h"), g.Guesses)
	assert.Equal(t, AlreadyGuessed, g.Tally.Status)
	assert.Equal(t, "h____", string(g.Tally.Letters))
	assert.Equal(t, 7, g.Tally.TurnsLeft)
}

func TestAlreadyWon(t *testing.T) {
	g := NewGame("hello")

	for _, l := range []rune("helox") {
		g.Guess(l)
	}

	assert.Equal(t, []rune("helo"), g.Guesses)
	assert.Equal(t, Won, g.Tally.Status)
	assert.Equal(t, "hello", string(g.Tally.Letters))
	assert.Equal(t, 7, g.Tally.TurnsLeft)
}

func TestAlreadyLost(t *testing.T) {
	g := NewGame("hello")

	for _, l := range []rune("abcdfgij") {
		g.Guess(l)
	}

	assert.Equal(t, []rune("abcdfgi"), g.Guesses)
	assert.Equal(t, Lost, g.Tally.Status)
	assert.Equal(t, "_____", string(g.Tally.Letters))
	assert.Equal(t, 0, g.Tally.TurnsLeft)
}
