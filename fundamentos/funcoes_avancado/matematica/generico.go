package matematica

//Multiplicador multiplica x vezes o y
func Multiplicador(x int, y int) int {
	return x * y
}

/*
Calculo executa qualquer tipo de calculo basta enviar a funcao desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Divisor efetua a divisão do numeroA pelo numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto efetua a divisão do numeroA pelo numeroB e retorna o resultado e o resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = Divisor(numeroA, numeroB)
	resto = numeroA % numeroB
	return
}
