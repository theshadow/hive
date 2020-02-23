package hived

import (
	"fmt"
	"testing"
)

func TestPlayer_HasZeroPieces(t *testing.T) {
	var player Player
	if !player.HasZeroPieces() {
		t.Log("a zero-value player has pieces")
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

func TestPlayer_Grasshoppers(t *testing.T) {
	tests := []struct {
		bug   string
		count int
	}{
		{bug: "Ant", count: 3},
		{bug: "Grasshopper", count: 3},
		{bug: "Beetle", count: 2},
		{bug: "Spider", count: 2},
	}

	fnCount := func(p Player, bug string) int {
		if bug == "Ant" {
			return p.Ants()
		} else if bug == "Grasshopper" {
			return p.Grasshoppers()
		} else if bug == "Beetle" {
			return p.Beetles()
		} else if bug == "Spider" {
			return p.Spiders()
		}
		return 0
	}

	fnTake := func(p Player, bug string) (Player, error) {
		if bug == "Ant" {
			return p.TakeAnAnt()
		} else if bug == "Grasshopper" {
			return p.TakeAGrasshopper()
		} else if bug == "Beetle" {
			return p.TakeABeetle()
		} else if bug == "Spider" {
			return p.TakeASpider()
		}
		return ZeroPlayer, fmt.Errorf("undefined bug handler %s", bug)
	}

	for _, test := range tests {
		pOrig := NewPlayer()

		if fnCount(pOrig, test.bug) != test.count {
			t.Logf("Player: %16b %ss: %d", pOrig, test.bug, fnCount(pOrig, test.bug))
			t.Logf("a new player doesn't have 3 %ss", test.bug)
			t.Fail()
		}

		for i := test.count; i > 0; i-- {
			pNew, err := fnTake(pOrig, test.bug)
			if err != nil {
				t.Logf("unable to take %sA: %s", test.bug, err)
				t.Fail()
			} else if fnCount(pNew, test.bug) != i-1 {
				t.Logf("Before: %16b After: %16b %s: %d", pOrig, pNew, test.bug, fnCount(pNew, test.bug))
				t.Logf("after taking an %s from a player there wasn't %d %s(s) left", test.bug, i-1, test.bug)
				t.Fail()
			}
			pOrig = pNew
		}

		_, err := fnTake(pOrig, test.bug)
		if err == nil {
			t.Logf("Before: %16b %s: %d", pOrig, test.bug, fnCount(pOrig, test.bug))
			t.Logf("expected an error when trying to take an %s with zero remaining", test.bug)
			t.Fail()
		}
	}
}
