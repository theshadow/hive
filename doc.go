// Copyright 2020 Xander Guzman. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package hive is a tiny game engine of the Hive (c) board game that aims to be
rule compliant with the base game plus the Ladybug, Mosquito, and Pill Bug
expansions. A typical use of this library is to implement either a client and/or
server pair for hosting your own games.

Basics

The library consists of an engine for managing the rules and state of a single game instance.
The engine attempts to be efficient and compact in its memory usage. It hides this behind a
layer of data types. The most expensive algorithm implemented is A* which can be used for
many things but is primarily used for validating the movement of pieces.

At the core you can instantiate a new Game instance and interact with the state of the game
using one of the player actions, Place or Move. If either action being performed would be in
violation of the game rules or state then the action will return an error. For more information
about the types of errors the action interface will return see the Game type.

To support the game rules and manage the state the game instance will use a collection of types
dedicated to tracking the location of pieces, player pieces, turn number, and turn history.

Types and Values

Game maintains the state of an instance of the game engine and is where the rules are implemented.
Board functions as the surface where a Piece is placed which is tracked with the Coordinate.
All together, a Player may perform one action per turn and that act is recorded as an Action.
A collection of Actions is stored in the state as the history.
*/
package hive
