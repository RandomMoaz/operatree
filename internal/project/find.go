package project

import (
	"fmt"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/subject"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	subjectTypes []string = []string{
		"EVENT",
		"TOPIC",
		"OBJECTIVE",
	}
)

// find subject(s) inside project tree using any string. Find uses metadata only not actual
// file contents or actual dir structure
func FindSubjects(p *Project, st string, term string) (subject.Subject, error) {

	types := fuzzy.FindFold(st, subjectTypes)

	var t subject.SubjectType
	if len(types) != 0 {
		t = subject.SubjectType(types[0])
	}

	db := BuildSearchDB(p)

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
			display := fmt.Sprintf("%-10s  %-30s  %s > ",
				string(db[i].Subject.Type),
				db[i].Subject.Name,
				db[i].Subject.DirName,
			)
			// pad display to fixed width, then append SearchStr for matching
			// fuzzyfinder matches against the full string but only displays what fits the terminal
			return fmt.Sprintf("%-120s  %s", display, db[i].SearchStr)
		},
		fuzzyfinder.WithQuery(term),
		fuzzyfinder.WithHeader("Matching Subjects"),
		fuzzyfinder.WithPromptString("Search subjects for > "),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := db[i].Subject
			return fmt.Sprintf(
				"Name:     %s\nDate:     %s\nLocation: %s\nNotes:    %s\nTags:     %s",
				s.Name, s.Date, s.Location, s.Notes,
				strings.Join(s.Tags, ", "),
			)
		}),
	)
	if err != nil {
		return subject.Subject{}, nil
	}

	return db[idx].Subject, nil
}
