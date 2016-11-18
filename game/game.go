package game

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Options represents a game options
type Options struct {
	Size    []int
	Caption string
}

type game struct{}

// New returns a new game
func New(opts *Options) *game {
	glfw.Init()
	defer glfw.Terminate()
	window, _ := glfw.CreateWindow(opts.Size[0], opts.Size[1], opts.Caption, nil, nil)
	window.MakeContextCurrent()
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
	return &game{}
}

func (*game) Draw() {}
