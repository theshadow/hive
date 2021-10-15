package hive

import "fmt"

// Player type tracks the color and remaining pieces that a player has.
//
//    . Unused
//    (C)olor
//    (Q)ueen
//    (A)nt         x 3
//    (G)rasshopper x 3
//    (B)eetle      x 2
//    (S)pider      x 2
//    (M)osquito
//    (L)adybug
//    (P)ill Bug
//
//    .CQA|AAGG|GBBS|SMLP
//    1111|1111|1111|1111
//
type Player uint16

func NewPlayer() *Player {
	p := ^Player(0)
	return &p
}
func (p *Player) HasZeroPieces() bool {
	return (*p << 2) == 0
}
func (p *Player) IsWhite() bool {
	return (*p & 0b0100000000000000) != 0
}
func (p *Player) IsBlack() bool {
	return (*p & 0b0100000000000000) == 0
}
func (p *Player) HasQueen() bool {
	return (*p & QueenMask) != 0
}

// Ants will return 3 or less for the count as there are only three ants per player
func (p *Player) Ants() (count int) {
	n := int((*p & AntsMask) >> 9)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Grasshoppers will return 3 or less for the count as there are only three grasshoppers per player
func (p *Player) Grasshoppers() (count int) {
	n := int((*p & GrasshoppersMask) >> 7)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Beetles will return 2 or less for the count as there are only two beetles per player
func (p *Player) Beetles() (count int) {
	n := int((*p & BeetlesMask) >> 5)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Spiders will return 2 or less for the count as there are only two spiders per player
func (p *Player) Spiders() (count int) {
	n := int((*p & SpidersMask) >> 3)
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func (p *Player) HasMosquito() bool {
	return ((*p & MosquitoMask) >> 2) != 0
}
func (p *Player) HasLadybug() bool {
	return ((*p & LadybugMask) >> 1) != 0
}
func (p *Player) HasPillBug() bool {
	return ((*p & PillBugMask) >> 0) != 0
}

// The Take* interface is the way you take a piece from a players inventory.
// As we're treating a player as a Location type it is immutable without mucking with memory.
// Instead, we accept that player is a Location type and say that any modifications are made via
// returning a modified version of the Location.

// TakeQueen will attempt to take a Queen piece from the players inventory and will return an ErrNoPieceAvailable if
// there aren't any pieces available.
func (p *Player) TakeQueen() error {
	if !p.HasQueen() {
		return ErrNoPieceAvailable
	}
	*p &^= QueenMask
	return nil
}

// TakeAnAnt will attempt to take an Ant piece from the players inventory and will return an ErrNoPieceAvailable if
// aren't any pieces available.
func (p *Player) TakeAnAnt() error {
	if p.Ants() == 3 {
		*p &^= AntABitMask
		return nil
	} else if p.Ants() == 2 {
		*p &^= AntBBitMask
		return nil
	} else if p.Ants() == 1 {
		*p &^= AntCBitMask
		return nil
	} else {
		return ErrNoPieceAvailable
	}
}

// TakeAGrasshopper will attempt to take a Grasshopper piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakeAGrasshopper() error {
	if p.Grasshoppers() == 3 {
		*p &^= GrasshopperAMask
		return nil
	} else if p.Grasshoppers() == 2 {
		*p &^= GrasshopperBMask
		return nil
	} else if p.Grasshoppers() == 1 {
		*p &^= GrasshopperCMask
		return nil
	} else {
		return ErrNoPieceAvailable
	}
}

// TakeABeetle will attempt to take a Beetle piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakeABeetle() error {
	if p.Beetles() == 2 {
		*p &^= BeetleAMask
		return nil
	} else if p.Beetles() == 1 {
		*p &^= BeetleBMask
		return nil
	} else {
		return ErrNoPieceAvailable
	}
}

// TakeASpider will attempt to take a Spider piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakeASpider() error {
	if p.Spiders() == 2 {
		*p &^= SpiderAMask
		return nil
	} else if p.Spiders() == 1 {
		*p &^= SpiderBMask
		return nil
	} else {
		return ErrNoPieceAvailable
	}
}

// TakeMosquito will attempt to take a Mosquito piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakeMosquito() error {
	if !p.HasMosquito() {
		return ErrNoPieceAvailable
	}
	*p &^= MosquitoMask
	return nil
}

// TakeLadybug will attempt to take a Ladybug piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakeLadybug() error {
	if !p.HasLadybug() {
		return ErrNoPieceAvailable
	}
	*p &^= LadybugMask
	return nil
}

// TakePillBug will attempt to take a PillBug piece from the players inventory and will return an
// ErrNoPieceAvailable if aren't any pieces available.
func (p *Player) TakePillBug() error {
	if !p.HasPillBug() {
		return ErrNoPieceAvailable
	}
	*p &^= PillBugMask
	return nil
}

// String
func (p *Player) String() string {
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

	GrasshopperAMask = 0b0000001000000000
	GrasshopperBMask = 0b0000000100000000
	GrasshopperCMask = 0b0000000010000000

	BeetleAMask = 0b0000000001000000
	BeetleBMask = 0b0000000000100000

	SpiderAMask = 0b0000000000010000
	SpiderBMask = 0b0000000000001000
)
