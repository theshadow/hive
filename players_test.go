package hived

import (
	"testing"
)

func TestPlayer_HasZeroPieces(t *testing.T) {
	var player Player
	if !player.HasZeroPieces() {
		t.Log("a zero-Location player has pieces")
		t.Fail()
	}

	player = NewPlayer()
	if player.HasZeroPieces() {
		t.Log(" a new player has zero pieces")
		t.Fail()
	}
}

func TestPlayer_Ants(t *testing.T) {
	pOrig := NewPlayer()

	if pOrig.Ants() != 3 {
		t.Logf("Player: %16b Ants: %d", pOrig, pOrig.Ants())
		t.Log("a new pOrig doesn't have 3 ants")
		t.Fail()
	}

	pNew, err := pOrig.TakeAnAnt()
	if err != nil {
		t.Logf("unable to take AntA: %s", err)
		t.Fail()
	} else if pNew.Ants() != 2 {
		t.Logf("Before: %16b After: %16b Ants: %d", pOrig, pNew, pNew.Ants())
		t.Log("after taking an ant from a new pOrig there wasn't two ants left")
		t.Fail()
	}

	pOrig = pNew
	pNew, err = pOrig.TakeAnAnt()
	if err != nil {
		t.Logf("unable to take AntB: %s", err)
		t.Fail()
	} else if pNew.Ants() != 1 {
		t.Logf("Before: %16b After: %16b Ants: %d", pOrig, pNew, pNew.Ants())
		t.Log("after taking an ant there wasn't one ant left")
		t.Fail()
	}

	pOrig = pNew
	pNew, err = pOrig.TakeAnAnt()
	if err != nil {
		t.Logf("unable to take AntC: %s", err)
		t.Fail()
	} else if pNew.Ants() != 0 {
		t.Logf("Before: %16b After: %16b Ants: %d", pOrig, pNew, pNew.Ants())
		t.Log("after taking an ant there wasn't zero ants left")
		t.Fail()
	}

	pOrig = pNew
	pNew, err = pOrig.TakeAnAnt()
	if err == nil {
		t.Logf("Before: %16b After: %16b Ants: %d", pOrig, pNew, pNew.Ants())
		t.Log("expected an error when trying to take an ant with zero remaining")
		t.Fail()
	}
}
