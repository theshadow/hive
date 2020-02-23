package hived

import "fmt"

/* A currentPlayer tracks the color and remaining cells the currentPlayer has.

 . Unused
(C)olor
(Q)ueen
(A)nt         x 3
(G)rasshopper x 3
(B)eetle      x 2
(S)pider      x 2
(M)osquito
(L)adybug
(P)ill Bug

.CQA|AAGG|GBBS|SMLP
1111|1111|1111|1111
*/
type Player uint16

func NewPlayer() (p Player) {
	// flip everything on and shift off the
	// unused bit just to be consistent and tidy.
	p = ^Player(0)
	return p
}
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

// Will return 3 or less for the count as there are only three ants per currentPlayer
func (p Player) Ants() (count int) {
	n := int((p & AntsMask) >> 9)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Will return 3 or less for the count as there are only three grasshoppers per currentPlayer
func (p Player) Grasshoppers() (count int) {
	n := int((p & GrasshoppersMask) >> 7)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Will return 2 or less for the count as there are only two beetles per currentPlayer
func (p Player) Beetles() (count int) {
	n := int((p & BeetlesMask) >> 5)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Will return 2 or less for the count as there are only two spiders per currentPlayer
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

/*
	The Take* interface is the way you take a piece from a players inventory.
	As we're treating a currentPlayer as a value type it is immutable without mucking with memory.
	Instead we accept that currentPlayer is a value type and say that any modifications are made via
	returning a modified version of the value.
*/
func (p Player) TakeQueen() (Player, error) {
	if !p.HasQueen() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p | QueenMask, nil
}
func (p Player) TakeAnAnt() (Player, error) {
	if p.Ants() == 3 {
		return p &^ AntABitMask, nil
	} else if p.Ants() == 2 {
		return p &^ AntBBitMask, nil
	} else if p.Ants() == 1 {
		return p &^ AntCBitMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeAGrasshopper() (Player, error) {
	if p.Grasshoppers() == 3 {
		return p &^ GrasshopperAMask, nil
	} else if p.Grasshoppers() == 2 {
		return p &^ GrasshopperBMask, nil
	} else if p.Grasshoppers() == 1 {
		return p &^ GrasshopperCMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeABeetle() (Player, error) {
	if p.Beetles() == 2 {
		return p &^ BeetleAMask, nil
	} else if p.Beetles() == 1 {
		return p &^ BeetleBMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeASpider() (Player, error) {
	if p.Spiders() == 2 {
		return p &^ SpiderAMask, nil
	} else if p.Spiders() == 1 {
		return p &^ SpiderBMask, nil
	} else {
		return ZeroPlayer, ErrNoPieceAvailable
	}
}
func (p Player) TakeMosquito() (Player, error) {
	if !p.HasMosquito() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p &^ MosquitoMask, nil
}
func (p Player) TakeLadybug() (Player, error) {
	if !p.HasLadybug() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p &^ LadybugMask, nil
}
func (p Player) TakePillBug() (Player, error) {
	if !p.HasPillBug() {
		return ZeroPlayer, ErrNoPieceAvailable
	}
	return p &^ PillBugMask, nil
}
func (p Player) String() string {
	color := "White"
	if p.IsBlack() {
		color = "Black"
	}
	return fmt.Sprintf("Color: %s, Queen: %t, Ant(s): %d, Grasshopper(s): %d, Beetle(s): %d, Spider(s): %d, Mosquitos: %t, Ladybug: %t, PillBug: %t",
		color, p.HasQueen(), p.Ants(), p.Grasshoppers(), p.Beetles(), p.Spiders(), p.HasMosquito(), p.HasLadybug(),
		p.HasPillBug())
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
