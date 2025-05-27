package ui

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

var (
	Bg color.NRGBA = color.NRGBA{R: 13, G: 27, B: 5, A: 255} // jedi night
	//Red color.NRGBA = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	//green color.NRGBA = color.NRGBA{R: 69, G: 179, B: 113, A: 255}
	//blue color.NRGBA = color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	//yellow color.NRGBA = color.NRGBA{R: 255, G: 165, B: 0, A: 255}

	//icons
	Next  string = "ui/icons/next.png"
	Pause string = "ui/icons/pause.png"
	Back  string = "ui/icons/back.png"
)

func DpPt(x, y float64) image.Point {
	return image.Pt(int(unit.Dp(x)), int(unit.Dp(y)))
}

func DrawSquare(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func DrawImage(ops *op.Ops, img image.Image) layout.Dimensions {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = paint.FilterNearest
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(1, 1))).Add(ops)
	paint.PaintOp{}.Add(ops)

	return layout.Dimensions{Size: img.Bounds().Max}
}
