package views

import (
	"app/internal/models"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GeraRelatorioPDF(vendas []models.Venda, nomeArquivo string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "RELATORIO DE VENDAS DE CARROS")
	pdf.Ln(12)

	modelos := models.ContarModelos(vendas)
	estados := models.ContarEstados(vendas)
	periodos := models.ContarPeriodos(vendas)

	mais, menos := models.TopEMenos(modelos)
	periodoTop, _ := models.TopEMenos(periodos)

	ranking := models.Ordenar(modelos)

	pdf.SetFont("Arial", "", 12)

	pdf.Cell(40, 10, "Modelo mais vendido: "+mais)
	pdf.Ln(8)

	pdf.Cell(40, 10, "Modelo menos vendido: "+menos)
	pdf.Ln(8)

	pdf.Cell(40, 10, "Periodo com mais vendas: "+periodoTop)
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Ranking de Modelos")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for i, item := range ranking {
		linha := fmt.Sprintf("%d. %s - %d vendas", i+1, item.Nome, item.Valor)
		pdf.Cell(40, 8, linha)
		pdf.Ln(6)
	}

	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Vendas por Estado")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for estado, qtd := range estados {
		linha := fmt.Sprintf("%s - %d vendas", estado, qtd)
		pdf.Cell(40, 8, linha)
		pdf.Ln(6)
	}

	return pdf.OutputFileAndClose(nomeArquivo)
}
