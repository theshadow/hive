package hive

import "testing"

func TestPiece_Set(t *testing.T) {
	var pExpected Piece
	pExpected |= Piece(uint32(WhiteColor) << 24)
	pExpected |= Piece(uint32(Beetle) << 16)
	pExpected |= Piece(uint32(PieceC) << 8)

	var pActual = NewPiece(WhiteColor, Beetle, PieceC)

	if pExpected != pActual {
		t.Errorf("the expected piece %s did not match the actual piece %s", pExpected, pActual)
	}
}

func TestPiece_Color(t *testing.T) {
	p := Piece(uint32(BlackColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.Color() != BlackColor {
		t.Error("Color didn't return black")
	}

	if !p.IsBlack() {
		t.Error("IsBlack didn't return true")
	}

	p = Piece(uint32(WhiteColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.Color() != WhiteColor {
		t.Error("Color didn't return white")
	}

	if !p.IsWhite() {
		t.Error("IsWhite didn't return true")
	}
}

func TestPiece_ColorS(t *testing.T) {
	p := Piece(uint32(BlackColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.ColorS() != colorLabels[BlackColor] {
		t.Error("ColorS didn't return black")
	}

	p = Piece(uint32(WhiteColor) << 24)
	t.Logf("Binary - %32b", p)
	t.Logf("Cell - %s", p)

	if p.ColorS() != colorLabels[WhiteColor] {
		t.Error("ColorS didn't return white")
		t.Fail()
	}
}

func TestPiece_Bug(t *testing.T) {
	for i := Queen; i < PillBug+1; i++ {
		p := Piece(uint32(i) << 16)

		t.Logf("Binary - %32b", p)
		t.Logf("Cell - %s", p)

		if p.Bug() != i {
			t.Errorf("Bug didn't return %d", i)
		}
		if p.BugS() != bugLabels[i] {
			t.Errorf("BugS didn't return %s", bugLabels[i])
		}
	}
}

func TestPiece_Piece(t *testing.T) {
	for i := PieceA; i < PieceC+1; i++ {
		p := Piece(uint32(i) << 8)

		t.Logf("Binary - %32b", p)
		t.Logf("Cell - %s", p)

		if int(p.Piece()) != i {
			t.Errorf("Cell didn't return %d", i)
		}
		if p.PieceS() != pieceLabels[i] {
			t.Errorf("Cell didn't return %s, returned %s", pieceLabels[i], p.PieceS())
		}
	}
}
