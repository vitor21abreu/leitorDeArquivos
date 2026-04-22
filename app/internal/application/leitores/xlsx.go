package leitores

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func LerXlsx(caminho string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(caminho)
	if err != nil {
		return nil, err
	}

	linhas, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	if len(linhas) == 0 {
		return nil, fmt.Errorf("arquivo vazio")
	}

	headers := linhas[0]
	var resultado []map[string]string

	for _, linha := range linhas[1:] {
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
