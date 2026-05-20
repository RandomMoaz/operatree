package project

import (
	"slices"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

type SearchDB struct {
	AbsPath    string
	SearchStr  string
	Subject    subject.Subject
	ModulePath []string // breadcrumb of module names from root to current
	SubjectIdx int
}

func BuildSearchDB(p *Project) []SearchDB {
	db := []SearchDB{}
	for _, m := range p.Modules {
		db = append(db, walkModule(m, []string{})...)
	}
	return db
}

func walkModule(m module.Module, path []string) []SearchDB {
	db := []SearchDB{}
	currentPath := append([]string{}, path...) // copy to avoid slice mutation
	currentPath = append(currentPath, m.Name)

	for j, s := range m.Subjects {
		a := slices.Concat(s.Tags, s.Paricipants)
		a = append(a, s.Name)
		a = append(a, s.Notes)
		a = append(a, s.Date)
		a = append(a, s.Location)

		db = append(db, SearchDB{
			AbsPath:    s.DirName,
			SearchStr:  strings.Join(a, ","),
			Subject:    s,
			ModulePath: currentPath,
			SubjectIdx: j,
		})
	}

	for _, sub := range m.Modules {
		db = append(db, walkModule(sub, currentPath)...)
	}

	return db
}
