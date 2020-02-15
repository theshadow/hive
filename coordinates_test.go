package hived

import "testing"

func TestCoordinate_Set(t *testing.T) {
	cActual := NewCoordinate(1, 2, 3, 4)

	var cExpected Coordinate
	cExpected |= Coordinate(int32(1) << 24)
	cExpected |= Coordinate(int32(2) << 16)
	cExpected |= Coordinate(int32(3) << 8)
	cExpected |= Coordinate(int32(4) << 0)

	if cActual != cExpected {
		t.Logf("actual coordinate did not match the expected coordinate")
		t.Fail()
	}
}

func TestCoordinate_Parts(t *testing.T) {
	c := NewCoordinate(1, 2, 3, 4)

	if c.X() != 1 {
		t.Logf("X didn't return the expected value")
		t.Fail()
	}

	if c.Y() != 2 {
		t.Logf("Y didn't return the expected value")
		t.Fail()
	}

	if c.Z() != 3 {
		t.Logf("Z didn't return the expected value")
		t.Fail()
	}

	if c.H() != 4 {
		t.Logf("H didn't return the expected value")
		t.Fail()
	}
}

func TestCoordinate_Add(t *testing.T) {
	origin := NewCoordinate(0, 0, 0, 0)
	location := origin.Add(NewCoordinate(-1, -2, -3, -4))
	if location.X() != -1 || location.Y() == -2 || location.Z() == -3 || location.H() != -4 {
		t.Logf("location doesn't match the expected location")
		t.Fail()
	}
}
