package main

import (
	"code-decrypter/code"
	"code-decrypter/game"
	"fmt"
	"math/rand"
)

func main() {
	numberOfGames := 10000

	maxGuesses := 0
	avgGuesses := make([]int, numberOfGames)
	possibleCodes := code.GenerateAllCodes()

	for gi := 0; gi < numberOfGames; gi++ {
		codeIndex := rand.Intn(len(possibleCodes) - 1)
		game := game.Game{CodeToGuess: possibleCodes[codeIndex]}
		count := 1
		firstBadCode := len(possibleCodes)

		for {
			guess := possibleCodes[0]
			result := guess.CompareTo(game.CodeToGuess)

			if result.Correct == len(game.CodeToGuess) {
				if count > maxGuesses {
					maxGuesses = count
				}
				avgGuesses[gi] = count
				break
			}

			// front load the good codes
			for i := 0; i < firstBadCode; i++ {
				if possibleCodes[i].CompareTo(guess) != result {
					// swap the first good code
					temp := possibleCodes[i]
					possibleCodes[i] = possibleCodes[firstBadCode-1]
					possibleCodes[firstBadCode-1] = temp
					firstBadCode--
				}

				// trim the end to avoid swapping
				for possibleCodes[firstBadCode-1].CompareTo(guess) != result {
					firstBadCode--
				}
			}

			count++
		}
	}

	var sum int
	for i := 0; i < len(avgGuesses); i++ {
		sum += avgGuesses[i]
	}
	average := float32(sum) / float32(len(avgGuesses))

	fmt.Printf("Max Guesses: %v\n", maxGuesses)
	fmt.Printf("Avg Guesses: %.2f\n", average)
}
