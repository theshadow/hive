package hived

import "fmt"

// Board represents a 4D hex grid (x, y, z, height). It works by storing
// the contents of a hex coordinate ("cell") in a slice and using a map
// to quickly reference the memory.
type Board struct {
	// used to quickly look up the piece in the cells
	locationMap map[Coordinate]int

	// maintains the state of the pieces
	cells []cell
}

func NewBoard() *Board {
	return &Board{
		locationMap: make(map[Coordinate]int),
		cells:       []cell{},
	}
}
func (brd *Board) Place(p Piece, c Coordinate) {
	cl := cell{p, c}
	brd.cells = append(brd.cells, cl)
	brd.locationMap[cl.Coordinate] = len(brd.cells) - 1
}
func (brd *Board) Move(a, b Coordinate) error {
	if idx, ok := brd.locationMap[a]; ok {
		// update the coordinate of the piece
		brd.cells[idx].Coordinate = b
		brd.locationMap[b] = idx
		delete(brd.locationMap, a)
		return nil
	}
	return ErrInvalidCoordinate
}
func (brd *Board) Cell(c Coordinate) (Piece, bool) {
	if idx, ok := brd.locationMap[c]; ok {
		return brd.cells[idx].Piece, true
	}
	return ZeroPiece, false
}

// Return an array with seven elements, each element represents one
// of the edges of the piece. We colloquially name these North, Northeast,
// Southeast, South, Southwest, Northwest, and Above respectively. By default it
// is always assumed that the top flat edge is always considered North and
// that the additional edges continue in a clockwise fashion around the piece.
//
// formation := [7]Piece{
//     // North,
//     // Northeast,
//     // Southeast,
//     // South,
//     // Southwest,
//     // Northwest,
//     // Above,
// }
//
// Will return an error when the supplied coordinate isn't a valid location
func (brd *Board) Neighbors(c Coordinate) (formation [7]Piece, err error) {
	if _, ok := brd.Cell(c); !ok {
		return formation, ErrInvalidCoordinate
	}

	for i, loc := range neighborsMatrix {
		loc = c.Add(loc)
		if p, ok := brd.Cell(loc); ok {
			formation[i] = p
		} else {
			formation[i] = ZeroPiece
		}
	}

	return formation, nil
}

func (brd *Board) Pieces() []cell {
	return brd.cells
}

var Origin = Coordinate(0)

var ErrInvalidCoordinate = fmt.Errorf("the specified coordinate is invalid")

var neighborsMatrix = []Coordinate{
	// North
	NewCoordinate(0, 1, -1, 0),
	// Northeast
	NewCoordinate(1, 0, -1, 0),
	// Southeast
	NewCoordinate(1, -1, 0, 0),
	// South
	NewCoordinate(0, -1, 1, 0),
	// Southwest
	NewCoordinate(-1, 0, 1, 0),
	// Northwest
	NewCoordinate(-1, 1, 0, 0),

	// Above
	NewCoordinate(0, 0, 0, 1),
}

// Represents a single cell of the hex grid. It's an internal type and shouldn't
// be used elsewhere beyond the Board type
type cell struct {
	Piece      Piece
	Coordinate Coordinate
}

const (
	North = iota
	Northeast
	Southeast
	South
	Southwest
	Northwest
	Above
)
