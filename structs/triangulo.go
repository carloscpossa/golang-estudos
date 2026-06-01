package estruturas

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}
