package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// CONFIGURATION
	attempts := 7
	maxNumber := 100

	randomInt := rand.IntN(maxNumber) + 1
	fmt.Printf("I have thought of an integer. Try to guess it! You have %d attempts.\n", attempts)

	for attempts != 0 {
		var answer int
		for {
			fmt.Print("Your guess: ")
			_, err := fmt.Scanln(&answer)
			if err != nil {
				fmt.Println("Write a valid integer!")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			break
		}

		if answer == randomInt {
			fmt.Println("You guessed it right! Congratulations!")
			return
		}

		var clue string
		if randomInt > answer {
			clue = "bigger"
		} else {
			clue = "smaller"
		}

		fmt.Printf("That's not right. My integer is %s\n", clue)
		attempts--
		fmt.Printf("Your attempts: %d\n", attempts)
	}

	fmt.Println("The attempts have ended! Better luck next time!")
	fmt.Printf("By the way, my integer was: %d\n", randomInt)
}
