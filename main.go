package main

import (
	definitions "chessengine/definitons"
	"fmt"
)

func main() {
	definitions.InitialiseSq120To64()

	for i := 0; i < definitions.BoardSize; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", definitions.Square120To64[i])
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", definitions.Square64To120[i])
	}
}
