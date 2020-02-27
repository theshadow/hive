package hived

import "testing"

func TestMove_Set(t *testing.T) {
	var p = NewPiece(WhiteColor, Grasshopper, PieceC)

	src := NewCoordinate(2, 3, 4, 5)
	dst := NewCoordinate(6, 7, 8, 9)

	var m = NewMove(Moved, p, src, dst)

	if m.Action() != Moved {
		t.Logf("the expected action %s did not match the actual action %s", actionLabels[Moved], m.ActionS())
		t.Fail()
	}

	if m.Piece() != p {
		t.Logf("the expected piece %s did not match the actual piece %s", p, m.Piece())
		t.Fail()
	}

	if m.Src() != src {
		t.Logf("the expected src coordinate %s did not match the actual coordinate %s", src, m.Src())
		t.Fail()
	}

	if m.Dst() != dst {
		t.Logf("the expected dst coordinate %s did not match the actual coordinate %s", src, m.Dst())
		t.Fail()
	}
}

func TestMove_Action(t *testing.T) {
	m := NewMove(
		Moved,
		NewPiece(WhiteColor, Beetle, PieceA),
		NewCoordinate(2, 3, 4, 5),
		NewCoordinate(6, 7, 8, 9))

	if m.Action() != Moved {
		t.Logf("Expected: %32b (%d), Actual: %32b (%d)", Moved, Moved, m.Action(), m.Action())
		t.Logf("Action didn't return the expected action")
		t.Fail()
	}
}

func TestMove_ActionS(t *testing.T) {
	m := NewMove(
		Moved,
		NewPiece(WhiteColor, Beetle, PieceA),
		NewCoordinate(2, 3, 4, 5),
		NewCoordinate(6, 7, 8, 9))

	if m.ActionS() != actionLabels[Moved] {
		t.Logf("Expected: %s, Actual: %s", actionLabels[Moved], m.ActionS())
		t.Logf("ActionS didn't return the expected action")
		t.Fail()
	}
}

func TestMove_Piece(t *testing.T) {
	pExpected := NewPiece(WhiteColor, Beetle, PieceA)
	m := NewMove(
		Moved,
		NewPiece(WhiteColor, Beetle, PieceA),
		NewCoordinate(2, 3, 4, 5),
		NewCoordinate(6, 7, 8, 9))

	if m.Piece() != pExpected {
		t.Logf("Expected: %32b (%d), Actual: %32b (%d)", pExpected, pExpected, m.Piece(), m.Piece())
		t.Logf("Cell didn't return the expected piece")
		t.Fail()
	}
}

func TestMove_Src(t *testing.T) {
	expected := NewCoordinate(2, 3, 4, 5)
	m := NewMove(
		Moved,
		NewPiece(WhiteColor, Beetle, PieceA),
		NewCoordinate(2, 3, 4, 5),
		NewCoordinate(6, 7, 8, 9))

	if m.Src() != expected {
		t.Logf("Expected: %s, Actual: %s", expected, m.Src())
		t.Logf("Src didn't return the expected coordinate")
		t.Fail()
	}
}

func TestMove_Dst(t *testing.T) {
	expected := NewCoordinate(6, 7, 8, 9)
	m := NewMove(
		Moved,
		NewPiece(WhiteColor, Beetle, PieceA),
		NewCoordinate(2, 3, 4, 5),
		NewCoordinate(6, 7, 8, 9))

	if m.Dst() != expected {
		t.Logf("Expected: %s, Actual: %s", expected, m.Src())
		t.Logf("Dst didn't return the expected coordinate")
		t.Fail()
	}
}
