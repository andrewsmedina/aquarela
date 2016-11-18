package game

// Options represents a game options
type Options struct {
	Size    []int
	Caption string
}

type game struct{}

// New returns a new game
func New(opts *Options) *game {
	return &game{}
}

func (*game) Draw() {}
