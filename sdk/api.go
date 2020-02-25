package sdk

import (
	"fmt"
	"plugin"

	"github.com/theshadow/hived"
)

var plug *plugin.Plugin

type Game interface {
	History() []hived.Move
	Move(a, b hived.Coordinate) error
	Over() bool
	Place(p hived.Piece, c hived.Coordinate) error
	Winner() (hived.Player, error)
}

func LibraryVersion() string {
	v, err := plug.Lookup("LibraryVersion")
	if err != nil {
		return ""
	}

	return v.(string)
}

type Feature uint64

func NewGame(features []Feature) (Game, error) {
	fn, err := plug.Lookup("NewGame")
	if err != nil {
		return nil, &ErrSDKFunctionLookupFailed{
			Err: err,
		}
	}

	var ftrs []hived.Feature
	for _, f := range features {
		ftrs = append(ftrs, hived.Feature(f))
	}

	return fn.(func(features []hived.Feature) (*hived.Game, error))(ftrs)
}

type ErrSDKFunctionLookupFailed struct {
	Err error
}
func (e *ErrSDKFunctionLookupFailed) Error() string {
	return fmt.Sprintf("failed to lookup SDK function")
}
func (e *ErrSDKFunctionLookupFailed) Unwrap() error { return e.Err }

type ErrSDKVarLookupFailed struct {
	Err error
}
func (e *ErrSDKVarLookupFailed) Error() string {
	return fmt.Sprintf("failed to lookup SDK variable")
}
func (e *ErrSDKVarLookupFailed) Unwrap() error { return e.Err }

type ErrSDKVersionMismatch struct {
	Expected string
	Actual string
}
func (e *ErrSDKVersionMismatch) Error() string {
	return fmt.Sprintf("the loaded hived.so version %s doesn't match the expected version %s for this sdk",
		e.Actual, e.Expected)
}

func init() {
	// TODO: figure out how to look for system libraries, the local path, and/or ENV var
	var err error
	if plug, err = plugin.Open("hived.so"); err != nil {
		panic(err)
	}

	ver := LibraryVersion()
	if LibraryVersion() != EngineVersion {
		panic(&ErrSDKVersionMismatch{ EngineVersion, ver})
	}
}
