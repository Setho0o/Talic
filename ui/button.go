package ui

import (
	"bytes"
	"image"
	"log"
	"os"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/widget/material"
)

type Button struct {
	pressed bool
}

func (b *Button) Layout(gtx layout.Context, path string) layout.Dimensions {
	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)
	event.Op(gtx.Ops, b)

	for {
		ev, ok := gtx.Event(pointer.Filter{
			Target: b,
			Kinds:  pointer.Press | pointer.Release,
		})
		if !ok {
			break
		}
		e, ok := ev.(pointer.Event)
		if !ok {
			continue
		}
		switch e.Kind {
		case pointer.Press:
			b.pressed = true
		case pointer.Release:
			b.pressed = false
		}
	}
	area.Pop()

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err, "failed to open file play.png")
	}
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err, "failed to decode play.png into a image ")
	}
	return DrawImage(gtx.Ops, img)
}

func FlexBox(gtx layout.Context, fb *[6]Button, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis:      layout.Horizontal,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Start,
	}.Layout(
		gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[0].pressed {
					log.Print("back")
					fb[0].pressed = false
				}
				return fb[0].Layout(gtx, Back)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[1].pressed {
					log.Print("Pause/Play")
					fb[1].pressed = false
				}
				return fb[1].Layout(gtx, Pause)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[2].pressed {
					log.Print("next")
					fb[2].pressed = false
				}
				return fb[2].Layout(gtx, Next)
			},
		),
		layout.Flexed(
			.100,
			layout.Spacer{}.Layout,
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[3].pressed {
					log.Print("shuffle")
					fb[3].pressed = false
				}
				return fb[3].Layout(gtx, Next)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[4].pressed {
					log.Print("restart")
					fb[4].pressed = false
				}
				return fb[4].Layout(gtx, Next)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				for fb[5].pressed {
					log.Print("playlist")
					fb[5].pressed = false
				}
				return fb[5].Layout(gtx, Next)
			},
		),
	)
}
