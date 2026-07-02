package maps

type Dicionario map[string]string

func (d Dicionario) Buscar(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrPalavraNaoEncontrada
	}

	return definicao, nil
}

func (d Dicionario) Adicionar(palavra, definicao string) error {

	_, existe := d[palavra]

	if existe {
		return ErrPalavraExistente
	}

	d[palavra] = definicao
	return nil
}

func (d Dicionario) Atualizar(palavra, novaDefinicao string) error {
	_, err := d.Buscar(palavra)

	switch err {
	case ErrPalavraNaoEncontrada:
		return ErrPalavraInexistente
	case nil:
		{
			d[palavra] = novaDefinicao
			return nil
		}
	default:
		return nil
	}
}

func (d Dicionario) Remover(palavra string) {
	delete(d, palavra)
}
