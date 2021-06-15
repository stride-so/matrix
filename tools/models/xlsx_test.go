package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stride-so/matrix/tools/test"
)

func TestXLSXReader(t *testing.T) {
	r := test.FixtureReader(t, "../testdata/matrix.xlsx")
	matrix, err := XLSXRead(r)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(matrix.Levels))
	assert.Equal(t, 5, len(matrix.Themes))
}

func TestXLSXWriter(t *testing.T) {
	r := test.FixtureReader(t, "../testdata/matrix.json")
	matrix, err := JSONRead(r)
	assert.NoError(t, err)
	w := test.TmpWriter(t, "out-*.xlsx")
	err = XLSXWrite(w, matrix)
	assert.NoError(t, err)
}
