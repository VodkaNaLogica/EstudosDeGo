package main

import "fmt"

var nome string = "Pablo"
var idade int = 15
var ativo bool = true

const pi float64 = 355 / 113

var familia string

const filosofia string = "Por que algo em vez de nada?"

func main() {
	familia = "amor"

	fmt.Println("Olá, Go")
	soma(99, 15)

	var inputIdade int
	fmt.Print("\nDigite sua idade: ")
	fmt.Scanln(&inputIdade)

	faseDaVida(inputIdade)
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
