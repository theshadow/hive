package hive

import (
	"errors"
	"testing"
)

func TestPlayer_NewPlayer(t *testing.T) {
	player := NewPlayer()
	if player.HasZeroPieces() {
		t.Logf("a new player should not have zero pieces: %s", player)
		t.Fail()
	}
}

func TestPlayer_TakeAnAnt(t *testing.T) {
	t.Log("A new player should start with 3 ants")
	p := NewPlayer()
	if cnt := p.Ants(); cnt != 3 {
		t.Logf("Expected a new player to have %d ants instead found %d.", 3, cnt)
		t.Fail()
	}

	t.Log("When taking an ant from a new player no error should be returned.")
	if err := p.TakeAnAnt(); err != nil {
		t.Logf("Unexpected error %#v while taking an ant from a new player.", err)
		t.Fail()
	} else if cnt := p.Ants(); cnt != 2 {
		t.Logf("Expected there to be %d ants after taking one from a new player, instead received %d.",
			2, cnt)
		t.Fail()
	}

	t.Logf("When taking an ant from a player with %d ants no error should be returned.", 2)
	if err := p.TakeAnAnt(); err != nil {
		t.Logf("Unexpected error %#v while taking an ant from a player with %d ants.", err, 2)
		t.Fail()
	} else if cnt := p.Ants(); cnt != 1 {
		t.Logf("Expected there to be %d ants after taking one from a new player, instead received %d.",
			1, cnt)
		t.Fail()
	}

	t.Logf("When taking an ant from a player with %d ants no error should be returned.", 1)
	if err := p.TakeAnAnt(); err != nil {
		t.Logf("Unexpected error %#v while taking an ant from a player with %d ant.", err, 1)
		t.Fail()
	} else if cnt := p.Ants(); cnt != 0 {
		t.Logf("Expected there to be %d ants after taking one from a new player, instead received %d.",
			0, cnt)
		t.Fail()
	}

	t.Logf("When taking an ant from a player with %d ants an error should be returned.", 0)
	if err := p.TakeAnAnt(); err == nil {
		t.Logf("Expected an error %#v while taking an ant from a player with %d ant.", err, 0)
		t.Fail()
	}
}

func TestPlayer_TakeAGrasshopper(t *testing.T) {
	t.Log("A new player should start with 3 grasshoppers")
	p := NewPlayer()
	if cnt := p.Grasshoppers(); cnt != 3 {
		t.Logf("Expected a new player to have %d grasshoppers instead found %d.", 3, cnt)
		t.Fail()
	}

	t.Log("When taking an grasshopper from a new player no error should be returned.")
	if err := p.TakeAGrasshopper(); err != nil {
		t.Logf("Unexpected error %#v while taking an grasshopper from a new player.", err)
		t.Fail()
	} else if cnt := p.Grasshoppers(); cnt != 2 {
		t.Logf("Expected there to be %d grasshoppers after taking one from a new player, instead received %d.",
			2, cnt)
		t.Fail()
	}

	t.Logf("When taking an grasshopper from a player with %d grasshoppers no error should be returned.", 2)
	if err := p.TakeAGrasshopper(); err != nil {
		t.Logf("Unexpected error %#v while taking an grasshopper from a player with %d grasshoppers.", err, 2)
		t.Fail()
	} else if cnt := p.Grasshoppers(); cnt != 1 {
		t.Logf("Expected there to be %d grasshoppers after taking one from a new player, instead received %d.",
			1, cnt)
		t.Fail()
	}

	t.Logf("When taking an grasshopper from a player with %d grasshoppers no error should be returned.", 1)
	if err := p.TakeAGrasshopper(); err != nil {
		t.Logf("Unexpected error %#v while taking an grasshopper from a player with %d grasshopper.", err, 1)
		t.Fail()
	} else if cnt := p.Grasshoppers(); cnt != 0 {
		t.Logf("Expected there to be %d grasshoppers after taking one from a player, instead received %d.",
			0, cnt)
		t.Fail()
	}

	t.Logf("When taking an grasshopper from a player with %d grasshoppers an error should be returned.", 0)
	if err := p.TakeAGrasshopper(); err == nil {
		t.Logf("Expected an error %#v while taking an grasshopper from a player with %d grasshopper.", err, 0)
		t.Fail()
	}
}

func TestPlayer_TakeASpider(t *testing.T) {
	t.Log("A new player should start with 3 spider")
	p := NewPlayer()
	if cnt := p.Spiders(); cnt != 2 {
		t.Logf("Expected a new player to have %d spider instead found %d.", 2, cnt)
		t.Fail()
	}

	t.Log("When taking an spider from a new player no error should be returned.")
	if err := p.TakeASpider(); err != nil {
		t.Logf("Unexpected error %#v while taking an spider from a new player.", err)
		t.Fail()
	} else if cnt := p.Spiders(); cnt != 1 {
		t.Logf("Expected there to be %d spider after taking one from a new player, instead received %d.",
			1, cnt)
		t.Fail()
	}

	t.Logf("When taking an spider from a player with %d spider no error should be returned.", 1)
	if err := p.TakeASpider(); err != nil {
		t.Logf("Unexpected error %#v while taking an spider from a player with %d spider.", err, 1)
		t.Fail()
	} else if cnt := p.Spiders(); cnt != 0 {
		t.Logf("Expected there to be %d spider after taking one from a player, instead received %d.",
			0, cnt)
		t.Fail()
	}

	t.Logf("When taking an spider from a player with %d spider an error should be returned.", 0)
	if err := p.TakeASpider(); err == nil {
		t.Logf("Expected an error %#v while taking an spider from a player with %d spider.", err, 0)
		t.Fail()
	}
}

func TestPlayer_TakeABeetle(t *testing.T) {
	t.Log("A new player should start with 3 beetle")
	p := NewPlayer()
	if cnt := p.Beetles(); cnt != 2 {
		t.Logf("Expected a new player to have %d beetle instead found %d.", 2, cnt)
		t.Fail()
	}

	t.Log("When taking an beetle from a new player no error should be returned.")
	if err := p.TakeABeetle(); err != nil {
		t.Logf("Unexpected error %#v while taking an beetle from a new player.", err)
		t.Fail()
	} else if cnt := p.Beetles(); cnt != 1 {
		t.Logf("Expected there to be %d beetle after taking one from a new player, instead received %d.",
			1, cnt)
		t.Fail()
	}

	t.Logf("When taking an beetle from a player with %d beetle no error should be returned.", 1)
	if err := p.TakeABeetle(); err != nil {
		t.Logf("Unexpected error %#v while taking an beetle from a player with %d beetle.", err, 1)
		t.Fail()
	} else if cnt := p.Beetles(); cnt != 0 {
		t.Logf("Expected there to be %d beetle after taking one from a player, instead received %d.",
			0, cnt)
		t.Fail()
	}

	t.Logf("When taking an beetle from a player with %d beetle an error should be returned.", 0)
	if err := p.TakeABeetle(); err == nil {
		t.Logf("Expected an error %#v while taking an beetle from a player with %d beetle.", err, 0)
		t.Fail()
	}
}

func TestPlayer_TakeQueen(t *testing.T) {
	t.Log("A new player should start with a queen")
	p := NewPlayer()
	if exists := p.HasQueen(); !exists {
		t.Log("Expected a new player to have a queen.")
		t.Fail()
	}

	t.Log("When taking an queen from a new player no error should be returned.")
	if err := p.TakeQueen(); err != nil {
		t.Logf("Unexpected error %#v while taking an queen from a new player.", err)
		t.Fail()
	} else if exists := p.HasQueen(); exists {
		t.Log("Expected there to be no queen after taking one from a new player.")
		t.Fail()
	}

	t.Log("When taking an queen from a player without a queen an error should be returned.")
	if err := p.TakeQueen(); p.HasQueen() || err == nil {
		t.Log("Expected an error while taking an queen from a player without a queen.")
		t.Fail()
	} else if !errors.Is(err, ErrNoPieceAvailable) {
		t.Logf("Exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
	}
}

func TestPlayer_TakeLadybug(t *testing.T) {
	t.Log("A new player should start with a ladybug")
	p := NewPlayer()
	if exists := p.HasLadybug(); !exists {
		t.Log("Expected a new player to have a ladybug.")
		t.Fail()
	}

	t.Log("When taking an ladybug from a new player no error should be returned.")
	if err := p.TakeLadybug(); err != nil {
		t.Logf("Unexpected error %#v while taking an ladybug from a new player.", err)
		t.Fail()
	} else if exists := p.HasLadybug(); exists {
		t.Log("Expected there to be no ladybug after taking one from a new player.")
		t.Fail()
	}

	t.Log("When taking an ladybug from a player without a ladybug an error should be returned.")
	if err := p.TakeLadybug(); p.HasLadybug() || err == nil {
		t.Log("Expected an error while taking an ladybug from a player without a ladybug.")
		t.Fail()
	} else if !errors.Is(err, ErrNoPieceAvailable) {
		t.Logf("Exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
	}
}

func TestPlayer_TakeMosquito(t *testing.T) {
	t.Log("A new player should start with a mosquito")
	p := NewPlayer()
	if exists := p.HasMosquito(); !exists {
		t.Log("Expected a new player to have a mosquito.")
		t.Fail()
	}

	t.Log("When taking an mosquito from a new player no error should be returned.")
	if err := p.TakeMosquito(); err != nil {
		t.Logf("Unexpected error %#v while taking an mosquito from a new player.", err)
		t.Fail()
	} else if exists := p.HasMosquito(); exists {
		t.Log("Expected there to be no mosquito after taking one from a new player.")
		t.Fail()
	}

	t.Log("When taking an mosquito from a player without a mosquito an error should be returned.")
	if err := p.TakeMosquito(); p.HasMosquito() || err == nil {
		t.Log("Expected an error while taking an mosquito from a player without a mosquito.")
		t.Fail()
	} else if !errors.Is(err, ErrNoPieceAvailable) {
		t.Logf("Exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
	}
}

func TestPlayer_TakePillBug(t *testing.T) {
	t.Log("A new player should start with a pill bug")
	p := NewPlayer()
	if exists := p.HasPillBug(); !exists {
		t.Log("Expected a new player to have a pill bug.")
		t.Fail()
	}

	t.Log("When taking an pill bug from a new player no error should be returned.")
	if err := p.TakePillBug(); err != nil {
		t.Logf("Unexpected error %#v while taking an pill bug from a new player.", err)
		t.Fail()
	} else if exists := p.HasPillBug(); exists {
		t.Log("Expected there to be no pill bug after taking one from a new player.")
		t.Fail()
	}

	t.Log("When taking an pill bug from a player without a pill bug an error should be returned.")
	if err := p.TakePillBug(); p.HasPillBug() || err == nil {
		t.Log("Expected an error while taking an pill bug from a player without a pill bug.")
		t.Fail()
	} else if !errors.Is(err, ErrNoPieceAvailable) {
		t.Logf("Exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
	}
}
