package main

import "fmt"

type FiguraGeometrica interface {
	area() float64
}

func escreverArea(f FiguraGeometrica) {
	fmt.Println("A área da figura é", f.area())
}

type (
	Retangulo struct {
		altura  float64
		largura float64
	}
	Circulo struct {
		raio float64
	}
)

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

func (c Circulo) area() float64 {
	return 3.14 * (c.raio * c.raio)
}

func main() {
	escreverArea(Retangulo{10, 15})
	escreverArea(Circulo{10})
}
