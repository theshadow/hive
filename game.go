package hived

import (
	"fmt"
)

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

	tie bool

	// current board state
	board *Board

	// TODO: Should the game know its history or should the session?
	// each move for this game
	history []Move

	// Track the pieces that are paralyzed by mapping the location of the piece to a
	// time till free value. When the value is zero, the piece is removed from the map
	// and freed.
	//
	// After each turn the the value is decremented by one.
	paralyzedPieces map[Coordinate]int

	// feature flags
	features map[Feature]bool
}

func NewGame(features []Feature) *Game {
	featureMap := copyFeatureMap()
	if features != nil {
		for _, f := range features {
			featureMap[f] = true
		}
	}
	return &Game{
		turns:           1, // makes math clearer and makes more sense to start at 1 instead of 0
		turn:            WhiteColor,
		white:           NewPlayer(),
		black:           NewPlayer(),
		board:           NewBoard(),
		history:         []Move{},
		paralyzedPieces: make(map[Coordinate]int),
		features:        featureMap,
	}

}

// TODO: Test this function
// Place will accept a piece and a coordinate and attempt
// to place it on the board at the specified coordinate
// if the specified coordinate is an invalid space due to
// game rules or if the player does not have the piece to
// place it will return an error.
func (g *Game) Place(p Piece, c Coordinate) error {
	// the first piece to be placed must be at origin
	if g.turns == FirstTurn && c != Origin {
		return ErrFirstPieceMustBeAtOrigin
	}

	// Is it this players turn to place a piece?
	if p.Color() != g.turn {
		return ErrNotPlayersTurn
	}

	player := g.currentPlayer()

	// take a piece
	if err := g.takeAPiece(p, player); err != nil {
		return err
	}

	// Is this the fourth turn and has the currentPlayer placed their queen or is this piece their queen?
	if g.turns == FourthTurn && player.HasQueen() {
		return ErrMustPlaceQueen
	}

	// Is this placement valid?
	//     - Is it on the surface? (H == 0)
	//     - Is it touching the opponents piece? (neighbors)
	//     no ErrMustPlacePieceOnSurface
	if c.H() > 0 {
		return ErrMustPlacePieceOnSurface
	}

	if g.turns == FirstTurn && g.featureEnabled(TournamentQueensRuleFeature) && p.IsQueen() {
		return ErrMayNotPlaceQueenOnFirstTurn
	} else if g.turns != FirstTurn {
		// we must allow the players to place pieces that touch each other on the first turn, but never again.
		if neighbors, _ := g.board.Neighbors(c); contactWithOpponentsPiece(p, neighbors) {
			return ErrMayNotPlaceTouchingOpponentsPiece
		}
	}

	// place the piece, we're not allowed to place two pieces at the same coordinate
	if err := g.board.Place(p, c); err != nil {
		return err
	}

	// update the history
	g.history = append(g.history, NewMove(Placed, p, 0, c))

	// turn management
	if p.IsQueen() {
		g.updatePlayerQueen(c)
	}

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

	// figure out which player we should be working with
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
	//     - Breaking the hive
	//     no: ErrPieceMayNotMove
	// If the formation of the neighbors is pinning the piece at the specified coordinate
	// then it may not move.
	if neighbors, err := g.board.Neighbors(a); err == nil && Formation(neighbors).isPinned() {
		return ErrPieceMayNotMove
	} else if err != nil {
		// There isn't a piece at (a).
		// TODO: ErrInvalidMove is way to vague here it failed for a reason the message doesn't announce
		//  this is a concern. Context!
		return ErrInvalidMove
	}

	// if the piece is paralyzed the player can't move it
	if g.pieceIsParalyzed(a) {
		return ErrPieceMayNotMove
	}

	// TODO: implement splitting hive on move
	// If it can slide, and there are four neighbors there is no split.

	// TODO: implement path validation
	// Is this move valid?
	//     - Can this piece move to this location (pathing)
	//     no: ErrInvalidMove

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

	// update the history
	g.history = append(g.history, NewMove(Moved, piece, a, b))

	// turn management
	if piece.IsQueen() {
		g.updatePlayerQueen(b)
	}
	g.toggleTurn()

	return nil
}

// Winner returns the player that won the game, if the game is not over
// this method will return an error.
//
// If there is a tie it will return a ZeroPlayer with a nil error.
func (g *Game) Winner() (Player, error) {
	if !g.Over() {
		return ZeroPlayer, ErrGameNotOver
	}

	if g.tie {
		return ZeroPlayer, nil
	}

	// As we toggle the player at the end of a turn we determine
	// the winner to be the person who the current player isn't.
	var winner Player
	if !g.currentPlayer().IsWhite() {
		winner = g.white
	} else if !g.currentPlayer().IsBlack() {
		winner = g.black
	}

	return winner, nil
}

// If either player has a suffocating queen then the game is over.
func (g *Game) Over() bool {
	// if both players have their queen then the game is not over.
	if g.black.HasQueen() && g.white.HasQueen() {
		return false
	}

	whiteSuffocating := false
	blackSuffocating := false

	// have they placed their queen?
	if g.black.HasQueen() == false {
		// I'm ignoring this error for a reason of long winded logic
		//
		// tl;dr — It should be impossible to reach this point and have a false victory.
		//
		// The only way Neighbors() can return an error is if the supplied coordinate
		// is invalid. Based on the game rules the first piece will be placed at the
		// origin so it would be impossible to reach this conditional while the player
		// also has a queen to place.
		//
		// Further, the only time where there may be a false victory is IF the
		// queen had a coordinate at the origin in the game state but the board
		// had a piece at origin that was not a queen. In that state, we would
		// have a false victory. However, we can't reach here without a queen being placed,
		// and the only way for a queen to have an origin coordinate is if the player
		// places or history their queen to origin.
		neighbors, _ := g.board.Neighbors(g.blackQueen)
		formation := Formation(neighbors)
		blackSuffocating = formation.IsSuffocating()
	}

	if g.white.HasQueen() == false {
		neighbors, _ := g.board.Neighbors(g.whiteQueen)
		formation := Formation(neighbors)
		whiteSuffocating = formation.IsSuffocating()
	}

	// tie
	if blackSuffocating && whiteSuffocating {
		g.tie = true
		return true
	} else if blackSuffocating || whiteSuffocating {
		return true
	}

	return false
}
func (g *Game) History() []Move {
	return g.history
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
		g.tickParalyzedPieces()

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

func (g *Game) featureEnabled(f Feature) bool {
	_, ok := g.features[f]
	return ok
}

func contactWithOpponentsPiece(p Piece, neighbors [7]Piece) bool {
	color := NoColor
	if p.Color() == WhiteColor {
		color = BlackColor
	} else {
		color = WhiteColor
	}

	for _, n := range neighbors {
		// don't care about zero pieces
		if n == ZeroPiece {
			continue
		}

		if n.Color() == color {
			return true
		}
	}

	return false
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

type Winner int

const (
	FirstTurn  = 1
	FourthTurn = 4

	Tie         Winner = 0
	BlackPlayer Winner = 1
	WhitePlayer Winner = 2
)

var ErrFirstPieceMustBeAtOrigin = fmt.Errorf("the first piece to be placed must be placed at origin")
var ErrGameNotOver = fmt.Errorf("there isn't a declared winner as the game is not over")
var ErrUnknownPiece = fmt.Errorf("an unknown piece was encountered")

var ErrMustPlacePieceOnSurface = fmt.Errorf("a piece must be placed on the surface of the board")
var ErrMayNotPlaceTouchingOpponentsPiece = fmt.Errorf("the player may not place a piece where it will touch an opponents piece after the first turn")

var ErrMayNotPlaceQueenOnFirstTurn = fmt.Errorf("tournament rules: a player may not place their queen on the first turn")

var ErrInvalidMove = fmt.Errorf("the specified move is invalid")
var ErrPieceMayNotMove = fmt.Errorf("this piece may not move")
var ErrNotPlayersTurn = fmt.Errorf("a player may only act on their turn")
var ErrMustPlaceQueen = fmt.Errorf("the player must place their queen by the fourth turns")
var ErrMustPlaceQueenToMove = fmt.Errorf("the players queen must be placed before a placed piece may move")
var ErrPieceAlreadyParalyzed = fmt.Errorf("the piece is already paralyzed and may not be stunned again this turn")

type Feature uint64

const (
	NoFeature Feature = iota
	LadybugPieceFeature
	PillBugPieceFeature
	MosquitoPieceFeature

	TournamentQueensRuleFeature
)

var featureMap = map[Feature]bool{
	LadybugPieceFeature:         false,
	PillBugPieceFeature:         false,
	MosquitoPieceFeature:        false,
	TournamentQueensRuleFeature: false,
}

func copyFeatureMap() (features map[Feature]bool) {
	features = make(map[Feature]bool)
	for k, v := range featureMap {
		features[k] = v
	}
	return features
}
