package main

import 	"fmt"

var Nomes = [] string {"Shé", "Fill", "Iga", "Kalel", "Tofu", "Potro"} 
var Idades = [] int {33, 35, 60, 7, 4, 4}
var Status = [] bool {false, true, true, false, false, true} 
var Ruas = [] string {"Rosa", "Cravo", "Dália", "Margarida", "Erva Doce", "Alecrim"} 
var Numeros = [] int {13, 03, 38, 68, 44, 129}
var Cidades = [] string {"São Paulo", "Barueri", "Salto", "São Paulo", "Mauá", "Mauá"} 
var Estados = [] string {"SP", "SP", "SP", "SP", "SP", "SP"} 


type Client struct {
	Nome string
	Idade int
	Ativo bool
	Endereco
}
type Endereco struct {
	Rua string
	Numero int
	Cidade string
	Estado string
}

func main() {

	for i := 0; i<len(Nomes); i++ {
		client := Client {
			Nome: Nomes[i],
			Idade: Idades[i],
			Ativo: Status[i],
			Endereco: Endereco{
				Rua: Ruas[i],
				Numero: Numeros[i],
				Cidade: Cidades[i],
				Estado: Estados[i],
			},
		}
		
		fmt.Printf("\nNome do clinte: %v\n \tIdade do cliente: %v\n \tEndereço do cliente: Rua %v, nº %v, %v - %v\n", Nomes[i], Idades[i], Ruas[i], Numeros[i], Cidades[i], Estados[i])
		if client.Ativo != true {
			fmt.Println(" \tCliente inativo")
		} else {
			fmt.Println(" \tCliente ativo")
		}
	}
}
