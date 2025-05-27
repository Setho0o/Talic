package ui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func Gui() {
	go func() {
		w := new(app.Window)
		w.Option(
			//app.Fullscreen.Option(),
			app.Maximized.Option(),
			app.Title("Talic"),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	var (
		ops            op.Ops
		FlexBoxButtons = [6]Button{}
	)

	theme := material.NewTheme()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e) // graphics ctx
			m := gtx.Constraints.Max

			layout.Stack{
				Alignment: layout.S,
			}.Layout(gtx,
				layout.Expanded(
					func(gtx layout.Context) layout.Dimensions {
						return DrawSquare(gtx, m, Bg)
					},
				),
				layout.Stacked(
					func(gtx layout.Context) layout.Dimensions {
						//return DrawSquare(gtx,DpPt(float64(m.X),200),blue)
						return FlexBox(gtx, &FlexBoxButtons, theme)
					},
				),
			)
			e.Frame(gtx.Ops) // pass ops to gpu
		}
	}
}
