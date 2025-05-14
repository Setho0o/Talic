package ui

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
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
		ops op.Ops
		b1 widget.Clickable
		b2 widget.Clickable

		maroon color.NRGBA = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	)
			
	theme := material.NewTheme()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			
			gtx := app.NewContext(&ops, e) // graphics ctx
			layout.Flex{
				Axis: layout.Horizontal,
				Spacing: layout.SpaceBetween,

			}.Layout(
				gtx,

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						sb := material.Button(theme, &b1,"start button")
						return sb.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						sb := material.Button(theme, &b2,"start button")
						return sb.Layout(gtx)
					},
				),
			)

			title := material.H1(theme, "Hello, Gio")
			title.Alignment = text.Middle
			title.Color = maroon
			

			

			title.Layout(gtx)
			e.Frame(gtx.Ops) // pass ops to gpu
		}
	}
}

