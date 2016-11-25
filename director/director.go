package director

import (
	"fmt"

	"github.com/andrewsmedina/aquarela/scene"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

type director struct {
	scene  *scene.Scene
	images *glutil.Images
	eng    sprite.Engine
}

func New() *director {
	return &director{}
}

func (d *director) Run(s *scene.Scene) {
	d.scene = s
	d.loop()
}

func (d *director) on(ctx gl.Context) {
	d.images = glutil.NewImages(ctx)
	d.eng = glsprite.Engine(d.images)
}

func (d *director) off() {
	d.eng.Release()
	d.images.Release()
}

func (d *director) draw(ctx gl.Context) {
	ctx.ClearColor(0, 0, 0, 1)
	ctx.Clear(gl.COLOR_BUFFER_BIT)
	d.scene.Draw(ctx)
}

func (d *director) loop() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					d.on(glctx)
					a.Send(paint.Event{})
					fmt.Println("on")
				case lifecycle.CrossOff:
					d.off()
					glctx = nil
					fmt.Println("off")
				}
			case paint.Event:
				fmt.Println("paint")
				d.draw(glctx)
				a.Publish()
				a.Send(paint.Event{})
			}
		}
	})
}
