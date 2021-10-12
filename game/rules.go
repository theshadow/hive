package game

import "fmt"

var ErrRuleFirstPieceMustBeAtOrigin = fmt.Errorf("the first piece to be placed must be placed at origin")
var ErrRuleMayNotPlaceAPieceOnAPiece = fmt.Errorf("may not place a piece where another piece exists")
var ErrRuleMustPlacePieceOnSurface = fmt.Errorf("a piece must be placed on the surface of the board")
var ErrRulePieceParalyzed = fmt.Errorf("this piece is paralyzed and may not move")
var ErrRuleMayNotPlaceTouchingOpponentsPiece = fmt.Errorf("the player may not place a piece where it will touch an opponents piece after the first turn")
var ErrRuleMayNotPlaceQueenOnFirstTurn = fmt.Errorf("tournament rules: a player may not place their queen on the first turn")
var ErrRuleNotPlayersTurn = fmt.Errorf("a player may only act on their turn")
var ErrRuleMustPlaceQueen = fmt.Errorf("the player must place their queen by the fourth turns")
var ErrRuleMustPlaceQueenToMove = fmt.Errorf("the players queen must be placed before a placed piece may move")
var ErrRulePieceAlreadyParalyzed = fmt.Errorf("the piece is already paralyzed and may not be stunned again this turn")
var ErrRuleMovementDistanceTooGreat = fmt.Errorf("the distance for the movement is too great for this piece")
