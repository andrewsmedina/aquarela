package layer

import (
	"github.com/andrewsmedina/aquarela/label"
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
