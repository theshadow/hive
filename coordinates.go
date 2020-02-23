package hived

import "fmt"

/* Used to represent a specific cell on the board

    X        Y        Z       H
11111111|11111111|11111111|11111111
*/
type Coordinate uint32

func NewCoordinate(x, y, z, h int8) Coordinate {
	var c Coordinate
	// We have a 50/50 shot per variable that this operation is redundant
	var ux, uy, uz, uh uint8

	// This is how we encode negative numbers into a uint8 (hint: by hand)
	//
	// The Bits library exists specifically because doing signed-bit storage
	// inside of unsigned types is complicated and will really test
	// your understanding of the rules in the type system. I didn't want to
	// include Bits because this is the only place where I have to do this
	// weird cross encoding.
	if x < 0 {
		ux = (0b10000000) | uint8(x*-1)
	} else {
		ux = uint8(x)
	}

	if y < 0 {
		uy = (0b10000000) | uint8(y*-1)
	} else {
		uy = uint8(y)
	}

	if z < 0 {
		uz = (0b10000000) | uint8(z*-1)
	} else {
		uz = uint8(z)
	}

	if h < 0 {
		uh = (0b10000000) | uint8(h*-1)
	} else {
		uh = uint8(h)
	}

	c |= Coordinate(uint32(ux) << 24) | Coordinate(uint32(uy) << 16) |
		Coordinate(uint32(uz) << 8) | Coordinate(uint32(uh))

	return c
}
func (c Coordinate) Add(loc Coordinate) Coordinate {
	return NewCoordinate(c.X()+loc.X(), c.Y()+loc.Y(), c.Z()+loc.Z(), c.H()+loc.H())
}
/*
  These functions rely on bit-masking to mask the required bits out and then
  bit operations to isolate and convert from the uint32 type to an int8
 */
func (c Coordinate) X() int8 {
	if uint8(c >> 24) & 0b10000000 > 0 {
		// shift to the bits, unset the high flag, cast, and add sign
		return int8(uint8(c >> 24) &^ 0b10000000) * -1
	}
	return int8(c >> 24)
}
func (c Coordinate) Y() int8 {
	if uint8(c & YMask >> 16) & 0b10000000 > 0 {
		return int8(uint8(c >> 16) &^ 0b10000000) * -1
	}
	return int8(c & YMask >> 16)
}
func (c Coordinate) Z() int8 {
	if uint8(c & ZMask >> 8) & 0b10000000 > 0 {
		return int8(uint8(c >> 8) &^ 0b10000000) * -1
	}
	return int8(c & ZMask >> 8)
}
func (c Coordinate) H() int8 {
	if (uint8(c) & 0b10000000) > 0 {
		return int8(uint8(c) &^ 0b10000000) * -1
	}
	return int8(c)
}
func (c Coordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d, H: %d", c.X(), c.Y(), c.Z(), c.H())
}

const (
	XMask = 0b11111111000000000000000000000000
	YMask = 0b00000000111111110000000000000000
	ZMask = 0b00000000000000001111111100000000
	HMask = 0b00000000000000000000000011111111
)
