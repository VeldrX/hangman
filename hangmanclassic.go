package hangman

import (
	"fmt"
	"os"
)

type HangManData struct {
	Word         []rune // Word composed of '_', ex: H_ll_
	WordToFind   []rune // Final word chosen by the program at the beginning. It is the word to find
	Attempts     int    // Number of attempts left
	WrongLetters []rune
	Asccitype    string
}

func Hangclassic() {
	imputcorect := true
	Myhangman := new(HangManData)
	argu := os.Args
	if len(argu) < 2 {
		fmt.Println("il manque un argument")
		os.Exit(1)
	}
	fileword := argu[1]
	Myhangman.WordToFind = GetWord(fileword, Myhangman)
	Myhangman.Word = GetInitLetters(Myhangman.WordToFind)
	Myhangman.Attempts = 10

	fmt.Println("!!! WELCOME  TO  THE  HANGED-MAN  GAME !!!")
	fmt.Println()
	Display(Myhangman.Word)

	var imput string

	for Myhangman.Attempts > 0 {
		fmt.Println("")
		fmt.Print("Type a Letter : ")
		fmt.Scan(&imput)
		imput_in_rune := []rune(imput)

		imput_in_rune, imputcorect = VerifyImput(imput_in_rune)

		if imputcorect {
			attemptbefor := Myhangman.Attempts
			AddLetter(Myhangman, imput_in_rune)
			Display(Myhangman.Word)
			if attemptbefor != Myhangman.Attempts {
				fmt.Println("")
				fmt.Println("")
				Josedisplay(*Myhangman)
			}
			fmt.Println()
			if Myhangman.Attempts > 0 {
				fmt.Println(Myhangman.Attempts)
			} else {
				fmt.Println(0)
			}

			compteur_egal := 0
			for i := 0; i < len(Myhangman.Word); i++ {
				if Myhangman.Word[i] == Myhangman.WordToFind[i] {
					compteur_egal += 1
				}
			}
			if compteur_egal == len(Myhangman.WordToFind) && Myhangman.Attempts > 0 {
				fmt.Println("BRAVO vous avez Gagné")
				os.Exit(0)
			}
		} else {
			fmt.Println()
			fmt.Println("Erreur il faut des lettres")
		}
	}
}
