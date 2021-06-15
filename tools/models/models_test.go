package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stride-so/matrix/tools/test"
)

func loadMatrix(t *testing.T, path string) Matrix {
	t.Helper()
	r := test.FixtureReader(t, path)
	matrix, err := JSONRead(r)
	if err != nil {
		t.Fail()
	}
	return matrix
}

func TestOrphanedSkills(t *testing.T) {
	matrix := loadMatrix(t, "../testdata/matrix.json")
	_, list := orphanedSkill(matrix)
	assert.Len(t, list, 0)
}
func TestMatrixLevelSkills(t *testing.T) {
	matrix := loadMatrix(t, "../testdata/matrix.json")
	skills := matrix.Themes.Skills()
	assert.NotEmpty(t, skills)
}
