package game

import (
	"errors"
	"testing"

	"github.com/theshadow/hive"
)

func TestGame_Place(t *testing.T) {

	t.Run("When a player attempts to place a piece not on their turn an error is returned", func(t *testing.T) {
		g := New(nil)
		p := hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(0, 0, 0, 0)
		if err := g.Place(p, coord); err == nil {
			t.Error("Attempted to place a black piece during the white players turn but didn't receive an error")
		} else {
			t.Log("When placing black piece during the white players turn the expected error is returned")
			if !errors.Is(err, ErrRuleNotPlayersTurn) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleNotPlayersTurn, err)
			}
		}
	})

	t.Run("When the first piece is not placed at origin an error is returned", func(t *testing.T) {
		g := New(nil)
		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(0, 0, 0, 1)
		if err := g.Place(p, coord); err == nil {
			t.Errorf("Attempted to place the first piece at %s and did not receive an error", coord)
		} else {
			t.Log("When first piece placed isn't at the origin and returned the expected error is returned")
			if !errors.Is(err, ErrRuleFirstPieceMustBeAtOrigin) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleFirstPieceMustBeAtOrigin, err)
			}
		}
	})

	t.Run("When a player places a piece and they do not have that piece available an error is returned", func(t *testing.T) {
		wplayer := hive.NewPlayer()
		_ = wplayer.TakeQueen()
		g := &Game{
			turns:           1,
			turn:            hive.WhiteColor,
			white:           wplayer,
			black:           hive.NewPlayer(),
			board:           hive.NewBoard(),
			history:         []hive.Action{},
			paralyzedPieces: make(map[hive.Coordinate]int),
			features:        featureMap,
		}
		wp := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		if err := g.Place(wp, hive.NewCoordinate(0, 0, 0, 0)); err == nil {
			t.Errorf("Attempted to place a '%s' piece that the player didn't have available and no error was returned",
				wp)
		}
	})

	t.Run("When a player attempts to place a piece that is not their queen by the fourth turn an error is returned", func(t *testing.T) {
		g := &Game{
			turns:           4,
			turn:            hive.WhiteColor,
			white:           hive.NewPlayer(),
			black:           hive.NewPlayer(),
			board:           hive.NewBoard(),
			history:         []hive.Action{},
			paralyzedPieces: make(map[hive.Coordinate]int),
			features:        featureMap,
		}
		wp := hive.NewPiece(hive.WhiteColor, hive.Ant, hive.PieceA)
		if err := g.Place(wp, hive.NewCoordinate(0, 0, 0, 0)); err == nil {
			t.Errorf("Attempted to place a %s piece on the fourth turn while not having placed a queen and didn't receive an error", wp)
		} else {
			t.Log("When placing a piece on the fourth turn while still having a queen to place the expected error is returned")
			if !errors.Is(err, ErrRuleMustPlaceQueen) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleMustPlaceQueen, err)
			}
		}
	})

	t.Run("When placing a piece above the surface and no pieces exist below an error is returned", func(t *testing.T) {
		g := &Game{
			turns:           3,
			turn:            hive.WhiteColor,
			white:           hive.NewPlayer(),
			black:           hive.NewPlayer(),
			board:           hive.NewBoard(),
			history:         []hive.Action{},
			paralyzedPieces: make(map[hive.Coordinate]int),
			features:        featureMap,
		}
		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(0, 0, 0, 1)
		if err := g.Place(p, coord); err == nil {
			t.Errorf("Attempted to place the first piece at %s and did not receive an error", coord)
		} else {
			t.Log("When a piece is placed floating above the surface of the board the expected error is returned")
			if !errors.Is(err, ErrRuleMustPlacePieceOnSurface) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleMustPlacePieceOnSurface, err)
			}
		}
	})

	t.Run("When placing a queen on the first turn while the tournaments feature flag is enabled an error is returned", func(t *testing.T) {
		g := New([]Feature{TournamentQueensRuleFeature})
		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(0, 0, 0, 0)
		if err := g.Place(p, coord); err == nil {
			t.Error("Attempted to place a queen on the first turn while the tournament rules were active and no error was returned")
		} else {
			t.Log("When placing a queen on the first turn while tournament rules are active the expected error is returned")
			if !errors.Is(err, ErrRuleMayNotPlaceQueenOnFirstTurn) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleMayNotPlaceQueenOnFirstTurn, err)
			}
		}
	})

	t.Run("When placing a piece after the first turn, placing a piece touching an opponents piece returns an error", func(t *testing.T) {
		g := New(nil)

		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(0, 0, 0, 0)
		if err := g.Place(p, coord); err != nil {
			t.Errorf("Unexpected error %#v while white was placing a piece", err)
		}

		p = hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
		coord = hive.NewCoordinate(1, -1, 0, 0)
		if err := g.Place(p, coord); err != nil {
			t.Errorf("Unexpected error %#v while black was placing a piece", err)
		}

		p = hive.NewPiece(hive.WhiteColor, hive.Ant, hive.PieceA)
		coord = hive.NewCoordinate(1, 0, -1, 0)
		if err := g.Place(p, coord); err == nil {
			t.Error("Expected an error while white was placing a piece that is adjacent to a black piece after the first turn")
			t.Fail()
		} else {
			t.Log("When white places a piece adjacent to a black piece after the first turn the expected error is returned")
			if !errors.Is(err, ErrRuleMayNotPlaceTouchingOpponentsPiece) {
				t.Errorf("Unexpected error received, expected %#v, instead received %#v",
					ErrRuleMayNotPlaceTouchingOpponentsPiece, err)
			}
		}
	})
}

func TestGame_Move(t *testing.T) {
	t.Run("When moving attempting to move a non-existing piece from an empty cell an error is returned", func(t *testing.T) {
		g := New(nil)

		if err := g.Move(hive.Origin, hive.Origin); err == nil {
			t.Error("Expected an error when attempting to move a piece that doesn't exist")
		} else {
			t.Log("When attempting to move a piece from an empty cell the expected error is returned")
			if !errors.Is(err, hive.ErrInvalidCoordinate) {
				t.Errorf("Expected an error of type %#v instead received %#v", hive.ErrInvalidCoordinate, err)
			}
		}
	})

	t.Run("When attempting to move a piece and the destination coordinate is the same as the source an error is returned", func(t *testing.T) {
		g := New(nil)
		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		if err := g.Place(p, hive.Origin); err != nil {
			t.Errorf("Unexpected error %#v while white was placing a piece", err)
		}

		p = hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(1, -1, 0, 0)
		if err := g.Place(p, coord); err != nil {
			t.Errorf("Unexpected error %#v while black was placing a piece", err)
		}

		if err := g.Move(hive.Origin, hive.Origin); err == nil {
			t.Error("When attempting to move a piece from the origin to the origin an error was expected")
		} else {
			t.Log("When attempting to move a piece from origin to the origin the expected error is returned")
			if !errors.Is(err, hive.ErrInvalidCoordinate) {
				t.Errorf("Expected an error of type %#v instead received %#v", hive.ErrInvalidCoordinate, err)
			}
		}
	})

	t.Run("When attempting to move an opponents piece an error is returned", func(t *testing.T) {
		g := New(nil)

		p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
		if err := g.Place(p, hive.Origin); err != nil {
			t.Errorf("Unexpected error %#v while white was placing a piece", err)
		}

		p = hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(1, -1, 0, 0)
		if err := g.Place(p, coord); err != nil {
			t.Errorf("Unexpected error %#v while black was placing a piece", err)
		}

		if err := g.Move(coord, hive.NewCoordinate(0, -1, 1, 0)); err == nil {
			t.Error("Expected an error to be returned when white attempted to move one of blacks pieces")
		} else {
			t.Log("When attempting to move an opponents piece the expected error is returned")
			if !errors.Is(err, ErrRuleNotPlayersTurn) {
				t.Errorf("Expected an error of type %#v instead received %#v", ErrRuleNotPlayersTurn, err)
			}
		}
	})

	t.Run("When attempting to move a piece without having placed their queen an error is returned", func(t *testing.T) {
		g := New(nil)

		p := hive.NewPiece(hive.WhiteColor, hive.Ant, hive.PieceA)
		if err := g.Place(p, hive.Origin); err != nil {
			t.Errorf("Unexpected error %#v while white was placing a piece", err)
		}

		p = hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
		coord := hive.NewCoordinate(1, -1, 0, 0)
		if err := g.Place(p, coord); err != nil {
			t.Errorf("Unexpected error %#v while black was placing a piece", err)
		}

		if err := g.Move(hive.Origin, hive.NewCoordinate(0, -1, 1, 0)); err == nil {
			t.Errorf("Expected an error to be returned when white attempted to move a piece without placing their queen")
		} else {
			t.Log("When attempting to move a piece without first having placed their queen the expected error is returned")
			if !errors.Is(err, ErrRuleMustPlaceQueenToMove) {
				t.Errorf("Expected an error of type %#v instead received %#v", ErrRuleMustPlaceQueenToMove, err)
			}
		}
	})

	// TODO When attempting to move a piece that is pinned an error is returned
	// TODO When attempting to move a piece that is paralyzed an error is returned
	// TODO When attempting to move a piece that would split the hive an error is returned
	// TODO When attempting to move a piece not following the piece's pathing rules an error is returned

}

func TestGame_History(t *testing.T) {

	// TODO come back and update the test after move is tested.
	t.Log("When requesting a copy of the game instance history all expected actions are returned")
	g := New(nil)

	p := hive.NewPiece(hive.WhiteColor, hive.Queen, hive.PieceA)
	coord := hive.NewCoordinate(0, 0, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Errorf("Unexpected error %#v while white was placing a piece", err)
	}

	p = hive.NewPiece(hive.BlackColor, hive.Queen, hive.PieceA)
	coord = hive.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Errorf("Unexpected error %#v while black was placing a piece", err)
	}

	actions := g.History()

	if len(actions) != 2 {
		t.Errorf("Expected there to be 2 actions instead received %d", len(actions))
	}
}
