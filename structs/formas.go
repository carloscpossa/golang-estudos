package estruturas

func Perimetro(retangulo Retangulo) float64 {
	return (retangulo.Altura + retangulo.Largura) * 2
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Altura * retangulo.Largura
}
