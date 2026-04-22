package application

import (
	"app/internal/application/leitores"
	"fmt"
	"path/filepath"
)

var Reader = map[string]func(string) ([]map[string]string, error){
	".csv":  leitores.LerCSV,
	".xlsx": leitores.LerXlsx,	
}

func GetLeitor(ext string) (func(string) ([]map[string]string, error), bool) {
	fn, ok := Reader[ext]
	return fn, ok
}

func LerArquivo(caminho string) ([]map[string]string, error) {
	ext := filepath.Ext(caminho)

	fn, ok := GetLeitor(ext)
	if !ok {
		return nil, fmt.Errorf("tipo não suportado")
	}

	return fn(caminho)
}

func LerXlsx(caminho string) ([]map[string]string, error) {
	// Implementação para ler arquivos Excel
	return nil, nil
}

func LerTXT(caminho string) ([]map[string]string, error) {
	// Implementação para ler arquivos TXT
	return nil, nil
}
