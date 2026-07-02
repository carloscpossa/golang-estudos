package maps

import (
	"testing"
)

func TestBusca(t *testing.T) {

	t.Run("palavra conhecida", func(t *testing.T) {
		dicionario := Dicionario{"teste": "isso é apenas um teste"}

		resultado, _ := dicionario.Buscar("teste")

		esperado := "isso é apenas um teste"

		ComparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		dicionario := Dicionario{"teste": "isso é apenas um teste"}

		_, err := dicionario.Buscar("desconhecida")

		ComparaErro(t, err, ErrPalavraNaoEncontrada)
	})

}

func ComparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dados '%s'", resultado, esperado, "teste")
	}
}

func ComparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func TestAdiciona(t *testing.T) {
	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}

		palavra := "teste"
		definicao := "isso é apenas um teste"

		err := dicionario.Adicionar(palavra, definicao)

		ComparaErro(t, err, nil)

		ComparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {

		palavra := "teste"
		definicao := "isso é apenas um teste"

		dicionario := Dicionario{palavra: definicao}

		err := dicionario.Adicionar(palavra, "teste novo")

		ComparaErro(t, err, ErrPalavraExistente)

		ComparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func TestAtualizacao(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"

		dicionario := Dicionario{palavra: definicao}

		novaDefinicao := "nova definição"

		err := dicionario.Atualizar(palavra, novaDefinicao)

		ComparaErro(t, err, nil)
		ComparaDefinicao(t, dicionario, palavra, novaDefinicao)

	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"

		dicionario := Dicionario{}

		err := dicionario.Atualizar(palavra, definicao)

		ComparaErro(t, err, ErrPalavraInexistente)
	})
}

func TestRemocao(t *testing.T) {

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"

		dicionario := Dicionario{palavra: definicao}

		dicionario.Remover(palavra)

		_, err := dicionario.Buscar(palavra)
		ComparaErro(t, err, ErrPalavraNaoEncontrada)
	})

}

func ComparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Buscar(palavra)

	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}

	ComparaStrings(t, resultado, definicao)
}
