package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const ingles = "inglês"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hello, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
