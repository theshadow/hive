package game

import "fmt"

var (
	ErrRuleFirstPieceMustBeAtOrigin          = fmt.Errorf("the first piece to be placed must be placed at origin")
	ErrRuleMayNotPlaceAPieceOnAPiece         = fmt.Errorf("may not place a piece where another piece exists")
	ErrRuleMustPlacePieceOnSurface           = fmt.Errorf("a piece must be placed on the surface of the board")
	ErrRulePieceParalyzed                    = fmt.Errorf("this piece is paralyzed and may not move")
	ErrRulePiecePinned                       = fmt.Errorf("this piece is pinned and may not move")
	ErrRuleMayNotPlaceTouchingOpponentsPiece = fmt.Errorf("the player may not place a piece where it will touch an opponents piece after the first turn")
	ErrRuleMayNotPlaceQueenOnFirstTurn       = fmt.Errorf("tournament rules: a player may not place their queen on the first turn")
	ErrRuleNotPlayersTurn                    = fmt.Errorf("a player may only act on their turn")
	ErrRuleMustPlaceQueen                    = fmt.Errorf("the player must place their queen by the fourth turns")
	ErrRuleMustPlaceQueenToMove              = fmt.Errorf("the players queen must be placed before a placed piece may move")
	ErrRulePieceAlreadyParalyzed             = fmt.Errorf("the piece is already paralyzed and may not be stunned again this turn")
	ErrRuleMovementDistanceTooGreat          = fmt.Errorf("the distance for the movement is too great for this piece")
)
