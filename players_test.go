package hived

import "testing"

func TestPiece_Set(t *testing.T) {
	var pExpected Piece
	pExpected |= Piece(uint32(WhiteColor) << 24)
	pExpected |= Piece(uint32(Beetle) << 16)
	pExpected |= Piece(uint32(PieceC) << 8)

	var pActual Piece
	(&pActual).Set(WhiteColor, Beetle, PieceC)

	if pExpected != pActual {
		t.Logf("the expected piece %s did not match the actual piece %s", pExpected, pActual)
		t.Fail()
	}
}

func TestPiece_Color(t *testing.T) {
	p := Piece(uint32(BlackColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.Color() != BlackColor {
		t.Log("Color didn't return black")
		t.Fail()
	}

	if !p.IsBlack() {
		t.Log("IsBlack didn't return true")
		t.Fail()
	}

	p = Piece(uint32(WhiteColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.Color() != WhiteColor {
		t.Log("Color didn't return white")
		t.Fail()
	}

	if !p.IsWhite() {
		t.Log("IsWhite didn't return true")
		t.Fail()
	}
}

func TestPiece_ColorS(t *testing.T) {
	p := Piece(uint32(BlackColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.ColorS() != colorLabels[BlackColor] {
		t.Log("ColorS didn't return black")
		t.Fail()
	}

	p = Piece(uint32(WhiteColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.ColorS() != colorLabels[WhiteColor] {
		t.Log("ColorS didn't return white")
		t.Fail()
	}
}

func TestPiece_Bug(t *testing.T) {
	for i := Queen; i < PillBug+1; i++ {
		p := Piece(uint32(i) << 16)

		t.Logf("Binary - %32b", p)
		t.Logf("Cell - %s", p)

		if p.Bug() != i {
			t.Logf("Bug didn't return %d", i)
			t.Fail()
		}
		if p.BugS() != bugLabels[i] {
			t.Logf("BugS didn't return %s", bugLabels[i])
		}
	}
}

func TestPiece_Piece(t *testing.T) {
	for i := PieceA; i < PieceC+1; i++ {
		p := Piece(uint32(i) << 8)

		t.Logf("Binary - %32b", p)
		t.Logf("Cell - %s", p)

		if int(p.Piece()) != i {
			t.Logf("Cell didn't return %d", i)
			t.Fail()
		}
		if p.PieceS() != pieceLabels[i] {
			t.Logf("Cell didn't return %s, returned %s", pieceLabels[i], p.PieceS())
			t.Fail()
		}
	}
}
