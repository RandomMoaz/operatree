package project

import (
	"fmt"

	"github.com/hanymamdouh82/operatree/pkg/subject"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	subjectTypes []string = []string{
		string(subject.SubjectEvent),
		string(subject.SubjectTask),
		string(subject.SubjectTopic),
		string(subject.SubjectObjective),
		string(subject.SubjectDataSource),
	}
)

// find subject(s) inside project tree using any string. Find uses metadata only not actual
// file contents or actual dir structure
func FindSubject(p *Project, st string, term string) (subject.Subject, error) {

	types := fuzzy.FindFold(st, subjectTypes)

	var t subject.SubjectType
	if len(types) != 0 {
		if len(types) != len(subjectTypes) {
			t = subject.SubjectType(types[0])
		} else {
			t = ""
		}
	}

	db := BuildSearchDB(p)
	if len(db) == 0 {
		return subject.Subject{}, fmt.Errorf("project doesn't contain any subjects yet")
	}

	// Optionally filter by type
	if t != "" {
		filtered := db[:0]
		for _, entry := range db {
			if entry.Subject.Type == t {
				filtered = append(filtered, entry)
			}
		}
		db = filtered
	}

	idx, err := fuzzyfinder.Find(
		db,
		func(i int) string {
			// display := fmt.Sprintf("%-10s  %-30s  %s > ",
			display := fmt.Sprintf("%-10s  %s > ",
				string(db[i].Subject.Type),
				db[i].Subject.Name,
				// db[i].Subject.DirName,
			)
			// pad display to fixed width, then append SearchStr for matching
			// fuzzyfinder matches against the full string but only displays what fits the terminal
			return fmt.Sprintf("%-120s  %s", display, db[i].SearchStr)
		},
		// fuzzyfinder.WithMode(fuzzyfinder.ModeCaseSensitive),
		fuzzyfinder.WithQuery(term),
		fuzzyfinder.WithHeader(fmt.Sprintf("Matching %s", t)),
		fuzzyfinder.WithPromptString("Search subjects for > "),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := db[i].Subject
			return formatPreview(s)
		}),
	)
	if err != nil {
		return subject.Subject{}, err
	}

	return db[idx].Subject, nil
}

// find subject(s) inside project tree using any string. Find uses metadata only not actual
// file contents or actual dir structure
func FindSubjectsSilent(p *Project, st string, term string) ([]subject.Subject, error) {

	types := fuzzy.FindFold(st, subjectTypes)

	var t subject.SubjectType
	if len(types) != 0 {
		if len(types) != len(subjectTypes) {
			t = subject.SubjectType(types[0])
		} else {
			t = ""
		}
	}

	db := BuildSearchDB(p)
	normalizedDB := make([]string, len(db))
	for i, v := range db {
		normalizedDB[i] = v.SearchStr
	}

	if len(db) == 0 {
		return nil, fmt.Errorf("project doesn't contain any subjects yet")
	}

	// Optionally filter by type
	if t != "" {
		filtered := db[:0]
		for _, entry := range db {
			if entry.Subject.Type == t {
				filtered = append(filtered, entry)
			}
		}
		db = filtered
	}

	// the real search using FindFold

	r := fuzzy.FindFold(term, normalizedDB)

	// converting returned string results to subjects by exact matching with original db slice
	// for exact matches, returns subject
	results := make([]subject.Subject, 0)

	for _, v := range r {
		for j, dbv := range db {
			if dbv.SearchStr == v {
				results = append(results, db[j].Subject)
			}
		}
	}

	return results, nil
}
