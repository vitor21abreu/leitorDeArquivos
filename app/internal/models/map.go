package models


func MapToVendas(data []map[string]string) []Venda {
	var vendas []Venda

	for _, item := range data {
		vendas = append(vendas, Venda{
			Modelo:  item["modelo"],
			Ano:     item["ano"],
			Estado:  item["estado"],
			Periodo: item["periodo"],
		})
	}

	return vendas
}