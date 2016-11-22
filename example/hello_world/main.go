package main

import (
	"github.com/andrewsmedina/aquarela/director"
	"github.com/andrewsmedina/aquarela/label"
	"github.com/andrewsmedina/aquarela/layer"
	"github.com/andrewsmedina/aquarela/scene"
)

func main() {
	l := label.New("Hello, World!", []int{320, 240})

	helloLayer := layer.New()
	helloLayer.Add(l)

	mainScene := scene.New(helloLayer)

	d := director.New()
	d.Run(mainScene)
}
