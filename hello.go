package main

import "fmt"

func main() {
	nome := "Dev"
	versao := 1.1
	fmt.Println("Olá, sr.", nome, ",sou uma aplicação em Go na versão", versao, "! Isso é incrível, não é?")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")


	var comando int
	fmt.Scanf("%d", &comando)


	fmt.Println("O comando escolhido foi", comando)
}