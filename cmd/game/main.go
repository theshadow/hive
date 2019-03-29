package main

import (
	"math"
	"syscall/js"

	"github.com/lucasb-eyer/go-colorful"

	. "github.com/theshadow/sge"
)

var (
	width      float64
	height     float64
	debugLabel = NewLabel("", "18px Arial", Point2D{900, 100})
	fpsLabel   = NewLabel("", "18px Arial", Point2D{900, 150})
)

func main() {
	// Init Canvas stuff
	doc := js.Global().Get("document")

	width = doc.Get("body").Get("clientWidth").Float()
	height = doc.Get("body").Get("clientHeight").Float()

	canvasEl := doc.Call("getElementById", "background-canvas")
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)

	ctx := canvasEl.Call("getContext", "2d")

	vp := NewViewPort(
		[]GraphObject{
			NewHexagon(HexagonFlatTop, Point2D{100, 100}, 25, colorful.FastHappyColor()),
			fpsLabel,
			debugLabel,
		},
		width,
		height)

	done := make(chan struct{}, 0)
	var (
		now float64
		tdiff float64
		tdiffSum float64
		tmark float64
		renderFrame js.Func
		markCount int
		curBodyW float64
		curBodyH float64
	)
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		now = args[0].Float()
		tdiff = now - tmark
		tdiffSum += now - tmark
		markCount++

		if markCount > 10 {
			fpsLabel.SetTextf("FPS: %.01f", 1000/(tdiffSum/float64(markCount)))
			tdiffSum, markCount = 0, 0
		}
		tmark = now

		// Pool window size to handle resize
		curBodyW = doc.Get("body").Get("clientWidth").Float()
		curBodyH = doc.Get("body").Get("clientHeight").Float()
		if curBodyW != width || curBodyH != height {
			width, height = curBodyW, curBodyH
			canvasEl.Set("width", width)
			canvasEl.Set("height", height)
			vp.SetDimensions(width, height)
		}

		vp.Tick(tdiff / 1000)
		vp.Render(CanvasCtx(ctx))

		js.Global().Call("requestAnimationFrame", renderFrame)

		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}

type AxialPoint2D [2]float64

func (p AxialPoint2D) Q() float64 {
	return p[0]
}
func (p AxialPoint2D) R() float64 {
	return p[1]
}

// hexToPixel accepts the coordinates of a hex in axial and the size and converts it into a pixel
// location
func hexToPixel(pt AxialPoint2D, size float64) Point2D {
	return Point2D{
		size * 3./2 * pt.R(),
		size * (math.Sqrt(3) * pt.Q() + math.Sqrt(3)/2 * pt.R()),
	}
}
