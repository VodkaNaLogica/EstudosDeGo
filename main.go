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

	//fmt.Println("Olá, Go")
	//soma(99, 15)

	//var inputIdade int
	//fmt.Print("\nDigite sua idade: ")
	//fmt.Scanln(&inputIdade)

	//faseDaVida(inputIdade)

	var espaços string = ""
	var tamanhoDaBase int = 10

	for i := 0; i < tamanhoDaBase; i++ {
		acumulado := "# "

		for n := tamanhoDaBase - 1; n >= i; n-- {
			espaços += " "
		}
		if i == 0 {
			resultado := espaços + "*"
			fmt.Println(resultado)
			espaços = strings.TrimSpace(espaços)
		} else {

			for j := 0; j < i; j++ {
				if rand.IntN(4) == 0 {
					acumulado += "O "
				} else {
					acumulado += "# "
				}
			}

			resultado := espaços + acumulado
			fmt.Println(resultado)
			espaços = strings.TrimSpace(espaços)
		}
	}

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
