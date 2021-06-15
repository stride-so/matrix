package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// JSONRead reads a matrix JSON file at the specified path and returns the Matrix
// model and any relevant error.
func JSONRead(r io.ReadCloser) (matrix Matrix, err error) {
	defer r.Close()
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	json.Unmarshal(buf, &matrix)
	return
}

func JSONWrite(w io.WriteCloser, matrix Matrix) error {
	defer w.Close()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(matrix)
}
