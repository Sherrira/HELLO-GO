package main

import (
	"fmt"
)

type Client struct {
	Nome string
	Idade int
	Ativo bool
}

func (c Client) desativarStatus() {
	c.Ativo = false
	fmt.Printf("O cliente %v foi desativado.", c.Nome)
}

func main() {
	clients := Client {
		Nome : "Shérrira", 
		Idade: 33,
		Ativo: true, 
	}

	fmt.Println(clients.Nome, clients.Idade)

	if clients.Ativo != true {
		fmt.Println("Cliente inativo")
	} else {
		fmt.Println("Cliente ativo")
	}
	
	clients.desativarStatus()


}
