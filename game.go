package hived

import "fmt"

// Game maintains the game instance state
type Game struct {
	// tracks the number of turns that have passed, a turn occurs after both players have performed an
	// action
	turns uint

	// tracks which players turn it is, WhiteColor or BlackColor.
	turn uint8

	// These track the state of each currentPlayer
	white Player
	black Player

	// Track the coordinate of each queen to quickly detect victory
	// states without having to perform an O(n) over all pieces
	// on the board to find the queen pieces.
	whiteQueen Coordinate
	blackQueen Coordinate

	// current board state
	board *Board

	// each move for this game
	moves []Move

	// Track the pieces that are paralyzed by mapping the location of the piece to a
	// time till free value. When the value is zero, the piece is removed from the map
	// and freed.
	//
	// After each turn the the value is decremented by one.
	paralyzedPieces map[Coordinate]int
}

func NewGame() *Game {
	return &Game{
		turns:           1, // makes math clearer and makes more sense to start at 1 instead of 0
		board:           NewBoard(),
		moves:           []Move{},
		paralyzedPieces: make(map[Coordinate]int),
	}
}

// TODO: Test this function
// Place will accept a piece and a coordinate and attempt
// to place it on the board at the specified coordinate
// if the specified coordinate is an invalid space due to
// game rules or if the player does not have the piece to
// place it will return an error.
func (g *Game) Place(p Piece, c Coordinate) error {
	// Is it this players turn to place a piece?
	//     no: ErrNotPlayersTurn
	if p.Color() != g.turn {
		return ErrNotPlayersTurn
	}

	// figure out which currentPlayer we should be working with
	if err := g.takeAPiece(p, g.currentPlayer()); err != nil {
		return err
	}

	// Is this the fourth turn and has the currentPlayer placed their queen or is this piece their queen?
	//     no: ErrMustPlaceQueen
	if g.nTurns() == FourthTurn && g.currentPlayer().HasQueen() {
		return ErrMustPlaceQueen
	}

	// Is this placement valid?
	//     - Is it on the surface? (H == 0)
	//     - Is it touching the opponents piece? (neighbors)
	//     no ErrInvalidPlacement
	// Place the piece
	if c.H() > 0 {
		return ErrInvalidPlacement
	}

	// Validate that where this piece is being placed doesn't touch an opponents piece
	if err := g.checkNeighbors(p, c); err != nil {
		return err
	}

	// place the piece
	g.board.Place(p, c)

	// turn management
	if p.IsQueen() {
		g.updatePlayerQueen(c)
	}
	g.tickParalyzedPieces()
	g.toggleTurn()

	return nil
}

// TODO: Implement rules for movement
// Move accepts two coordinates and attempts to move the piece found at (a) to (b).
// It will return an error if the movement violates any game rules or if the specified
// coordinate for (a) is invalid.
func (g *Game) Move(a, b Coordinate) error {
	// Is this a valid piece to move?
	piece, ok := g.board.Cell(a)
	if !ok {
		return ErrInvalidCoordinate
	}

	// figure out which currentPlayer we should be working with
	player := g.currentPlayer()

	// Is this currentPlayer allowed to move?
	//     no: ErrNotPlayersTurn
	if piece.Color() != g.turn {
		return ErrNotPlayersTurn
	}

	// Has this color placed their queen?
	//     no: ErrMustPlaceQueenToMove
	if player.HasQueen() {
		return ErrMustPlaceQueenToMove
	}

	// Is this piece allowed to move?
	//     - Rule of sliding
	//     - Paralyzed after Pill Bug action
	//     no: ErrPieceMayNotMove
	// If the formation of the neighbors is pinning the piece at the specified coordinate
	// then it may not move.
	if neighbors, err := g.board.Neighbors(a); err == nil && Formation(neighbors).isPinned() {
		return ErrPieceMayNotMove
	} else if err != nil {
		// There isn't a piece at (a).
		// TODO: ErrInvalidMove is way to vague here it failed for a reason the message doesn't announce
		//  this is a concern.
		return ErrInvalidMove
	}

	// if the piece is paralyzed the player can't move it
	if g.pieceIsParalyzed(a) {
		return ErrPieceMayNotMove
	}

	// TODO: implement path validation
	// Is this move valid?
	//     - Breaking the hive
	//     - Can this piece move to this location (pathing)
	//     no: ErrInvalidMove

	// TODO: implement splitting hive on move
	// If it can slide, and there are four neighbors there is no split.
	
	// TODO: How does ladybug movement work?
	// Probably a modified path algorithm where any cell with a piece within a distance of
	// 2 from the ladybug is considered to have a height of zero?
	// TODO: How does pill bug movement work? Could these be implemented as custom path rules?
	// What if part of the pathing rules allowed the bug to modify the terrain? None of
	// the pieces have height limits for their movement so we could create a terrain
	// mask that made the pathing algorithm see those cells as empty. Thus a lady
	// bug can path over pieces. That might work, I think there are some edge cases.
	//
	// If A is a piece touching a pill bug / (mosquito:pill bug) of the right color
	// and A is not paralyzed, and the pill bug / (mosquito:pill bug) is not paralyzed,
	// and B is an empty cell. Move A to B and return nil.

	// TODO: How does mosquito movement work? This may be another custom Move function,
	//       can I generalize?
	// If A is a mosquito, calculate for each bug type adjacent if B is a valid point
	// in that pieces path algorithm move the piece return nil

	// Move the piece
	if err := g.board.Move(a, b); err != nil {
		// TODO: perhaps wrap this? Does it matter? What context do I gain or lose?
		return err
	}

	// turn management
	if piece.IsQueen() {
		g.updatePlayerQueen(b)
	}
	g.tickParalyzedPieces()
	g.toggleTurn()

	return nil
}
// Winner returns the player that won the game, if the game is not over
// this method will return an error.
func (g *Game) Winner() (Player, error) {
	if !g.Over() {
		return ZeroPlayer, ErrGameNotOver
	}
	return g.currentPlayer(), nil
}
// If either player has a suffocating queen then the game is over.
func (g *Game) Over() bool {
	// have they placed their queen?
	if g.black.HasQueen() == false {
		// I'm ignoring this error for a reason of long winded logic
		//
		// tldr; It should be impossible to reach this point and have a false victory.
		//
		// The only way Neighbors() can return an error is if the supplied coordinate
		// is invalid. Based on the game rules the first piece will be placed at the
		// origin so it would be impossible to reach this conditional while the player
		// also has a queen to place.
		//
		// Further, the only time where there may be a false victory is IF the
		// queen had a coordinate at the origin in the game state but the board
		// had the piece at origin that was not a queen. In that state, we would
		// have a false victory. However, we can't reach here without a queen being placed,
		// and the only way for a queen to have an origin coordinate is if the player
		// places or moves their queen to origin.
		neighbors, _ := g.board.Neighbors(g.blackQueen)
		formation := Formation(neighbors)
		if formation.IsSuffocating() {
			return true
		}
	}

	if g.white.HasQueen() == false {
		neighbors, _ := g.board.Neighbors(g.whiteQueen)
		formation := Formation(neighbors)
		if formation.IsSuffocating() {
			return true
		}
	}

	return false
}
func (g *Game) updatePlayerQueen(c Coordinate) {
	if g.currentPlayer().IsWhite() {
		g.whiteQueen = c
	} else {
		g.blackQueen = c
	}
}
func (g *Game) takeAPiece(p Piece, player Player) error {
	player, err := takeAPiece(p, g.currentPlayer())
	if err != nil {
		return err
	}
	if player.IsBlack() {
		g.black = player
	} else {
		g.white = player
	}
	return nil
}

// checks all of the neighbors and if any of them don't match the color of the
// piece being placed it returns an error.
func (g *Game) checkNeighbors(p Piece, c Coordinate) error {
	neighbors, err := g.board.Neighbors(c)
	if err != nil {
		return err
	}

	for _, n := range neighbors {
		// don't care about zero pieces
		if n == ZeroPiece {
			continue
		}

		// TODO: Are there any other rules about neighbors and placement???
		if p.Color() != n.Color() {
			// TODO: Log which piece is offending?
			return ErrInvalidPlacement
		}
	}
	return nil
}
func (g *Game) currentPlayer() Player {
	// figure out which currentPlayer we should be working with
	if g.turn == WhiteColor {
		return g.white
	}
	return g.black
}

func (g *Game) toggleTurn() {
	if g.turn == WhiteColor {
		g.turn = BlackColor
	} else {
		g.turn = WhiteColor
		g.turns++
	}
}

func (g *Game) pieceIsParalyzed(c Coordinate) bool {
	_, ok := g.paralyzedPieces[c]
	return ok
}

// used when a Pill Bug paralyzes a piece or itself
func (g *Game) paralyzePiece(c Coordinate) error {
	if _, ok := g.paralyzedPieces[c]; ok {
		return ErrPieceAlreadyParalyzed
	}
	g.paralyzedPieces[c] = 1
	return nil
}

// When called decrements each paralyzed piece's Time-till-freed value by one.
// When the piece reaches zero, it is freed.
func (g *Game) tickParalyzedPieces() {
	// Time till Freed
	for c, ttf := range g.paralyzedPieces {
		if ttf-1 == 0 {
			delete(g.paralyzedPieces, c)
		} else {
			g.paralyzedPieces[c]--
		}
	}
}

func (g *Game) nTurns() int {
	return int(g.turns)
}

func takeAPiece(p Piece, player Player) (Player, error) {
	if p.IsQueen() {
		return player.TakeQueen()
	} else if p.IsAnt() {
		return player.TakeAnAnt()
	} else if p.IsGrasshopper() {
		return player.TakeAGrasshopper()
	} else if p.IsSpider() {
		return player.TakeASpider()
	} else if p.IsBeetle() {
		return player.TakeABeetle()
	} else if p.IsLadybug() {
		return player.TakeLadybug()
	} else if p.IsMosquito() {
		return player.TakeMosquito()
	} else if p.IsPillBug() {
		return player.TakePillBug()
	}
	return ZeroPlayer, ErrUnknownPiece
}

const (
	FourthTurn = 4
)

var ErrGameNotOver = fmt.Errorf("there isn't a declared winner as the game is not over")
var ErrUnknownPiece = fmt.Errorf("an unknown piece was encountered")
var ErrInvalidPlacement = fmt.Errorf("the specified placement is invalid")
var ErrInvalidMove = fmt.Errorf("the specified move is invalid")
var ErrPieceMayNotMove = fmt.Errorf("this piece may not move")
var ErrNotPlayersTurn = fmt.Errorf("a currentPlayer may only move a piece on their turns")
var ErrMustPlaceQueen = fmt.Errorf("the currentPlayer must place their queen by the fourth turns")
var ErrMustPlaceQueenToMove = fmt.Errorf("the players queen must be placed before a placed piece may move")
var ErrPieceAlreadyParalyzed = fmt.Errorf("the piece is already paralyzed and may not be stunned again this turn")

