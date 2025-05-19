package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rillmind/navGameGO/assets"
)

type Game struct {
	score int

	player  *Player
	lasers  []*Laser
	meteors []*Meteor
	stars   []*Star

	starSpawTimer   *Timer
	meteorSpawTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawTimer: NewTimer(24),
		starSpawTimer:   NewTimer(28),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

func (g *Game) Update() error {
	for _, l := range g.lasers {
		l.Update()
	}

	g.player.Update()

	g.meteorSpawTimer.Update()

	if g.meteorSpawTimer.IsReady() {
		g.meteorSpawTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	for i, m := range g.meteors {
		for ii, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:ii], g.lasers[ii+1:]...)

				g.score += 1
			}
		}
	}

	g.starSpawTimer.Update()

	if g.starSpawTimer.IsReady() {
		g.starSpawTimer.Reset()

		s := NewStar()
		g.stars = append(g.stars, s)
	}

	for _, s := range g.stars {
		s.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, s := range g.stars {
		s.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 20, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.stars = nil
	g.meteorSpawTimer.Reset()
	g.score = 0
}
