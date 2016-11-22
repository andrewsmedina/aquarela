package image

type rect struct {
	Left   int
	Right  int
	Bottom int
	Top    int
}

func (r *rect) Move(xy []int) *rect {
	return &rect{}
}

type img struct{}

func (i *img) Rect() *rect {
	return &rect{}
}

// Load load an image
func Load(src string) *img {
	file, _ := os.Open(src)
	i, _, _ := image.Decode(file)
	rgba := image.NewRGBA(i.Bounds())
	draw.Draw(rgba, rgba.Bounds(), i, image.Point{0, 0}, draw.Src)
	return &img{}
}
