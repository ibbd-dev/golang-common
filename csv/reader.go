package csv

import (
	"os"

	goCsv "encoding/csv"
)

type Reader struct {
	Filename string

	Csv *goCsv.Reader

	file *os.File
}

func NewReader(filename string, params *goCsv.Reader) (r *Reader, err error) {
	r = &Reader{}

	r.file, err = os.Open(filename)
	if err != nil {
		return r, err
	}
	r.Csv = goCsv.NewReader(r.file)
	r.Csv.Comma = params.Comma
	r.Csv.LazyQuotes = params.LazyQuotes
	r.Csv.TrimLeadingSpace = params.TrimLeadingSpace
	r.Csv.FieldsPerRecord = params.FieldsPerRecord

	return r, nil
}

func (r *Reader) Close() {
	r.file.Close()
}
