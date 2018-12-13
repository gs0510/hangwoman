package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wins := 0
	loses := 0
	numLetters := rand.Intn(11) + 4 //generates random number from 4 to 15
	again, hasWon := playHangwoman(numLetters)
	for {
		if hasWon == true {
			wins++
			numLetters = rand.Intn(11) + 4
		} else {
			loses++
			numLetters = rand.Intn(11) + 4
		}
		if again == "y" {
			clearscreen()
			fmt.Printf("------------------------\n")
			fmt.Printf("    Current Score\n")
			fmt.Printf("  %d: wins, %d: loses\n", wins, loses)
			fmt.Printf("------------------------\n")
			again, hasWon = playHangwoman(numLetters)
		} else if again == "n" {
			break
		}
	}
}

func playHangwoman(numLetters int) (playagain string, isWinner bool) {
	stageOfDeath := 0
	gamemode := 0
	hasGuessedOneLetter := false
	hasWon := false
	guess := ""
	guessedLetters := ""
	again := ""
	dashes := ""
	newdashes := ""
	fmt.Printf("H A N G M A N\n")
	for {
		fmt.Println("Select game mode:")
		fmt.Println("1. Only use Common words (easy mode)")
		fmt.Println("2. Use all words (hard mode)")
		fmt.Scanln(&gamemode)
		if (gamemode == 1) || (gamemode == 2) {
			clearscreen()
			break
		} else {
			fmt.Println("Please type 1 or 2")
		}
	}
	word := randomWord(numLetters, gamemode)
	for {

		hangWoman(stageOfDeath)
		if stageOfDeath == 9 {
			fmt.Printf("Oh dear hangwoman is dead\n")
			fmt.Printf("The word that could have saved him was %s\n", word)
			for {
				fmt.Printf("Play again? (y/n) \n")
				fmt.Scanln(&again)
				isYorN, err := regexp.MatchString("^y|Y|n|N", again)
				if err != nil {
					fmt.Printf("Something has gone horribly wrong. ")
					fmt.Printf("exiting with error can not regex match %v", again)
					return
				}
				if isYorN == false {
					fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
				} else if len(again) > 1 {
					fmt.Printf("You entered more than 1 character! Try again\n")
				} else if strings.ToLower(again) == "y" {
					return "y", false
				} else {
					return "n", false
				}

			}
		}
		if hasGuessedOneLetter == false {
			dashes = hideWord(len(word))
			fmt.Printf("%s\n", dashes)
		} else {
			fmt.Printf("%s\n", newdashes)
		}
		fmt.Printf("Guess a letter: ")
		fmt.Scanln(&guess)

		isALetter, err := regexp.MatchString("^[a-zA-Z]", guess)
		if err != nil {
			clearscreen()
			fmt.Printf("Something has gone horribly wrong. ")
			fmt.Printf("exiting with error can not regex match %v", guess)
			return
		}

		if isALetter == false {
			clearscreen()
			fmt.Printf("That's not a letter! Try again\n")
		} else if len(guess) > 1 {
			clearscreen()
			fmt.Printf("You entered more than 1 character! Try again\n")
		} else if strings.Contains(guessedLetters, guess) {
			clearscreen()
			fmt.Printf("You have already guessed that letter! Try again\n")
		} else if strings.Contains(word, guess) {
			clearscreen()
			fmt.Printf("The letter you guessed is in the word\n")
			guessedLetters += guess

			if hasGuessedOneLetter == false {
				updateddashes := revealDashes(word, guess, dashes)
				newdashes = updateddashes
			} else {
				updateddashes := revealDashes(word, guess, newdashes)
				newdashes = updateddashes
			}
			hasGuessedOneLetter = true
			if newdashes == word {
				hasWon = true
			}
			if hasWon == true {
				clearscreen()
				fmt.Printf("-= C O N G R A T U L A T I O N S =-\n")
				fmt.Printf("You won the game! The word was %s\n", word)
				for {
					fmt.Printf("Play again? (y/n) \n")
					fmt.Scanln(&again)
					isYorN, err := regexp.MatchString("^y|Y|n|N", again)
					if err != nil {
						fmt.Printf("Something has gone horribly wrong. ")
						fmt.Printf("exiting with error can not regex match %v", again)
						return
					}
					if isYorN == false {
						fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
					} else if len(again) > 1 {
						fmt.Printf("You entered more than 1 character! Try again\n")
					} else if strings.ToLower(again) == "y" {
						return "y", true
					} else {
						return "n", true
					}
				}
			}
		} else {
			clearscreen()
			fmt.Printf("The letter you guessed is not in the word\n")
			stageOfDeath++
			guessedLetters += guess
		}
	}
}
func hangWoman(stageOfDeath int) {
	switch stageOfDeath {
	case 0:
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 1:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 2:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 3:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 4:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 5:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 6:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 7:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 8:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 9:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("R.I.P |\n")
		fmt.Printf("========\n")
	}
}
func hideWord(wordlen int) string {
	dashes := ""
	for i := 0; i < wordlen; i++ {
		dashes += "_"
	}
	return dashes
}
func revealDashes(word string, guess string, dashes string) string {
	newdashes := ""
	for i, r := range dashes {
		if c := string(r); c != "_" {
			newdashes += c

		} else {
			var letter = string(word[i])
			if guess == letter {
				newdashes += guess
			} else {
				newdashes += "_"
			}
		}
	}
	return newdashes
}

func checkIfWon(newdashes string, word string) bool {
	if newdashes == word {
		return true
	}
	return false
}

func randomWord(numLetters int, gamemode int) string {
	switch gamemode {
	case 1:
		var dataletters []byte
		var err error
		if numLetters == 4 {
			dataletters, err = ioutil.ReadFile("words/commonFourLetterWords.txt")
		} else if numLetters == 5 {
			dataletters, err = ioutil.ReadFile("words/commonFiveLetterWords.txt")
		} else if numLetters >= 6 {
			dataletters, err = ioutil.ReadFile("words/commonSixLetterWords.txt")
		}

		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword

	case 2:
		var dataletters []byte
		var err error
		if numLetters == 4 {
			dataletters, err = ioutil.ReadFile("words/allFourLetterWords.txt")
		} else if numLetters == 5 {
			dataletters, err = ioutil.ReadFile("words/allFiveLetterWords.txt")
		} else if numLetters == 6 {
			dataletters, err = ioutil.ReadFile("words/allSixLetterWords.txt")
		} else if numLetters == 7 {
			dataletters, err = ioutil.ReadFile("words/allSevenLetterWords.txt")
		} else if numLetters == 8 {
			dataletters, err = ioutil.ReadFile("words/allEightLetterWords.txt")
		} else if numLetters == 9 {
			dataletters, err = ioutil.ReadFile("words/allNineLetterWords.txt")
		} else if numLetters == 10 {
			dataletters, err = ioutil.ReadFile("words/allTenLetterWords.txt")
		} else if numLetters == 11 {
			dataletters, err = ioutil.ReadFile("words/allElevenLetterWords.txt")
		} else if numLetters == 12 {
			dataletters, err = ioutil.ReadFile("words/allTwelveLetterWords.txt")
		} else if numLetters == 13 {
			dataletters, err = ioutil.ReadFile("words/allThirteenLetterWords.txt")
		} else if numLetters == 14 {
			dataletters, err = ioutil.ReadFile("words/allFourteenLetterWords.txt")
		} else if numLetters == 15 {
			dataletters, err = ioutil.ReadFile("words/allFifteenLetterWords.txt")
		}

		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword

	}

	return "Error"
}

func clearscreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
