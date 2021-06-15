package models

import (
	"fmt"
	"strings"

	"github.com/stride-so/matrix/tools/unique"
)

// Author is the original author or owner of this document
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// License is the license of the competency matrix
type License string

// Company is the organization that this matrix belongs to
type Company struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Tracks is a collection of Track types
type Tracks []string

// Themes is a collection of labeled skills
type Themes []Theme

func (t Themes) Skills() []string {
	set := unique.StringSlice{}
	for _, tt := range t {
		for _, s := range tt.Skills {
			set.Add(s)
		}
	}
	return set.Values()
}

// Theme is a named set of skills
type Theme struct {
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
}

// Level describes the position in the matrix, the titles of those positions
// and the core skills required for those roles.
type Level struct {
	Name   string `json:"name"`
	Track  string `json:"track"`
	Titles Titles `json:"titles"`
	Detail string `json:"detail"`
	Skills Skills `json:"skills"`
}
type Titles []string

func (t Titles) String() string {
	return fmt.Sprintf("[%v]", strings.Join(t, ", "))
}

// Levels is a collection of Level types.
type Levels []Level

func (l Levels) Map() map[string]Level {
	set := make(map[string]Level)
	for _, v := range l {
		set[v.Name] = Level{
			Name:   v.Name,
			Detail: v.Detail,
			Track:  v.Track,
			Titles: v.Titles,
			Skills: v.Skills,
		}
	}
	return set
}

func (l Levels) Skills() []string {
	set := unique.StringSlice{}
	for _, ll := range l {
		for _, s := range ll.Skills {
			set.Add(s.Name)
		}
	}
	return set.Values()
}

// Skill is a specific named core competency.
type Skill struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

// Skills is a collection
type Skills []Skill

func (s Skills) Map() map[string]string {
	set := make(map[string]string)
	for _, v := range s {
		set[v.Name] = v.Body
	}
	return set
}

// Matrix is the data structure containing an open comptency matrix
type Matrix struct {
	Version int64   `json:"version"`
	Author  Author  `json:"author"`
	Company Company `json:"company"`
	Tracks  Tracks  `json:"tracks"`
	Themes  Themes  `json:"themes"`
	Levels  Levels  `json:"levels"`
}

// Validate returns all the issues with the data structure.
// TODO: ensure all skills belong to one theme.
func (m *Matrix) Validate() (ok bool, errors []error) {
	return true, []error{}
}

func orphanedSkill(m Matrix) (ok bool, orphaned []string) {
	// check to see if any level skills are _not_ defined in the theme skills.
	l := unique.NewStringSlice(m.Themes.Skills())
	r := unique.NewStringSlice(m.Levels.Skills())
	for _, v := range r.Values() {
		if !l.Has(v) {
			orphaned = append(orphaned, v)
		}
	}
	return len(orphaned) == 0, orphaned
}

func (m *Matrix) Table() {
	fmt.Println("OK")
}
