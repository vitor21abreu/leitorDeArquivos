package models

import (
    "sort"
)

type Item struct {
    Nome  string
    Valor int
}


func ContarModelos(vendas []Venda) map[string]int {
    count := make(map[string]int)
    for _, v := range vendas {
        count[v.Modelo]++
    }
    return count
}

func ContarEstados(vendas []Venda) map[string]int {
    count := make(map[string]int)
    for _, v := range vendas {
        count[v.Estado]++
    }
    return count
}

func ContarPeriodos(vendas []Venda) map[string]int {
    count := make(map[string]int)
    for _, v := range vendas {
        count[v.Periodo]++
    }
    return count
}

func TopEMenos(count map[string]int) (string, string) {
    max := -1
    min := int(^uint(0) >> 1) // Representa o maior valor possível de um int

    var mais, menos string

    for chave, qtd := range count {
        if qtd > max {
            max = qtd
            mais = chave
        }
        if qtd < min {
            min = qtd
            menos = chave
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