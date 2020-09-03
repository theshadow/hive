// +build example

package main

import "math/bits"

func main() {

}

func canMove(field uint8) bool {
	cnt := bits.OnesCount(field)
	if cnt == 5 || (cnt == 4 && isPinned(field)) {
		return false
	} else if cnt == 3 && causesSplit(field, []uint8{}) {
		return false
	} else {
		return true
	}
}

// isPinned will rotate the passed in value up to eight times attempting to locate a match of the passed in mask
func isPinned(in uint8, mask uint8) (pinned bool) {
	for iter := 0; (in & mask) != mask && iter < 8; iter++ {
		in = bits.RotateLeft8(in, 1)
	}

	if (in & mask) == mask {
		pinned = true
	}

	return pinned
}

func causesSplit(in uint8, board []uint8) bool {
	return false
}