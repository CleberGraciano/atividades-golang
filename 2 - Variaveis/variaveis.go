package main

import "fmt"

func main() {

	var variavel1 string = "Variavel1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"

	fmt.Println(variavel5, variavel6)

	const constan1 string = "Constant 1"
	fmt.Println(constan1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}
