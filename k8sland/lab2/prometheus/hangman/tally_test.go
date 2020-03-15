package hangman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLettersMiss(t *testing.T) {
	res := updateLetters([]rune("hello"), []rune("x"))

	assert.Equal(t, "_____", string(res))
}

func TestUpdateLettersHit(t *testing.T) {
	res := updateLetters([]rune("hello"), []rune("l"))

	assert.Equal(t, "__ll_", string(res))
}
