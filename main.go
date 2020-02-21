package hived

import (
	"fmt"
	"os"
)

func main() {
	game := NewGame()

	if err := game.Place(NewBlackPiece(Grasshopper, PieceA), Origin); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if game.Over() {
		winner, err := game.Winner()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		} else if winner == ZeroPlayer {
			fmt.Println("Winner: Tie")
		} else {
			fmt.Printf("Winner: %s", winner)
		}
	} else {
		fmt.Println("Game didn't end, and that's OK!")
	}
}
