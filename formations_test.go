package hived

import "testing"

func TestFormation_CanSlide(t *testing.T) {

}

func TestFormation_IsPinned(t *testing.T) {
	t.Log("When there is a piece above in the formation the result is a boolean TRUE value.")
	formation := Formation{
		ZeroPiece,                         // N
		ZeroPiece,                         // NE
		ZeroPiece,                         // SE
		ZeroPiece,                         // S
		ZeroPiece,                         // SW
		ZeroPiece,                         // NW
		NewPiece(BlackColor, Ant, PieceA), // A
	}

	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when there is a piece above.")
		t.Fail()
	}

	t.Log("When the formation is in a Chevron A a boolean TRUE value is returned.")
	formation = Formation{
		ZeroPiece,                         // N
		NewPiece(BlackColor, Ant, PieceA), // NE
		ZeroPiece,                         // SE
		NewPiece(BlackColor, Ant, PieceB), // S
		ZeroPiece,                         // SW
		NewPiece(BlackColor, Ant, PieceC), // NW
		ZeroPiece,                         // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Chevron A.")
		t.Fail()
	}

	t.Log("When the formation is in a Chevron B a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA), // NE
		ZeroPiece,                         // N
		NewPiece(BlackColor, Ant, PieceB), // S
		ZeroPiece,                         // SE
		NewPiece(BlackColor, Ant, PieceC), // NW
		ZeroPiece,                         // SW
		ZeroPiece,                         // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Chevron B.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship A a boolean TRUE value is returned.")
	formation = Formation{
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceA),    // NE
		ZeroPiece,                            // SE
		NewPiece(BlackColor, Ant, PieceB),    // S
		NewPiece(BlackColor, Ant, PieceC),    // NW
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship A.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship B a boolean TRUE value is returned.")
	formation = Formation{
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceA),    // NE
		NewPiece(BlackColor, Ant, PieceB),    // S
		NewPiece(BlackColor, Ant, PieceC),
		ZeroPiece,                            // SE
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship B.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship C a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceB),    // S
		ZeroPiece,
		NewPiece(BlackColor, Ant, PieceC),
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship C.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship D a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceB),    // S
		NewPiece(BlackColor, Ant, PieceC),
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship D.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship E a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		NewPiece(BlackColor, Ant, PieceB),    // S
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceC),
		ZeroPiece,
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship E.")
		t.Fail()
	}

	t.Log("When the formation is in a Spaceship F a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		NewPiece(BlackColor, Ant, PieceB),    // S
		NewPiece(BlackColor, Ant, PieceC),
		ZeroPiece,                            // N
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Spaceship F.")
		t.Fail()
	}

	t.Log("When the formation is in a Butterfly A a boolean TRUE value is returned.")
	formation = Formation{
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceA),    // NE
		NewPiece(BlackColor, Ant, PieceB),    // S
		ZeroPiece,
		NewPiece(BlackColor, Ant, PieceC),
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Butterfly A.")
		t.Fail()
	}

	t.Log("When the formation is in a Butterfly B a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceB),    // S
		NewPiece(BlackColor, Ant, PieceC),
		ZeroPiece,
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Butterfly B.")
		t.Fail()
	}

	t.Log("When the formation is in a Butterfly C a boolean TRUE value is returned.")
	formation = Formation{
		NewPiece(BlackColor, Ant, PieceA),    // NE
		NewPiece(BlackColor, Ant, PieceB),    // S
		ZeroPiece,                            // N
		NewPiece(BlackColor, Ant, PieceC),
		NewPiece(BlackColor, Spider, PieceA), // SW
		ZeroPiece,
		ZeroPiece,                            // A
	}
	if !formation.IsPinned() {
		t.Logf("Expected a boolean TRUE value when the formation is in Butterfly C.")
		t.Fail()
	}
}
