package arrays

func Soma(numeros []int) int {

	soma := 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {

	var somas []int

	for _, slice := range numerosParaSomar {
		somas = append(somas, Soma(slice))
	}

	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int

	for _, slice := range numerosParaSomar {
		if len(slice) > 1 {
			somas = append(somas, Soma(slice[1:]))
		} else {
			somas = append(somas, 0)
		}

	}

	return somas
}
