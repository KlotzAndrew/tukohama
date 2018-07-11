package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaths(t *testing.T) {
	matrix := [][]float64{
		[]float64{1, 2, 4},
		[]float64{0.5, 1, 3},
		[]float64{0.25, 0.5, 1},
	}
	expected := []sequence{
		sequence{[]int{0, 1, 2, 0}, float64(1.5)},
		sequence{[]int{1, 2, 1}, float64(1.5)},
		sequence{[]int{2, 1, 2}, float64(1.5)},
	}

	paths := GetSequences(matrix)
	assert.Equal(t, expected, paths, "sequences incorrect")
}
