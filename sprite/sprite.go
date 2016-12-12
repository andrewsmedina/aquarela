package sprite

import (
	"golang.org/x/mobile/gl"
)

// Sprite represents a sprite
type Sprite struct {
	image    string
	position []int
}

// New creates a new sprite
func New(image string, position []int) *Sprite {
	return &Sprite{image: image, position: position}
}

// Draw draws the sprite in the gl context
func (s *Sprite) Draw(ctx gl.Context) {}
