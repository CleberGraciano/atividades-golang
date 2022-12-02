package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.56
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//mais perto do char que tem no go

	//Exibe o numero que esse caractere representa na tbela ASCII
	char := 'B'
	fmt.Println(char)

	var booleano1 bool
	fmt.Println(booleano1)

	//Por default o tipo error retorna Nil de zero
	var erro error
	fmt.Println(erro)
	var erro2 error = errors.New("Erro interno encontrado")
	fmt.Println(erro2)

}
