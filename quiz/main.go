package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")

	// var name string = "Benni", same as:
	// name := "Benni"
	var name string

	fmt.Printf("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Printf("Hello %v!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	_, err = fmt.Scan(&age)
	if err != nil {
		return
	}

	ageCheck := age >= 12
	if ageCheck == false {
		fmt.Println("You are to young")
		return
	}
	fmt.Println("You can play")
	score := 0
	numQuestions := 0

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer1 string
	var answer2 string

	_, err = fmt.Scan(&answer1, &answer2)
	if err != nil {
		return
	}
	answer1 = strings.ToLower(answer1)

	numQuestions++
	if answer1+" "+answer2 == "rtx 3090" ||
		answer1+" "+answer2 == "Rtx 3090" ||
		answer1+" "+answer2 == "RTx 3090" ||
		answer1+" "+answer2 == "RTX 3090" ||
		answer1+" "+answer2 == "rTx 3090" ||
		answer1+" "+answer2 == "rTX 3090" ||
		answer1+" "+answer2 == "rtX 3090" {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	_, err = fmt.Scan(&cores)
	if err != nil {
		return
	}
	numQuestions++
	if cores == 12 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Correct is 12 Cores.")
	}

	percent := (float64(score) / float64(numQuestions)) * 100
	fmt.Printf("You got %v from a total of %v questions correct, that is %v%%.\n"+
		"Bye have a nice day.",
		score, numQuestions, percent)

}
