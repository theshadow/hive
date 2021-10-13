package hive

import "testing"

func TestCoordinate(t *testing.T) {
	tests := []struct {
		actual   Coordinate
		expected uint32
	}{
		{
			actual:   NewCoordinate(1, 2, 3, 4),
			expected: uint32(1<<24) | uint32(2<<16) | uint32(3<<8) | uint32(4),
		},
		{
			actual:   NewCoordinate(-2, -3, -4, -5),
			expected: uint32((0b10000000|2)<<24) | uint32((0b10000000|3)<<16) | uint32((0b10000000|4)<<8) | uint32(0b10000000|5),
		},
	}

	for i, test := range tests {
		t.Logf("Test[%d]  Actual: %d (%32b)", i+1, test.actual, test.actual)
		t.Logf("Test[%d] Expected: %d (%32b)", i+1, test.expected, test.expected)

		if test.actual != Coordinate(test.expected) {
			t.Error("actual coordinate did not match the expected coordinate")
		}
	}

}

func TestCoordinate_Parts(t *testing.T) {

	testCases := []struct {
		c          Coordinate
		x, y, z, h int8
	}{
		{NewCoordinate(1, 2, 3, 4), 1, 2, 3, 4},
		{NewCoordinate(-2, -3, -4, -5), -2, -3, -4, -5},
	}

	for _, test := range testCases {
		c := test.c
		t.Logf("Coordinate: %d (%32b), X: %d, Y: %d, Z: %d, H: %d", c, c, test.x, test.y, test.z, test.h)

		if c.X() != test.x {
			t.Errorf("expected X to return %d instead returned %d (%8b)", test.x, c.X(), c.X())
		}
		if c.Y() != test.y {
			t.Errorf("expected Y to return %d instead returned %d (%8b)", test.y, c.Y(), c.Y())
		}
		if c.Z() != test.z {
			t.Errorf("expected Z to return %d instead returned %d (%8b)", test.z, c.Z(), c.Z())
		}
		if c.H() != test.h {
			t.Logf("expected H to return %d instead returned %d (%8b)", test.h, c.H(), c.H())
		}
	}
}
