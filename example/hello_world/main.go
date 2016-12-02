package main

import (
	"github.com/andrewsmedina/aquarela/director"
	"github.com/andrewsmedina/aquarela/layer"
	"github.com/andrewsmedina/aquarela/scene"
	"github.com/andrewsmedina/aquarela/sprite"
)

func main() {
	s := sprite.New("image.png", []int{320, 240})

	helloLayer := layer.New()
	helloLayer.Add(s)

	mainScene := scene.New(helloLayer)

	d := director.New()
	d.Run(mainScene)
}
