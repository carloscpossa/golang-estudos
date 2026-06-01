package iteracao

func Repetir(caractere string, quantidadeRepeticoes int) (repeticao string) {
	for range quantidadeRepeticoes {
		repeticao += caractere
	}
	return
}
