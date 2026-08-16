package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var nome string = "Pablo"
var idade int = 15
var ativo bool = false

const pi float64 = 355.0 / 113.0

var arrayA = [5]int{10, 20, 30, 40, 50}
var sliceA = []string{"A", "B", "C"}

var idades = map[string]int{
	"João":  15,
	"Maria": 16,
	"Pedro": 17,
}

var familia string

const filosofia string = "Por que algo em vez de nada?"

func main() {
	familia = "amor"

	fmt.Println("Olá, Go")
	fmt.Println("Meu nome é", nome)
	fmt.Println("Eu tenho", idade, "anos de idade")
	fmt.Println("Esse programa vai acabar nessa apresentação:", ativo)
	fmt.Printf("PI e aproximadamente: %.7f\n", pi)

	for i := 0; i < len(arrayA); i++ {
		println(arrayA[i])
		arrayA[i] = arrayA[i] * rand.IntN(1000)
	}

	for i := 0; i < len(arrayA); i++ {
		println(arrayA[i])
	}

	sliceA = append(sliceA, "Oi no slice do Go")
	sliceA = append(sliceA, "Olá denovo mas diferente no slice do Go")

	fmt.Println(sliceA)

	fmt.Println(idades)

	idades["Ana"] = 14
	idades["João"] = 18
	fmt.Println(idades)

	delete(idades, "Pedro")
	fmt.Println(idades)

	idade, existe := idades["Carlos"]
	if existe {
		fmt.Println("Encontrado", idade)
	} else {
		fmt.Println("Não encontrado")
	}

	var numA int = rand.IntN(1000)
	var numB int = rand.IntN(1000)

	fmt.Println(numA, " + ", numB, " = ", soma(numA, numB))

	for {
		fmt.Print("\nDigite sua idade: ")

		var inputIdade float32
		_, err := fmt.Scanln(&inputIdade)

		if err == nil {
			faseDaVida(inputIdade)
			break
		}
		var lixo string
		fmt.Scanln(&lixo)

		println("Digite um número válido")
	}

	var inputSimbolosDoTriangulo string

	fmt.Print("\nDigite o simbolo desejado para a construção do triangulo: ")
	fmt.Scanln(&inputSimbolosDoTriangulo)

	var inputBaseDoTriangulo int

	for {
		fmt.Print("\nDigite o tamanho da base do tringulo desejada: ")

		_, err := fmt.Scanln(&inputBaseDoTriangulo)

		if err == nil {
			break
		}
		var lixo string
		fmt.Scanln(&lixo)

		println("Digite um número válido")
	}

	CriarTriangulo(inputBaseDoTriangulo, inputSimbolosDoTriangulo)

	var burriceDoUsuario int = 0

LoopVermelho:
	for {
		var gostarDeVermelho string

		fmt.Print("\nVocê gosta de vermelho? ")
		fmt.Scanln(&gostarDeVermelho)

		switch strings.ToLower(gostarDeVermelho) {
		case "sim":
			println("Você é amigo")
			break LoopVermelho
		case "não", "nao":
			println("Você é inimigo")
			break LoopVermelho
		default:
			if burriceDoUsuario == 2 {
				println("\nVocê é burro? ")
			} else if burriceDoUsuario == 4 {
				println("Eu não mereço isso, adeus")
				break LoopVermelho
			} else {
				println("Digite uma resposta de verdade eu só entendo sim ou não")
			}
			burriceDoUsuario += 1
		}
	}

	fmt.Scanln()
}

func soma(a int, b int) int {
	return a + b
}

func faseDaVida(idade float32) {
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
		if idade < 1 {
			println("Como você ta me respondendo?")
		} else if idade < 10 {
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
