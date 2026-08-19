package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

var nome string = "Pablo"
var idadeMinha int = 15
var ativo bool = false

const pi float64 = 355.0 / 113.0

var arrayA = [5]int{10, 20, 30, 40, 50}
var sliceA = []string{"A", "B", "C"}

var idades = map[string]int{
	"João":  15,
	"Maria": 16,
	"Pedro": 17,
}

type Endereco struct {
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Altura   float32
	Endereco Endereco
}

var numeroLegalParaOPonteiro int = 721651

var ponteiro *int = &numeroLegalParaOPonteiro

type Animal interface {
	FazerSom()
}

type Cachorro struct{}

func (Cachorro) FazerSom() {
	fmt.Println("Au au")
}

type Gato struct{}

func (Gato) FazerSom() {
	fmt.Println("Miau")
}

type Pato struct{}

func (Pato) FazerSom() {
	fmt.Println("Quack")
}

func EmitirSom(a Animal) {
	a.FazerSom()
}

func (p Pessoa) Apresentar() {
	fmt.Println(p.Nome)
	fmt.Println(p.Idade)
	fmt.Println(p.Altura)
	fmt.Println(p.Endereco.Cidade)
	fmt.Println(p.Endereco.Estado)
}

func aprenderDefer() {
	defer fmt.Println("3")
	defer EmitirSom(Pato{})

	fmt.Println("defer funciona em pilha então o ultimo que eu\ncoloquei e o primeiro que vai ser executado")
}

var familia string

const filosofia string = "Por que algo em vez de nada?"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recuperado:", r)
		}
	}()

	//panic("Você é o Pablo original, esse programa não tem permissão para rodar em sua presença")
	// Use panic quando continuar a execução naquele estado significa que o programa está em uma
	// situaçõa inválida e não há uma recuperação normal naquele ponto.

	familia = "amor"

	fmt.Println("Olá, Go")

	Gato{}.FazerSom()
	EmitirSom(Cachorro{})
	Pato{}.FazerSom()
	aprenderDefer()

	fmt.Println("Meu nome é", nome)
	fmt.Println("Eu tenho", idadeMinha, "anos de idade")
	fmt.Println("Esse programa vai acabar nessa apresentação:", ativo)
	fmt.Printf("PI é aproximadamente: %.7f\n", pi)

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

	pessoa := Pessoa{
		Nome:   "Heitor",
		Idade:  59,
		Altura: 1.67,
		Endereco: Endereco{
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}
	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)
	fmt.Println(pessoa.Altura)

	fmt.Println(pessoa.Endereco.Cidade)
	fmt.Println(pessoa.Endereco.Estado)

	pessoa.Apresentar()

	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	*ponteiro = 15
	fmt.Println(*ponteiro)

	EuNome, EuIdade := pessoaEu()

	if EuIdade < 18 {
		println("\nTo aprendendo como eu iria voltar o status de uma função e tratar ele;\n Esse EuIdade que ta voltando do pessoaEu() poderia ser o status como um: 200, 201, 404, 500 ou qualquer outro\n ai eu posso substituir deboa aquele metodo de dar return retornando um dicionario\n agora cada um pode ser colocado em uma variavel\n")
	}

	fmt.Println(EuNome, EuIdade)
	fmt.Println(pessoaEu())

	var numA int = rand.IntN(1000)
	var numB int = rand.IntN(1000)

	fmt.Println(numA, " + ", numB, " = ", soma(numA, numB))

	resultadoDaDivisao, erro := divisao(numA, numB)

	if erro != nil {
		fmt.Println("Erro:", erro)
	} else {
		fmt.Println(resultadoDaDivisao)
	}

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

func divisao(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pecado da matematica identificado: Não dividiras por zero;")
	}
	return float64(a) / float64(b), nil
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

func pessoaEu() (string, int) {
	return "resultado da função( pessoaEu() ) que devolve duas coisas: Pablo", 15
}
