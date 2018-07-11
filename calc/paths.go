package calc

import (
	"math"
)

const inf = 2147483647

type sequence struct {
	Path        []int
	ReturnValue float64
}

type rate struct {
	Value    float64
	HasValue bool
}

func newRate(v float64) rate { return rate{v, true} }
func newRateNoop() rate      { return rate{0, false} }

func GetSequences(matrix [][]rate) []sequence {
	var sequences []sequence
	paths := getPaths(matrix)

	for _, path := range paths {
		s := sequence{path, returnValue(matrix, path)}
		sequences = append(sequences, s)
	}

	return sequences
}

func returnValue(matrix [][]rate, path []int) float64 {
	value := float64(1)

	for e := 0; e < len(path)-1; e++ {
		i := path[e]
		j := path[e+1]
		value = value * matrix[i][j].Value
	}

	return value
}

func getPaths(m [][]rate) [][]int {
	matrix := toLog(m)
	length := len(matrix)
	pre := make([]int, length)
	dist := make([]float64, length)

	for i := 0; i < length; i++ {
		dist[i] = inf
		pre[i] = -1
	}

	dist[0] = 0

	// relax nodes
	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if isBetter(matrix, dist, i, j) {
					dist[j] = dist[i] + matrix[i][j].Value
					pre[j] = i
				}
			}
		}
	}

	// check for cycles
	cyclic := make([]bool, length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if isBetter(matrix, dist, i, j) {
				dist[j] = dist[i] + matrix[i][j].Value
				cyclic[j] = true
			}
		}
	}

	// calc sequences
	var paths [][]int
	for i := 0; i < len(cyclic); i++ {
		if cyclic[i] != true {
			continue
		}
		visited := make([]bool, length)
		var seq []int
		p := i

		for p != -1 && visited[p] != true {
			seq = append(seq, p)
			visited[p] = true
			p = pre[p]
		}
		seq = append([]int{i}, reverse(seq)...)
		paths = append(paths, seq)
	}
	return paths
}

func isBetter(matrix [][]rate, dist []float64, i, j int) bool {
	if (matrix[i][j].HasValue == true) &&
		(dist[i]+matrix[i][j].Value) < dist[j] {
		return true
	}
	return false
}

func toLog(matrix [][]rate) [][]rate {
	m := make([][]rate, len(matrix))
	for i := 0; i < len(m); i++ {
		m[i] = make([]rate, len(matrix[i]))
		copy(m[i], matrix[i])
		for j := 0; j < len(m); j++ {
			m[i][j].Value = -math.Log(m[i][j].Value)
		}
	}
	return m
}

func reverse(n []int) []int {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}
