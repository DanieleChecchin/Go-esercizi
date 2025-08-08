package main

import "fmt"

// Stampa i numeri da 1 a 10 con un ciclo for.

func oneToTen() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

// Chiedi un numero N e calcola la somma da 1 a N.

func sumOfNumber() {

	var n int
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	// Se n = 10, 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10
	fmt.Println("The sum of the", n, "numbers is:", sum)
}

// Chiedi un numero all’utente e stampa la sua tabellina (es. 5×1 fino a 5×10).

func multiplicationTable() {

	var n int
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)
	for i := 0; i <= 10; i++ {
		s := n * i
		fmt.Printf("%v * %v = %v \n", n, i, s)

	}
}
