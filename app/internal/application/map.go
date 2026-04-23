package application

import "app/internal/domain"

func MapToVendas(data []map[string]string) []domain.Venda {
	var vendas []domain.Venda

	for _, item := range data {
		vendas = append(vendas, domain.Venda{
			Modelo:  item["modelo"],
			Ano:     item["ano"],
			Estado:  item["estado"],
			Periodo: item["periodo"],
		})
	}

	return vendas
}