package main

import "fmt"

type Conta struct {	nome string
	saldo int
}
func NewConta() *Conta {
	return &Conta{saldo:0}
}
func (c *Conta) Simulador(valor int) int {
	c.saldo += valor
	fmt.Println(c.saldo)
	return c.saldo
}


func main() {
	conta := Conta{saldo: 100}
	conta.Simulador(200)
	}


