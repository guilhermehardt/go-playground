// Working with interfaces in GO
package main

import "fmt"

type veiculo interface {
	abastecer()
	trocarOleo() bool
}

type automovel struct {
	modelo string
	ano    int
	km     float64
}

type carro struct {
	automovel
	portas   int
	completo bool
}

type pickup struct {
	automovel
	tipo      string
	tracao4x4 bool
}

// métodos
func (c carro) abastecer() {
	fmt.Println("O carro ", c.modelo, "foi abastecido")
}

func (c carro) trocarOleo() bool {
	if c.km > 8.000 {
		return true
	}
	return false
}

func (p pickup) abastecer() {
	fmt.Println("A Pick-up ", p.modelo, "foi abastecida")
}

func (p pickup) trocarOleo() bool {
	if p.km > 10.000 {
		return true
	}
	return false
}

func manutencao(v veiculo) {
	fmt.Println("# Manutenção ")
	v.abastecer()

	if v.trocarOleo() {
		fmt.Println("e o óleo foi trocado")
	}
}

func main() {
	veiculo1 := carro{
		automovel: automovel{"Audi A3 Sportback", 2019, 7.225},
		portas:    4,
		completo:  true,
	}

	veiculo2 := pickup{
		automovel: automovel{"Amarok TDI", 2020, 12.582},
		tipo:      "Offroad",
		tracao4x4: true,
	}

	manutencao(veiculo1)
	manutencao(veiculo2)
}
