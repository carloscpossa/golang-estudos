package maps

const (
	ErrPalavraNaoEncontrada = ErrDicionario("Não foi possível encontrar a palavra que você procura")
	ErrPalavraInexistente   = ErrDicionario("Palavra inexistente no dicionário")
	ErrPalavraExistente     = ErrDicionario("Palavra já existente no dicionário")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}
