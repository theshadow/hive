package hive

import (
	"log"
	"math"

	"github.com/lucasb-eyer/go-colorful"

	. "github.com/theshadow/sge"
)

type HexGrid struct {
	w int
	h int
	center Point2D
	grid []GraphObject
}

// NewHexGrid returns a new grid in the dimensions of w and h. Pos represents the center of the grid and where
// all hexes will be rendered offset from.
func NewHexGrid(w, h int, pos Point2D) *HexGrid {
	count := w * h
    grid := make([]GraphObject, count)
    size := float64(50)

    for r, q := 0, 0; r * h + q < count; q++ {
    	center := hexToPixel(AxialPoint2D{float64(q), float64(r)}, size).Add(pos)
		grid[r * h + q] = NewHexagon(
			HexagonFlatTop,
			center,
			size,
			colorful.FastHappyColor())
		log.Printf("Hex Point %#v", grid[r*h+q].Pos())
		if q == w {
			q = 0
			r++
		}
	}

	return &HexGrid{
		grid: grid,
		w: w,
		h: h,
		center: pos,
	}
}

// Render will run through the grid in O(n) form and call Tick() on each one.
func (g *HexGrid) Tick(time float64) {
	for i := 0; i < g.w * g.h; i++ {
		g.grid[i].Tick(time)
	}
}

// Render will run through the grid in O(n) form and call Render() on each one.
func (g *HexGrid) Render(c CanvasCtx) {
	for i := 0; i < g.w * g.h; i++ {
		g.grid[i].Render(c)
	}
}

func (g *HexGrid) Pos() Point2D {
	return g.center
}

func (g *HexGrid) SetPos(pos Point2D) {
	g.center = pos
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
		size * 3./2 * float64(pt.R()),
		size * (math.Sqrt(3) * pt.Q() + math.Sqrt(3)/2 * pt.R()),
	}
}

