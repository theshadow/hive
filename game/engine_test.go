package game

import (
	"errors"
	"testing"

	"github.com/theshadow/hived"
)

func TestGame_Place(t *testing.T) {
	t.Log("When a player attempts to place a piece not on their turn an error is returned.")
	g := New(nil)
	p := hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord := hived.NewCoordinate(0, 0, 0, 0)
	if err := g.Place(p, coord); err == nil {
		t.Log("Attempted to place a black piece during the white players turn but didn't receive an error.")
		t.Fail()
	} else {
		t.Log("When placing black piece during the white players turn the expected error is returned")
		if !errors.Is(err, ErrRuleNotPlayersTurn) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleNotPlayersTurn, err)
			t.Fail()
		}
	}

	t.Log("When the first piece is not placed at origin an error is returned.")
	g = New(nil)
	p = hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(0, 0, 0, 1)
	if err := g.Place(p, coord); err == nil {
		t.Logf("Attempted to place the first piece at %s and did not receive an error.", coord)
		t.Fail()
	} else {
		t.Log("When first piece placed isn't at the origin and returned the expected error is returned.")
		if !errors.Is(err, ErrRuleFirstPieceMustBeAtOrigin) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleFirstPieceMustBeAtOrigin, err)
			t.Fail()
		}
	}

	t.Log("When a player places a piece and they do not have that piece available an error is returned.")
	wplayer := hived.NewPlayer()
	_ = wplayer.TakeQueen()
	g = &Game{
		turns:           1,
		turn:            hived.WhiteColor,
		white:           wplayer,
		black:           hived.NewPlayer(),
		board:           hived.NewBoard(),
		history:         []hived.Action{},
		paralyzedPieces: make(map[hived.Coordinate]int),
		features:        featureMap,
	}
	wp := hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	if err := g.Place(wp, hived.NewCoordinate(0, 0, 0, 0)); err == nil {
		t.Logf("Attempted to place a '%s' piece that the player didn't have available and no error was returned.",
			wp)
		t.Fail()
	}

	t.Log("When a player places their piece on the fourth turn and the player hasn't placed their queen and the piece being placed isn't a queen an error is returned.")
	g = &Game{
		turns:           4,
		turn:            hived.WhiteColor,
		white:           hived.NewPlayer(),
		black:           hived.NewPlayer(),
		board:           hived.NewBoard(),
		history:         []hived.Action{},
		paralyzedPieces: make(map[hived.Coordinate]int),
		features:        featureMap,
	}
	wp = hived.NewPiece(hived.WhiteColor, hived.Ant, hived.PieceA)
	if err := g.Place(wp, hived.NewCoordinate(0, 0, 0, 0)); err == nil {
		t.Logf("Attempted to place a %s piece on the fourth turn while not having placed a queen and didn't receive an error.",
			wp)
		t.Fail()
	} else {
		t.Log("When placing a piece on the fourth turn while still having a queen to place the expected error is returned.")
		if !errors.Is(err, ErrRuleMustPlaceQueen) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleMustPlaceQueen, err)
			t.Fail()
		}
	}

	t.Log("When placing a piece above the surface and no pieces exist below an error is returned.")
	g = &Game{
		turns:           3,
		turn:            hived.WhiteColor,
		white:           hived.NewPlayer(),
		black:           hived.NewPlayer(),
		board:           hived.NewBoard(),
		history:         []hived.Action{},
		paralyzedPieces: make(map[hived.Coordinate]int),
		features:        featureMap,
	}
	p = hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(0, 0, 0, 1)
	if err := g.Place(p, coord); err == nil {
		t.Logf("Attempted to place the first piece at %s and did not receive an error.", coord)
		t.Fail()
	} else {
		t.Log("When a piece is placed floating above the surface of the board the expected error is returned.")
		if !errors.Is(err, ErrRuleMustPlacePieceOnSurface) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleMustPlacePieceOnSurface, err)
			t.Fail()
		}
	}

	t.Log("When placing a queen on the first turn while the tournaments feature flag is enabled an error is returned.")
	g = New([]Feature{TournamentQueensRuleFeature})
	p = hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(0, 0, 0, 0)
	if err := g.Place(p, coord); err == nil {
		t.Log("Attempted to place a queen on the first turn while the tournament rules were active and no error was returned.")
		t.Fail()
	} else {
		t.Log("When placing a queen on the first turn while tournament rules are active the expected error is returned.")
		if !errors.Is(err, ErrRuleMayNotPlaceQueenOnFirstTurn) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleMayNotPlaceQueenOnFirstTurn, err)
			t.Fail()
		}
	}

	/////////////////////////////////////////////////////

	t.Log("When placing a piece after the first turn, placing a piece touching an opponents piece returns an error.")
	g = New(nil)

	p = hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(0, 0, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while white was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while black was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.WhiteColor, hived.Ant, hived.PieceA)
	coord = hived.NewCoordinate(1, 0, -1, 0)
	if err := g.Place(p, coord); err == nil {
		t.Log("Expected an error while white was placing a piece that is adjacent to a black piece after the first turn.")
		t.Fail()
	} else {
		t.Log("When white places a piece adjacent to a black piece after the first turn the expected error is returned.")
		if !errors.Is(err, ErrRuleMayNotPlaceTouchingOpponentsPiece) {
			t.Logf("Unexpected error received, expected %#v, instead received %#v.",
				ErrRuleMayNotPlaceTouchingOpponentsPiece, err)
			t.Fail()
		}
	}
}

func TestGame_Move(t *testing.T) {
	t.Log("When moving attempting to move a non-existing piece from an empty cell an error is returned.")
	g := New(nil)

	if err := g.Move(hived.Origin, hived.Origin); err == nil {
		t.Log("Expected an error when attempting to move a piece that doesn't exist.")
		t.Fail()
	} else {
		t.Log("When attempting to move a piece from an empty cell the expected error is returned.")
		if !errors.Is(err, hived.ErrInvalidCoordinate) {
			t.Logf("Expected an error of type %#v instead received %#v.", hived.ErrInvalidCoordinate, err)
			t.Fail()
		}
	}

	t.Log("When attempting to move a piece and the destination coordinate is the same as the source an error is returned.")
	g = New(nil)
	p := hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	if err := g.Place(p, hived.Origin); err != nil {
		t.Logf("Unexpected error %#v while white was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord := hived.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while black was placing a piece", err)
		t.Fail()
	}

	if err := g.Move(hived.Origin, hived.Origin); err == nil {
		t.Log("When attempting to move a piece from the origin to the origin an error was expected")
		t.Fail()
	} else {
		t.Log("When attempting to move a piece from origin to the origin the expected error is returned.")
		if !errors.Is(err, hived.ErrInvalidCoordinate) {
			t.Logf("Expected an error of type %#v instead received %#v.", hived.ErrInvalidCoordinate, err)
			t.Fail()
		}
	}

	t.Log("When attempting to move an opponents piece an error is returned.")
	g = New(nil)

	p = hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	if err := g.Place(p, hived.Origin); err != nil {
		t.Logf("Unexpected error %#v while white was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while black was placing a piece", err)
		t.Fail()
	}

	if err := g.Move(coord, hived.NewCoordinate(0, -1, 1, 0)); err == nil {
		t.Log("Expected an error to be returned when white attempted to move one of blacks pieces.")
		t.Fail()
	} else {
		t.Log("When attempting to move an opponents piece the expected error is returned.")
		if !errors.Is(err, ErrRuleNotPlayersTurn) {
			t.Logf("Expected an error of type %#v instead received %#v.", ErrRuleNotPlayersTurn, err)
			t.Fail()
		}
	}

	t.Log("When attempting to move a piece without having placed their queen an error is returned.")
	g = New(nil)

	p = hived.NewPiece(hived.WhiteColor, hived.Ant, hived.PieceA)
	if err := g.Place(p, hived.Origin); err != nil {
		t.Logf("Unexpected error %#v while white was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while black was placing a piece", err)
		t.Fail()
	}

	if err := g.Move(hived.Origin, hived.NewCoordinate(0, -1, 1, 0)); err == nil {
		t.Log("Expected an error to be returned when white attempted to move a piece without placing their queen.")
		t.Fail()
	} else {
		t.Log("When attempting to move a piece without first having placed their queen the expected error is returned.")
		if !errors.Is(err, ErrRuleMustPlaceQueenToMove) {
			t.Logf("Expected an error of type %#v instead received %#v.", ErrRuleMustPlaceQueenToMove, err)
			t.Fail()
		}
	}



	// TODO When attempting to move a piece that is pinned an error is returned
	// TODO When attempting to move a piece that is paralyzed an error is returned
	// TODO When attempting to move a piece that woul d split the hive an error is returned
	// TODO When attempting to move a piece not following the pieces pathing rules an error is returned

}

func TestGame_History(t *testing.T) {

	// TODO come back and update the test after move is tested.
	t.Log("When requesting a copy of the game instance history all expected actions are returned")
	g := New(nil)

	p := hived.NewPiece(hived.WhiteColor, hived.Queen, hived.PieceA)
	coord := hived.NewCoordinate(0, 0, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while white was placing a piece", err)
		t.Fail()
	}

	p = hived.NewPiece(hived.BlackColor, hived.Queen, hived.PieceA)
	coord = hived.NewCoordinate(1, -1, 0, 0)
	if err := g.Place(p, coord); err != nil {
		t.Logf("Unexpected error %#v while black was placing a piece", err)
		t.Fail()
	}

	actions := g.History()

	if len(actions) != 2 {
		t.Logf("Expected there to be 2 actions instead received %d", len(actions))
		t.Fail()
	}
}