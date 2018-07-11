package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSequences(t *testing.T) {
	matrix := [][]rate{
		[]rate{newRate(1), newRate(2), newRate(4)},
		[]rate{newRate(0.5), newRate(1), newRate(3)},
		[]rate{newRate(0.25), newRate(0.5), newRate(1)},
	}
	expected := []sequence{
		sequence{[]int{0, 1, 2, 0}, float64(1.5)},
		sequence{[]int{1, 2, 1}, float64(1.5)},
		sequence{[]int{2, 1, 2}, float64(1.5)},
	}

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}

func TestNoArbitrage(t *testing.T) {
	matrix := [][]rate{
		[]rate{newRate(1), newRate(2), newRate(4)},
		[]rate{newRate(0.5), newRate(1), newRate(2)},
		[]rate{newRate(0.25), newRate(0.5), newRate(1)},
	}
	var expected []sequence

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}

func TestIncompleteGraph(t *testing.T) {
	matrix := [][]rate{
		[]rate{newRate(1), newRate(2), newRate(4)},
		[]rate{newRate(0.5), newRate(1), newRate(3)},
		[]rate{newRate(0.25), newRateNoop(), newRate(1)},
	}
	expected := []sequence{
		sequence{[]int{0, 1, 2, 0}, float64(1.5)},
		sequence{[]int{1, 2, 0, 1}, float64(1.5)},
		sequence{[]int{2, 0, 1, 2}, float64(1.5)},
	}

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}
