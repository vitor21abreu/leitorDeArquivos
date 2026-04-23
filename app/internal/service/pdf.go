package service

import (
	"app/internal/application"
	"app/internal/domain"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GeraRelatorioPDF(vendas []domain.Venda, nomeArquivo string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// ===== HEADER =====
	pdf.SetFillColor(40, 60, 90)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(190, 12, "DASHBOARD DE VENDAS", "", 1, "C", true, 0, "")
	pdf.Ln(5)

	pdf.SetTextColor(0, 0, 0)

	// ===== PROCESSAMENTO =====
	modelos := application.ContarModelos(vendas)

	mais, menos := application.TopEMenos(modelos)
	ranking := application.Ordenar(modelos)

	total := 0
	for _, v := range modelos {
		total += v
	}

	// ===== CARDS =====
	pdf.SetFont("Arial", "B", 12)

	drawCard := func(x, y float64, titulo, valor string) {
		pdf.SetXY(x, y)
		pdf.SetFillColor(230, 230, 230)
		pdf.CellFormat(60, 25, "", "1", 0, "", true, 0, "")

		pdf.SetXY(x+2, y+5)
		pdf.SetFont("Arial", "", 10)
		pdf.Cell(50, 5, titulo)

		pdf.SetXY(x+2, y+12)
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(50, 5, valor)
	}

	drawCard(10, 25, "Total de Vendas", fmt.Sprintf("%d", total))
	drawCard(75, 25, "Mais Vendido", mais)
	drawCard(140, 25, "Menos Vendido", menos)

	pdf.Ln(35)

	// ===== TABELA =====
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Ranking")
	pdf.Ln(8)

	pdf.SetFillColor(180, 200, 230)
	pdf.SetFont("Arial", "B", 12)

	pdf.CellFormat(90, 8, "Modelo", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 8, "Vendas", "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 12)

	for _, item := range ranking {
		if item.Nome == mais {
			pdf.SetFillColor(200, 255, 200) // destaque
		} else {
			pdf.SetFillColor(255, 255, 255)
		}

		pdf.CellFormat(90, 8, item.Nome, "1", 0, "", true, 0, "")
		pdf.CellFormat(40, 8, fmt.Sprintf("%d", item.Valor), "1", 1, "C", true, 0, "")
	}

	// ===== GRÁFICO =====
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Grafico de Vendas")
	pdf.Ln(10)

	max := 0
	for _, v := range ranking {
		if v.Valor > max {
			max = v.Valor
		}
	}

	barWidth := 30.0

	for _, item := range ranking {
		height := float64(item.Valor) / float64(max) * 40

		x := pdf.GetX()
		y := pdf.GetY()

		pdf.SetFillColor(100, 149, 237)
		pdf.Rect(x, y-height, barWidth, height, "F")

		pdf.SetY(y + 2)
		pdf.SetX(x)
		pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(barWidth, 5, item.Nome, "", 0, "C", false, 0, "")

		pdf.SetXY(x+barWidth+5, y)
	}

	return pdf.OutputFileAndClose(nomeArquivo)
}
