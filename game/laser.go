package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rillmind/navGameGO/assets"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfY

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7

	l.position.Y += -float64(speed)
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
