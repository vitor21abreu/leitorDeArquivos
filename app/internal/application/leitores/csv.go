package leitores

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LerCSV(caminho string) ([]map[string]string, error) {
	file, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	registros, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(registros) == 0 {
		return nil, fmt.Errorf("arquivo vazio")
	}

	var resultado []map[string]string

	headers := registros[0]

	for _, linha := range registros[1:] {
		item := make(map[string]string)

		for i, valor := range linha {
			if i < len(headers) {
				item[headers[i]] = valor
			}
		}

		resultado = append(resultado, item)
	}

	return resultado, nil
}
