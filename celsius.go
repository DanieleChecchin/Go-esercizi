package main

import "fmt"

func celsiusToFahrenheit() {

	// Dichiaro la variabile celsius e indico di che tipo è
	var celsius float64

	// Chiedo all'utente di inserire un valore in gradi Celsius
	fmt.Print("Inserisci la temperatura in gradi Celsius: ")

	// Leggo il valore inserito dall'utente e lo memorizzo nella variabile celsius -- &= Passa l'informazione alla variabile
	fmt.Scanln(&celsius)

	var fahrenheit float64 = celsius*1.8 + 32

	// Stampo il risultato della conversione
	fmt.Printf("La temperatura di %v gradi celsius è di %v gradi fahrenheit !\n", celsius, fahrenheit)
}
