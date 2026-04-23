package views

import (
	"app/internal/models"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GeraRelatorioPDF(vendas []models.Venda, nomeArquivo string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// --- CABEÇALHO ---
	pdf.SetFillColor(40, 60, 90)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(190, 15, "RELATÓRIO ANALÍTICO DE VENDAS", "", 1, "C", true, 0, "")
	pdf.Ln(5)

	pdf.SetTextColor(0, 0, 0)

	// --- DADOS DO MODEL ---
	modelos := models.ContarModelos(vendas)
	estados := models.ContarEstados(vendas)
	periodos := models.ContarPeriodos(vendas)

	mais, menos := models.TopEMenos(modelos)
	periodoTop, _ := models.TopEMenos(periodos)
	ranking := models.Ordenar(modelos)

	total := 0
	maiorValor := 0
	for _, v := range modelos {
		total += v
		if v > maiorValor {
			maiorValor = v
		}
	}

	// --- CARDS DE INDICADORES (KPIs) ---
	drawCard := func(x, y float64, titulo, valor string) {
		pdf.SetXY(x, y)
		pdf.SetFillColor(245, 245, 245)
		pdf.CellFormat(44, 22, "", "1", 0, "", true, 0, "")
		pdf.SetXY(x+2, y+4)
		pdf.SetFont("Arial", "", 9)
		pdf.Cell(40, 5, titulo)
		pdf.SetXY(x+2, y+10)
		pdf.SetFont("Arial", "B", 11)
		if len(valor) > 16 {
			valor = valor[:13] + "..."
		}
		pdf.Cell(40, 5, valor)
	}

	drawCard(10, 30, "Total Vendas", fmt.Sprintf("%d", total))
	drawCard(56, 30, "Mais Vendido", mais)
	drawCard(102, 30, "Menos Vendido", menos)
	drawCard(148, 30, "Melhor Período", periodoTop)

	// --- TABELA DE RANKING ---
	pdf.SetXY(10, 65)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Ranking de Modelos")
	pdf.Ln(8)

	pdf.SetFillColor(180, 200, 230)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(120, 8, "Modelo", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 8, "Qtd Vendas", "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 12)
	for _, item := range ranking {
		if item.Nome == mais {
			pdf.SetFillColor(200, 255, 200) // Verde: Top 1
		} else if item.Nome == menos {
			pdf.SetFillColor(255, 200, 200) // Vermelho: Último
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		pdf.CellFormat(120, 8, item.Nome, "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 8, fmt.Sprintf("%d", item.Valor), "1", 1, "C", true, 0, "")
	}

	// --- RESUMO POR ESTADO ---
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Distribuição Regional:")
	pdf.SetFont("Arial", "", 10)
	strEstados := ""
	for est, qtd := range estados {
		strEstados += fmt.Sprintf("%s (%d)   ", est, qtd)
	}
	pdf.Ln(8)
	pdf.MultiCell(0, 5, strEstados, "", "L", false)

	// --- GRÁFICO DE BARRAS ---
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Gráfico de Performance (Top 5)")
	pdf.Ln(12)

	if maiorValor == 0 {
		maiorValor = 1
	}
	baseY := pdf.GetY() + 35
	startX := 15.0
	barWidth := 25.0
	spacing := 10.0
	maxBarHeight := 30.0

	for i, item := range ranking {
		if i > 4 {
			break
		} // Limite de 5 barras para não estourar a folha
		height := (float64(item.Valor) / float64(maiorValor)) * maxBarHeight
		currentX := startX + float64(i)*(barWidth+spacing)

		pdf.SetFillColor(100, 149, 237) // Azul Cornflower
		pdf.Rect(currentX, baseY-height, barWidth, height, "F")

		pdf.SetFont("Arial", "", 7)
		pdf.SetXY(currentX, baseY+2)
		pdf.CellFormat(barWidth, 5, item.Nome, "", 0, "C", false, 0, "")

		pdf.SetXY(currentX, baseY-height-5)
		pdf.CellFormat(barWidth, 5, fmt.Sprintf("%d", item.Valor), "", 0, "C", false, 0, "")
	}

	return pdf.OutputFileAndClose(nomeArquivo)
}
