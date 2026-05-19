package subject

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/metadata"
)

// creates new Subject of type event, add it to module subjects and updates metadata file
// `ppth` is the parent dir abs path
func SubjectFactory(st SubjectType, ppth string) (Subject, error) {

	s := Subject{
		Type:    st,
		SubDirs: SubDirs[st],
		Files:   Files[st],
	}

	// call interactive to collect subject properties and fill subject object
	// CLI uses switch to decide propmpt fields
	// This should run before any other function since rest of function depends on
	// captured inputs from user
	if err := interactiveCLI(st, &s); err != nil {
		return Subject{}, err
	}

	// reset name and path since path depends on name
	s.Name = nameFactory(s)
	s.DirName = path.Join(ppth, s.Name)

	return s, nil
}

// Name factory return name based on user input name after sanitizing it and joins with another properties
// based on type. For example, event name should follow pattern `yyyy-MM-dd-user-sanitized-name`
func nameFactory(s Subject) string {

	switch s.Type {
	case SubjectEvent:
		sn := metadata.FormatName(s.Name)
		return fmt.Sprintf("%s-%s", s.Date, sn)
	case SubjectTask:
		sn := metadata.FormatName(s.Name)
		return fmt.Sprintf("%s-%s", s.Date, sn)
	case SubjectTopic:
		sn := metadata.FormatName(s.Name)
		return sn
	default:
		return s.Name
	}
}
