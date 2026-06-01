package estruturas

import "testing"

func TestPerimetro(t *testing.T) {

	verificaPerimetro := func(t *testing.T, resultado, esperado float64) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("Perímetro do retângulo", func(t *testing.T) {
		retangulo := Retangulo{10.0, 10.0}
		resultado := Perimetro(retangulo)
		esperado := 40.0

		verificaPerimetro(t, resultado, esperado)
	})

}

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("%#v resultado %.2f, esperado %.2f", forma, resultado, esperado)
		}
	}

	t.Run("Area do retângulo", func(t *testing.T) {
		retangulo := Retangulo{Largura: 12.0, Altura: 6.0}
		esperado := 72.0

		verificaArea(t, retangulo, esperado)
	})

	t.Run("Área do círculo", func(t *testing.T) {
		circulo := Circulo{Raio: 10}
		esperado := 314.1592653589793

		verificaArea(t, circulo, esperado)
	})

	t.Run("Área do triângulo", func(t *testing.T) {
		triangulo := Triangulo{12, 6}
		esperado := 36.0

		verificaArea(t, triangulo, esperado)
	})
}
