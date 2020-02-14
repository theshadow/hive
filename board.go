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
	return Piece(0), false
}
// TODO: Test this method
func (brd *Board) Neighbors(c Coordinate) ([]Piece, error) {
	if _, ok := brd.Cell(c); !ok {
		return nil, ErrInvalidCoordinate
	}

	var neighbors []Piece

	for _, loc := range neighborsCoordinates {
		loc = c.Add(loc)
		if p, ok := brd.Cell(loc); ok {
			neighbors = append(neighbors, p)
		}
	}

	return neighbors, nil
}

func (brd *Board) Pieces() []cell {
	return brd.cells
}

var ErrInvalidCoordinate = fmt.Errorf("the specified coordinate is invalid")

var neighborsCoordinates = []Coordinate{
	NewCoordinate(0, 1, -1, 0),
	NewCoordinate(1, 0, -1, 0),
	NewCoordinate(1, -1, 0, 0),
	NewCoordinate(0, -1, 1, 0),
	NewCoordinate(-1, 0, 1, 0),
	NewCoordinate(-1, 1, 0, 0),
	NewCoordinate(0, 0, 0, 1),
}

/*
      X        Y        Z       H
  11111111|11111111|11111111|11111111
*/
type Coordinate uint32

func NewCoordinate(x, y, z, h int8) Coordinate {
	var c Coordinate
	(&c).Set(x, y, z, h)

	return c
}
func (c *Coordinate) Set(x, y, z, h int8) {
	*c |= Coordinate(int32(x) << 24)
	*c |= Coordinate(int32(y) << 16)
	*c |= Coordinate(int32(z) << 8)
	*c |= Coordinate(int32(h) << 0)
}
func (c Coordinate) Add(loc Coordinate) Coordinate {
	return NewCoordinate(c.X() + loc.X(), c.Y() + loc.Y(), c.Z() + loc.Z(), c.H() + loc.H())
}
func (c Coordinate) X() int8 {
	return int8(uint32(c) & XMask >> 24)
}
func (c Coordinate) Y() int8 {
	return int8(uint32(c) & YMask >> 16)
}
func (c Coordinate) Z() int8 {
	return int8(uint32(c) & ZMask >> 8)
}
func (c Coordinate) H() int8 {
	return int8(uint32(c) & HMask >> 0)
}
func (c Coordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d, H: %d", c.X(), c.Y(), c.Z(), c.H())
}

const (
	XMask = 0b11111111000000000000000000000000
	YMask = 0b00000000111111110000000000000000
	ZMask = 0b00000000000000001111111100000000
	HMask = 0b00000000000000000000000011111111
)

// Represents a single cell of the hex grid. It's an internal type and shouldn't
// be used elsewhere beyond the Board type
type cell struct {
	Piece      Piece
	Coordinate Coordinate
}
