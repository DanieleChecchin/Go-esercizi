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

// Data una lista di numeri (in uno slice), trova il valore minimo e massimo.

func findMinMax() {
	// Creo uno slice vuoto per salvare i numeri inseriti
	var numeri []int

	fmt.Println("Inserisci 5 numeri interi:")

	for i := 0; i < 5; i++ {
		var valore int
		fmt.Printf("Numero %d: ", i+1)
		fmt.Scanln(&valore)
		numeri = append(numeri, valore)
	}

	// Supponiamo che il primo valore sia sia il minimo che il massimo
	min := numeri[0]
	max := numeri[0]

	// uso il ciclo for per trovare il minimo e massimo
	for _, num := range numeri {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println("Il numero minimo è:", min)
	fmt.Println("Il numero massimo è:", max)
}
