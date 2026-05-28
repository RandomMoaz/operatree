package project

import "github.com/hanymamdouh82/operatree/pkg/subject"

// Returns subjects by type. It takes care of recursive walk for nested modules
func ListSubjects(p *Project, st subject.SubjectType) []subject.Subject {

	db := BuildSearchDB(p)

	if string(st) != "" {
		filtered := db[:0]
		for _, entry := range db {
			if entry.Subject.Type == st {
				filtered = append(filtered, entry)
			}
		}
		db = filtered
	}

	ss := make([]subject.Subject, 0, len(db))
	for _, s := range db {
		ss = append(ss, s.Subject)
	}

	return ss
}
