package label

type Label struct {
	text     string
	position []int
}

func New(text string, position []int) *Label {
	return &Label{text: text, position: position}
}
