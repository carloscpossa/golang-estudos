package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Go!`

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}

		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})

}
