package game

import (
	"math"

	. "github.com/theshadow/hived"
)

// Distance will return the Manhattan distance between two coordinates.
// It does not validate if either coordinate is a cell with a valid piece
// as it's mostly here for path algorithms
func distance(a, b Coordinate) int {
	return int(math.Round((math.Abs(float64(a.X()-b.X())) + math.Abs(float64(a.X()-b.X())) +
		math.Abs(float64(a.X()-b.X()))) / 2))
}

func neighbors(c Coordinate) []Coordinate {
	var neighbors []Coordinate
	for _, loc := range NeighborsMatrix {
		neighbors = append(neighbors, c.Add(loc))
	}
	return neighbors
}

func heuristic(a, b Coordinate) int {
	return int(math.Round(math.Abs(float64(a.X()-b.X())) +
		math.Abs(float64(a.Y()-b.Y()))))
}

const (
	beetleMaxDistance  = 1
	ladybugMaxDistance = 3
	pillBugMaxDistance = 1
	queenMaxDistance   = 1
	spiderMaxDistance  = 3
)
