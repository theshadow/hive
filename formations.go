package hived

// Used to track the neighbors around a piece. Specifically adds functionality for quickly validating
// position of pieces around a center piece
//
// Detecting Specific Formations
//
// There are four major formations that we're interested in, which are the Spaceship, Butterfly, Chevron,
// and Broken Butterfly. As all other formations are either easily tested for or are discarded as
// non-relevant
//
// We can reduce our complexity at the start by ignoring the Above contact as any piece with something
// above it is already pinned. With the remaining six contacts if we took the six positions and represented
// them in base-2 then we'd see the following.
//
// WHERE:
//    Cardinal Directions (N, NE, SE, S, SW, NW) are the six contact points around the center piece.
//    Ones represent a cardinal direction with a piece. Dec is the decimal representation of the formation.
//
//
//            N  NE SE S  SW NW  DEC
//   Chevron: 1  0  1  0  1  0   42
// Spaceship: 1  1  1  0  1  0   58
// Butterfly: 1  1  0  1  1  0   54
//
// We also know that each of these formations have multiple permutations where as long as the spacing between
// the pieces remains the same the formation still prevents the piece from moving. Below we can see the various
// permutations. An easy way to imagine this is to treat the bitfield as a matrix that we'll be performing
// operations against. Specifically we want to rotate the field through each of the available permutations.
//
// A permutation is defined as moving either the first or the last element of the matrix and to the opposite
// end of the matrix. See the functional-pseudo example below:
//
//  M = [ 1, 0, 1, 0, 1, 0 ]
//
//  first  = f(M) -> int   = 1
//  last   = f(M) -> int   = 0
//  head   = f(M) -> []int = [ 0, 1, 0, 1, 0 ]
//  tail   = f(M) -> []int = [ 1, 0, 1, 0, 1 ]
//
//  ∴
//
//  lRotation = f(M) -> []int = last(M) + head(M)  = [ 1, 0, 1, 0, 1, 0 ]
//  rRotation = f(M) -> []int = tail(M) + first(M) = [ 0, 1, 0, 1, 0, 1 ]
//
//  By rotating each field through all of the permutations we can see the following options, numbers within parenthesis
//  represent the decimal form of the matrix.
//
//   CHEVRON        SPACESHIP      BUTTERFLY
//  010101 (21)    010111 (23)    011011 (27)
//  101010 (42)    011101 (29)    101101 (45)
//                 101011 (43)    110110 (54)
//                 101110 (46)
//                 110101 (53)
//                 111010 (58)
//
// True to matrix math we can quickly identify that in each column all of the permutations contain reflections. That
// is, "101010" is the mirror image of "010101". We may be able to use this information during our checks to reduce
// the number of operations if we can find a cheap way to create the reflection of an integer Location. If we can't I believe
// we can simply create a map[int]int where the key is the decimal Location of the bitfield and the mapped to integer is
// the type of formation, be it Chevron, Butterfly, or Spaceship.
//
type Formation [7]Piece

func (f Formation) CanSlide() bool {
	return !f.IsPinned()
}

func (f Formation) IsPinned() bool {
	if f.above() != ZeroPiece {
		return true
	}
	if f.contacts() >= 5 {
		return true
	}
	return isPinned(f.bitField())
}

// contacts returns the number of edges with pieces ignoring Above
// as it's not necessary for any algorithms and makes checks
// further on more complicated.
func (f Formation) contacts() (count int) {
	for _, p := range f[:6] {
		if p != ZeroPiece {
			count++
		}
	}
	return count
}

// bitField returns the formation information encoded into an integer by
// representing each contact point as a bit. This allows us to derive an
// integer value for a given formation and use that to quickly determine
// if the current formation is one where the piece is pinned or not.
func (f Formation) bitField() (field int) {
	// i tracks which bit to set starting with the highest
	// bit as we defined N as the highest bit above.
	//
	// j tracks the position in the Formation([7]Piece) array from the lowest element
	// because we wrote the Neighbors function to store N from that point.
	//
	// Once again Above has no value in this algorithm we just ignore it by starting
	// from 5.
	//
	// Was this a poor idea to define them this way? I'm not sure. On the one hand
	// it makes this loop a little more complicated. On the other hand the order
	// feels natural to humans and the formations are kind of a human construct?
	// I dunno, I just code here.
	for i, j := 5, 0; i >= 0; i, j = i-1, j+1 {
		v := f[j]
		if v == ZeroPiece {
			continue
		}
		field |= 1 << i
	}
	return field
}
func (f Formation) above() Piece {
	return f[Above]
}

func (f Formation) IsSuffocating() bool {
	return f.contacts() == 6
}

// given an integer form of a formation reference
// the formationMap and if the formation exists
// return true, otherwise return false.
func isPinned(formation int) bool {
	if _, ok := formationMap[formation]; !ok {
		return false
	}
	return true
}

// Internal map used to detect specific formations.
var formationMap = map[int]int{
	// Chevron
	21: Chevron,
	42: Chevron,

	// Spaceship
	23: Spaceship,
	29: Spaceship,
	43: Spaceship,
	46: Spaceship,
	53: Spaceship,
	58: Spaceship,

	// Butterfly
	27: Butterfly,
	45: Butterfly,
	54: Butterfly,
}

const (
	NoFormation = iota
	Chevron
	Spaceship
	Butterfly
)
