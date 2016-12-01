package sprite

import (
	"golang.org/x/mobile/gl"
)

type Sprite struct {
	image    string
	position []int
}

func New(image string, position []int) *Sprite {
	return &Sprite{image: image, position: position}
}

func (s *Sprite) Draw(ctx gl.Context) {}
