package application

import (
	"app/internal/application/leitores"
	"fmt"
	"path/filepath"
	"strings"
)

var Readers = map[string]func(string) ([]map[string]string, error){
	".csv":  leitores.LerCSV,
	".xlsx": leitores.LerXlsx,
}

func GetLeitor(ext string) (func(string) ([]map[string]string, error), error) {
	fn, ok := Readers[ext]
	if !ok {
		return nil, fmt.Errorf("leitor não encontrado para extensão: %s", ext)
	}

	return fn, nil
}

func LerArquivo(caminho string) ([]map[string]string, error) {
	ext := strings.ToLower(filepath.Ext(caminho))

	fn, err := GetLeitor(ext)
	if err != nil {
		return nil, err
	}

	return fn(caminho)
}
