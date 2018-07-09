package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPaths(t *testing.T) {
	matrix := [][]float64{{1, 2, 4}, {0.5, 1, 3}, {0.25, 0.5, 1}}
	paths := Get(matrix)
	expected := 41
	assert.Equal(t, expected, paths, "they should be equaal!")
}
