package hived

import "fmt"

type Game struct {
	turns uint

	turn uint8

	white Player
	black Player

	// current board state
	board *Board

	// each move for this game
	moves []Move

	// Track the pieces that are paralyzed by mapping the location of the piece to a
	// time till free value. When the value is zero, the piece is removed from the map
	// and freed.
	//
	// After each turn the the value is decremented
	paralyzedPieces map[Coordinate]int
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(),
		moves: []Move{},
		paralyzedPieces: make(map[Coordinate]int),
	}
}

func (g *Game) Place(p Piece, c Coordinate) error {
	// TODO: Implement game rules for placement
	// Is it this players turn to place a piece?
	//     no: ErrNotPlayersTurn
	if p.Color() != g.turn {
		return ErrNotPlayersTurn
	}

	// Is this the fourth turn and has the player placed their queen or is this piece their queen?
	//     no: ErrMustPlaceQueen
	if g.Turns() == 4 && g.Player().HasQueen() && !p.IsQueen() {
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
	neighbors, err := g.board.Neighbors(c)
	if err != nil {
		return err
	}
	if len(neighbors) == 0 {
		// must be placed next to a piece
		return ErrInvalidPlacement
	}
	for _, n := range neighbors {
		// TODO: What if the piece above is a beetle of a different color?
		//		this may not matter because when placing, the cell must be empty.
		//      This means that there should never be, in this instance,
		//      a neighbor with a positive height.
		if n.Color() != n.Color() {
			// TODO: Log which piece is offending?
			return ErrInvalidPlacement
		}
	}

	// place the piece
	g.board.Place(p, c)

	g.turns++

	return nil
}

func (g *Game) Move(a, b Coordinate) error {
	// TODO: Implement game rules for movement
	// Is this a valid piece to move?
	piece, ok := g.board.Cell(a)
	if !ok {
		return ErrInvalidCoordinate
	}

	// Is this player allowed to move?
	//     no: ErrNotPlayersTurn
	if piece.Color() != g.turn {
		return ErrNotPlayersTurn
	}

	// figure out which player we should be working with
	var player Player
	if g.turn == WhiteColor {
		player = g.white
	} else {
		player = g.black
	}

	// Has this color placed their queen?
	//     no: ErrMustPlaceQueenToMove
	if player.HasQueen() {
		return ErrMustPlaceQueenToMove
	}

	// TODO: Build rule functions to make this function
	// Is this piece allowed to move?
	//     - Rule of sliding
	//     - Paralyzed after Pill Bug action
	//     no: ErrPieceMayNotMove
	if g.pieceParalyzed(a) {
		return ErrPieceMayNotMove
	}

	// TODO: implement the hive breaking and pathing rules
	// Is this move valid?
	//     - Breaking the hive
	//     - Can this piece move to this location (pathing)
	//     no: ErrInvalidMove

	// Move the piece
	if err := g.board.Move(a, b); err != nil {
		// TODO: perhaps wrap this? Does it matter? What context do I gain or lose?
		return err
	}

	// free pieces after the paralyzation ends
	g.updateParalyzedPieces()

	// increase the turns
	g.turns++

	return nil
}
func (g *Game) Player() Player {
	// figure out which player we should be working with
	if g.turn == WhiteColor {
		return g.white
	}
	return g.black
}
func (g *Game) pieceParalyzed(c Coordinate) bool {
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
func (g *Game) updateParalyzedPieces() {
	for c, ttf := range g.paralyzedPieces {
		if ttf - 1 == 0 {
			delete(g.paralyzedPieces, c)
		} else {
			g.paralyzedPieces[c]--
		}
	}
}

func (g *Game) Turns() int {
	return int(g.turns)
}

var ErrInvalidPlacement = fmt.Errorf("the specified placement is invalid")
var ErrInvalidMove = fmt.Errorf("the specified move is invalid")
var ErrPieceMayNotMove = fmt.Errorf("this piece may not move")
var ErrNotPlayersTurn = fmt.Errorf("a player may only move a piece on their turns")
var ErrMustPlaceQueen = fmt.Errorf("the player must place their queen by the fourth turns")
var ErrMustPlaceQueenToMove = fmt.Errorf("the players queen must be placed before a placed piece may move")
var ErrPieceAlreadyParalyzed = fmt.Errorf("the piece is already paralyzed and may not be stunned again this turn")

/* A player tracks the color and remaining cells the player has.

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

	.CQA|AAGG|GBBS|SMLP
	1111|1111|1111|1111


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
	return (p & QueenMask) != 0
}
func (p Player) Ants() (count int) {
	n := int((p & AntsMask) >> 11)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Grasshoppers() (count int) {
	n := int((p & GrasshoppersMask) >> 7)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Beetles() (count int) {
	n := int((p & BeetlesMask) >> 5)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) Spiders() (count int) {
	n := int((p & SpidersMask) >> 3)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p Player) HasMosquito() bool {
	return ((p & MosquitoMask) >> 2) != 0
}
func (p Player) HasLadybug() bool {
	return ((p & LadybugMask) >> 1) != 0
}
func (p Player) HasPillBug() bool {
	return ((p & PillBugMask) >> 0) != 0
}

func (p Player) TakeQueen() (Player, error) {
	if !p.HasQueen() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p | QueenMask, nil
}
func (p Player) TakeAnAnt() (Player, error) {
	if p.Ants() == 3 {
		return p | AntABitMask, nil
	} else if p.Ants() == 2 {
		return p | AntBBitMask, nil
	} else if p.Ants() == 1 {
		return p | AntCBitMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeAGrasshopper() (Player, error) {
	if p.Grasshoppers() == 3 {
		return p | GrasshopperAMask, nil
	} else if p.Grasshoppers() == 2 {
		return p | GrasshopperBMask, nil
	} else if p.Grasshoppers() == 1 {
		return p | GrasshopperCMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeABeetle() (Player, error) {
	if p.Beetles() == 2 {
		return p | BeetleAMask, nil
	} else if p.Beetles() == 1 {
		return p | BeetleBMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeASpider() (Player, error) {
	if p.Spiders() == 2 {
		return p | SpiderAMask, nil
	} else if p.Spiders() == 1 {
		return p | SpiderBMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeMosquito() (Player, error) {
	if !p.HasMosquito() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p | MosquitoMask, nil
}
func (p Player) TakeLadybug() (Player, error) {
	if !p.HasLadybug() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p | LadybugMask, nil
}
func (p Player) TakePillBug() (Player, error) {
	if !p.HasPillBug() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p | PillBugMask, nil
}

var ErrNoPieceAvailable = fmt.Errorf("attempted to take piece that has none left to take")

var ZeroPlayer = Player(0)

const (
	QueenMask        = 0b0010000000000000
	AntsMask         = 0b0001110000000000
	GrasshoppersMask = 0b0000001110000000
	BeetlesMask      = 0b0000000001100000
	SpidersMask      = 0b0000000000011000
	MosquitoMask     = 0b0000000000000100
	LadybugMask      = 0b0000000000000010
	PillBugMask      = 0b0000000000000001

	AntABitMask = 0b0001000000000000
	AntBBitMask = 0b0000100000000000
	AntCBitMask = 0b0000010000000000

	GrasshopperAMask = 0b0000100000000000
	GrasshopperBMask = 0b0000010000000000
	GrasshopperCMask = 0b0000001000000000

	BeetleAMask = 0b0000010000000000
	BeetleBMask = 0b0000001000000000

	SpiderAMask = 0b0000001000000000
	SpiderBMask = 0b0000000100000000
)
