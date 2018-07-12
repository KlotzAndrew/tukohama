package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSequences(t *testing.T) {
	matrix := [][]Rate{
		[]Rate{NewRate(1), NewRate(2), NewRate(4)},
		[]Rate{NewRate(0.5), NewRate(1), NewRate(3)},
		[]Rate{NewRate(0.25), NewRate(0.5), NewRate(1)},
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
	matrix := [][]Rate{
		[]Rate{NewRate(1), NewRate(2), NewRate(4)},
		[]Rate{NewRate(0.5), NewRate(1), NewRate(2)},
		[]Rate{NewRate(0.25), NewRate(0.5), NewRate(1)},
	}
	var expected []sequence

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}

func TestIncompleteGraph(t *testing.T) {
	matrix := [][]Rate{
		[]Rate{NewRate(1), NewRate(2), NewRate(4)},
		[]Rate{NewRate(0.5), NewRate(1), NewRate(3)},
		[]Rate{NewRate(0.25), NewRateNoop(), NewRate(1)},
	}
	expected := []sequence{
		sequence{[]int{0, 1, 2, 0}, float64(1.5)},
		sequence{[]int{1, 2, 0, 1}, float64(1.5)},
		sequence{[]int{2, 0, 1, 2}, float64(1.5)},
	}

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}
