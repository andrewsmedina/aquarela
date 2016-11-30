package layer

import (
	"golang.org/x/mobile/gl"
)

type drawer interface {
	Draw(gl.Context)
}

type Layer struct {
	d drawer
}

func New() *Layer {
	return &Layer{}
}

func (l *Layer) Add(d drawer) {
	l.d = d
}

func (l *Layer) Draw(ctx gl.Context) {
	l.d.Draw(ctx)
}
