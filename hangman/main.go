package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func doContains(slc []int, pos int) bool {
	for _, num := range slc {
		if pos == num {
			return true
		}
	}

	return false
}

func replaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func generateRandomIndexes(size int) []int {
	var unique = make(map[int]bool)
	var positionsToShow []int

	for i := 0; len(positionsToShow) < 3; i++ {
		random := rand.Intn(size)
		if unique[random] {
			continue
		}
		positionsToShow = append(positionsToShow, random)
		unique[random] = true
	}

	return positionsToShow
}

func movieToGuess(mov string) string {
	positionsToShow := generateRandomIndexes(len(mov))

	var encodedMovie string

	for pos, char := range mov {
		if doContains(positionsToShow, pos) {
			encodedMovie = encodedMovie + string(char)
		} else {
			encodedMovie = encodedMovie + "*"
		}
	}

	return encodedMovie
}

func showEncodedMovie(mov string) {
	fmt.Println("\n######################################################\n")
	for _, ltr := range mov {
		fmt.Print(string(ltr), " ")
	}
	fmt.Println("\n######################################################\n")
}

func checkGuess(guess string, movieName string, encodedMov string) (bool, int) {
	for pos, char := range movieName {
		if string(char) == guess && string(encodedMov[pos]) == "*" {
			return true, pos
		}
	}

	return false, 0
}

func main() {

	movies := []string{"ANTMAN", "BATMAN", "AVENGERS", "SPIDERMAN", "AVATAR", "DOCTORSTRANGE"}

	rand.Seed(time.Now().Unix())
	movie := movies[rand.Intn(len(movies))]

	// fmt.Println(movie)
	encodedMov := movieToGuess(movie)

	fmt.Println("\nHeyy, Let's play a guessing game.\nI will prompt a movie name with some missing characters and you've to guess the missing characters one character at a time.\nYou will only have 3 incorrect guesses\n\n")

	incorrectGuesses := 3

	for incorrectGuesses > 0 {
		showEncodedMovie(encodedMov)

		if encodedMov == movie {
			fmt.Println("\n\nYou are the winner!!!\n\n")
			showEncodedMovie(encodedMov)
			return
		}

		var guess string
		fmt.Println("Enter your guess: \n")
		fmt.Scanln(&guess)
		guess = strings.ToUpper(guess)
		res, idx := checkGuess(guess, movie, encodedMov)

		if res {
			fmt.Println("You guessed one of itt, more to go\n")
			encodedMov = replaceAtIndex(encodedMov, guess, idx)
		} else {
			fmt.Println("OOPS!, Your guess is wrong try again!")
			incorrectGuesses = incorrectGuesses - 1
		}
	}

	fmt.Println("\n\nOOPS!, You lost the game!!!\n")
	fmt.Println("The correct movie name is: \n")
	showEncodedMovie(movie)
}
