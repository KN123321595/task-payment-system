package iban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateIban(t *testing.T) {
	iban1 := GenerateIban()
	assert.Equal(t, 28, len([]rune(iban1)))

	iban2 := GenerateIban()
	assert.NotEqual(t, iban1, iban2)
}
