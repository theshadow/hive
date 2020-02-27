package hived

import "fmt"

/*
    Action |          Cell
   11111111|11111111|11111111|11111111

            Source Coordinate         |          Dst Coordinate
   11111111|11111111|11111111|11111111|11111111|11111111|11111111|11111111
*/
type Move struct {
	action     uint32
	transition uint64
}

func NewMove(action uint8, p Piece, src Coordinate, dst Coordinate) Move {
	var m Move
	m.action |= uint32(action) << 24
	m.action |= uint32(p) >> 8
	m.transition |= uint64(src) << 32
	m.transition |= uint64(dst)
	return m
}
func (m *Move) Set(action uint8, p Piece, src Coordinate, dst Coordinate) {

}
func (m Move) Piece() Piece {
	return Piece(m.action << 8)
}
func (m Move) Src() Coordinate {
	return Coordinate(m.transition >> 32)
}
func (m Move) Dst() Coordinate {
	return Coordinate(m.transition & DstMask)
}
func (m Move) WasPlaced() bool {
	return m.Action() == Placed
}
func (m Move) WasMoved() bool {
	return m.Action() == Moved
}
func (m Move) Action() uint8 {
	return uint8(uint64(m.action) & ActionMask >> 24)
}
func (m Move) ActionS() string {
	return actionLabels[m.Action()]
}
func (m Move) String() string {
	return fmt.Sprintf("Action: %s, Cell: %s, Src: %s, Dst: %s", m.ActionS(), m.Piece(), m.Src(), m.Dst())
}

const (
	Placed uint8 = iota
	Moved

	ActionMask = 0b11111111000000000000000000000000
	DstMask    = 0b0000000000000000000000000000000011111111111111111111111111111111
)

var actionLabels = []string{
	"Placed",
	"Moved",
}
