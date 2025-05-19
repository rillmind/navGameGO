package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player  *Player
	lasers  []*Laser
	meteors []*Meteor

	meteorSpawTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawTimer: NewTimer(24),
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
			}
		}
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
	g.meteorSpawTimer.Reset()
}
