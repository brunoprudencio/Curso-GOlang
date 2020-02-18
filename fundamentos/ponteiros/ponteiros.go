package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //null

	p = &i // pegando enedereço de i e atribuindo para o ponteiro
	*p++ // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
