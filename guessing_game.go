package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Prints the necessary prompts to get user input, gets it, parses it, returns it
func getUserGuess() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your guess: ")
	scanner.Scan()
	guess := scanner.Text()
	i, err := strconv.ParseInt(guess, 10, 0)
	userGuess := int(i)
	if err != nil {
		fmt.Println("There was an error processing your number")
		fmt.Println("Please try again")
		userGuess = getUserGuess()
	}
	return userGuess

}

// Function that seeds the rand generator and then returns a random number
func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess a number between 0 and 100")
	randNum := getRandomNum()
	numGuesses := 0
	for {
		guess := getUserGuess()
		if guess < randNum {
			fmt.Println("Your guess was too low")
		} else if guess > randNum {
			fmt.Println("Your guess was too high")
		} else {
			fmt.Println("Congrats you guessed the number!")
			break
		}
		numGuesses++
	}
	fmt.Println("Thank you for playing")
	fmt.Println("It took you", numGuesses, "guesses")

}
