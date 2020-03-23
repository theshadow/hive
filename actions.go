package hived

import "fmt"

/*
   | Act    |          Piece           |
   |11111111|11111111|11111111|11111111|

   |                              Transition                               |
   |        Source Coordinate         |          Dst Coordinate            |
   |11111111|11111111|11111111|11111111|11111111|11111111|11111111|11111111|
*/
type Action struct {
	act        uint32
	transition uint64
}

func NewAction(action uint8, p Piece, src Coordinate, dst Coordinate) Action {
	var m Action
	m.act |= uint32(action) << 24
	m.act |= uint32(p) >> 8
	m.transition |= uint64(src) << 32
	m.transition |= uint64(dst)
	return m
}
func (m Action) Piece() Piece {
	return Piece(m.act << 8)
}
func (m Action) Src() Coordinate {
	return Coordinate(m.transition >> 32)
}
func (m Action) Dst() Coordinate {
	return Coordinate(m.transition & DstMask)
}
func (m Action) WasPlaced() bool {
	return m.Act() == Placed
}
func (m Action) WasMoved() bool {
	return m.Act() == Moved
}
func (m Action) Act() uint8 {
	return uint8(uint64(m.act) & ActMask >> 24)
}
func (m Action) ActS() string {
	return actLabels[m.Act()]
}
func (m Action) String() string {
	return fmt.Sprintf("Act: %s, Cell: %s, Src: %s, Dst: %s", m.ActS(), m.Piece(), m.Src(), m.Dst())
}

const (
	Placed uint8 = iota
	Moved

	ActMask = 0b11111111000000000000000000000000
	DstMask = 0b0000000000000000000000000000000011111111111111111111111111111111
)

var actLabels = []string{
	"Placed",
	"Moved",
}
