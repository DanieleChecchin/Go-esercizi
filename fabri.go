package main

import "fmt"

func fabriMap(n *[]string) {
	*n = append(*n, "luca")
}

func fabri() {
	s := []string{"dani"}
	fabriMap(&s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s) //più usato
	fmt.Printf("%#v\n", s)
}

// & passo l'indirizzo in memoria a fabriMap e lui modifica il contenuto aggiungendo Luca
// senza * e & passo solo il contenuto, di fatto creo una copia si s ma non la modifico realmente
