package main

//funcoes_basico

import (
	"cursodego/fundamentos/funcoes_basica/matematica"
	"fmt"
)

func main() {

	resultado := matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da calculo foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 5, 4)
	fmt.Printf("O resultado do calculo foi: %d\r\n", resultado)

}
