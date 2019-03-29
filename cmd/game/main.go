package main

import (
	"github.com/theshadow/hive"
	"syscall/js"

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

	center := Point2D{
		width/2,
		height/2,
	}

	debugLabel.SetTextf("Center %#v", center)

	vp := NewViewPort(
		[]GraphObject{
			hive.NewHexGrid(5, 5, center),
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

