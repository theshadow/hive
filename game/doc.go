// Copyright 2020 Xander Guzman. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
/*
Package game contains the implementation of the state management and rules engine type.
The Game type contains manages the state and provides the rules engine interface. This
interface is described as the two actions a player may take each turn. Those are Place
and Move. A Place action is where a player takes a piece from their pool and sets it
at a particular coordinate on the board while Move is where the player updates the
location of a piece that has already been placed.

Basics

During each turn, each player may perform one of the actions if the action they attempt
to perform generates an error the action receiver function will return an error. If the
error is a violation of the game rules a Rules Error will be returned. If the error is
due to a logic failure or a state corruption then other error types may be returned.

Features

The engine has implemented feature flags for rules beyond the base game. These rules
may be toggled on and off at the instantiation of the game type.

Types and Values

The Game type should act as the primary interface for the library if you want to just
provide a client or server wrapper around the the game. If you're looking to implement
your own game the rest of the types within the hived package are at your disposal.

Errors

The Game type will return two types of errors Rule and State. Rule errors are returned
when the player action made violates a game rule. A state error is returned when the
attempted interaction with the game state is invalid.

State Errors

- ErrGameNotOver : Returned when using the Winner interface and the game hasn't reached an end state.
- ErrUnknownPiece : Returned when attempting to place a piece that isn't recognized by the engine.
- ErrUnknownBoardError : Returned if there is an unexpected error while updating the state of the board.

Rule Errors

Returned when either a Place or a Move action violates a rule of the game. For more information see
the rules file. These errors should be very specific and clear when compared to the rules of the game.
*/
package game
