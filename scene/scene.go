package scene

import (
	"github.com/andrewsmedina/aquarela/layer"
	"golang.org/x/mobile/gl"
)

type Scene struct {
	layer *layer.Layer
}

func New(l *layer.Layer) *Scene {
	return &Scene{layer: l}
}

func (s *Scene) Draw(ctx gl.Context) {
	s.layer.Draw(ctx)
}
