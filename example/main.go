package main

import (
	"github.com/andrewsmedina/aquarela/game"
	"github.com/andrewsmedina/aquarela/image"
)

func main() {
	width := 320
	height := 240
	size := []int{320, 240}

	opts := &game.Options{
		Caption: "Pong",
		Size:    size,
	}
	g := game.New(opts)

	speed := []int{2, 2}

	ball := image.Load("ball.bmp")
	ballRect := ball.Rect()

	for {
		// game loop
		ballRect = ballRect.Move(speed)

		if ballRect.Left < 0 || ballRect.Right > width {
			speed[0] = -speed[0]
		}

		if ballRect.Top < 0 || ballRect.Bottom > height {
			speed[1] = -speed[1]
		}

		g.Draw()
	}
}
