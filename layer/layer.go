package layer

import (
	"github.com/andrewsmedina/aquarela/label"
	"golang.org/x/mobile/gl"
)

type Layer struct {
	label *label.Label
}

func New() *Layer {
	return &Layer{}
}

func (l *Layer) Add(lb *label.Label) {
	l.label = lb
}

func (l *Layer) Draw(ctx gl.Context) {
	l.label.Draw(ctx)
}
