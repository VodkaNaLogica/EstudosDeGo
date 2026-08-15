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
	soma(2, 2)
	fmt.Scanln()
}

func soma(a int, b int) int {
	fmt.Println(a + b)
	return 0
}
