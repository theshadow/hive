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
	t.Run("A new player should start with 3 ants", func(t *testing.T) {
		p := NewPlayer()
		if cnt := p.Ants(); cnt != 3 {
			t.Errorf("Expected a new player to have %d ants instead found %d.", 3, cnt)
		}
	})

	t.Run("When taking an ant from a player with ants no error should be returned.", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeAnAnt(); err != nil {
			t.Errorf("Unexpected error %#v while taking an ant from a new player.", err)
		} else if cnt := p.Ants(); cnt != 2 {
			t.Errorf("Expected there to be %d ants after taking one from a new player, instead received %d.",
				2, cnt)
		}
	})

	t.Run("When taking an ant from a player without any ants an error should be returned.", func(t *testing.T) {
		p := NewPlayer()

		// Take three ants
		_ = p.TakeAnAnt()
		_ = p.TakeAnAnt()
		_ = p.TakeAnAnt()

		if err := p.TakeAnAnt(); err == nil {
			t.Error("Expected an error while taking an ant from a player without any ants")
		}
	})
}

func TestPlayer_TakeAGrasshopper(t *testing.T) {
	t.Run("A new player should start with 3 grasshoppers", func(t *testing.T) {
		p := NewPlayer()
		if cnt := p.Grasshoppers(); cnt != 3 {
			t.Errorf("Expected a new player to have %d grasshoppers instead found %d.", 3, cnt)
		}
	})

	t.Run("When taking a grasshopper from a player without any left an error is returned", func(t *testing.T) {
		p := NewPlayer()
		_ = p.TakeAGrasshopper()
		_ = p.TakeAGrasshopper()
		_ = p.TakeAGrasshopper()
		if err := p.TakeAGrasshopper(); err == nil {
			t.Error("expected an error to be returned and instead received nil")
		}
	})
}

func TestPlayer_TakeASpider(t *testing.T) {
	t.Run("A new player should start with 2 spiders", func(t *testing.T) {
		p := NewPlayer()
		if cnt := p.Spiders(); cnt != 2 {
			t.Errorf("Expected a new player to have %d spider instead found %d.", 2, cnt)
		}

	})

	t.Run("When taking an spider from a new player with spiders no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeASpider(); err != nil {
			t.Errorf("Unexpected error %#v while taking an spider from a new player.", err)
		} else if cnt := p.Spiders(); cnt != 1 {
			t.Errorf("Expected there to be %d spider after taking one from a new player, instead received %d.",
				1, cnt)
		}
	})

	t.Run("When taking an spider from a player without any spiders an error is returned", func(t *testing.T) {
		p := NewPlayer()
		_ = p.TakeASpider()
		_ = p.TakeASpider()
		if err := p.TakeASpider(); err == nil {
			t.Error("expected an error instead received nill")
		}
	})
}

func TestPlayer_TakeABeetle(t *testing.T) {
	t.Run("A new player should start with 2 beetles", func(t *testing.T) {
		p := NewPlayer()
		if cnt := p.Beetles(); cnt != 2 {
			t.Errorf("Expected a new player to have %d beetle instead found %d.", 2, cnt)
		}

	})

	t.Run("When taking an beetle from a new player with beetles no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeABeetle(); err != nil {
			t.Errorf("Unexpected error %#v while taking an beetle from a new player.", err)
		} else if cnt := p.Beetles(); cnt != 1 {
			t.Errorf("Expected there to be %d beetle after taking one from a new player, instead received %d.",
				1, cnt)
		}
	})

	t.Run("When taking an beetle from a player without any beetles an error is returned", func(t *testing.T) {
		p := NewPlayer()
		_ = p.TakeABeetle()
		_ = p.TakeABeetle()
		if err := p.TakeABeetle(); err == nil {
			t.Error("expected an error instead received nill")
		}
	})
}

func TestPlayer_TakeQueen(t *testing.T) {
	t.Run("A new player should start with a queen", func(t *testing.T) {
		p := NewPlayer()
		if exists := p.HasQueen(); !exists {
			t.Error("expected a new player to have a queen")
		}
	})

	t.Run("When taking an queen from a new player no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeQueen(); err != nil {
			t.Errorf("Unexpected error %#v while taking an queen from a new player.", err)
		} else if exists := p.HasQueen(); exists {
			t.Error("expected there to be no queen after taking one from a new player")
		}
	})

	t.Run("When taking an queen from a player without a queen an error should be returned.", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeQueen(); err != nil && !errors.Is(err, ErrNoPieceAvailable) {
			t.Errorf("exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
		}
	})
}

func TestPlayer_TakeLadybug(t *testing.T) {
	t.Run("A new player should start with a ladybug", func(t *testing.T) {
		p := NewPlayer()
		if exists := p.HasLadybug(); !exists {
			t.Error("expected a new player to have a ladybug")
		}
	})

	t.Run("When taking an ladybug from a new player no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeLadybug(); err != nil {
			t.Errorf("Unexpected error %#v while taking an ladybug from a new player.", err)
		} else if exists := p.HasLadybug(); exists {
			t.Error("expected there to be no ladybug after taking one from a new player")
		}
	})

	t.Run("When taking an ladybug from a player without a ladybug an error should be returned.", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeLadybug(); err != nil && !errors.Is(err, ErrNoPieceAvailable) {
			t.Errorf("exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
		}
	})
}

func TestPlayer_TakeMosquito(t *testing.T) {
	t.Run("A new player should start with a mosquito", func(t *testing.T) {
		p := NewPlayer()
		if exists := p.HasMosquito(); !exists {
			t.Error("expected a new player to have a mosquito")
		}
	})

	t.Run("When taking an mosquito from a new player no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeMosquito(); err != nil {
			t.Errorf("Unexpected error %#v while taking an mosquito from a new player.", err)
		} else if exists := p.HasMosquito(); exists {
			t.Error("expected there to be no mosquito after taking one from a new player")
		}
	})

	t.Run("When taking an mosquito from a player without a mosquito an error should be returned.", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakeMosquito(); err != nil && !errors.Is(err, ErrNoPieceAvailable) {
			t.Errorf("exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
		}
	})
}

func TestPlayer_TakePillBug(t *testing.T) {
	t.Run("A new player should start with a pillbug", func(t *testing.T) {
		p := NewPlayer()
		if exists := p.HasPillBug(); !exists {
			t.Error("expected a new player to have a pillbug")
		}
	})

	t.Run("When taking an pillbug from a new player no error should be returned", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakePillBug(); err != nil {
			t.Errorf("Unexpected error %#v while taking an pillbug from a new player.", err)
		} else if exists := p.HasPillBug(); exists {
			t.Error("expected there to be no pillbug after taking one from a new player")
		}
	})

	t.Run("When taking an pillbug from a player without a pillbug an error should be returned.", func(t *testing.T) {
		p := NewPlayer()
		if err := p.TakePillBug(); err != nil && !errors.Is(err, ErrNoPieceAvailable) {
			t.Errorf("exepceted an error of type %#v instead received %#v", ErrNoPieceAvailable, err)
		}
	})
}
