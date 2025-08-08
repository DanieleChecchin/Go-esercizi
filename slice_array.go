package main

import "fmt"

// Chiedi all’utente di inserire 5 numeri interi, salvali in uno slice e poi stampali.

func insertFiveNumbers() {

	var numbers []int // Slice vuoto per salvare i numeri

	fmt.Println("Inserisci 5 numeri interi:")

	for i := 0; i < 5; i++ {

		//Creo una variabile k per ogni numero inserito
		var k int
		// Chiedo all'utente di inserire un numero e uso +, così salto lo 0
		fmt.Printf("Numero %d: ", i+1)
		// metto il valore dentro la variabile k
		fmt.Scanln(&k)
		// Aggiungo il numero inserito allo slice
		numbers = append(numbers, k)
	}

	fmt.Println("Hai inserito i seguenti numeri:", numbers)

}
