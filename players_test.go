package hived

import (
	"testing"
)

func TestPlayer_NewPlayer(t *testing.T) {
	player := NewPlayer()
	if player.HasZeroPieces() {
		t.Logf("a new player should not have zero pieces: %s", player)
		t.Fail()
	}
}

func TestPlayer_TakeInterface(t *testing.T) {
	cases := []struct {
		name    string
		count   int
		countFn func(p *Player) (count int)
		takeFn  func(p *Player) error
	}{
		{
			"Ants",
			3,
			func(p *Player) (count int) {
				return p.Ants()
			},
			func(p *Player) error {
				return p.TakeAnAnt()
			},
		},
		{
			"Grasshoppers",
			3,
			func(p *Player) (count int) {
				return p.Grasshoppers()
			},
			func(p *Player) error {
				return p.TakeAGrasshopper()
			},
		},
		{
			"Spiders",
			2,
			func(p *Player) (count int) {
				return p.Spiders()
			},
			func(p *Player) error {
				return p.TakeASpider()
			},
		},
		{
			"Beetles",
			2,
			func(p *Player) (count int) {
				return p.Beetles()
			},
			func(p *Player) error {
				return p.TakeABeetle()
			},
		},
		{
			"Queen",
			1,
			func(p *Player) (count int) {
				if p.HasQueen() {
					return 1
				}
				return 0
			},
			func(p *Player) error {
				return p.TakeQueen()
			},
		},
		{
			"Ladybug",
			1,
			func(p *Player) (count int) {
				if p.HasLadybug() {
					return 1
				}
				return 0
			},
			func(p *Player) error {
				return p.TakeLadybug()
			},
		},
		{
			"Mosquito",
			1,
			func(p *Player) (count int) {
				if p.HasMosquito() {
					return 1
				}
				return 0
			},
			func(p *Player) error {
				return p.TakeMosquito()
			},
		},
		{
			"PillBug",
			1,
			func(p *Player) (count int) {
				if p.HasPillBug() {
					return 1
				}
				return 0
			},
			func(p *Player) error {
				return p.TakePillBug()
			},
		},
	}

	for _, test := range cases {
		player := NewPlayer()

		for cnt := test.count; cnt >= 0; cnt-- {
			if cnt == test.count && test.countFn(player) != cnt {
				t.Logf("Player: %16b %s: %d", player, test.name, test.countFn(player))
				t.Logf("a new player doesn't have %d %s", test.count, test.name)
				t.Fail()
			}

			if err := test.takeFn(player); err != nil && test.countFn(player) > 0 {
				t.Logf("with %s there is expected to be %d pieces, receieved an error when attempting to take the %d piece",
					test.name, test.count, cnt)
				t.Fail()
			} else if err == nil && cnt == 0 && test.countFn(player) == 0 {
				t.Logf("expected an error when attempting to take %s when there were %d left", test.name, cnt)
				t.Fail()
			}
		}
	}
}
