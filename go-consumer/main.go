package main

import (
	"fmt"
	"os"
)

func main() {

	ler("Relatorio 2024-04-20")

}

func ler(texto string) {
	f, err := os.Create("relatorio.pdf")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	tamanho, err := f.WriteString(texto)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho = %d\n", tamanho)

}
