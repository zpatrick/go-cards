package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFace(t *testing.T) {
	cases := map[Rank]bool{
		Two:   false,
		Three: false,
		Four:  false,
		Five:  false,
		Six:   false,
		Seven: false,
		Eight: false,
		Nine:  false,
		Ten:   false,
		Jack:  true,
		Queen: true,
		King:  true,
		Ace:   true,
	}

	for rank, expected := range cases {
		t.Run(string(rank), func(t *testing.T) {
			assert.Equal(t, expected, rank.IsFace())
		})
	}
}

func TestStandardAceHigh(t *testing.T) {
	expected := map[Rank]int{
		Two:   2,
		Three: 3,
		Four:  4,
		Five:  5,
		Six:   6,
		Seven: 7,
		Eight: 8,
		Nine:  9,
		Ten:   10,
		Jack:  11,
		Queen: 12,
		King:  13,
		Ace:   14,
	}

	for rank, value := range expected {
		assert.Equal(t, value, StandardAceHigh("", rank))
	}
}

func TestStandardAceLow(t *testing.T) {
	expected := map[Rank]int{
		Ace:   1,
		Two:   2,
		Three: 3,
		Four:  4,
		Five:  5,
		Six:   6,
		Seven: 7,
		Eight: 8,
		Nine:  9,
		Ten:   10,
		Jack:  11,
		Queen: 12,
		King:  13,
	}

	for rank, value := range expected {
		assert.Equal(t, value, StandardAceLow("", rank))
	}
}
