package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rillmind/navGameGO/assets"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]
	speed := (rand.Float64() * 10)

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.X, s.position.Y)
	screen.DrawImage(s.image, op)
}
