package main

import (
	"code-decrypter/code"
	"code-decrypter/game"
	"fmt"
)

func main() {
	possibleCodes := code.GenerateAllCodes()

	game := game.Game{CodeToGuess: code.Code{6, 1, 4, 3, 5}}
	count := 1

	for {
		guess := possibleCodes[len(possibleCodes)-1]
		result := guess.CompareTo(game.CodeToGuess)
		fmt.Printf("Guess %v: %v\n", count, guess)

		if result.Correct == len(game.CodeToGuess) {
			fmt.Printf("Correct %v", guess)
			break
		}

		newPossibleCodes := make([]code.Code, 1)
		for i := 0; i < len(possibleCodes); i++ {
			if possibleCodes[i].CompareTo(guess) == result {
				newPossibleCodes = append(newPossibleCodes, possibleCodes[i])
			}
		}

		possibleCodes = newPossibleCodes
		count++
	}
}
