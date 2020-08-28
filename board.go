package hived

import (
	"fmt"
)

// TODO how do I marshall this type
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

// Place will return an error if there is already a piece at the specified coordinate.
//
// Should the board know about the game rules?
//
// With this definition the board will return an error if there is already a piece at
// the specified coordinate. This is due to the fact that I'm accepting a trade off.
// While there is merit and value in creating a more robust board that can manage
// multiple pieces at a given location we will loose a valuable and in my opinion
// cheap safety net that will help us validate game rules. Details about the
// two arguments are outlined below.
//
// Hive doesn't allow two pieces to occupy the same coordinate, however, in a
// more generic sense a hex board could, in theory, have multiple pieces in a cell
// based on the game rules. To make a more generic board would require refactoring
// the cells of the board to be able to contain multiple pieces and then devising
// a system for exposing what pieces already exist and letting the game decide
// what to do. However, there are concerns about the increase in complexity for the
// memory management of the type.
//
// By returning an error I make validating game rules drastically more reliable as
// the board will help guarantee the game rules don't do something stupid.
// It also greatly simplifies the memory management of the type. Tracking multiple
// pieces and deciding on a clear set of behavior when you Place or Act a piece
// sounds like a daunting task and really not worth the effort.
//
// So, in conclusion, the value of making cells more robust at this juncture is
// out stripped by the value of having a safety net.
func (brd *Board) Place(p Piece, c Coordinate) error {
	if _, ok := brd.Cell(c); ok {
		return ErrPauliExclusionPrinciple
	}

	cl := cell{p, c}
	brd.cells = append(brd.cells, cl)
	brd.locationMap[cl.Coordinate] = len(brd.cells) - 1

	return nil
}

// Act will accept a source (A) coordinate and a destination (B) coordinate
// and attempt to move the piece to that location. It will return an error
// if a piece doesn't exist at the source or if a piece does exists at the
// destination.
func (brd *Board) Move(a, b Coordinate) error {
	if idx, ok := brd.locationMap[a]; ok {
		if _, ok := brd.locationMap[b]; ok {
			return ErrPauliExclusionPrinciple
		}
		// update the coordinate of the piece
		brd.cells[idx].Coordinate = b
		brd.locationMap[b] = idx
		delete(brd.locationMap, a)
		return nil
	}
	return ErrInvalidCoordinate
}

// Cell will return true when there is a piece at that coordinate
//
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


	for i, loc := range NeighborsMatrix {
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
var ErrPauliExclusionPrinciple = fmt.Errorf("two pieces may not occupy the same coordinate")

var NeighborsMatrix = []Coordinate{
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
// TODO move to internal
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
