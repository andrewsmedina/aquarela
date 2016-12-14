package layer

import (
	"golang.org/x/mobile/gl"
)

type drawer interface {
	Draw(gl.Context)
}

// Layer represents a layer
type Layer struct {
	d drawer
}

// New returns a new Layer
func New() *Layer {
	return &Layer{}
}

// Add adds a drawer component to layer
func (l *Layer) Add(d drawer) {
	l.d = d
}

 // Draw draws layer elements in the given context
func (l *Layer) Draw(ctx gl.Context) {
	l.d.Draw(ctx)
}
