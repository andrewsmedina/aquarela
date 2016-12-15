package director

import (
	"fmt"
	"time"

	"github.com/andrewsmedina/aquarela/scene"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

type director struct {
	scene  *scene.Scene
	images *glutil.Images
	eng    sprite.Engine
	n *sprite.Node
	startTime time.Time
}

func New() *director {
	return &director{startTime: time.Now()}
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

func (d *director) draw(ctx gl.Context, sz size.Event) {
	ctx.ClearColor(0, 0, 0, 1)
	ctx.Clear(gl.COLOR_BUFFER_BIT)
	d.scene.Draw(ctx)
	now := clock.Time(time.Since(d.startTime) * 60 / time.Second)
	d.eng.Render(d.n, now, sz)
}

func (d *director) loop() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
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
			case size.Event:
				sz = e
			case paint.Event:
				fmt.Println("paint")
				d.draw(glctx, sz)
				a.Publish()
				a.Send(paint.Event{})
			}
		}
	})
}
