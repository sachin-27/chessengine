package bitboards

import (
	definitions "chessengine/definitions"
	"fmt"
)

func PrintBitBoard(bitboard uint64) {
	shiftNum := uint64(1)
	zeroVal := uint64(0)
	for rank := definitions.Rank8; rank >= definitions.Rank1; rank-- {
		for file := definitions.FileA; file <= definitions.FileH; file++ {
			sq120 := definitions.FileAndRank2Square(file, rank)
			sq64 := definitions.Square120To64[sq120]
			if (shiftNum<<sq64)&bitboard != zeroVal {
				fmt.Print("X")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
