package hived

// Used to track the neighbors around a piece. Specifically adds functionality for quickly validating
// position of pieces around a center piece
//
// Detecting specific formations
//
// There are four major formations that we're interested in, which are the Spaceship, Butterfly, Chevron,
// and Broken Butterfly. As all other formations can quickly be identified by the number of contacts.
//
// We can reduce our complexity at the start by ignoring the Above contact as any piece with something
// above it is already pinned. With the remaining six contacts if we took the six positions and represented
// them in base-2 then we'd see the following.
//
// WHERE:
//    Cardinal Directions (N, NE, SE, S, SW, NW) are the six contact points around the center piece.
//    Ones represent a cardinal direction with a piece. Dec is the decimal representation of the formation.
//
//            N  NE SE S  SW NW  DEC
//   Chevron: 1  0  1  0  1  0   42
// Spaceship: 1  1  1  0  1  0   58
// Butterfly: 1  1  0  1  1  0   54
//
// We also know that each of these formations have multiple permutations where as long as the spacing between
// the pieces remains the same the formation still prevents the piece from moving. Below we can see the various
// permutations. An easy way to imagine this is to treat the bitfield is a matrix that we'll be performing
// operations against. Specifically we want to rotate the field through each of the available permutations.
//
// A permutation is defined as moving either the head or the tail of the matrix and appending it to the opposite
// end. See the functional-pseudo example below:
//
//  matrix = [ 1, 0, 1, 0, 1, 0 ]
//  head = f(matrix) = 1
//  body = f(matrix) = [ 0, 1, 0, 1, 0 ]
//
//  ∴
//
//  rotation = f(matrix) = body(matrix) + head(matrix) = [ 0, 1, 0, 1, 0, 1 ]
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
// the number of checks if we can find a cheap way to create the reflection of an integer value. If we can't I believe
// we can simply create a map[int]int where the key is the decimal value of the bitfield and the mapped to integer is
// the type of formation, be it Chevron, Butterfly, or Spaceship.
//
// TODO: Can I create a linear function that when provided a formation in decimal form, that it can validate if its
//       a part of the formation set?
//
type Formation [7]Piece
// Contacts returns the number of edges with pieces ignoring Above
// as it's not necessary for any algorithms and makes checks
// further on more complicated.
func (f Formation) Contacts() (count int) {
	for _, p := range f[:6] {
		if p != ZeroPiece {
			count++
		}
	}
	return count
}
func (f Formation) CanSlide() bool {
	// beetle on top
	if f.above() != ZeroPiece {
		return false
	}

	var formation Formation

	// Shell formation always blocks movement
	if contacts := f.Contacts(); contacts >= 5 {
		return false
	} else if isPinned(formation) {
		return false
	}

	return true
}
func (f Formation) MaySplitHive() bool {
	return f.inBrokenButterfly()
}
func (f Formation) HasNoNeighbors() bool {
	for _, p := range f {
		if p != ZeroPiece {
			return false
		}
	}
	return true
}

func (f Formation) inBrokenButterfly() bool {
	return false
}

func (f Formation) inSpaceship() bool {
	return false
}

func (f Formation) isPinned() bool {
	// TODO: Create the NeighborBitField() function on the board that returns the neighbors as an integer
	return isPinned(0)
}

func isPinned(formation int) bool {
	if _, ok := formationMap[formation]; !ok {
		return false
	}
	return true
}

// Internal map used to detect specific formations.
var formationMap = map[int]int{
	// Chevron
	20: Chevron,
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
