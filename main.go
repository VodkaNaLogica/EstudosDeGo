package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var nome string = "Pablo"
var idade int = 15
var ativo bool = true

const pi float64 = 355 / 113

var familia string

const filosofia string = "Por que algo em vez de nada?"

func main() {
	familia = "amor"

	fmt.Println("Olá, Go")

	var numA int = rand.IntN(1000)
	var numB int = rand.IntN(1000)

	fmt.Println(numA, " + ", numB)

	soma(numA, numB)

	var inputIdade int
	fmt.Print("\nDigite sua idade: ")
	fmt.Scanln(&inputIdade)

	faseDaVida(inputIdade)

	var inputSimbolosDoTriangulo string

	fmt.Print("\nDigite o simbolo desejado para a construção do triangulo: ")
	fmt.Scanln(&inputSimbolosDoTriangulo)

	var inputBaseDoTriangulo int

	fmt.Print("\nDigite o tamanho da base do tringulo desejada: ")
	fmt.Scanln(&inputBaseDoTriangulo)

	CriarTriangulo(inputBaseDoTriangulo, inputSimbolosDoTriangulo)

	fmt.Scanln()
}

func soma(a int, b int) {
	fmt.Println(a + b)
}

func faseDaVida(idade int) {
	if idade >= 18 {
		if idade < 21 {
			println("Adolecência tardia")
		} else if idade < 39 {
			println("Jovem adulto")
		} else if idade < 59 {
			println("Adulto")
		} else {
			println("Idoso")
		}

	} else {

		if idade < 10 {
			println("Criança")
		} else if idade < 13 {
			println("Pre-Adolecente")
		} else {
			println("Adolecente")
		}

	}

}

func CriarTriangulo(tamanhoDaBase int, SimbolosDoTriangulo string) {
	var simbolo string = strings.TrimSpace(SimbolosDoTriangulo)

	for i := 0; i < tamanhoDaBase; i++ {
		espaçoAlinhador := strings.Repeat(strings.Repeat(" ", len(simbolo)), tamanhoDaBase-i)
		espaçoEntreSimbolo := strings.Repeat(" ", len(simbolo))
		var acumulado strings.Builder

		for j := 0; j < i+1; j++ {
			acumulado.WriteString(simbolo)
			acumulado.WriteString(espaçoEntreSimbolo)
		}

		fmt.Println(espaçoAlinhador, acumulado.String())
	}
}
