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

/*
      QA|AAGG|GBBS|SMLP
	1111|1111|1111|1111
*/
type Player uint16
