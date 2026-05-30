package config

import (
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

func Find(cfg Config) (Project, error) {

	db := cfg.Projects

	idx, err := fuzzyfinder.Find(
		db,
		func(i int) string {
			display := fmt.Sprintf("[%-10s]  %s",
				strings.ToUpper(db[i].Template),
				db[i].Name,
			)
			// pad display to fixed width, then append SearchStr for matching
			// fuzzyfinder matches against the full string but only displays what fits the terminal
			return fmt.Sprintf("%-120s", display)
		},
		// fuzzyfinder.WithMode(fuzzyfinder.ModeCaseSensitive),
		fuzzyfinder.WithPromptString("Search projects for > "),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := fmt.Sprintf("Name: %s\nPath: %s\nTemplate: %s\n", db[i].Name, db[i].AbsPath, db[i].Template)
			return s
		}),
	)
	if err != nil {
		return Project{}, err
	}

	return db[idx], nil
}
