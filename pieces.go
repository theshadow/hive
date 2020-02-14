package hived

import "fmt"

/*
    Color |  Bug   | Cell  | Unused
  11111111|11111111|11111111|11111111
*/
type Piece uint32
func NewPiece(color, bug, piece uint8) Piece {
	var p Piece
	(&p).Set(color, bug, piece)
	return p
}
func (p *Piece) Set(color, bug, piece uint8) {
	*p |= Piece(uint32(color) << 24)
	*p |= Piece(uint32(bug) << 16)
	*p |= Piece(uint32(piece) << 8)
}
func (p Piece) IsWhite() bool {
	return p.Color() == WhiteColor
}
func (p Piece) IsBlack() bool {
	return p.Color() == BlackColor
}
func (p Piece) Color() uint8 {
	return uint8(uint32(p)&ColorMask>>24)
}
func (p Piece) ColorS() string {
	return colorLabels[p.Color()]
}
func (p Piece) IsQueen() bool {
	return p.Bug() == Queen
}
func (p Piece) IsBeetle() bool {
	return p.Bug() == Beetle
}
func (p Piece) IsGrasshopper() bool {
	return p.Bug() == Grasshopper
}
func (p Piece) IsSpider() bool {
	return p.Bug() == Spider
}
func (p Piece) IsAnt() bool {
	return p.Bug() == Ant
}
func (p Piece) IsMosquito() bool {
	return p.Bug() == Mosquito
}
func (p Piece) IsLadybug() bool {
	return p.Bug() == Ladybug
}
func (p Piece) IsPillBug() bool {
	return p.Bug() == PillBug
}
func (p Piece) Bug() uint8 {
	return uint8(uint32(p)&BugMask>>16)
}
func (p Piece) BugS() string {
	return bugLabels[p.Bug()]
}
func (p Piece) IsPieceA() bool {
	return p.Piece() == PieceA
}
func (p Piece) IsPieceB() bool {
	return p.Piece() == PieceB
}
func (p Piece) IsPieceC() bool {
	return p.Piece() == PieceC
}
func (p Piece) Piece() uint8 {
	return uint8(uint32(p) & PieceMask >> 8)
}
func (p Piece) PieceS() string {
	return pieceLabels[p.Piece()]
}
func (p Piece) String() string {
	return fmt.Sprintf("Color: %s, Bug: %s, Cell: %s", p.ColorS(), p.BugS(), p.PieceS())
}

const (
	NoBug uint8 = iota
	Queen
	Beetle
	Grasshopper
	Spider
	Ant
	Mosquito
	Ladybug
	PillBug

	BugMask = 0b0000000011111110000000000000000

	NoPiece uint8 = 0
	PieceA = 1
	PieceB = 2
	PieceC = 3

	PieceMask = 0b00000000000000001111111100000000

	NoColor uint8 = 0
	BlackColor = 1
	WhiteColor = 2

	ColorMask = 0b1111111000000000000000000000000
)

var ZeroPiece = Piece(0)

var pieceLabels = []string{
	"No Cell",
	"Cell A",
	"Cell B",
	"Cell C",
}

var colorLabels = []string{
	"No Color",
	"Black",
	"White",
}

var bugLabels = []string{
	"No Bug",
	"Queen",
	"Beetle",
	"Grasshopper",
	"Spider",
	"Ant",
	"Mosquito",
	"Ladybug",
	"PillBug",
}
