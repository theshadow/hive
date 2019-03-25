package hive

type HivePiece int
type PieceColor int

const (
	QueenPiece HivePiece = iota
	GrassHopperPiece
	SoldierAntPiece
	BeetlePiece
	SpiderPiece
	MosquitoPiece
	LadybugPiece
	PillBugPiece

	BlackPieceColor PieceColor = iota
	WhitePieceColor
)



type Piece struct {
	h *sge.Hexagon
}
