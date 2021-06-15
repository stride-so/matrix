package models

import (
	"fmt"
	"io"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	MatrixSheet = "matrix"
)

var (
	ErrNoSheet = fmt.Errorf("no matrix sheet. there must be a single sheet named '%s'", MatrixSheet)
)

func makeStyles(f *excelize.File) (map[string]int, error) {
	result := make(map[string]int)
	styles := map[string]*excelize.Style{
		"level": {
			Alignment: &excelize.Alignment{
				WrapText: true,
			},
		},
		"skill": {
			Alignment: &excelize.Alignment{
				Vertical: "top",
			},
		},
	}
	for k, v := range styles {
		id, err := f.NewStyle(v)
		if err != nil {
			return result, err
		}
		result[k] = id
	}
	return result, nil
}

// this should also check for a single or multiple sheets named matrix at some point.
func hasMatrixSheet(f *excelize.File) bool {
	for _, v := range f.GetSheetList() {
		if strings.ToLower(v) == MatrixSheet {
			return true
		}
	}
	return false
}

func XLSXRead(r io.Reader) (Matrix, error) {
	matrix := Matrix{}
	f, err := excelize.OpenReader(r)
	if err != nil {
		return matrix, err
	}
	if !hasMatrixSheet(f) {
		return matrix, ErrNoSheet
	}
	rows, err := f.GetRows(MatrixSheet)
	if err != nil {
		return matrix, err
	}

	// read level names
	levelRow := rows[0]
	for i := 2; i < len(levelRow); i++ {
		name := levelRow[i]
		if name == "" {
			continue
		}
		matrix.Levels = append(matrix.Levels, Level{
			Name: name,
		})
	}

	// read title
	// TODO handle more than one + array format [one, two, three]
	titleRow := rows[1]
	for i := 2; i < len(titleRow); i++ {
		title := titleRow[i]
		if title == "" {
			continue
		}
		matrix.Levels[i-2].Titles = Titles{title}
	}

	// read detail value
	detailRow := rows[2]
	for i := 2; i < len(detailRow); i++ {
		detail := detailRow[i]
		if detail == "" {
			continue
		}
		matrix.Levels[i-2].Detail = detail
	}

	// skills
	for i := 3; i < len(rows); i++ {
		row := rows[i]
		skill := row[1]
		for i := 2; i < len(row); i++ {
			body := row[i]
			if skill == "" || body == "" {
				continue
			}
			matrix.Levels[i-2].Skills = append(matrix.Levels[i-2].Skills, Skill{
				Name: skill,
				Body: row[i],
			})
		}
	}

	// the ability to add to a map should all be part of the theme model instead
	// doing this here is just going to be more complex and harder to maintain.
	themes := make(map[string]map[string]struct{})
	for i := 3; i < len(rows); i++ {
		row := rows[i]
		theme := row[0]
		skill := row[1]
		if theme == "" {
			continue
		}
		if _, ok := themes[theme]; !ok {
			themes[theme] = map[string]struct{}{skill: {}}
		} else {
			themes[theme][skill] = struct{}{}
		}
	}
	for k, v := range themes {
		skills := []string{}
		for s := range v {
			skills = append(skills, s)
		}
		matrix.Themes = append(matrix.Themes, Theme{
			Title:  k,
			Skills: skills,
		})
	}

	return matrix, nil
}

func XLSXWrite(w io.WriteCloser, matrix Matrix) error {
	f := excelize.NewFile()
	styles, err := makeStyles(f)
	if err != nil {
		return err
	}

	f.SetSheetName("Sheet1", MatrixSheet)
	f.SetCellValue(MatrixSheet, "A1", "Themes")
	f.SetCellValue(MatrixSheet, "A2", "Titles")
	f.SetCellValue(MatrixSheet, "B1", "Skills")
	f.SetCellValue(MatrixSheet, "A3", "Detail")

	set := (func(f *excelize.File, sheet string) func(string, int, string) {
		return func(col string, row int, value string) {
			f.SetCellValue(sheet, fmt.Sprintf("%v%v", col, row), value)
		}
	})(f, MatrixSheet)

	// Write all the Themes and Skills
	row := 3
	skills := []string{}
	for _, theme := range matrix.Themes {
		for _, skill := range theme.Skills {
			row += 1
			skills = append(skills, skill)
			set("A", row, theme.Title)
			set("B", row, skill)
		}
	}
	err = f.SetCellStyle(MatrixSheet, "A1", "B99", styles["skill"])
	if err != nil {
		return err
	}
	err = f.SetColWidth(MatrixSheet, "A", "B", 20)
	if err != nil {
		return err
	}

	// Write all the Levels
	for i, level := range matrix.Levels {
		col, err := excelize.ColumnNumberToName(3 + i)
		if err != nil {
			return err
		}
		set(col, 1, level.Name)
		set(col, 2, level.Titles.String())
		set(col, 3, level.Detail)
		row = 3
		levelSkills := level.Skills.Map()
		for _, v := range skills {
			row += 1
			skill, ok := levelSkills[v]
			if !ok {
				continue
			}
			set(col, row, skill)
		}
		if err := f.SetColWidth(MatrixSheet, col, col, 50); err != nil {
			return err
		}
		if err := f.SetCellStyle(
			MatrixSheet,
			fmt.Sprintf("%v3", col),
			fmt.Sprintf("%v99", col),
			styles["level"],
		); err != nil {
			return err
		}
	}

	err = f.Write(w)
	return err
}
