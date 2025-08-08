package main

import "fmt"

// Chiedi all’utente due numeri e stampa quale dei due è maggiore (o se sono uguali).

func highestNumber() {

	var firstNumber int
	var secondNumber int

	fmt.Println("Enter the first number:")
	fmt.Scanln(&firstNumber)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Printf("The highest number is : %v \n", firstNumber)
	} else {
		fmt.Printf("The highest number is : %v \n", secondNumber)
	}

}

// Chiedi l’età all’utente. Se ha 18 anni o più, scrivi "Puoi guidare", altrimenti "Non puoi guidare".

func canDrive() {
	var age int

	fmt.Println("How old are you?")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive yet")
	}
}

// Chiedi un numero e stampa se è positivo, negativo o zero.

func checkNumber() {
	var number int

	fmt.Println("Enter a number:")
	fmt.Scanln(&number)

	if number > 0 {
		fmt.Printf("The number %v in positive \n", number)
	} else if number < 0 {
		fmt.Printf("The number %v is negative \n", number)
	} else {
		fmt.Printf("The number %v is zero \n", number)
	}
}
