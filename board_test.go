package hive

import "testing"

func TestBoard_Place(t *testing.T) {
	board := NewBoard()

	pieceA := NewPiece(WhiteColor, Grasshopper, PieceA)
	cPieceA := NewCoordinate(0, 0, 0, 0)
	_ = board.Place(pieceA, cPieceA)

	pieceB := NewPiece(BlackColor, Grasshopper, PieceA)
	cPieceB := NewCoordinate(1, 1, 1, 0)
	_ = board.Place(pieceB, cPieceB)

	pieceC := NewPiece(WhiteColor, Queen, PieceA)
	cPieceC := NewCoordinate(2, 2, 2, 0)
	_ = board.Place(pieceC, cPieceC)

	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Error("Cell didn't return expected piece")
	}
	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Error("Cell didn't return expected piece")
	}
	if p, ok := board.Cell(cPieceA); !ok || p != pieceA {
		t.Error("Cell didn't return expected piece")
	}

	if _, ok := board.Cell(NewCoordinate(100, 100, 100, 0)); ok {
		t.Error("Cell returned an unexpected piece")
	}

	if err := board.Place(pieceC, cPieceC); err == nil {
		t.Error("expected place to return an ErrPauliExclusionPrinciple when trying to place a piece in a cell where a piece exists")
	}
}

func TestBoard_Move(t *testing.T) {
	board := NewBoard()

	p := NewPiece(WhiteColor, Grasshopper, PieceA)
	cA := NewCoordinate(0, 0, 0, 0)
	cB := NewCoordinate(1, 1, 1, 0)

	_ = board.Place(p, cA)
	if err := board.Move(cA, cB); err != nil {
		t.Error("couldn't move piece on the board")
	}

	if _, ok := board.Cell(cA); ok {
		t.Error("found a piece at the source coordinate")
	}

	if _, ok := board.Cell(cB); !ok {
		t.Error("found no piece at the destination coordinate")
	}

	p = NewPiece(WhiteColor, Ant, PieceA)
	cA = NewCoordinate(0, 0, 0, 0)
	_ = board.Place(p, cA)
	if err := board.Move(cA, cB); err == nil {
		t.Error("expected an ErrPauliExclusionPrinciple when trying to move a piece to a cell with a piece")
	}
}

// Test that Neighbors can return a piece on all sides
//
// Place a piece at origin and use the NeighborsMatrix to place the pieces
// on all sides of the origin piece.
//
// Use Neighbors() to retrieve the neighbors from origin and then validate
// that the pieces match.
//
func TestBoard_Neighbors(t *testing.T) {
	board := NewBoard()
	origin := NewCoordinate(0, 0, 0, 0)

	// origin piece
	pWhiteGrasshopperA := NewPiece(WhiteColor, Grasshopper, PieceA)

	// surrounding pieces
	otherPieces := []Piece{
		NewPiece(BlackColor, Grasshopper, PieceA),
		NewPiece(BlackColor, Grasshopper, PieceA),
		NewPiece(BlackColor, Ant, PieceA),
		NewPiece(BlackColor, Ant, PieceB),
		NewPiece(BlackColor, Beetle, PieceA),
		NewPiece(BlackColor, Ladybug, PieceA),
		NewPiece(BlackColor, Beetle, PieceB),
	}

	_ = board.Place(pWhiteGrasshopperA, origin)

	// use the neighbors matrix to place each piece manually
	for idx, op := range otherPieces {
		coord := origin.Add(NeighborsMatrix[idx])
		err := board.Place(op, coord)
		if err != nil {
			t.Errorf("failed to place origin: %s location: %s error: %s", origin, coord, err)
			break
		}
	}

	// grab the neighbors and compare
	neighbors := board.Neighbors(origin)
	for i, neighbor := range neighbors {
		if otherPieces[i] != neighbor {
			t.Errorf("expected piece %s at %d, found %s", otherPieces[i], i, neighbor)
		}
	}
}
