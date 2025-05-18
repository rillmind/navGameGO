package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rillmind/navGameGO/assets"
)

type Player struct {
	image    *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:    image,
		position: position,
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}
