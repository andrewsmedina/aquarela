package label

import (
	"golang.org/x/mobile/gl"
)

type Label struct {
	text     string
	position []int
}

func New(text string, position []int) *Label {
	return &Label{text: text, position: position}
}

func (l *Label) Draw(ctx gl.Context) {}
