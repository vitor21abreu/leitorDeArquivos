package application

import (
	"app/internal/domain"
	"sort"
)

type Item struct {
	Nome  string
	Valor int
}

func ContarModelos(vendas []domain.Venda) map[string]int {
	count := make(map[string]int)
	for _, v := range vendas {
		count[v.Modelo]++
	}
	return count
}

func ContarEstados(vendas []domain.Venda) map[string]int {
	count := make(map[string]int)
	for _, v := range vendas {
		count[v.Estado]++
	}
	return count
}

func ContarPeriodos(vendas []domain.Venda) map[string]int {
	count := make(map[string]int)
	for _, v := range vendas {
		count[v.Periodo]++
	}
	return count
}

func TopEMenos(count map[string]int) (string, string) {
	max := -1
	min := int(^uint(0) >> 1)

	var mais, menos string

	for modelo, qtd := range count {
		if qtd > max {
			max = qtd
			mais = modelo
		}
		if qtd < min {
			min = qtd
			menos = modelo
		}
	}

	return mais, menos
}

func Ordenar(count map[string]int) []Item {
	var lista []Item

	for k, v := range count {
		lista = append(lista, Item{k, v})
	}

	sort.Slice(lista, func(i, j int) bool {
		return lista[i].Valor > lista[j].Valor
	})

	return lista
}
