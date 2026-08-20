package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {

	t.Run("valida o servidor mais rápido", func(t *testing.T) {
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)

		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, _ := Corredor(URLLenta, URLRapida)

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10 segundos", func(t *testing.T) {
		servidorA := criarServidorComAtraso(21 * time.Millisecond)

		servidorB := criarServidorComAtraso(22 * time.Millisecond)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Configuravel(servidorA.URL, servidorB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(atraso)
			w.WriteHeader(http.StatusOK)
		}))
}
