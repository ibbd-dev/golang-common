package csv

import (
	"os"

	goCsv "encoding/csv"
)

type Writer struct {
	Filename string

	Csv *goCsv.Writer

	file *os.File
}

func NewWriter(filename string) (w *Writer, err error) {
	w = &Writer{}

	w.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return w, err
	}
	w.Csv = goCsv.NewWriter(w.file)

	return w, nil
}

func (w *Writer) Close() {
	w.file.Close()
}
