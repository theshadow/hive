package hived

// Used to track the neighbors around a piece. Specifically adds functionality for quickly validating
// state logic
type Formation [7]Piece
func (f Formation) Contacts() (count int) {
	for _, p := range f {
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

	// Shell formation always blocks movement
	if contacts := f.Contacts(); contacts >= 5 {
		return false
	} else if contacts == 4 && (f.inButterfly() || f.inSpaceship()) {
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
func (f Formation) inButterfly() bool {
	if f.Contacts() != 4 {
		return false
	}

	// TODO: What if I created a bitmask of the neighbors and rotated against
	//       patterns for butterfly, broken butterfly, and spaceship
	// butterfly = [ 0, 1, 1, 0, 1, 1 ]
	// spaceship = [ 1, 1, 1, 0, 1, 0 ]
	// broken    = [ 0, 1, 1, 0, 1, 0 ]

	return false
}

func (f Formation) inBrokenButterfly() bool {
	return false
}

func (f Formation) inSpaceship() bool {
	return false
}

func (f Formation) north() Piece {
	return f[0]
}
func (f Formation) northeast() Piece {
	return f[1]
}
func (f Formation) southeast() Piece {
	return f[2]
}
func (f Formation) south() Piece {
	return f[3]
}
func (f Formation) southwest() Piece {
	return f[4]
}
func (f Formation) northwest() Piece {
	return f[5]
}
func (f Formation) above() Piece {
	return f[6]
}
