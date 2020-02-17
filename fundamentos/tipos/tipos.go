package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Rune, tipo que representa o valor inteiro de um caracter da tabela unicode
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32 e float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é:", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é:", bo)
	fmt.Println(!bo)

	//string
	s1 := "Olá, eu sou o Goku"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é:", len(s1))

	//string com múltiplas linhas
	s2 := `Olá
	eu
	sou
	o
	Goku`
	fmt.Println("O tamanho da string é:", len(s2))

	//char? não existe!
	//var x char = 'b' -- não existe tipo char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //retorna int32, representação na tabela unicode
	fmt.Println(char)


}
