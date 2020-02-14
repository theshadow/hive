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

func TestBoard_Place(t *testing.T) {
	board := NewBoard()

	pieceA := NewPiece(WhiteColor, Grasshopper, PieceA)
	cPieceA := NewCoordinate(0, 0, 0, 0)
	board.Place(pieceA, cPieceA)

	pieceB := NewPiece(BlackColor, Grasshopper, PieceA)
	cPieceB := NewCoordinate(1, 1, 1, 0)
	board.Place(pieceB, cPieceB)

	pieceC := NewPiece(WhiteColor, Queen, PieceA)
	cPieceC := NewCoordinate(2, 2, 2, 0)
	board.Place(pieceC, cPieceC)

	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Logf("Cell didn't return expected piece")
		t.Fail()
	}
	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Logf("Cell didn't return expected piece")
		t.Fail()
	}
	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Logf("Cell didn't return expected piece")
		t.Fail()
	}

	if _, ok := board.Cell(NewCoordinate(100, 100, 100, 0)); ok {
		t.Logf("Cell returned an unexpected piece")
		t.Fail()
	}
}

func TestBoard_Move(t *testing.T) {
	board := NewBoard()

	p := NewPiece(WhiteColor, Grasshopper, PieceA)
	cA := NewCoordinate(0, 0, 0, 0)
	cB := NewCoordinate(1, 1, 1, 0)

	board.Place(p, cA)
	if err := board.Move(cA, cB); err != nil {
		t.Logf("couldn't move piece on the board")
		t.Fail()
	}

	if _, ok := board.Cell(cA); ok {
		t.Logf("found a piece at the source coordinate")
		t.Fail()
	}

	if _, ok := board.Cell(cB); !ok {
		t.Logf("found no piece at the destination coordinate")
		t.Fail()
	}
}
