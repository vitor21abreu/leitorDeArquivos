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

	pdf.SetFillColor(40, 60, 90)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(190, 12, "DASHBOARD DE VENDAS", "", 1, "C", true, 0, "")
	pdf.Ln(5)

	pdf.SetTextColor(0, 0, 0)

	modelos := application.ContarModelos(vendas)
	mais, menos := application.TopEMenos(modelos)
	ranking := application.Ordenar(modelos)

	total := 0
	maiorValor := 0
	for _, v := range modelos {
		total += v
		if v > maiorValor {
			maiorValor = v
		}
	}

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

	pdf.SetXY(10, 60)

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
			pdf.SetFillColor(200, 255, 200)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}

		pdf.CellFormat(90, 8, item.Nome, "1", 0, "", true, 0, "")
		pdf.CellFormat(40, 8, fmt.Sprintf("%d", item.Valor), "1", 1, "C", true, 0, "")
	}

	pdf.Ln(15)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Grafico de Vendas")
	pdf.Ln(12)

	if maiorValor == 0 {
		maiorValor = 1
	}

	baseY := pdf.GetY() + 40
	startX := 15.0
	barWidth := 25.0
	spacing := 10.0
	maxBarHeight := 35.0

	for i, item := range ranking {
		height := (float64(item.Valor) / float64(maiorValor)) * maxBarHeight
		currentX := startX + float64(i)*(barWidth+spacing)

		pdf.SetFillColor(100, 149, 237)
		pdf.Rect(currentX, baseY-height, barWidth, height, "F")

		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(currentX, baseY+2)
		pdf.CellFormat(barWidth, 5, item.Nome, "", 0, "C", false, 0, "")

		pdf.SetXY(currentX, baseY-height-5)
		pdf.CellFormat(barWidth, 5, fmt.Sprintf("%d", item.Valor), "", 0, "C", false, 0, "")
	}

	pdf.SetY(baseY + 15)

	return pdf.OutputFileAndClose(nomeArquivo)
}
