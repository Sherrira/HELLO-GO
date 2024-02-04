package main

import "fmt"

type Cliente struct {
	nome string
	passos int
	saldo float64 
}

func (c *Cliente) CalculaSaldo() {
	c.saldo = float64(c.passos) * 0.015
}

func (c *Cliente) Andou() {
	fmt.Printf("Parabéns, %v!\nVocê cumpriu a sua meta diária com %v passos!\nIsso gerou um saldo de R$%v a ser adicionado em sua conta!\nContinue assim!\n", c.nome, c.passos, c.saldo)
}

func main() {
	sherr := Cliente{
		nome: "Shérrira",
		passos: 16879,
	}

	sherr.CalculaSaldo()

	sherr.Andou()
	fmt.Printf("O valor da struct com nome é: %v", sherr.nome)
}
