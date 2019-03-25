package main

import (
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

	background := doc.Call("getElementById", "background-image")
	sprite := doc.Call("getElementById", "sprite-image")

	ctx := canvasEl.Call("getContext", "2d")

	vp := NewViewPort(
		[]GraphObject{
			NewImage(ImageCtx(background), Point2D{0, 0}),
			NewSegmentFollower(
				NewEmitter(15, 50, Point2D{100, 100}, 5, 1, 5, 3),
				.10, LineSegment{Point2D{100, 100}, Point2D{100, 500}}),
			NewSegmentFollower(
				NewHexagon(HexagonPointyTop, Point2D{200, 100}, 25, colorful.FastHappyColor()),
				.10, LineSegment{Point2D{200, 100}, Point2D{200, 500}}),
			NewSegmentFollower(
				NewDebuggedObject(NewPolygon(colorful.FastWarmColor(),
					Point2D{300, 100}, Point2D{400, 200},
					Point2D{400, 300}, Point2D{200, 200}),
					debugLabel,
					func (o GraphObject, l *Label) {
						vec := o.(*Polygon).Vector()
						l.SetTextf("Pos: {%.3f, %.3f} Vec: [0]{%.3f, %.3f} [1]{%.3f, %.3f} [2]{%.3f, %.3f} [3]{%.3f, %.3f}",
							o.Pos().X(), o.Pos().Y(), vec[0].X(), vec[0].Y(), vec[1].X(),
							vec[1].Y(), vec[2].X(), vec[2].Y(), vec[3].X(), vec[3].Y())
					}),
				.10, LineSegment{Point2D{300, 100}, Point2D{300, 500}}),
			NewSegmentFollower(
				NewTriangle(Point2D{400, 100}, Point2D{500, 200}, Point2D{300, 200}, colorful.FastHappyColor()),
				.10, LineSegment{Point2D{400, 100}, Point2D{400, 500}}),
			NewSegmentFollower(
				NewRectangle(Point2D{500, 100}, 10, 10, true),
				.10, LineSegment{Point2D{500, 100}, Point2D{500, 500}}),
			NewSegmentFollower(
				NewRectangle(Point2D{600, 100}, 10, 10, false),
				.10, LineSegment{Point2D{600, 100}, Point2D{600, 500}}),
			NewSegmentFollower(
				NewCircle(Point2D{700, 100}, colorful.FastHappyColor(), float64(25), float64(.5)),
				.10, LineSegment{Point2D{700, 100}, Point2D{700, 500}}),
			NewSegmentFollower(
				NewSprite(ImageCtx(sprite), Point2D{800, 100}, 1000, 100, 10),
				.10, LineSegment{Point2D{800, 100}, Point2D{800, 500}}),
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
