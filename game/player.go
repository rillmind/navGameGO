package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rillmind/navGameGO/assets"
)

type Player struct {
	image            *ebiten.Image
	position         Vector
	game             *Game
	laserLoadingTime *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:            image,
		position:         position,
		game:             game,
		laserLoadingTime: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	p.laserLoadingTime.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTime.IsReady() {
		p.laserLoadingTime.Reset()

		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfY := float64(bounds.Dy()) / 2

		spawPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfY/2,
		}

		laser := NewLaser(spawPos)
		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}
