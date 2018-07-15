package calc

import (
	"math"
	"strconv"
)

const inf = 2147483647

type Sequence struct {
	Path        []int
	ReturnValue float64
}

type Rate struct {
	Value    float64
	HasValue bool
}

func NewRate(v float64) Rate { return Rate{v, true} }
func NewRateNoop() Rate      { return Rate{0, false} }

func GetSequences(matrix [][]Rate) []Sequence {
	var sequences []Sequence
	paths := getPaths(matrix)

	for _, path := range paths {
		s := Sequence{path, returnValue(matrix, path)}
		sequences = append(sequences, s)
	}

	return sequences
}

func returnValue(matrix [][]Rate, path []int) float64 {
	value := float64(1)

	for e := 0; e < len(path)-1; e++ {
		i := path[e]
		j := path[e+1]
		value = value * matrix[i][j].Value
	}

	return value
}

func getPaths(m [][]Rate) [][]int {
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
	paths := cyclicPaths(cyclic, pre)
	return paths
}

func cyclicPaths(cyclic []bool, pre []int) [][]int {
	var paths [][]int
	pathPermutations := map[string]bool{}

	for i := 0; i < len(cyclic); i++ {
		if cyclic[i] != true {
			continue
		}
		visited := make([]bool, len(cyclic))
		var seq []int
		p := i

		for p != -1 && visited[p] != true {
			seq = append(seq, p)
			visited[p] = true
			p = pre[p]
			if p == i {
				seq = append([]int{i}, reverse(seq)...)
				canAdd := canAddNewSeq(seq, pathPermutations)
				if canAdd {
					pathPermutations = addPerm(seq, pathPermutations)
					paths = append(paths, seq)
				}
				continue
			}
		}
	}
	return paths
}

// NOTE: 2-1-2 is the same as 1-2-1
func canAddNewSeq(seq []int, permuations map[string]bool) bool {
	name := seqName(seq[0 : len(seq)-1])
	return !permuations[name]
}

func addPerm(seq []int, permuations map[string]bool) map[string]bool {
	path2x := append(seq, seq[1:]...)
	endIndex := len(seq) - 1
	for i := 0; i < len(seq)-1; i++ {
		key := seqName(path2x[i:endIndex])
		permuations[key] = true
		endIndex += 1
	}
	return permuations
}

func seqName(seq []int) string {
	var name string
	for _, i := range seq {
		name = name + strconv.Itoa(i) + "-"
	}
	return name
}

func isBetter(matrix [][]Rate, dist []float64, i, j int) bool {
	if (matrix[i][j].HasValue == true) &&
		(dist[i]+matrix[i][j].Value) < dist[j] {
		return true
	}
	return false
}

func toLog(matrix [][]Rate) [][]Rate {
	m := make([][]Rate, len(matrix))
	for i := 0; i < len(m); i++ {
		m[i] = make([]Rate, len(matrix[i]))
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
