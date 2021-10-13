package hive

import "testing"

func TestFormation_CanSlide(t *testing.T) {

}

func TestFormation_IsPinned(t *testing.T) {
	t.Run("When there is a piece above in the formation the result is a boolean TRUE value", func(t *testing.T) {
		formation := Formation{
			ZeroPiece,
			ZeroPiece,
			ZeroPiece,
			ZeroPiece,
			ZeroPiece,
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceA),
		}

		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when there is a piece above")
		}
	})

	t.Run("When the formation is in a Chevron A a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Chevron A")
		}
	})

	t.Run("When the formation is in a Chevron B a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Chevron B")
		}
	})


	t.Run("When the formation is in a Spaceship A a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			NewPiece(BlackColor, Ant, PieceC),
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship A")
		}
	})

	t.Run("When the formation is in a Spaceship B a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceA),
			NewPiece(BlackColor, Ant, PieceB),
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship B")
		}
	})

	t.Run("When the formation is in a Spaceship C a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship C")
		}
	})

	t.Run("When the formation is in a Spaceship D a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			NewPiece(BlackColor, Ant, PieceC),
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship D")
		}
	})

	t.Run("When the formation is in a Spaceship E a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship E")
		}
	})


	t.Run("When the formation is in a Spaceship F a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			NewPiece(BlackColor, Ant, PieceB),
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Spaceship F")
		}
	})

	t.Run("When the formation is in a Butterfly A a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceA),
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Butterfly A")
		}
	})

	t.Run("When the formation is in a Butterfly B a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceB),
			NewPiece(BlackColor, Ant, PieceC),
			ZeroPiece,
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Butterfly B")
		}
	})

	t.Run("When the formation is in a Butterfly C a boolean TRUE value is returned", func(t *testing.T) {
		formation := Formation{
			NewPiece(BlackColor, Ant, PieceA),
			NewPiece(BlackColor, Ant, PieceB),
			ZeroPiece,
			NewPiece(BlackColor, Ant, PieceC),
			NewPiece(BlackColor, Spider, PieceA),
			ZeroPiece,
			ZeroPiece,
		}
		if !formation.IsPinned() {
			t.Error("Expected a boolean TRUE value when the formation is in Butterfly C")
		}
	})

}
