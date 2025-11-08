package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎯 Welcome to the Number Guessing Game!")
	fmt.Println("--------------------------------------")
	fmt.Println("Rules:")
	fmt.Println("1. The computer will choose a number between 1 and 100.")
	fmt.Println("2. You must guess it within the allowed attempts.")
	fmt.Println("3. You’ll be told if your guess is too high or too low.")
	fmt.Println()

	// Choose difficulty
	var level string
	var maxAttempts int

	for {
		fmt.Print("Select difficulty (easy / medium / hard): ")
		fmt.Scan(&level)
		level = strings.ToLower(level)

		switch level {
		case "easy":
			maxAttempts = 10
		case "medium":
			maxAttempts = 7
		case "hard":
			maxAttempts = 5
		default:
			fmt.Println("Invalid choice. Please enter easy, medium, or hard.")
			continue
		}
		break
	}

	// Generate random number
	secretNumber := rand.Intn(100) + 1
	var guess int

	fmt.Printf("\nI’ve picked a number between 1 and 100. You have %d chances!\n", maxAttempts)

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("\nAttempt %d/%d — Enter your guess: ", attempts, maxAttempts)
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Printf("\n🎉 Congratulations! You guessed the number %d in %d attempts!\n", secretNumber, attempts)
			return
		} else if guess < secretNumber {
			fmt.Println("📉 Too low! Try again.")
		} else {
			fmt.Println("📈 Too high! Try again.")
		}
	}

	fmt.Printf("\n😞 Game over! You ran out of chances. The number was %d.\n", secretNumber)
}
