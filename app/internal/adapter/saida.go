package adapter

import (
	"app/internal/application"
	"app/internal/service"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PastaSaidaComData() string {
	agora := time.Now()
	return filepath.Join("saida", agora.Format("2006-01"))
}

func ListarArquivos() ([]string, error) {
	dir := "../arquivos"

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var caminhos []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))

		if ext == ".csv" || ext == ".xlsx" {
			caminho := filepath.Join(dir, file.Name())
			caminhos = append(caminhos, caminho)
		}
	}

	return caminhos, nil
}

func ProcessarArquivos() {
	arquivos, err := ListarArquivos()
	if err != nil {
		panic(err)
	}

	dir := PastaSaidaComData()
	os.MkdirAll(dir, os.ModePerm)

	for _, caminho := range arquivos {

		data, err := application.LerArquivo(caminho)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		vendas := application.MapToVendas(data)

		nomeBase := filepath.Base(caminho)
		nomeSemExt := strings.TrimSuffix(nomeBase, filepath.Ext(nomeBase))

		nomePDF := filepath.Join(dir, nomeSemExt+".pdf")

		err = service.GeraRelatorioPDF(vendas, nomePDF)
		if err != nil {
			fmt.Println("Erro ao gerar PDF:", err)
			continue
		}

		fmt.Println("Relatório gerado:", nomePDF)
	}
}
