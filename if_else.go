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
