package hived

import "fmt"

type Game struct {
	// each move for this game
	moves []Move

	// current board state
	board Board
}
func NewGame() *Game {
	return new(Game)
}

func (g *Game) Place(p Piece, c Coordinate) error {
	// TODO: Implement game rules for placement
	// Is this player allowed to move?
	//     no: ErrNotPlayersTurn
	// Is this the fourth turn and has the player placed their queen or is this piece their queen?
	//     no: ErrMustPlaceQueen
	// Is this placement valid?
	//     no ErrInvalidPlacement
	// Place the piece

	return nil
}

func (g *Game) Move(a, b Coordinate) error {
	// TODO: Implement game rules for movement
	// Is this player allowed to move?
	//     no: ErrNotPlayersTurn
	// Has this color placed their queen?
	//     no: ErrMustPlaceQueenToMove
	// Is this piece allowed to move?
	//     no: ErrPieceMayNotMove
	// Is this move valid?
	//     no: ErrInvalidMove
	// Move the piece

	return nil
}

var ErrInvalidPlacement = fmt.Errorf("the specified placement is invalid")
var ErrInvalidMove = fmt.Errorf("the specified move is invalid")
var ErrPieceMayNotMove = fmt.Errorf("this piece may not move")
var ErrNotPlayersTurn = fmt.Errorf("a player may only move a piece on their turn")
var ErrMustPlaceQueen = fmt.Errorf("the player must place their queen by the fourth turn")
var ErrMustPlaceQueenToMove = fmt.Errorf("the players queen must be placed before a placed piece may move")

/* A player tracks the color and remaining cells the player has.

	.CQA|AAGG|GBBS|SMLP
	1111|1111|1111|1111

     . Unused
	(C)olor
	(Q)ueen
	(A)nt
	(G)rasshopper
	(B)eetle
	(S)pider
	(M)osquito
	(L)adybug
	(P)ill Bug
*/
type Player uint16
func (p Player) HasZeroPieces() bool {
	return (p << 2) == 0
}
func (p Player) IsWhite() bool {
	return (p & 0b0100000000000000) != 0
}
func (p Player) IsBlack() bool {
	return (p & 0b0100000000000000) == 0
}
func (p Player) HasQueen() bool {
	return (p & 0b0010000000000000) != 0
}
func (p Player) Ants() (count int) {
	n := int((p & 0b0001110000000000) >> 11)
	for ; n > 0; {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Grasshoppers() (count int) {
	n := int((p & 0b0000001110000000) >> 7)
	for ; n > 0; {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Beetles() (count int) {
	n := int((p & 0b0000000001100000) >> 5)
	for ; n > 0; {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Spiders() (count int) {
	n := int((p & 0b0000000000011000) >> 3)
	for ; n > 0; {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) HasMosquito() bool {
	return ((p & 0b0000000000000100) >> 2) != 0
}
func (p Player) HasLadybug() bool {
	return ((p & 0b0000000000000010) >> 1) != 0
}
func (p Player) HasPillBug() bool {
	return ((p & 0b0000000000000001) >> 0) != 0
}