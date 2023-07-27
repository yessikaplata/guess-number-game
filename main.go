package main

import (
	"fmt"
	"math/rand"
)

const (
	MAX_ATTEMPTS = 10
)

func main() {
	fmt.Println("******Guess the magician's secret number!******")
	fmt.Printf("A Magician will pick a secret number between 1 to 100 and put it in his hat.\nYou guess what number it is.\nIf your guess is too high or too low, a Magician will give you a hint.\nYou have %v attempts to win!\n", MAX_ATTEMPTS)
	isPlay := true
	for isPlay {
		playgame()
		isPlay = evaluatePlayAgain()
	}
	fmt.Println("Good bye!!!")
}

func playgame() {
	attempts, userNumber := 1, 0
	number := rand.Intn(100)
	for attempts <= MAX_ATTEMPTS {
		fmt.Printf("Attempt number %v. Enter your number: ", attempts)
		fmt.Scan(&userNumber)
		if userNumber == number {
			fmt.Printf("You have won in %v attempts. The number is %v\n", attempts, number)
			return
		}
		if attempts >= MAX_ATTEMPTS {
			fmt.Printf("Your attempts are over. The number was %v", number)
			return
		}
		if number < userNumber {
			fmt.Printf("Try Again. The number is less than %v\n", userNumber)
		} else {
			fmt.Printf("Try Again. The number is greater than %v\n", userNumber)
		}
		attempts++
	}
}

func evaluatePlayAgain() bool {
	var option string
	fmt.Print("Do you want to play again? (y/n): ")
	fmt.Scan(&option)
	if option == "y" {
		return true
	}
	return false
}
