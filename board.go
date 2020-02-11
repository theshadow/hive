package hived

import "fmt"

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

type PlacedPiece struct {
	Piece      Piece
	Coordinate Coordinate
}

type Board struct {
	// used to quickly move pieces around in the pieces array.
	locationMap map[Coordinate]int

	// used for rendering the board by allowing the rendering engine to loop over all of the pieces
	pieces []PlacedPiece
}

func (brd Board) Place(pp PlacedPiece) {
	brd.pieces = append(brd.pieces, pp)
	brd.locationMap[pp.Coordinate] = len(brd.pieces) - 1
}
func (brd Board) Move(a, b Coordinate) {
	if idx, ok := brd.locationMap[a]; ok {
		brd.pieces[idx].Coordinate = b
		delete(brd.locationMap, a)
	}
}
func (brd Board) Pieces() []PlacedPiece {
	return brd.pieces
}

