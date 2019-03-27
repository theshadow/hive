package hive

import "github.com/theshadow/sge"

type PieceType int
type PieceColor int

const (
	QueenPiece PieceType = iota
	GrassHopperPiece
	SoldierAntPiece
	BeetlePiece
	SpiderPiece
	MosquitoPiece
	LadybugPiece
	PillBugPiece
)

// Point3D
type Point3D [3]float64

func (p Point3D) X() float64 {
	return p[0]
}

func (p Point3D) Y() float64 {
	return p[1]
}

func (p Point3D) Z() float64 {
	return p[2]
}

// Piece
type Piece struct {
	// where on the board is this piece located
	loc    Point3D
	// the renderable object
	obj    sge.GraphObject
	// the hive piece
	typ    PieceType
}

func NewPiece(loc Point3D, o sge.GraphObject, t PieceType) *Piece {
	return &Piece{
		loc: loc,
		obj: o,
		typ: t,
	}
}

type Grid struct {
	pieces []Piece
}
