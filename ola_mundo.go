package main

import (
	"fmt"
)

func main() {
	/*
		Comentário 1
	*/

	// Definição de variável com tipagem estática
	var text_start string = "Inicío do olá mundo"
	var text_end string
	var count int = 4546846 * 3

	// Definição de variável com tipagem automática
	author := "Vagner Praia"
	print_author := true

	text_end = "Fim do olá mundo"

	// Definindo constantes
	const ano int = 2018

	fmt.Println("Texto inicial: " + text_start)
	fmt.Println(count)
	fmt.Println("Texto final: " + text_end)
	fmt.Println(print_author)
	fmt.Println("Author: " + author)
	fmt.Println(ano)

	// Comentário 2
}
