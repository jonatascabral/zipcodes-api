package services

import (
	"encoding/csv"
	"io"
)

type Csv struct {
	File io.Reader
	Header [][]string
	Data [][]string
}

func NewCsv(file io.Reader) (*Csv, error) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	csvFile := &Csv{File:file}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(csvFile.Header) == 0 {
			csvFile.Header = append(csvFile.Header, line)
		} else {
			csvFile.Data = append(csvFile.Data, line)
		}
	}
	return csvFile, nil
}