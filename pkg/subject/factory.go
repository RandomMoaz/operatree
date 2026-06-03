package subject

import (
	"fmt"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/metadata"
)

func SubjectFactory(s Subject, ppth string, pss []Subject) (Subject, error) {

	// requires interactive
	if s.Name == "" {
		ns, err := interactive(s.Type, ppth, pss)
		if err != nil {
			return s, err
		}

		return ns, nil
	}

	// Regular path is silent
	ns, err := silent(s, ppth)
	if err != nil {
		return s, err
	}

	return ns, nil
}

// creates new Subject of type event in silent mode add it to module subjects and updates metadata file.
// s is the subject, the caller is responsible to fill the least required fields .
// this function is commonly used with `--silent` flag.
// `ppth` is the parent dir abs path.
// `pss` all project subjects. Use in interactive CLI to attach related.
func silent(s Subject, ppth string) (Subject, error) {

	s.SubDirs = SubDirs[s.Type]
	s.Files = Files[s.Type]

	// reset name and path since path depends on name
	s.Name = nameFactory(s)
	s.DirName = filepath.Join(ppth, s.Name)

	return s, nil
}

// creates new Subject of type event, add it to module subjects and updates metadata file
// `ppth` is the parent dir abs path
// `pss` all project subjects. Use in interactive CLI to attach related
func interactive(st SubjectType, ppth string, pss []Subject) (Subject, error) {

	s := Subject{
		Type:    st,
		SubDirs: SubDirs[st],
		Files:   Files[st],
	}

	// call interactive to collect subject properties and fill subject object
	// CLI uses switch to decide propmpt fields
	// This should run before any other function since rest of function depends on
	// captured inputs from user
	if err := interactiveCLI(st, &s, pss); err != nil {
		return Subject{}, err
	}

	// reset name and path since path depends on name
	s.Name = nameFactory(s)
	s.DirName = filepath.Join(ppth, s.Name)

	return s, nil
}

// Name factory return name based on user input name after sanitizing it and joins with another properties
// based on type. For example, event name should follow pattern `yyyy-MM-dd-user-sanitized-name`
func nameFactory(s Subject) string {

	switch s.Type {
	case SubjectEvent:
		sn := metadata.FormatName(s.Name)
		if s.Date != "" {
			sn = fmt.Sprintf("%s-%s", s.Date, sn)
		}
		return sn
	case SubjectTask:
		sn := metadata.FormatName(s.Name)
		return fmt.Sprintf("%s-%s", s.Date, sn)
	case SubjectTopic:
		sn := metadata.FormatName(s.Name)
		return sn
	case SubjectDataSource:
		sn := metadata.FormatName(s.Name)
		return sn
	default:
		return s.Name
	}
}
